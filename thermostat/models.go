package thermostat

import (
	"time"
)

type APIInfo struct {
	Version  int    `json:"api_ver"`
	Model    string `json:"model"`
	Firmware string `json:"firmware"`
	Type     string `json:"type"`
}

type QueryInfo struct {
	Name               string  `json:"name"`
	Mode               int     `json:"mode"`
	State              int     `json:"state"`
	Fan                int     `json:"fan"`
	Fanstate           int     `json:"fanstate"`
	Tempunits          int     `json:"tempunits"`
	Schedule           int     `json:"schedule"`
	Schedulepart       int     `json:"schedulepart"`
	Away               int     `json:"away"`
	Holiday            int     `json:"holiday"`
	Override           int     `json:"override"`
	OverrideTime       int     `json:"overridetime"`
	Forceunocc         int     `json:"forceunocc"`
	SpaceTemp          float64 `json:"spacetemp"`
	HeatTemp           float64 `json:"heattemp"`
	CoolTemp           float64 `json:"cooltemp"`
	CoolTempMin        float64 `json:"cooltempmin"`
	CoolTempMax        float64 `json:"cooltempmax"`
	HeatTempMin        float64 `json:"heattempmin"`
	HeatTempMax        float64 `json:"heattempmax"`
	ActiveStage        int     `json:"activestage,omitempty"`
	HumidityEnabled    int     `json:"hum_active,omitempty"`
	Humidity           int     `json:"hum,omitempty"`
	HumidifySetPoint   int     `json:"hum_setpoint,omitempty"`
	DehumidifySetPoint int     `json:"dehum_setpoint,omitempty"`
	SetPointDelta      float64 `json:"setpointdelta"`
	AvailableModes     int     `json:"availablemodes"`
}

type QueryResponse struct {
	Sensors  []*Sensor  `json:"sensors,omitempty"`
	Runtimes []*Runtime `json:"runtimes,omitempty"`
	Alerts   []*Alert   `json:"alerts,omitempty"`
}

type Sensor struct {
	Name string `json:"name"`
	Temp int    `json:"temp"`
}

type Runtime struct {
	Time  time.Time `json:"ts"`
	Heat1 int       `json:"heat1"`
	Heat2 int       `json:"heat2"`
	Cool1 int       `json:"cool1"`
	Cool2 int       `json:"cool2"`
	Aux1  int       `json:"aux1"`
	Aux2  int       `json:"aux2"`
	FC    int       `json:"fc,omitempty"`
	OV    int       `json:"ov,omitempty"`
}

type Alert struct {
	Name   string `json:"alert"`
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
	Away               *int `json:"away,omitempty"`
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
	sr.Away = new(int)
	if away {
		*sr.Away = 1
	} else {
		*sr.Away = 0
	}
	return sr
}

// func (sr *SettingsRequest) Away() *SettingsRequest {
// 	return sr.SetAway(true)
// }

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
