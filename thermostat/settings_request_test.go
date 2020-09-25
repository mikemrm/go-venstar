package thermostat

import (
	"io/ioutil"
	"net/http"
	"testing"
)

func TestClientUpdateSettings(t *testing.T) {
	t.Run("errors get returned", func(t *testing.T) {
		tstat := &Thermostat{
			client: &fakeThermostatClient{
				error: true,
			},
		}

		err := tstat.UpdateSettings(NewSettingsRequest())
		if err == nil {
			t.Fatal("error expected but no error returned")
		}
		want := "processing update settings request: requesting /settings: this is an error"
		if err.Error() != want {
			t.Error("error expected, got:", err.Error(), "want:", want)
		}
	})
	t.Run("response errors get returned", func(t *testing.T) {
		wantErr := `Settings Request update error: bad reason`
		tstat := &Thermostat{
			client: &fakeThermostatClient{
				body: `{"error": true, "reason": "bad reason"}`,
			},
		}

		err := tstat.UpdateSettings(NewSettingsRequest())
		if err == nil {
			t.Fatal("error expected, got: nil want:", wantErr)
		}
		if err.Error() != wantErr {
			t.Fatal("error invalid, got:", err.Error(), "want:", wantErr)
		}
	})
	t.Run("unknown errors captured", func(t *testing.T) {
		wantErr := `Settings Request unknown error`
		tstat := &Thermostat{
			client: &fakeThermostatClient{
				body: `{"success": false}`,
			},
		}

		err := tstat.UpdateSettings(NewSettingsRequest())
		if err == nil {
			t.Fatal("error expected, got: nil want:", wantErr)
		}
		if err.Error() != wantErr {
			t.Fatal("error invalid, got:", err.Error(), "want:", wantErr)
		}
	})
	t.Run("successful response", func(t *testing.T) {
		tstat := &Thermostat{
			client: &fakeThermostatClient{
				body: `{"success": true}`,
			},
		}

		err := tstat.UpdateSettings(NewSettingsRequest())
		if err != nil {
			t.Fatal("error unexpected, got:", err.Error(), "want: nil")
		}
	})
}

func TestSettingsRequestSetTempUnits(t *testing.T) {
	cr := NewSettingsRequest()
	want := 1
	cr.SetTempUnits(want)
	if cr.TempUnits == nil {
		t.Fatal("TempUnits invalid, got: nil want:", want)
	}
	if *cr.TempUnits != want {
		t.Error("TempUnits invalid, got:", *cr.TempUnits, "want:", want)
	}
}

func TestSettingsRequestFahrenheit(t *testing.T) {
	cr := NewSettingsRequest()
	want := 0
	cr.Fahrenheit()
	if cr.TempUnits == nil {
		t.Fatal("TempUnits invalid, got: nil want:", want)
	}
	if *cr.TempUnits != want {
		t.Error("TempUnits invalid, got:", *cr.TempUnits, "want:", want)
	}
}

func TestSettingsRequestCelsius(t *testing.T) {
	cr := NewSettingsRequest()
	want := 1
	cr.Celsius()
	if cr.TempUnits == nil {
		t.Fatal("TempUnits invalid, got: nil want:", want)
	}
	if *cr.TempUnits != want {
		t.Error("TempUnits invalid, got:", *cr.TempUnits, "want:", want)
	}
}

func TestSettingsRequestSetAway(t *testing.T) {
	t.Run("away", func(t *testing.T) {
		cr := NewSettingsRequest()
		want := 1
		cr.SetAway(true)
		if cr.IsAway == nil {
			t.Fatal("Away invalid, got: nil want:", want)
		}
		if *cr.IsAway != want {
			t.Error("Away invalid, got:", *cr.IsAway, "want:", want)
		}
	})
	t.Run("home", func(t *testing.T) {
		cr := NewSettingsRequest()
		want := 0
		cr.SetAway(false)
		if cr.IsAway == nil {
			t.Fatal("Away invalid, got: nil want:", want)
		}
		if *cr.IsAway != want {
			t.Error("Away invalid, got:", *cr.IsAway, "want:", want)
		}
	})
}

func TestSettingsRequestAway(t *testing.T) {
	cr := NewSettingsRequest()
	want := 1
	cr.Away()
	if cr.IsAway == nil {
		t.Fatal("Away invalid, got: nil want:", want)
	}
	if *cr.IsAway != want {
		t.Error("Away invalid, got:", *cr.IsAway, "want:", want)
	}
}

