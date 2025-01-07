package gomeas

import (
	"errors"
	"strings"
	jsoniter "github.com/json-iterator/tinygo"
)

type Message interface {
	Marshal() ([]byte, error)
	Topic() ([]byte, error)
}

type DeviceConfig interface {
	GetConfig() ([]byte, error)
}

// //go:generate go run github.com/json-iterator/tinygo/gen
// type ComponentMap = map[string]Component

//go:generate go run github.com/json-iterator/tinygo/gen
type Config struct {
	ConfigTopic   string       `json:"-"`
	Name          string       `json:"name"`
	DeviceClass   string       `json:"device_class"`
	StateTopic    *string       `json:"stat_t,omitempty"`
	ComamandTopic *string       `json:"cmd_t,omitempty"`
	UniqueID      string       `json:"unique_id"`
	Device        *Device       `json:"dev,omitempty"`
	Components    map[string]*Component `json:"cmps,omitempty"`
}

//go:generate go run github.com/json-iterator/tinygo/gen
type Device struct {
	ConfigurationURL *string   `json:"cu,omitempty"`
	Connections      []string `json:"cns,omitempty"`
	HardwareVersion  *string   `json:"hw,omitempty"`
	Identifiers      []string `json:"ids,omitempty"`
	Manufacturer     *string   `json:"mf,omitempty"`
	Model            *string   `json:"mdl,omitempty"`
	ModelID          *string   `json:"mdl_id,omitempty"`
	Name             *string   `json:"name,omitempty"`
	SerialNumber     *string   `json:"sn,omitempty"`
	SuggestedArea    *string   `json:"sa,omitempty"`
	SoftwareVersion  *string   `json:"sw,omitempty"`
	ViaDevice        *string   `json:"via_device,omitempty"`
}

const (
	CELSIUS    = `°C`
	FAHRENHEIT = `°F`
)

//go:generate go run github.com/json-iterator/tinygo/gen
type Component struct {
	P                 string `json:"p"`
	DeviceClass       string `json:"device_class"`
	UnitOfMeasurement string `json:"unit_of_measurement"`
	ValueTemplate     string `json:"value_template"`
	UniqueID          string `json:"unique_id"`
}

func (config Config) Marshal() ([]byte, error) {
	json := jsoniter.CreateJsonAdapter(Config_json{}, Device_json{}, Component_json{})
  return json.Marshal(&config)
}

func (config Config) Topic() ([]byte, error) {
	if config.ConfigTopic == "" {
		return nil, errors.New("must set configuration topic to initialize device discovery")
	}
	elements := strings.Split(config.ConfigTopic, "/")
	println(len(elements))
	if len(elements) != 4 && len(elements) != 5 {
		return nil, errors.New("topic must be in the format: <discovery_prefix>/<component>/[<node_id>]/<object_id>/config")
	}
	if elements[1] != "device" {
		return nil, errors.New(`second field of topic must be set to "device"`)
	}
	if elements[len(elements)-1] != "config" {
		return nil, errors.New(`last field of the topic must be the word "config"`)
	}
	return []byte(config.ConfigTopic), nil
}
