package gomeas

import (
	"context"
	"errors"
	"io"
	"time"

	mqtt "github.com/soypat/natiu-mqtt"
)

type MQTT struct {
	Username string
	Password string
	ClientID string
	client   *mqtt.Client
}

type Options struct {
	Username string
	Password string
	ClientID string
}

// ping pings the server every 30 seconds to maintain connection
func (mq *MQTT) ping(ctx context.Context) {
	println("info:starting background ping service")
	ticker := time.NewTicker(30 * time.Second)
	go func() {
		for {
			select {
			case <-ticker.C:
				if mq.client.IsConnected() {
					println("pinging server")
					err := mq.client.StartPing()
					if err != nil {
						println(err.Error())
					}
				}
			case <-ctx.Done():
				return
			}
		}
	}()
}

// handle manages the loop for receiving messages from mqtt
func (mq *MQTT) handle(ctx context.Context, tcon func() error) {
	go func() {
		for {
			if !mq.client.IsConnected() {
				time.Sleep(time.Second)
				if err := tcon(); err != nil {
					println(err.Error())
					continue
				}
				println("info:reconnected to mqtt broker")
				continue
			}
			select {
			case <-ctx.Done():
				println("info:received done signal, exiting handler")
				return
			default:
				println("info:handling inbound request")
				err := mq.client.HandleNext()
				if err != nil {
					mq.client.Disconnect(err)
				}
			}
		}
	}()
}

// TODO: Handle the dependency injection better
func NewMQTTConnection(ctx context.Context, rwc io.ReadWriteCloser, options *Options) (*MQTT, <-chan []byte, error) {
	if options.ClientID == "" {
		options.ClientID = GenerateId(8)
	}

	mq := &MQTT{
		Username: options.Username,
		Password: options.Password,
		ClientID: options.ClientID,
	}

	data, err := mq.start(ctx, rwc)
	if err != nil {
		return nil, nil, err
	}
	return mq, data, nil
}

// Start connects to MQTT using the ReadWriteCloser you specifiy.  Likely a tinygo network connection
func (mq *MQTT) start(ctx context.Context, rwc io.ReadWriteCloser) (<-chan []byte, error) {
	receiver := make(chan []byte, 2)
	cfg := mqtt.ClientConfig{
		Decoder: mqtt.DecoderNoAlloc{UserBuffer: make([]byte, 1500)}, OnPub: func(_ mqtt.Header, _ mqtt.VariablesPublish, r io.Reader) error {
			message, _ := io.ReadAll(r)
			if len(message) > 0 {
				select {
				case receiver <- message:
				case <-ctx.Done():
					return nil
				default:
					println("warn:writing to a full channel")
					//Ignores message if buffer full
				}
			}
			println("info:received message")
			return nil
		},
	}
	client := mqtt.NewClient(cfg)
	mq.client = client

	var varConn mqtt.VariablesConnect
	varConn.SetDefaultMQTT([]byte(mq.ClientID))
	varConn.Username = []byte(mq.Username)
	varConn.Password = []byte(mq.Password)
	println("info:connecting with username: " + mq.Username)

	tryconnect := func() error {
		println("info:trying connect")
		ctxwt, cancel := context.WithTimeout(ctx, 4*time.Second)
		defer cancel()
		// return client.StartConnect(rwc, &varConn)
		return client.Connect(ctxwt, rwc, &varConn)
	}

	err := tryconnect()
	if err != nil {
		println("error:connect attempt failed")
		return nil, err
	}
	println("info:connected to mqtt broker")
	mq.ping(ctx)
	mq.handle(ctx, tryconnect)

	return receiver, nil
}

func (mq *MQTT) Publisher(ctx context.Context, msgs <-chan Message) error {
	if !mq.client.IsConnected() {
		time.Sleep(time.Second)
		return errors.New("client is not connected")
	}

	go func() {
		println("info:starting mqtt publisher")
		pflags, err := mqtt.NewPublishFlags(mqtt.QoS0, false, false)
		if err != nil {
			panic(err)
		}
		for {
			if !mq.client.IsConnected() {
				println("error:client is disconnected, sleeping for 1 second and retrying")
				time.Sleep(time.Second)
				continue
			}

			select {
			case msg := <-msgs:
				println("info:received message to publish")

				topic, err := msg.GetTopic()
				if err != nil {
					println(err.Error())
					continue
				}
				println("info:topic:" + string(topic))

				vpub := mqtt.VariablesPublish{
					TopicName:        topic,
					PacketIdentifier: randInt16(),
				}
				data, err := msg.Marshal()
				if err != nil {
					println(err.Error())
					continue
				}
				println("info:publishing message")
				err = mq.client.PublishPayload(pflags, vpub, data)
				if err != nil {
					println("error:" + err.Error())
				}
				time.Sleep(time.Second)
			case <-ctx.Done():
				println("info:received done signal, exiting")
				return
			}
		}
	}()
	return nil
}

func (mq *MQTT) Subscribe(ctx context.Context, topic string) error {
	var vsub mqtt.VariablesSubscribe
	vsub.TopicFilters = []mqtt.SubscribeRequest{{TopicFilter: []byte(topic), QoS: mqtt.QoS0}}
	vsub.PacketIdentifier = randInt16()
	err := mq.client.Subscribe(ctx, vsub)
	if err != nil {
		return err
	}
	return nil
}
