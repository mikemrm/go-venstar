package thermostat

import (
	"errors"
	"io"
	"net/http"
	"testing"
)

func TestClientUpdateControls(t *testing.T) {
	t.Run("errors get returned", func(t *testing.T) {
		tstat := &Thermostat{
			client: &fakeThermostatClient{
				error: true,
			},
		}

		err := tstat.UpdateControls(NewControlRequest())
		if err == nil {
			t.Fatal("error expected but no error returned")
		}
		want := "processing update control request: requesting /control: this is an error"
		if err.Error() != want {
			t.Error("error expected, got:", err.Error(), "want:", want)
		}
	})
	t.Run("response errors get returned", func(t *testing.T) {
		wantErr := `Control Request update error: bad reason`
		tstat := &Thermostat{
			client: &fakeThermostatClient{
				body: `{"error": true, "reason": "bad reason"}`,
			},
		}

		err := tstat.UpdateControls(NewControlRequest())
		if err == nil {
			t.Fatal("error expected, got: nil want:", wantErr)
		}
		if err.Error() != wantErr {
			t.Fatal("error invalid, got:", err.Error(), "want:", wantErr)
		}
	})
	t.Run("unknown errors captured", func(t *testing.T) {
		wantErr := `Control Request unknown error`
		tstat := &Thermostat{
			client: &fakeThermostatClient{
				body: `{"success": false}`,
			},
		}

		err := tstat.UpdateControls(NewControlRequest())
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

		err := tstat.UpdateControls(NewControlRequest())
		if err != nil {
			t.Fatal("error unexpected, got:", err.Error(), "want: nil")
		}
	})
}

func TestControlRequestSetMode(t *testing.T) {
	cr := NewControlRequest()
	want := 2
	cr.SetMode(want)
	if cr.Mode == nil {
		t.Fatal("Mode invalid, got: nil want:", want)
	}
	if *cr.Mode != want {
		t.Error("Mode invalid, got:", *cr.Mode, "want:", want)
	}
}

func TestControlRequestSetFan(t *testing.T) {
	cr := NewControlRequest()
	want := 1
	cr.SetFan(want)
	if cr.Fan == nil {
		t.Fatal("Fan invalid, got: nil want:", want)
	}
	if *cr.Fan != want {
		t.Error("Fan invalid, got:", *cr.Fan, "want:", want)
	}
}

func TestControlRequestFanAuto(t *testing.T) {
	cr := NewControlRequest()
	want := 0
	cr.FanAuto()
	if cr.Fan == nil {
		t.Fatal("Fan invalid, got: nil want:", want)
	}
	if *cr.Fan != want {
		t.Error("Fan invalid, got:", *cr.Fan, "want:", want)
	}
}

func TestControlRequestFanOn(t *testing.T) {
	cr := NewControlRequest()
	want := 1
	cr.FanOn()
	if cr.Fan == nil {
		t.Fatal("Fan invalid, got: nil want:", want)
	}
	if *cr.Fan != want {
		t.Error("Fan invalid, got:", *cr.Fan, "want:", want)
	}
}

func TestControlRequestSetHeatTemp(t *testing.T) {
	cr := NewControlRequest()
	want := 0
	cr.SetHeatTemp(want)
	if cr.HeatTemp == nil {
		t.Fatal("HeatTemp invalid, got: nil want:", want)
	}
	if *cr.HeatTemp != want {
		t.Error("HeatTemp invalid, got:", *cr.HeatTemp, "want:", want)
	}
}

func TestControlRequestSetCoolTemp(t *testing.T) {
	cr := NewControlRequest()
	want := 0
	cr.SetCoolTemp(want)
	if cr.CoolTemp == nil {
		t.Fatal("CoolTemp invalid, got: nil want:", want)
	}
	if *cr.CoolTemp != want {
		t.Error("CoolTemp invalid, got:", *cr.CoolTemp, "want:", want)
	}
}

func TestControlRequestOff(t *testing.T) {
	cr := NewControlRequest()
	wantMode := 0
	wantCool := 70
	wantHeat := 65
	cr.Off(wantCool, wantHeat)
	if cr.Mode == nil {
		t.Fatal("Mode invalid, got: nil want:", wantMode)
	}
	if cr.CoolTemp == nil {
		t.Fatal("Cool invalid, got: nil want:", wantCool)
	}
	if cr.HeatTemp == nil {
		t.Fatal("Heat invalid, got: nil want:", wantHeat)
	}
	if *cr.Mode != wantMode {
		t.Error("Mode invalid, got:", *cr.Mode, "want:", wantMode)
	}
	if *cr.CoolTemp != wantCool {
		t.Error("CoolTemp invalid, got:", *cr.CoolTemp, "want:", wantCool)
	}
	if *cr.HeatTemp != wantHeat {
		t.Error("HeatTemp invalid, got:", *cr.HeatTemp, "want:", wantHeat)
	}
}

