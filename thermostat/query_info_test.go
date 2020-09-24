package thermostat

import (
	"strconv"
	"testing"
)

func TestGetQueryInfo(t *testing.T) {
	t.Run("errors get returned", func(t *testing.T) {
		tstat := &Thermostat{
			client: &fakeThermostatClient{
				error: true,
			},
		}

		_, err := tstat.GetQueryInfo()
		if err == nil {
			t.Fatal("error expected but no error returned")
		}
		want := "processing query info request: requesting /query/info: this is an error"
		if err.Error() != want {
			t.Error("error expected, got:", err.Error(), "want:", want)
		}
	})
	tests := []struct {
		name  string
		field string
		value string
		typ   string
	}{
		{"Name is set", "name", `"namevalue"`, "string"},
		{"Mode is set", "mode", "5", "int"},
		{"State is set", "state", "6", "int"},
		{"Fan is set", "fan", "7", "int"},
		{"FanState is set", "fanstate", "8", "int"},
		{"TempUnits is set", "tempunits", "9", "int"},
		{"Schedule is set", "schedule", "10", "int"},
		{"SchedulePart is set", "schedulepart", "11", "int"},
		{"Away is set", "away", "12", "int"},
		{"Holiday is set", "holiday", "13", "int"},
		{"Override is set", "override", "14", "int"},
		{"OverrideRemaining is set", "overridetime", "15", "int"},
		{"ForceUnoccupied is set", "forceunocc", "16", "int"},
		{"SpaceTemp is set", "spacetemp", "17.5", "float"},
		{"HeatTemp is set", "heattemp", "18.5", "float"},
		{"CoolTemp is set", "cooltemp", "19.5", "float"},
		{"CoolTempMin is set", "cooltempmin", "20.5", "float"},
		{"CoolTempMax is set", "cooltempmax", "21.5", "float"},
		{"HeatTempMin is set", "heattempmin", "22.5", "float"},
		{"HeatTempMax is set", "heattempmax", "23.5", "float"},
		{"ActiveStage is set", "activestage", "24", "int"},
		{"HumidityEnabled is set", "hum_active", "25", "int"},
		{"Humidity is set", "hum", "26", "int"},
		{"HumidifySetPoint is set", "hum_setpoint", "27", "int"},
		{"DehumidifySetPoint is set", "dehum_setpoint", "28", "int"},
		{"SetPointDelta is set", "setpointdelta", "29.5", "float"},
		{"AvailableModes is set", "availablemodes", "30", "int"},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			tstat := &Thermostat{
				client: &fakeThermostatClient{
					body: `{"` + test.field + `": ` + test.value + `}`,
				},
			}

			info, err := tstat.GetQueryInfo()
			if err != nil {
				t.Fatal("error unexpected, got:", err)
			}
			swant := test.value
			var iwant int
			var fwant float64
			if test.typ == "int" {
				iwant, err = strconv.Atoi(swant)
				if err != nil {
					t.Fatal("Failed to convert test value to int:", test.value)
				}
			} else if test.typ == "float" {
				fwant, err = strconv.ParseFloat(swant, 64)
				if err != nil {
					t.Fatal("Failed to convert test value to float:", test.value)
				}
			} else if test.typ == "string" {
				swant = swant[1 : len(swant)-1]
			}
			if test.field == "name" && info.Name != swant {
				t.Error("Name incorrect, got:", info.Name, "want:", iwant)
			}
			if test.field == "mode" && int(info.Mode) != iwant {
				t.Error("Mode incorrect, got:", info.Mode, "want:", iwant)
			}
			if test.field == "state" && int(info.State) != iwant {
				t.Error("State incorrect, got:", info.State, "want:", iwant)
			}
			if test.field == "fan" && int(info.Fan) != iwant {
				t.Error("Fan incorrect, got:", info.Fan, "want:", iwant)
			}
			if test.field == "fanstate" && int(info.FanState) != iwant {
				t.Error("FanState incorrect, got:", info.FanState, "want:", iwant)
			}
			if test.field == "tempunits" && int(info.TempUnits) != iwant {
				t.Error("TempUnits incorrect, got:", info.TempUnits, "want:", iwant)
			}
			if test.field == "schedule" && int(info.Schedule) != iwant {
				t.Error("Schedule incorrect, got:", info.Schedule, "want:", iwant)
			}
			if test.field == "schedulepart" && int(info.SchedulePart) != iwant {
				t.Error("SchedulePart incorrect, got:", info.SchedulePart, "want:", iwant)
			}
			if test.field == "away" && int(info.Away) != iwant {
				t.Error("Away incorrect, got:", info.Away, "want:", iwant)
			}
			if test.field == "holiday" && int(info.Holiday) != iwant {
				t.Error("Holiday incorrect, got:", info.Holiday, "want:", iwant)
			}
			if test.field == "override" && int(info.Override) != iwant {
				t.Error("Override incorrect, got:", info.Override, "want:", iwant)
			}
			if test.field == "overridetime" && int(info.OverrideRemaining) != iwant {
				t.Error("OverrideRemaining incorrect, got:", info.OverrideRemaining, "want:", iwant)
			}
			if test.field == "forceunocc" && int(info.ForceUnoccupied) != iwant {
				t.Error("ForceUnoccupied incorrect, got:", info.ForceUnoccupied, "want:", iwant)
			}
			if test.field == "spacetemp" && info.SpaceTemp != fwant {
				t.Error("SpaceTemp incorrect, got:", info.SpaceTemp, "want:", fwant)
			}
			if test.field == "heattemp" && info.HeatTemp != fwant {
				t.Error("HeatTemp incorrect, got:", info.HeatTemp, "want:", fwant)
			}
			if test.field == "cooltemp" && info.CoolTemp != fwant {
				t.Error("CoolTemp incorrect, got:", info.CoolTemp, "want:", fwant)
			}
			if test.field == "cooltempmin" && info.CoolTempMin != fwant {
				t.Error("CoolTempMin incorrect, got:", info.CoolTempMin, "want:", fwant)
			}
			if test.field == "cooltempmax" && info.CoolTempMax != fwant {
				t.Error("CoolTempMax incorrect, got:", info.CoolTempMax, "want:", fwant)
			}
			if test.field == "heattempmin" && info.HeatTempMin != fwant {
				t.Error("HeatTempMin incorrect, got:", info.HeatTempMin, "want:", fwant)
			}
			if test.field == "heattempmax" && info.HeatTempMax != fwant {
				t.Error("HeatTempMax incorrect, got:", info.HeatTempMax, "want:", fwant)
			}
			if test.field == "activestage" && info.ActiveStage != iwant {
				t.Error("ActiveStage incorrect, got:", info.ActiveStage, "want:", iwant)
			}
			if test.field == "hum_active" && int(info.HumidityEnabled) != iwant {
				t.Error("HumidityEnabled incorrect, got:", info.HumidityEnabled, "want:", iwant)
			}
			if test.field == "hum" && info.Humidity != iwant {
				t.Error("Humidity incorrect, got:", info.Humidity, "want:", iwant)
			}
			if test.field == "hum_setpoint" && info.HumidifySetPoint != iwant {
				t.Error("HumidifySetPoint incorrect, got:", info.HumidifySetPoint, "want:", iwant)
			}
			if test.field == "dehum_setpoint" && info.DehumidifySetPoint != iwant {
				t.Error("DehumidifySetPoint incorrect, got:", info.DehumidifySetPoint, "want:", iwant)
			}
			if test.field == "setpointdelta" && info.SetPointDelta != fwant {
				t.Error("SetPointDelta incorrect, got:", info.SetPointDelta, "want:", fwant)
			}
			if test.field == "availablemodes" && int(info.AvailableModes) != iwant {
				t.Error("AvailableModes incorrect, got:", info.AvailableModes, "want:", iwant)
			}
		})
	}
}

