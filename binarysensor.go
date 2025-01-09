package gomeas

import (
	"errors"

	"github.com/mailru/easyjson"
)

// homeassistant.com/integrations/binary_sensor.mqtt/
//
//go:generate go run github.com/json-iterator/tinygo/gen ./generated/
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

func (b BinarySensor) Topic() ([]byte, error) {
	if b.ConfigTopic == "" {
		return nil, errors.New("ConfigTopic is empty")
	}
	return []byte(b.ConfigTopic), nil
}

func (b BinarySensor) Marshal() ([]byte, error) {

	if len(b.Availability) > 0 && b.AvailabilityTopic != nil {
		return nil, errors.New("avilability and availability topic are both set")
	}
	if len(b.Availability) > 0 {
		for _, av := range b.Availability {
			if av.Topic == "" {
				return nil, errors.New("availability topic is empty")
			}
		}
	}
	b.Platform = "binary_sensor" //enforced by documentation
	if b.StateTopic == "" {
		return nil, errors.New("state topic is empty")
	}
	// json := jsoniter.CreateJsonAdapter(BinarySensor_json{}, Device_json{})
	return easyjson.Marshal(b)
}

//go:generate go run github.com/json-iterator/tinygo/gen
type BinarySensorState struct {
	ConfigTopic string `json:"-"`
	State       string `json:"state"`
}

func (bss BinarySensorState) Topic() ([]byte, error) {
	if bss.ConfigTopic == "" {
		return nil, errors.New("ConfigTopic is empty")
	}
	return []byte(bss.ConfigTopic), nil
}

func (bss BinarySensorState) Marshal() ([]byte, error) {
	switch bss.State {
	case "ON", "OFF":
		// json := jsoniter.CreateJsonAdapter(BinarySensorState_json{})
		return easyjson.Marshal(bss)
	default:
		return nil, errors.New("state is not ON or OFF")
	}
}
