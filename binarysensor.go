package gomeas

import (
	"errors"

	"github.com/mailru/easyjson"
)

// homeassistant.com/integrations/binary_sensor.mqtt/
type BinarySensor struct {
	Config
	Encoding               string `json:"encoding,omitempty"`
	EntityCategory         string `json:"entity_category,omitempty"`
	EntityPicture          string `json:"entity_picture,omitempty"`
	ExpireAfter            int    `json:"expire_after,omitempty"`
	ForceUpdate            bool   `json:"force_update,omitempty"`
	Icon                   string `json:"icon,omitempty"`
	JsonAttributesTemplate string `json:"json_attributes_template,omitempty"`
	JsonAttributesTopic    string `json:"json_attributes_topic,omitempty"`
	Name                   string `json:"name,omitempty"`
	ObjectID               string `json:"object_id,omitempty"`
	OffDelay               int    `json:"off_delay,omitempty"`
	PayloadAvailable       string `json:"payload_available,omitempty"`
	PayloadNotAvailable    string `json:"payload_not_available,omitempty"`
	PayloadOff             string `json:"payload_off,omitempty"`
	PayloadOn              string `json:"payload_on,omitempty"`
	Platform               string `json:"platform"`
	Qos                    int    `json:"qos,omitempty"`
	StateTopic             string `json:"state_topic"`
	UniqueID               string `json:"unique_id,omitempty"`
	ValueTemplate          string `json:"value_template,omitempty"`
}

func NewBinarySensor(config Config) (*BinarySensor, error) {
	return &BinarySensor{Config: config}, nil
}

func (b *BinarySensor) GetTopic() ([]byte, error) {
	err := b.validateComponent("binary_sensor")
	if err != nil {
		return nil, err
	}
	return []byte(b.ConfigTopic), nil
}

func (b *BinarySensor) Marshal() ([]byte, error) {
	b.Platform = "binary_sensor" //enforced by documentation
	if b.StateTopic == "" {
		return nil, errors.New("error:state topic is empty")
	}
	return easyjson.Marshal(b)
}

type BinarySensorState struct {
	StateTopic string `json:"-"`
	State      string `json:"state"`
}

func (bss *BinarySensorState) GetTopic() ([]byte, error) {
	if bss.StateTopic == "" {
		return nil, errors.New("error: state topic is empty")
	}
	return []byte(bss.StateTopic), nil
}

func (bss *BinarySensorState) Marshal() ([]byte, error) {
	switch bss.State {
	case "ON", "OFF":
		return easyjson.Marshal(bss)
	default:
		return nil, errors.New("error:state is not ON or OFF")
	}
}
