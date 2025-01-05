package gomeas

import (
	"encoding/json"
	"errors"
)

type AlarmState int

const (
	DISARMED AlarmState = iota
	ARMED_HOME
	ARMED_AWAY
	ARMED_NIGHT
	ARMED_VACATION
	ARMED_CUSTOM_BYPASS
	PENDING
	TRIGGERED
	ARMING
	DISARMING
)

func (s AlarmState) String() string {
	switch s {
	case DISARMED:
		return "disarmed"
	case ARMED_HOME:
		return "armed_home"
	case ARMED_AWAY:
		return "armed_away"
	case ARMED_NIGHT:
		return "armed_night"
	case ARMED_VACATION:
		return "armed_vacation"
	case ARMED_CUSTOM_BYPASS:
		return "armed_custom_bypass"
	case PENDING:
		return "pending"
	case TRIGGERED:
		return "triggered"
	case ARMING:
		return "arming"
	case DISARMING:
		return "disarming"
	default:
		return ""
	}
}

func (s AlarmState) MarshalJSON() ([]byte, error) {
	if s.String() == "" {
		return nil, errors.New("alarm state cannot be empty")
	}
	return json.Marshal(s.String())
}

type AlarmControlPanel struct {
}
