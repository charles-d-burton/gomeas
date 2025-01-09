package gomeas

import (
  "errors"
  "strings"
	jsoniter "github.com/json-iterator/tinygo"
)


//go:generate go run github.com/json-iterator/tinygo/gen
type Sensor struct {
  Config
  Platform string `json:"platform"`
}

func (sensor Sensor) Marshal() ([]byte, error) {
  if sensor.UniqueID == "" {
    return nil, errors.New("sensor:unique_id-undefined")
  }
  if sensor.Name == "" {
    return nil, errors.New("sensor:name-undefined")
  }
	json := jsoniter.CreateJsonAdapter(Sensor_json{}, Device_json{}, Components_json{}, Component_json{}, Availabilities_json{})
  return json.Marshal(sensor)
}

func (sensor Sensor) Topic() ([]byte, error) {
	if sensor.ConfigTopic == "" {
		return nil, errors.New("must set topic to initialize device discovery")
	}
	elements := strings.Split(sensor.ConfigTopic, "/")
	if len(elements) != 4 && len(elements) != 5 {
		return nil, errors.New("topic must be in the format: <discovery_prefix>/<component>/[<node_id>]/<object_id>/config")
	}
	if elements[len(elements)-1] != "config" {
		return nil, errors.New(`last field of the topic must be the word "config"`)
	}
  sensor.Platform = "sensor"
  if elements[1] != "sensor" {
    return nil, errors.New(`second field of topic must be the word "sensor"`)
  }
	return []byte(sensor.ConfigTopic), nil
}

