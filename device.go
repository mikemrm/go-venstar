package venstar

import (
	"errors"

	"github.com/mikemrm/go-venstar/thermostat"
)

type Device struct {
	Type       string
	Address    string
	thermostat *thermostat.Thermostat
}

func (d *Device) Thermostat() *thermostat.Thermostat {
	if d.thermostat == nil {
		d.thermostat = thermostat.New(d.Address)
	}
	return d.thermostat
}

func NewDevice(typ, address string) (*Device, error) {
	if typ == "" || address == "" {
		return nil, errors.New("type or address empty")
	}
	return &Device{
		Type:    typ,
		Address: address,
	}, nil
}
