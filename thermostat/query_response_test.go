package thermostat

import (
	"fmt"
	"strconv"
	"strings"
	"testing"
)

func TestGetQuerySensors(t *testing.T) {
	t.Run("errors get returned", func(t *testing.T) {
		tstat := &Thermostat{
			client: &fakeThermostatClient{
				error: true,
			},
		}

		_, err := tstat.GetQuerySensors()
		if err == nil {
			t.Fatal("error expected but no error returned")
		}
		want := "processing query sensors request: requesting /query/sensors: this is an error"
		if err.Error() != want {
			t.Error("error expected, got:", err.Error(), "want:", want)
		}
	})
	tests := []struct {
		name       string
		sensorName string
		sensorTemp float64
		expErr     string
	}{
		{"empty", "", 0.0, ""},
		{"name missing", "", 15.5, ""},
		{"temp missing", "sensorname", 0.0, ""},
		{"valid", "sensorname", 5.5, ""},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			var fields []string
			if test.sensorName != "" {
				fields = append(fields, fmt.Sprintf(`"name": "%s"`, test.sensorName))
			}
			if test.sensorTemp != 0.0 {
				fields = append(fields, fmt.Sprintf(`"temp": %f`, test.sensorTemp))
			}
			sensors := "{" + strings.Join(fields, ",") + "}"
			tstat := &Thermostat{
				client: &fakeThermostatClient{
					body: fmt.Sprintf(`{"sensors": [%s]}`, sensors),
				},
			}

			info, err := tstat.GetQuerySensors()

			if test.expErr != "" && err == nil {
				t.Fatal("error expected, got: nil want:", test.expErr)
			}
			if test.expErr != "" && err != nil && err.Error() != test.expErr {
				t.Fatal("error expected, got:", err.Error(), "want:", test.expErr)
			}
			if len(info) != 1 {
				t.Fatal("unexpected count of sensors returned, got:", len(info), "want: 1))")
			}
			if test.sensorName != info[0].Name {
				t.Error("sensor name incorrect, got:", info[0].Name, "want:", test.sensorName)
			}
			if test.sensorTemp != info[0].Temp {
				t.Error("sensor name incorrect, got:", info[0].Temp, "want:", test.sensorTemp)
			}
		})
	}
}

