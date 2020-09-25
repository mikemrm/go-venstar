package thermostat

import (
	"encoding/json"
	"strings"
	"time"
)

// QueryResponse encompasses the results for sensor, runtime and alert results.
type QueryResponse struct {
	Sensors  []*Sensor  `json:"sensors,omitempty"`
	Runtimes []*Runtime `json:"runtimes,omitempty"`
	Alerts   []*Alert   `json:"alerts,omitempty"`
}

// Sensor represents the thermostat sensor readings
type Sensor struct {
	Name string  `json:"name"`
	Temp float64 `json:"temp"`
}

// Alert represents the thermostat alert values
type Alert struct {
	Name   string `json:"name"`
	Active bool   `json:"active"`
}

// Runtime represents the thermostat runtime results
type Runtime struct {
	Timestamp   time.Time
	Heaters     map[string]time.Duration
	Coolers     map[string]time.Duration
	Aux         map[string]time.Duration
	FreeCooling time.Duration
	Override    time.Duration
}

func (r *Runtime) UnmarshalJSON(data []byte) error {
	jdata := make(map[string]int)
	if err := json.Unmarshal(data, &jdata); err != nil {
		return err
	}
	r.Heaters = make(map[string]time.Duration)
	r.Coolers = make(map[string]time.Duration)
	r.Aux = make(map[string]time.Duration)
	for k, v := range jdata {
		if k == "ts" {
			r.Timestamp = time.Unix(int64(v), 0)
		} else if k == "fc" {
			r.FreeCooling = time.Duration(v) * time.Minute
		} else if k == "ov" {
			r.Override = time.Duration(v) * time.Minute
		} else if strings.HasPrefix(k, "heat") {
			r.Heaters[strings.TrimPrefix(k, "heat")] = time.Duration(v) * time.Minute
		} else if strings.HasPrefix(k, "cool") {
			r.Coolers[strings.TrimPrefix(k, "cool")] = time.Duration(v) * time.Minute
		} else if strings.HasPrefix(k, "aux") {
			r.Aux[strings.TrimPrefix(k, "aux")] = time.Duration(v) * time.Minute
		}
	}
	return nil
}
