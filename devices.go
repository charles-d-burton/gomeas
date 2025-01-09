package gomeas

import (
	"errors"
	jsoniter "github.com/json-iterator/tinygo"
	"strings"
)

type Message interface {
	Marshal() ([]byte, error)
	Topic() ([]byte, error)
}

type DeviceConfig interface {
	GetConfig() ([]byte, error)
}

//go:generate go run github.com/json-iterator/tinygo/gen
type Components = map[string]Component

//go:generate go run github.com/json-iterator/tinygo/gen
type Config struct {
	ConfigTopic          string          `json:"-"`
	Availability         []*Availability `json:"availability,omitempty"`
	AvailibilityMode     *string         `json:"availibility_mode,omitempty"`
	AvailabilityTemplate *string         `json:"availability_template,omitempty"`
	AvailabilityTopic    *string         `json:"availability_topic,omitempty"`
	Name                 string          `json:"name"`
	DeviceClass          string          `json:"device_class"`
	StateTopic           *string         `json:"state_topic,omitempty"`
	ComamandTopic        *string         `json:"command_topic,omitempty"`
	UniqueID             string          `json:"unique_id"`
	Device               *Device         `json:"device,omitempty"`
	Components           Components      `json:"components,omitempty"`
	EnabledByDefault     bool            `json:"enabled_by_default,omitempty"`
}

//go:generate go run github.com/json-iterator/tinygo/gen
type Device struct {
	ConfigurationURL *string  `json:"configuration_url,omitempty"`
	Connections      []string `json:"connections,omitempty"`
	HardwareVersion  *string  `json:"hw_version,omitempty"`
	Identifiers      []string `json:"identifiers,omitempty"`
	Manufacturer     *string  `json:"manufacturer,omitempty"`
	Model            *string  `json:"model,omitempty"`
	ModelID          *string  `json:"model_id,omitempty"`
	Name             *string  `json:"name,omitempty"`
	SerialNumber     *string  `json:"serial_number,omitempty"`
	SuggestedArea    *string  `json:"suggested_area,omitempty"`
	SoftwareVersion  *string  `json:"software_version,omitempty"`
	ViaDevice        *string  `json:"via_device,omitempty"`
}

//go:generate go run github.com/json-iterator/tinygo/gen
type Availability struct {
	PayloadAvailable    string `json:"payload_available,omitempty"`
	PayloadNotAvailable string `json:"payload_not_available,omitempty"`
	Topic               string `json:"topic"`
	ValueTemplate       string `json:"value_template,omitempty"`
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
	if config.UniqueID == "" {
		return nil, errors.New("config:unique_id-undefined")
	}
	if config.Name == "" {
		return nil, errors.New("config:name-undefined")
	}
	if config.DeviceClass == "" {
		return nil, errors.New("config:device_class-undefined")
	}
	json := jsoniter.CreateJsonAdapter(Config_json{}, Device_json{}, Components_json{}, Component_json{})
	return json.Marshal(config)
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
	if elements[len(elements)-1] != "config" {
		return nil, errors.New(`last field of the topic must be the word "config"`)
	}
	return []byte(config.ConfigTopic), nil
}
