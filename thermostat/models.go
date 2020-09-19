package thermostat

import (
	"encoding/json"
	"strings"
	"time"
)

type APIInfo struct {
	Version  int    `json:"api_ver"`
	Model    string `json:"model"`
	Firmware string `json:"firmware"`
	Type     string `json:"type"`
}

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

type QueryResponse struct {
	Sensors  []*Sensor  `json:"sensors,omitempty"`
	Runtimes []*Runtime `json:"runtimes,omitempty"`
	Alerts   []*Alert   `json:"alerts,omitempty"`
}

type Sensor struct {
	Name string  `json:"name"`
	Temp float64 `json:"temp"`
}

type Runtime struct {
	Timestamp   time.Time
	Heaters     map[string]time.Duration
	Coolers     map[string]time.Duration
	Aux         map[string]time.Duration
	FreeCooling time.Duration
	Override    time.Duration
}

func (r *Runtime) UnmarshalJSON(data []byte) error {
	jdata := make(map[string]int)
	if err := json.Unmarshal(data, &jdata); err != nil {
		return err
	}
	r.Heaters = make(map[string]time.Duration)
	r.Coolers = make(map[string]time.Duration)
	r.Aux = make(map[string]time.Duration)
	for k, v := range jdata {
		if k == "ts" {
			r.Timestamp = time.Unix(int64(v), 0)
		} else if k == "fc" {
			r.FreeCooling = time.Duration(v) * time.Minute
		} else if k == "ov" {
			r.Override = time.Duration(v) * time.Minute
		} else if strings.HasPrefix(k, "heat") {
			r.Heaters[strings.TrimPrefix(k, "heat")] = time.Duration(v) * time.Minute
		} else if strings.HasPrefix(k, "cool") {
			r.Coolers[strings.TrimPrefix(k, "cool")] = time.Duration(v) * time.Minute
		} else if strings.HasPrefix(k, "aux") {
			r.Aux[strings.TrimPrefix(k, "aux")] = time.Duration(v) * time.Minute
		}
	}
	return nil
}

type Alert struct {
	Name   string `json:"name"`
	Active bool   `json:"active"`
}

type ControlRequest struct {
	Mode     *int `json:"mode,omitempty"`
	Fan      *int `json:"fan,omitempty"`
	HeatTemp *int `json:"heattemp,omitempty"`
	CoolTemp *int `json:"cooltemp,omitempty"`
}

func (cr *ControlRequest) SetMode(value int) *ControlRequest {
	cr.Mode = new(int)
	*cr.Mode = value
	return cr
}

func (cr *ControlRequest) SetFan(value int) *ControlRequest {
	cr.Fan = new(int)
	*cr.Fan = value
	return cr
}

func (cr *ControlRequest) FanAuto() *ControlRequest {
	return cr.SetFan(0)
}

func (cr *ControlRequest) FanOn() *ControlRequest {
	return cr.SetFan(1)
}

func (cr *ControlRequest) SetHeatTemp(value int) *ControlRequest {
	cr.HeatTemp = new(int)
	*cr.HeatTemp = value
	return cr
}

func (cr *ControlRequest) SetCoolTemp(value int) *ControlRequest {
	cr.CoolTemp = new(int)
	*cr.CoolTemp = value
	return cr
}

func (cr *ControlRequest) Off(cool, heat int) *ControlRequest {
	return cr.SetMode(0).SetCoolTemp(cool).SetHeatTemp(heat)
}

func (cr *ControlRequest) Heat(temp int) *ControlRequest {
	return cr.SetMode(1).SetHeatTemp(temp)
}

func (cr *ControlRequest) Cool(temp int) *ControlRequest {
	return cr.SetMode(2).SetCoolTemp(temp)
}

func (cr *ControlRequest) Auto(cool, heat int) *ControlRequest {
	return cr.SetMode(3).SetCoolTemp(cool).SetHeatTemp(heat)
}

type SettingsRequest struct {
	TempUnits          *int `json:"tempunits,omitempty"`
	IsAway             *int `json:"away,omitempty"`
	Schedule           *int `json:"schedule,omitempty"`
	HumidifySetPoint   *int `json:"hum_setpoint,omitempty"`
	DehumidifySetPoint *int `json:"dehum_setpoint,omitempty"`
}

func (sr *SettingsRequest) SetTempUnits(value int) *SettingsRequest {
	sr.TempUnits = new(int)
	*sr.TempUnits = value
	return sr
}

func (sr *SettingsRequest) Fahrenheit() *SettingsRequest {
	return sr.SetTempUnits(0)
}

func (sr *SettingsRequest) Celsius() *SettingsRequest {
	return sr.SetTempUnits(1)
}

func (sr *SettingsRequest) SetAway(away bool) *SettingsRequest {
	sr.IsAway = new(int)
	if away {
		*sr.IsAway = 1
	} else {
		*sr.IsAway = 0
	}
	return sr
}

func (sr *SettingsRequest) Away() *SettingsRequest {
	return sr.SetAway(true)
}

func (sr *SettingsRequest) Home() *SettingsRequest {
	return sr.SetAway(false)
}

func (sr *SettingsRequest) SetSchedule(value int) *SettingsRequest {
	sr.Schedule = new(int)
	*sr.Schedule = value
	return sr
}

func (sr *SettingsRequest) SetHumidifySetPoint(value int) *SettingsRequest {
	sr.HumidifySetPoint = new(int)
	*sr.HumidifySetPoint = value
	return sr
}

func (sr *SettingsRequest) SetDehumidifySetPoint(value int) *SettingsRequest {
	sr.DehumidifySetPoint = new(int)
	*sr.DehumidifySetPoint = value
	return sr
}

type ErrorResponse struct {
	Error  bool   `json:"error"`
	Reason string `json:"reason,omitempty"`
}
