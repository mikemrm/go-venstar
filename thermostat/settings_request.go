package thermostat

import (
	"io"
	"net/http"
	"net/url"
	"strconv"
	"strings"
)

// SettingsResquest is the object used to update thermostat settings.
// Any attribute set to null will not be included in the request.
type SettingsRequest struct {
	TempUnits          *int `json:"tempunits,omitempty"`
	IsAway             *int `json:"away,omitempty"`
	Schedule           *int `json:"schedule,omitempty"`
	HumidifySetPoint   *int `json:"hum_setpoint,omitempty"`
	DehumidifySetPoint *int `json:"dehum_setpoint,omitempty"`
}

// SetTempUnits sets the temperature units 0:fahrenheit 1:celsius
func (sr *SettingsRequest) SetTempUnits(value int) *SettingsRequest {
	sr.TempUnits = new(int)
	*sr.TempUnits = value
	return sr
}

// Fahrenheit is a shortcut to `SetTempUnits(0)`
func (sr *SettingsRequest) Fahrenheit() *SettingsRequest {
	return sr.SetTempUnits(0)
}

// Celsius is a shortcut to `SetTempUnits(1)`
func (sr *SettingsRequest) Celsius() *SettingsRequest {
	return sr.SetTempUnits(1)
}

// SetAway sets whether the conditioned space is occupied or not
func (sr *SettingsRequest) SetAway(away bool) *SettingsRequest {
	sr.IsAway = new(int)
	if away {
		*sr.IsAway = 1
	} else {
		*sr.IsAway = 0
	}
	return sr
}

// Away is a shortcut to `SetAway(true)`
func (sr *SettingsRequest) Away() *SettingsRequest {
	return sr.SetAway(true)
}

// Home is a shortcut to `SetAway(false)`
func (sr *SettingsRequest) Home() *SettingsRequest {
	return sr.SetAway(false)
}

// SetSchedule sets whether the thermostat should be following the predefined
// schedule
func (sr *SettingsRequest) SetSchedule(value int) *SettingsRequest {
	sr.Schedule = new(int)
	*sr.Schedule = value
	return sr
}

// ScheduleOff is a shortcut to `SetSchedule(0)`
func (sr *SettingsRequest) ScheduleOff() *SettingsRequest {
	return sr.SetSchedule(0)
}

// ScheduleOn is a shortcut to `SetSchedule(1)`
func (sr *SettingsRequest) ScheduleOn() *SettingsRequest {
	return sr.SetSchedule(1)
}

// SetHumidifySetPoint sets the percent humidity set point
func (sr *SettingsRequest) SetHumidifySetPoint(value int) *SettingsRequest {
	sr.HumidifySetPoint = new(int)
	*sr.HumidifySetPoint = value
	return sr
}

// SetDehumidifySetPoint sets the percent dehumidify set point
func (sr *SettingsRequest) SetDehumidifySetPoint(value int) *SettingsRequest {
	sr.DehumidifySetPoint = new(int)
	*sr.DehumidifySetPoint = value
	return sr
}

// BuildRequest applys the necessary request changes to the provided request.
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
	req.Body = io.NopCloser(strings.NewReader(body))
	req.ContentLength = int64(len(body))
	return nil
}

// NewSettingsRequest initializes a new SettingsRequest object
func NewSettingsRequest() *SettingsRequest {
	return &SettingsRequest{}
}
