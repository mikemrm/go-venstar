package venstar

import (
	"errors"

	"github.com/mikemrm/go-venstar/thermostat"
)

// Device encompasses any Venstar device
type Device struct {
	Type       string
	Address    string
	thermostat *thermostat.Thermostat
}

// Thermostat initializes a Thermostat instance if an instance has not already
// been created. Otherwise it returns the previous reference.
func (d *Device) Thermostat() *thermostat.Thermostat {
	if d.thermostat == nil {
		d.thermostat = thermostat.New(d.Address)
	}
	return d.thermostat
}

// NewDevice creates a new Device instance with the provided Type and Address
func NewDevice(typ, address string) (*Device, error) {
	if typ == "" || address == "" {
		return nil, errors.New("type or address empty")
	}
	return &Device{
		Type:    typ,
		Address: address,
	}, nil
}