func TestSettingsRequestHome(t *testing.T) {
	cr := NewSettingsRequest()
	want := 0
	cr.Home()
	if cr.IsAway == nil {
		t.Fatal("Away invalid, got: nil want:", want)
	}
	if *cr.IsAway != want {
		t.Error("Away invalid, got:", *cr.IsAway, "want:", want)
	}
}

func TestSettingsRequestSetSchedule(t *testing.T) {
	cr := NewSettingsRequest()
	want := 1
	cr.SetSchedule(want)
	if cr.Schedule == nil {
		t.Fatal("Schedule invalid, got: nil want:", want)
	}
	if *cr.Schedule != want {
		t.Error("Schedule invalid, got:", *cr.Schedule, "want:", want)
	}
}

func TestSettingsRequestScheduleOff(t *testing.T) {
	cr := NewSettingsRequest()
	want := 0
	cr.ScheduleOff()
	if cr.Schedule == nil {
		t.Fatal("Schedule invalid, got: nil want:", want)
	}
	if *cr.Schedule != want {
		t.Error("Schedule invalid, got:", *cr.Schedule, "want:", want)
	}
}

func TestSettingsRequestScheduleOn(t *testing.T) {
	cr := NewSettingsRequest()
	want := 1
	cr.ScheduleOn()
	if cr.Schedule == nil {
		t.Fatal("Schedule invalid, got: nil want:", want)
	}
	if *cr.Schedule != want {
		t.Error("Schedule invalid, got:", *cr.Schedule, "want:", want)
	}
}

func TestSettingsRequestSetHumidifySetPoint(t *testing.T) {
	cr := NewSettingsRequest()
	want := 30
	cr.SetHumidifySetPoint(want)
	if cr.HumidifySetPoint == nil {
		t.Fatal("HumidifySetPoint invalid, got: nil want:", want)
	}
	if *cr.HumidifySetPoint != want {
		t.Error("HumidifySetPoint invalid, got:", *cr.HumidifySetPoint, "want:", want)
	}
}

func TestSettingsRequestSetDehumidifySetPoint(t *testing.T) {
	cr := NewSettingsRequest()
	want := 75
	cr.SetDehumidifySetPoint(want)
	if cr.DehumidifySetPoint == nil {
		t.Fatal("DehumidifySetPoint invalid, got: nil want:", want)
	}
	if *cr.DehumidifySetPoint != want {
		t.Error("DehumidifySetPoint invalid, got:", *cr.DehumidifySetPoint, "want:", want)
	}
}

func TestSettingsRequestBuildRequest(t *testing.T) {
	tests := []struct {
		name       string
		units      int
		away       int
		schedule   int
		humidify   int
		dehumidify int
		wantLen    int
		errExp     string
	}{
		{"no error when empty", -1, -1, -1, -1, -1, 0, ""},
		{"units", 0, -1, -1, -1, -1, 11, ""},
		{"units,away", 0, 1, -1, -1, -1, 18, ""},
		{"units,away,schedule", 0, 1, 2, -1, -1, 29, ""},
		{"units,away,schedule,humidify", 0, 1, 2, 3, -1, 44, ""},
		{"units,away,schedule,humidify,dehumidify", 0, 1, 2, 3, 4, 61, ""},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			cr := NewSettingsRequest()
			if test.units != -1 {
				cr.SetTempUnits(test.units)
			}
			if test.away != -1 {
				cr.SetAway(test.away == 1)
			}
			if test.schedule != -1 {
				cr.SetSchedule(test.schedule)
			}
			if test.humidify != -1 {
				cr.SetHumidifySetPoint(test.humidify)
			}
			if test.dehumidify != -1 {
				cr.SetDehumidifySetPoint(test.dehumidify)
			}
			req := &http.Request{
				Header: make(http.Header),
			}
			err := cr.BuildRequest(req)
			if test.errExp != "" && err == nil {
				t.Fatal("Error invalid, got: nil want:", test.errExp)
			}
			if test.errExp == "" && err != nil {
				t.Fatal("Error invalid, got:", err.Error(), "want: nil")
			}
			if test.errExp != "" && err != nil && err.Error() != test.errExp {
				t.Fatal("Error invalid, got:", err.Error(), "want:", test.errExp)
			}
			if test.errExp != "" {
				return
			}

			body, err := ioutil.ReadAll(req.Body)
			if err != nil {
				t.Fatal("Unexpected body read error:", err)
			}
			if len(body) != test.wantLen {
				t.Error("Body length invalid, got:", string(body))
			}
			if req.ContentLength != int64(test.wantLen) {
				t.Error("ContentLength invalid, got:", req.ContentLength, "want:", test.wantLen)
			}
		})
	}
}
