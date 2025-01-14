package gomeas

import (
	"errors"
	"strings"
)

type Components = map[string]Component

/*
Config is a struct that holds the configuration for a device

	These are the general defaults for most of the different auto-discovery
	messages that Homeassistant accepts
*/
type Config struct {
	ConfigTopic          string          `json:"-"`
	Availability         []*Availability `json:"availability,omitempty"`
	AvailibilityMode     *string         `json:"availibility_mode,omitempty"`
	AvailabilityTemplate *string         `json:"availability_template,omitempty"`
	AvailabilityTopic    *string         `json:"availability_topic,omitempty"`
	Name                 string          `json:"name"`
	DeviceClass          string          `json:"device_class"`
	StateTopic           *string         `json:"state_topic,omitempty"`
	CommandTopic         *string         `json:"command_topic,omitempty"`
	UniqueID             string          `json:"unique_id"`
	Device               *Device         `json:"device,omitempty"`
	Components           Components      `json:"components,omitempty"`
	EnabledByDefault     bool            `json:"enabled_by_default,omitempty"`
}

// Devices maps to the device specifications for Homeasisstant
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

// Availability is a struct that holds the availability configuration for a device
type Availability struct {
	PayloadAvailable    string `json:"payload_available,omitempty"`
	PayloadNotAvailable string `json:"payload_not_available,omitempty"`
	Topic               string `json:"topic"`
	ValueTemplate       string `json:"value_template,omitempty"`
}

type Component struct {
	P                 string `json:"p"`
	DeviceClass       string `json:"device_class"`
	UnitOfMeasurement string `json:"unit_of_measurement"`
	ValueTemplate     string `json:"value_template"`
	UniqueID          string `json:"unique_id"`
}

func (config *Config) validateComponent(component string) error {
	if config.ConfigTopic == "" {
		return errors.New("must set topic to initialize device discovery")
	}

	elements := strings.Split(config.ConfigTopic, "/")
	if len(elements) != 4 && len(elements) != 5 {
		return errors.New("error:topic must be in the format: <discovery_prefix>/<component>/[<node_id>]/<object_id>/config")
	}

	if elements[len(elements)-1] != "config" {
		return errors.New(`error:last field of the topic must be the word "config"`)
	}

	if elements[1] != component {
		return errors.New(`error:second field of topic must be the word ` + component)
	}
	return nil
}

func (config *Config) validateAvailability() error {
	if len(config.Availability) > 0 && config.AvailabilityTopic != nil {
		return errors.New("error:avilability and availability topic are both set")
	}
	if len(config.Availability) > 0 {
		for _, av := range config.Availability {
			if av.Topic == "" {
				return errors.New("error:availability topic is empty")
			}
		}
	}
	return nil
}
