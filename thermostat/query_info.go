package thermostat

import (
	"time"
)

// QueryInfo is the struct returned for `/query/info` requests.
type QueryInfo struct {
	Name               string            `json:"name"`
	Mode               Mode              `json:"mode"`
	State              State             `json:"state"`
	Fan                Fan               `json:"fan"`
	FanState           FanState          `json:"fanstate"`
	TempUnits          TempUnits         `json:"tempunits"`
	Schedule           Schedule          `json:"schedule"`
	SchedulePart       SchedulePart      `json:"schedulepart"`
	Away               Away              `json:"away"`
	Holiday            Holiday           `json:"holiday"`
	Override           Override          `json:"override"`
	OverrideRemaining  OverrideRemaining `json:"overridetime"`
	ForceUnoccupied    ForceUnoccupied   `json:"forceunocc"`
	SpaceTemp          float64           `json:"spacetemp"`
	HeatTemp           float64           `json:"heattemp"`
	CoolTemp           float64           `json:"cooltemp"`
	CoolTempMin        float64           `json:"cooltempmin"`
	CoolTempMax        float64           `json:"cooltempmax"`
	HeatTempMin        float64           `json:"heattempmin"`
	HeatTempMax        float64           `json:"heattempmax"`
	ActiveStage        int               `json:"activestage,omitempty"`
	HumidityEnabled    HumidityEnabled   `json:"hum_active,omitempty"`
	Humidity           int               `json:"hum,omitempty"`
	HumidifySetPoint   int               `json:"hum_setpoint,omitempty"`
	DehumidifySetPoint int               `json:"dehum_setpoint,omitempty"`
	SetPointDelta      float64           `json:"setpointdelta"`
	AvailableModes     AvailableModes    `json:"availablemodes"`
}

// Mode allows for a string representation of the value to be returned.
type Mode int

// String returns a string representation of the value.
func (m Mode) String() string {
	switch m {
	case 0:
		return "off"
	case 1:
		return "heat"
	case 2:
		return "cool"
	case 3:
		return "auto"
	}
	return ""
}

// State allows for a string representation of the value to be returned.
type State int

// String returns a string representation of the value.
func (s State) String() string {
	switch s {
	case 0:
		return "idle"
	case 1:
		return "heating"
	case 2:
		return "cooling"
	case 3:
		return "lockout"
	case 4:
		return "error"
	}
	return ""
}

// Fan allows for a string representation of the value to be returned.
type Fan int

// String returns a string representation of the value.
func (f Fan) String() string {
	switch f {
	case 0:
		return "auto"
	case 1:
		return "on"
	}
	return ""
}

// FanState allows for a string representation of the value to be returned.
type FanState int

// String returns a string representation of the value.
func (f FanState) String() string {
	switch f {
	case 0:
		return "off"
	case 1:
		return "on"
	}
	return ""
}

// TempUnits allows for a string representation of the value to be returned.
type TempUnits int

// String returns a string representation of the value.
func (f TempUnits) String() string {
	switch f {
	case 0:
		return "fahrenheit"
	case 1:
		return "celsius"
	}
	return ""
}

// Schedule allows for a string representation of the value to be returned.
type Schedule int

// String returns a string representation of the value.
func (f Schedule) String() string {
	switch f {
	case 0:
		return "inactive"
	case 1:
		return "active"
	}
	return ""
}

// SchedulePart allows for a string representation of the value to be returned.
type SchedulePart int

// String returns a string representation of the value.
func (f SchedulePart) String() string {
	switch f {
	case 0:
		return "morning"
	case 1:
		return "day"
	case 2:
		return "evening"
	case 3:
		return "night"
	case 255:
		return "inactive"
	}
	return ""
}

// Away allows for a string representation of the value to be returned.
type Away int

// String returns a string representation of the value.
func (f Away) String() string {
	switch f {
	case 0:
		return "home"
	case 1:
		return "away"
	}
	return ""
}

// Holiday allows for a string representation of the value to be returned.
type Holiday int

// String returns a string representation of the value.
func (f Holiday) String() string {
	switch f {
	case 0:
		return "not observing"
	case 1:
		return "observing"
	}
	return ""
}

// Override allows for a string representation of the value to be returned.
type Override int

// String returns a string representation of the value.
func (f Override) String() string {
	switch f {
	case 0:
		return "off"
	case 1:
		return "on"
	}
	return ""
}

// OverrideRemaining allows for a string representation of the value to be
// returned.
type OverrideRemaining int

// String returns a string representation of the value.
func (f OverrideRemaining) String() string {
	return (time.Duration(f) * time.Minute).String()
}

// ForceUnoccupied allows for a string representation of the value to be
// returned.
type ForceUnoccupied int

// String returns a string representation of the value.
func (f ForceUnoccupied) String() string {
	switch f {
	case 0:
		return "off"
	case 1:
		return "on"
	}
	return ""
}

// HumidityEnabled allows for a string representation of the value to be
// returned.
type HumidityEnabled int

// String returns a string representation of the value.
func (f HumidityEnabled) String() string {
	switch f {
	case 0:
		return "disabled"
	case 1:
		return "enabled"
	}
	return ""
}

// AvailableModes allows for a string representation of the value to be
// returned.
type AvailableModes int

// String returns a string representation of the value.
func (f AvailableModes) String() string {
	switch f {
	case 0:
		return "all"
	case 1:
		return "heat/cool"
	case 2:
		return "heat"
	case 3:
		return "cool"
	}
	return ""
}
