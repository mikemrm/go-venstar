package thermostat

import (
	"io/ioutil"
	"net/http"
	"net/url"
	"strconv"
	"strings"
)

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

func (sr *SettingsRequest) ScheduleOff() *SettingsRequest {
	return sr.SetSchedule(0)
}

func (sr *SettingsRequest) ScheduleOn() *SettingsRequest {
	return sr.SetSchedule(1)
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

func (sr *SettingsRequest) BuildRequest(req *http.Request) error {
	params := make(url.Values)
	if sr.TempUnits != nil {
		params.Set("tempunits", strconv.Itoa(*sr.TempUnits))
	}
	if sr.IsAway != nil {
		params.Set("away", strconv.Itoa(*sr.IsAway))
	}
	if sr.Schedule != nil {
		params.Set("schedule", strconv.Itoa(*sr.Schedule))
	}
	if sr.HumidifySetPoint != nil {
		params.Set("hum_setpoint", strconv.Itoa(*sr.HumidifySetPoint))
	}
	if sr.DehumidifySetPoint != nil {
		params.Set("dehum_setpoint", strconv.Itoa(*sr.DehumidifySetPoint))
	}
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	body := params.Encode()
	req.Body = ioutil.NopCloser(strings.NewReader(body))
	req.ContentLength = int64(len(body))
	return nil
}

func NewSettingsRequest() *SettingsRequest {
	return &SettingsRequest{}
}