func TestModeString(t *testing.T) {
	tests := []struct {
		name  string
		value int
		want  string
	}{
		{"off", 0, "off"},
		{"heat", 1, "heat"},
		{"cool", 2, "cool"},
		{"auto", 3, "auto"},
		{"default", 4, ""},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			got := Mode(test.value).String()
			if got != test.want {
				t.Error("unexpected string for value", test.value, "got:", got, "want:", test.want)
			}
		})
	}
}

func TestStateString(t *testing.T) {
	tests := []struct {
		name  string
		value int
		want  string
	}{
		{"idle", 0, "idle"},
		{"heating", 1, "heating"},
		{"cooling", 2, "cooling"},
		{"lockout", 3, "lockout"},
		{"error", 4, "error"},
		{"default", 5, ""},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			got := State(test.value).String()
			if got != test.want {
				t.Error("unexpected string for value", test.value, "got:", got, "want:", test.want)
			}
		})
	}
}

func TestFanString(t *testing.T) {
	tests := []struct {
		name  string
		value int
		want  string
	}{
		{"auto", 0, "auto"},
		{"on", 1, "on"},
		{"default", 2, ""},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			got := Fan(test.value).String()
			if got != test.want {
				t.Error("unexpected string for value", test.value, "got:", got, "want:", test.want)
			}
		})
	}
}

