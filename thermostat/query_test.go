package thermostat

import (
	"strconv"
	"testing"
)

func TestGetAPIInfo(t *testing.T) {
	t.Run("errors get returned", func(t *testing.T) {
		tstat := &Thermostat{
			client: &fakeThermostatClient{
				error: true,
			},
		}

		_, err := tstat.GetAPIInfo()
		if err == nil {
			t.Fatal("error expected but no error returned")
		}
		want := "processing api info request: requesting /: this is an error"
		if err.Error() != want {
			t.Error("error expected, got:", err.Error(), "want:", want)
		}
	})
	tests := []struct {
		name  string
		field string
		value string
		isstr bool
	}{
		{"Version is set", "api_ver", "5", false},
		{"Model is set", "model", `"modelvalue"`, true},
		{"Firmware is set", "firmware", `"firmwarevalue"`, true},
		{"Type is set", "type", `"typevalue"`, true},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			tstat := &Thermostat{
				client: &fakeThermostatClient{
					body: `{"` + test.field + `": ` + test.value + `}`,
				},
			}

			info, err := tstat.GetAPIInfo()
			if err != nil {
				t.Fatal("error unexpected, got:", err)
			}
			var iwant int
			want := test.value
			if test.isstr {
				want = want[1 : len(want)-1]
			} else {
				iwant, err = strconv.Atoi(want)
				if err != nil {
					t.Fatal("Failed to convert test value to int:", test.value)
				}
			}
			if test.field == "api_ver" && info.Version != iwant {
				t.Error("Version incorrect, got:", info.Version, "want:", iwant)
			}
			if test.field == "model" && info.Model != want {
				t.Error("Model incorrect, got:", info.Model, "want:", want)
			}
			if test.field == "firmware" && info.Firmware != want {
				t.Error("Firmware incorrect, got:", info.Firmware, "want:", want)
			}
			if test.field == "type" && info.Type != want {
				t.Error("Type incorrect, got:", info.Type, "want:", want)
			}
		})
	}
}
