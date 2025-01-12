package gomeas

import (
	"errors"
	"github.com/mailru/easyjson"
)

// https://www.home-assistant.io/integrations/button.mqtt/
type Button struct {
	Config
	CommandTemplate string `json:"command_template,omitempty"`
	Platform        string `json:"platform"`
}

func NewButton(config Config) (*Button, error) {
	return &Button{Config: config}, nil
}

func (b *Button) GetConfigTopic() ([]byte, error) {
	err := b.validateComponent("button")
	if err != nil {
		return nil, err
	}
	return []byte(b.ConfigTopic), nil
}

func (b *Button) Marshal() ([]byte, error) {
	b.Platform = "button"
	if b.CommandTopic == nil || *b.CommandTopic == "" {
		return nil, errors.New("error:command topic is empty")
	}
	return easyjson.Marshal(b)
}