func TestControlRequestHeat(t *testing.T) {
	cr := NewControlRequest()
	wantMode := 1
	wantCool := 70
	wantHeat := 65
	cr.Heat(wantHeat, wantCool)
	if cr.Mode == nil {
		t.Fatal("Mode invalid, got: nil want:", wantMode)
	}
	if cr.CoolTemp == nil {
		t.Fatal("Cool invalid, got: nil want:", wantCool)
	}
	if cr.HeatTemp == nil {
		t.Fatal("Heat invalid, got: nil want:", wantHeat)
	}
	if *cr.Mode != wantMode {
		t.Error("Mode invalid, got:", *cr.Mode, "want:", wantMode)
	}
	if *cr.CoolTemp != wantCool {
		t.Error("CoolTemp invalid, got:", *cr.CoolTemp, "want:", wantCool)
	}
	if *cr.HeatTemp != wantHeat {
		t.Error("HeatTemp invalid, got:", *cr.HeatTemp, "want:", wantHeat)
	}
}

func TestControlRequestCool(t *testing.T) {
	cr := NewControlRequest()
	wantMode := 2
	wantCool := 70
	wantHeat := 65
	cr.Cool(wantCool, wantHeat)
	if cr.Mode == nil {
		t.Fatal("Mode invalid, got: nil want:", wantMode)
	}
	if cr.CoolTemp == nil {
		t.Fatal("Cool invalid, got: nil want:", wantCool)
	}
	if cr.HeatTemp == nil {
		t.Fatal("Heat invalid, got: nil want:", wantHeat)
	}
	if *cr.Mode != wantMode {
		t.Error("Mode invalid, got:", *cr.Mode, "want:", wantMode)
	}
	if *cr.CoolTemp != wantCool {
		t.Error("CoolTemp invalid, got:", *cr.CoolTemp, "want:", wantCool)
	}
	if *cr.HeatTemp != wantHeat {
		t.Error("HeatTemp invalid, got:", *cr.HeatTemp, "want:", wantHeat)
	}
}

func TestControlRequestAuto(t *testing.T) {
	cr := NewControlRequest()
	wantMode := 3
	wantCool := 70
	wantHeat := 65
	cr.Auto(wantCool, wantHeat)
	if cr.Mode == nil {
		t.Fatal("Mode invalid, got: nil want:", wantMode)
	}
	if cr.CoolTemp == nil {
		t.Fatal("Cool invalid, got: nil want:", wantCool)
	}
	if cr.HeatTemp == nil {
		t.Fatal("Heat invalid, got: nil want:", wantHeat)
	}
	if *cr.Mode != wantMode {
		t.Error("Mode invalid, got:", *cr.Mode, "want:", wantMode)
	}
	if *cr.CoolTemp != wantCool {
		t.Error("CoolTemp invalid, got:", *cr.CoolTemp, "want:", wantCool)
	}
	if *cr.HeatTemp != wantHeat {
		t.Error("HeatTemp invalid, got:", *cr.HeatTemp, "want:", wantHeat)
	}
}

