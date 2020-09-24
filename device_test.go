package venstar

import (
	"testing"
)

func TestNewDevice(t *testing.T) {
	tests := []struct {
		name    string
		typ     string
		address string
		expErr  string
	}{
		{"both empty", "", "", "type or address empty"},
		{"type empty", "", "filled", "type or address empty"},
		{"address empty", "filled", "", "type or address empty"},
		{"valid", "type", "address", ""},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			device, err := NewDevice(test.typ, test.address)

			if test.expErr != "" && err == nil {
				t.Fatal("error expected but no error returned")
			}
			if test.expErr == "" && err != nil {
				t.Fatal("error unexpectedly returned: ", err)
			}
			if test.expErr != "" && err != nil && err.Error() != test.expErr {
				t.Fatal("error expected, got:", err.Error(), "want:", test.expErr)
			}
			if test.expErr != "" {
				return
			}
			if test.typ != device.Type {
				t.Error("device type invalid, got:", device.Type, "want:", test.typ)
			}
			if test.address != device.Address {
				t.Error("device address invalid, got:", device.Address, "want:", test.address)
			}
		})
	}
}
