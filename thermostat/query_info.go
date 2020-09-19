package thermostat

import (
	"time"
)

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

type Mode int

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

type State int

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

type Fan int

func (f Fan) String() string {
	switch f {
	case 0:
		return "auto"
	case 1:
		return "on"
	}
	return ""
}

type FanState int

func (f FanState) String() string {
	switch f {
	case 0:
		return "off"
	case 1:
		return "on"
	}
	return ""
}

type TempUnits int

func (f TempUnits) String() string {
	switch f {
	case 0:
		return "fahrenheit"
	case 1:
		return "celsius"
	}
	return ""
}

type Schedule int

func (f Schedule) String() string {
	switch f {
	case 0:
		return "active"
	case 1:
		return "inactive"
	}
	return ""
}

type SchedulePart int

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

type Away int

func (f Away) String() string {
	switch f {
	case 0:
		return "home"
	case 1:
		return "away"
	}
	return ""
}

type Holiday int

func (f Holiday) String() string {
	switch f {
	case 0:
		return "not observing"
	case 1:
		return "observing"
	}
	return ""
}

type Override int

func (f Override) String() string {
	switch f {
	case 0:
		return "off"
	case 1:
		return "on"
	}
	return ""
}

type OverrideRemaining int

func (f OverrideRemaining) String() string {
	return (time.Duration(f) * time.Minute).String()
}

type ForceUnoccupied int

func (f ForceUnoccupied) String() string {
	switch f {
	case 0:
		return "off"
	case 1:
		return "on"
	}
	return ""
}

type HumidityEnabled int

func (f HumidityEnabled) String() string {
	switch f {
	case 0:
		return "disabled"
	case 1:
		return "enabled"
	}
	return ""
}

type AvailableModes int

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