func TestControlRequestValidate(t *testing.T) {
	tests := []struct {
		name   string
		mode   int
		fan    int
		heat   int
		cool   int
		errExp string
	}{
		{"no error when empty", -1, -1, -1, -1, ""},
		{"Mode off, missing temps", 0, -1, -1, -1, "HeatTemp must be defined when Mode is defined"},
		{"Mode off, missing heat temp", 0, -1, -1, 70, "HeatTemp must be defined when Mode is defined"},
		{"Mode off, missing cool temp", 0, -1, 65, -1, "CoolTemp must be defined when Mode is defined"},
		{"Mode off, cool higher than heat", 0, -1, 70, 65, ""},
		{"Mode off", 0, -1, 65, 70, ""},
		{"Mode off, Fan auto", 0, 0, 65, 70, ""},
		{"Mode off, Fan on", 0, 1, 65, 70, ""},
		{"Mode heat, missing temps", 1, -1, -1, -1, "HeatTemp must be defined when Mode is defined"},
		{"Mode heat, missing heat temp", 1, -1, -1, 70, "HeatTemp must be defined when Mode is defined"},
		{"Mode heat, missing cool temp", 1, -1, 65, -1, "CoolTemp must be defined when Mode is defined"},
		{"Mode heat, cool higher than heat", 1, -1, 70, 65, ""},
		{"Mode heat", 1, -1, 65, 70, ""},
		{"Mode heat, Fan auto", 1, 0, 65, 70, ""},
		{"Mode heat, Fan on", 1, 1, 65, 70, ""},
		{"Mode cool, missing temps", 2, -1, -1, -1, "HeatTemp must be defined when Mode is defined"},
		{"Mode cool, missing heat temp", 2, -1, -1, 70, "HeatTemp must be defined when Mode is defined"},
		{"Mode cool, missing cool temp", 2, -1, 65, -1, "CoolTemp must be defined when Mode is defined"},
		{"Mode cool, cool higher than heat", 2, -1, 70, 65, ""},
		{"Mode cool", 2, -1, 65, 70, ""},
		{"Mode cool, Fan auto", 2, 0, 65, 70, ""},
		{"Mode cool, Fan on", 2, 1, 65, 70, ""},
		{"Mode auto, missing temps", 3, -1, -1, -1, "HeatTemp must be defined when Mode is defined"},
		{"Mode auto, missing heat temp", 3, -1, -1, 70, "HeatTemp must be defined when Mode is defined"},
		{"Mode auto, missing cool temp", 3, -1, 65, -1, "CoolTemp must be defined when Mode is defined"},
		{"Mode auto, cool higher than heat", 3, -1, 70, 65, "CoolTemp must be greater than HeatTemp when Mode is Auto"},
		{"Mode auto", 3, -1, 65, 70, ""},
		{"Mode auto, Fan auto", 3, 0, 65, 70, ""},
		{"Mode auto, Fan on", 3, 1, 65, 70, ""},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			cr := NewControlRequest()
			if test.mode != -1 {
				cr.SetMode(test.mode)
			}
			if test.fan != -1 {
				cr.SetFan(test.fan)
			}
			if test.heat != -1 {
				cr.SetHeatTemp(test.heat)
			}
			if test.cool != -1 {
				cr.SetCoolTemp(test.cool)
			}
			err := cr.Validate()
			if test.errExp != "" && err == nil {
				t.Fatal("Error invalid, got: nil want:", test.errExp)
			}
			if test.errExp == "" && err != nil {
				t.Fatal("Error invalid, got:", err.Error(), "want: nil")
			}
			if test.errExp != "" && err != nil && err.Error() != test.errExp {
				t.Fatal("Error invalid, got:", err.Error(), "want:", test.errExp)
			}
		})
	}
}

func TestControlRequestBuildRequest(t *testing.T) {
	tests := []struct {
		name    string
		mode    int
		fan     int
		heat    int
		cool    int
		wantLen int
		errExp  string
	}{
		{"no error when empty", -1, -1, -1, -1, 0, ""},
		{"validation error returned", -1, -1, -1, -1, -1, "validation error"},
		{"mode", 0, -1, -1, -1, 6, ""},
		{"mode,fan", 0, 1, -1, -1, 12, ""},
		{"mode,fan,heat", 0, 1, 2, -1, 23, ""},
		{"mode,fan,heat,cool", 0, 1, 2, 3, 34, ""},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			cr := NewControlRequest()
			if test.mode != -1 {
				cr.SetMode(test.mode)
			}
			if test.fan != -1 {
				cr.SetFan(test.fan)
			}
			if test.heat != -1 {
				cr.SetHeatTemp(test.heat)
			}
			if test.cool != -1 {
				cr.SetCoolTemp(test.cool)
			}
			validateCalled := false
			cr.validator = func(_ *ControlRequest) error {
				validateCalled = true
				if test.wantLen == -1 {
					return errors.New("validation error")
				}
				return nil
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

			body, err := io.ReadAll(req.Body)
			if err != nil {
				t.Fatal("Unexpected body read error:", err)
			}
			if !validateCalled {
				t.Error("Validate not called")
			}
			if len(body) != test.wantLen {
				t.Error("Body length invalid, got:", req.Body, "want:", test.wantLen)
			}
			if req.ContentLength != int64(test.wantLen) {
				t.Error("ContentLength invalid, got:", req.ContentLength, "want:", test.wantLen)
			}
		})
	}
}