func TestGetQueryRuntimes(t *testing.T) {
	t.Run("errors get returned", func(t *testing.T) {
		tstat := &Thermostat{
			client: &fakeThermostatClient{
				error: true,
			},
		}

		_, err := tstat.GetQueryRuntimes()
		if err == nil {
			t.Fatal("error expected but no error returned")
		}
		want := "processing query runtime request: requesting /query/runtimes: this is an error"
		if err.Error() != want {
			t.Error("error expected, got:", err.Error(), "want:", want)
		}
	})

	tests := []struct {
		name   string
		field  string
		value  int
		want   string
		expErr string
	}{
		{"bad value returns error", "float", 0, "", "processing query runtime request: decoding /query/runtimes response: decoding json: json: cannot unmarshal number 1.5 into Go struct field QueryResponse.runtimes of type int"},
		{"timestamp", "ts", 1600984738, "2020-09-24 21:58:58 +0000 UTC", ""},
		{"free cooling", "fc", 10, "10m0s", ""},
		{"override", "ov", 20, "20m0s", ""},
		{"1 heater", "heat", 1, "10m0s", ""},
		{"2 heaters", "heat", 2, "20m0s", ""},
		{"1 cooler", "cool", 1, "10m0s", ""},
		{"2 coolers", "cool", 2, "20m0s", ""},
		{"1 aux", "aux", 1, "10m0s", ""},
		{"2 aux", "aux", 2, "20m0s", ""},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			var parts []string

			if test.field == "heat" || test.field == "cool" || test.field == "aux" {
				for i := 0; i < test.value; i++ {
					parts = append(parts, fmt.Sprintf(`"%s%d": %d`, test.field, i+1, (i+1)*10))
				}
			} else if test.field == "float" {
				parts = append(parts, `"float": 1.5`)
			} else {
				parts = append(parts, fmt.Sprintf(`"%s": %d`, test.field, test.value))
			}

			tstat := &Thermostat{
				client: &fakeThermostatClient{
					body: fmt.Sprintf(`{"runtimes": [{%s}]}`, strings.Join(parts, ",")),
				},
			}

			info, err := tstat.GetQueryRuntimes()

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
			if len(info) != 1 {
				t.Fatal("unexpected count of runtimes returned, got:", len(info), "want: 1")
			}

			if test.field == "ts" && info[0].Timestamp.String() != test.want {
				t.Error("Timestamp invalid, got:", info[0].Timestamp.String(), "want:", test.want)
			}
			if test.field == "fc" && info[0].FreeCooling.String() != test.want {
				t.Error("FreeCooling invalid, got:", info[0].FreeCooling.String(), "want:", test.want)
			}
			if test.field == "ov" && info[0].Override.String() != test.want {
				t.Error("Override invalid, got:", info[0].Override.String(), "want:", test.want)
			}
			if test.field == "heat" && info[0].Heaters[strconv.Itoa(test.value)].String() != test.want {
				t.Error("Heater invalid, got:", info[0].Heaters[strconv.Itoa(test.value)].String(), "want:", test.want)
			}
			if test.field == "cool" && info[0].Coolers[strconv.Itoa(test.value)].String() != test.want {
				t.Error("Cooler invalid, got:", info[0].Coolers[strconv.Itoa(test.value)].String(), "want:", test.want)
			}
			if test.field == "aux" && info[0].Aux[strconv.Itoa(test.value)].String() != test.want {
				t.Error("Aux invalid, got:", info[0].Aux[strconv.Itoa(test.value)].String(), "want:", test.want)
			}
		})
	}
}

func TestGetQueryAlerts(t *testing.T) {
	t.Run("errors get returned", func(t *testing.T) {
		tstat := &Thermostat{
			client: &fakeThermostatClient{
				error: true,
			},
		}

		_, err := tstat.GetQueryAlerts()
		if err == nil {
			t.Fatal("error expected but no error returned")
		}
		want := "processing query alerts request: requesting /query/alerts: this is an error"
		if err.Error() != want {
			t.Error("error expected, got:", err.Error(), "want:", want)
		}
	})
	tests := []struct {
		name        string
		alertName   string
		alertActive int
		expErr      string
	}{
		{"empty", "", -1, ""},
		{"name missing", "", 0, ""},
		{"temp missing", "alertName", -1, ""},
		{"valid", "alertName", 1, ""},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			var fields []string
			if test.alertName != "" {
				fields = append(fields, fmt.Sprintf(`"name": "%s"`, test.alertName))
			}
			if test.alertActive == 0 {
				fields = append(fields, `"active": false`)
			}
			if test.alertActive == 1 {
				fields = append(fields, `"active": true`)
			}
			tstat := &Thermostat{
				client: &fakeThermostatClient{
					body: fmt.Sprintf(`{"alerts": [{%s}]}`, strings.Join(fields, ",")),
				},
			}

			info, err := tstat.GetQueryAlerts()

			if test.expErr != "" && err == nil {
				t.Fatal("error expected, got: nil want:", test.expErr)
			}
			if test.expErr != "" && err != nil && err.Error() != test.expErr {
				t.Fatal("error expected, got:", err.Error(), "want:", test.expErr)
			}
			if len(info) != 1 {
				t.Fatal("unexpected count of alerts returned, got:", len(info), "want: 1))")
			}
			if test.alertName != info[0].Name {
				t.Error("sensor name incorrect, got:", info[0].Name, "want:", test.alertName)
			}
			if test.alertActive == 0 && info[0].Active {
				t.Error("sensor name incorrect, got:", info[0].Active, "want: false")
			}
			if test.alertActive == 1 && !info[0].Active {
				t.Error("sensor name incorrect, got:", info[0].Active, "want: true")
			}
		})
	}
}
