package gomeas

import (
	"errors"

	"github.com/mailru/easyjson"
)

type Sensor struct {
	Config
	Platform string `json:"platform"`
}

func NewSensor(config Config) (*Sensor, error) {
	return &Sensor{Config: config}, nil
}

func (sensor *Sensor) GetConfigTopic() ([]byte, error) {
	err := sensor.validateComponent("sensor")
	if err != nil {
		return nil, err
	}
	sensor.Platform = "sensor"
	return []byte(sensor.ConfigTopic), nil
}

func (sensor *Sensor) Marshal() ([]byte, error) {
	err := sensor.validateAvailability()
	if err != nil {
		return nil, err
	}
	if sensor.UniqueID == "" {
		return nil, errors.New("error:unique_id-undefined")
	}
	if sensor.Name == "" {
		return nil, errors.New("error:name-undefined")
	}
	return easyjson.Marshal(sensor)
}
