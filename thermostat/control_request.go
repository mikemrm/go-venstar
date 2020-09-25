package thermostat

import (
	"errors"
	"io/ioutil"
	"net/http"
	"net/url"
	"strconv"
	"strings"
)

type ControlRequest struct {
	Mode      *int `json:"mode,omitempty"`
	Fan       *int `json:"fan,omitempty"`
	HeatTemp  *int `json:"heattemp,omitempty"`
	CoolTemp  *int `json:"cooltemp,omitempty"`
	validator func(*ControlRequest) error
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

func (cr *ControlRequest) Heat(heat, cool int) *ControlRequest {
	return cr.SetMode(1).SetHeatTemp(heat).SetCoolTemp(cool)
}

func (cr *ControlRequest) Cool(cool, heat int) *ControlRequest {
	return cr.SetMode(2).SetCoolTemp(cool).SetHeatTemp(heat)
}

func (cr *ControlRequest) Auto(cool, heat int) *ControlRequest {
	return cr.SetMode(3).SetCoolTemp(cool).SetHeatTemp(heat)
}

func (cr *ControlRequest) Validate() error {
	return cr.validator(cr)
}

func (cr *ControlRequest) BuildRequest(req *http.Request) error {
	err := cr.Validate()
	if err != nil {
		return err
	}
	params := make(url.Values)
	if cr.Mode != nil {
		params.Set("mode", strconv.Itoa(*cr.Mode))
	}
	if cr.Fan != nil {
		params.Set("fan", strconv.Itoa(*cr.Fan))
	}
	if cr.HeatTemp != nil {
		params.Set("heattemp", strconv.Itoa(*cr.HeatTemp))
	}
	if cr.CoolTemp != nil {
		params.Set("cooltemp", strconv.Itoa(*cr.CoolTemp))
	}
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	body := params.Encode()
	req.Body = ioutil.NopCloser(strings.NewReader(body))
	req.ContentLength = int64(len(body))
	return nil
}

func defaultControlRequestValidator(cr *ControlRequest) error {
	// All control calls with mode must include heattemp and cooltemp parameters.
	if cr.Mode != nil {
		if cr.HeatTemp == nil {
			return errors.New("HeatTemp must be defined when Mode is defined")
		}
		if cr.CoolTemp == nil {
			return errors.New("CoolTemp must be defined when Mode is defined")
		}
	}
	// When setting mode to Auto, cooltemp must be greater than heattemp and the setpointdelta from "/query/info" needs to be respected
	if cr.Mode != nil && *cr.Mode == 3 {
		if *cr.CoolTemp <= *cr.HeatTemp {
			return errors.New("CoolTemp must be greater than HeatTemp when Mode is Auto")
		}
	}
	return nil
}

func NewControlRequest() *ControlRequest {
	return &ControlRequest{
		validator: defaultControlRequestValidator,
	}
}