func TestFanStateString(t *testing.T) {
	tests := []struct {
		name  string
		value int
		want  string
	}{
		{"off", 0, "off"},
		{"on", 1, "on"},
		{"default", 2, ""},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			got := FanState(test.value).String()
			if got != test.want {
				t.Error("unexpected string for value", test.value, "got:", got, "want:", test.want)
			}
		})
	}
}

func TestTempUnitsString(t *testing.T) {
	tests := []struct {
		name  string
		value int
		want  string
	}{
		{"fahrenheit", 0, "fahrenheit"},
		{"celsius", 1, "celsius"},
		{"default", 2, ""},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			got := TempUnits(test.value).String()
			if got != test.want {
				t.Error("unexpected string for value", test.value, "got:", got, "want:", test.want)
			}
		})
	}
}

func TestScheduleString(t *testing.T) {
	tests := []struct {
		name  string
		value int
		want  string
	}{
		{"inactive", 0, "inactive"},
		{"active", 1, "active"},
		{"default", 2, ""},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			got := Schedule(test.value).String()
			if got != test.want {
				t.Error("unexpected string for value", test.value, "got:", got, "want:", test.want)
			}
		})
	}
}

func TestSchedulePart(t *testing.T) {
	tests := []struct {
		name  string
		value int
		want  string
	}{
		{"morning", 0, "morning"},
		{"day", 1, "day"},
		{"evening", 2, "evening"},
		{"night", 3, "night"},
		{"inactive", 255, "inactive"},
		{"default", 4, ""},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			got := SchedulePart(test.value).String()
			if got != test.want {
				t.Error("unexpected string for value", test.value, "got:", got, "want:", test.want)
			}
		})
	}
}

func TestAway(t *testing.T) {
	tests := []struct {
		name  string
		value int
		want  string
	}{
		{"home", 0, "home"},
		{"away", 1, "away"},
		{"default", 2, ""},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			got := Away(test.value).String()
			if got != test.want {
				t.Error("unexpected string for value", test.value, "got:", got, "want:", test.want)
			}
		})
	}
}

func TestHoliday(t *testing.T) {
	tests := []struct {
		name  string
		value int
		want  string
	}{
		{"not observing", 0, "not observing"},
		{"observing", 1, "observing"},
		{"default", 2, ""},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			got := Holiday(test.value).String()
			if got != test.want {
				t.Error("unexpected string for value", test.value, "got:", got, "want:", test.want)
			}
		})
	}
}

func TestOverride(t *testing.T) {
	tests := []struct {
		name  string
		value int
		want  string
	}{
		{"off", 0, "off"},
		{"on", 1, "on"},
		{"default", 2, ""},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			got := Override(test.value).String()
			if got != test.want {
				t.Error("unexpected string for value", test.value, "got:", got, "want:", test.want)
			}
		})
	}
}

func TestOverrideRemaining(t *testing.T) {
	tests := []struct {
		name  string
		value int
		want  string
	}{
		{"none", 0, "0s"},
		{"1m", 1, "1m0s"},
		{"1h30m", 90, "1h30m0s"},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			got := OverrideRemaining(test.value).String()
			if got != test.want {
				t.Error("unexpected string for value", test.value, "got:", got, "want:", test.want)
			}
		})
	}
}

func TestForceUnoccupied(t *testing.T) {
	tests := []struct {
		name  string
		value int
		want  string
	}{
		{"off", 0, "off"},
		{"on", 1, "on"},
		{"default", 2, ""},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			got := ForceUnoccupied(test.value).String()
			if got != test.want {
				t.Error("unexpected string for value", test.value, "got:", got, "want:", test.want)
			}
		})
	}
}

func TestHumidityEnabled(t *testing.T) {
	tests := []struct {
		name  string
		value int
		want  string
	}{
		{"disabled", 0, "disabled"},
		{"enabled", 1, "enabled"},
		{"default", 2, ""},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			got := HumidityEnabled(test.value).String()
			if got != test.want {
				t.Error("unexpected string for value", test.value, "got:", got, "want:", test.want)
			}
		})
	}
}

func TestAvailableModes(t *testing.T) {
	tests := []struct {
		name  string
		value int
		want  string
	}{
		{"all", 0, "all"},
		{"heat/cool", 1, "heat/cool"},
		{"heat", 2, "heat"},
		{"cool", 3, "cool"},
		{"default", 4, ""},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			got := AvailableModes(test.value).String()
			if got != test.want {
				t.Error("unexpected string for value", test.value, "got:", got, "want:", test.want)
			}
		})
	}
}
