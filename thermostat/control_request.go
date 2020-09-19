package thermostat

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
