package monitoring

import (
	"time"

	"github.com/mikemrm/go-venstar/thermostat"
)

type Results struct {
	Timestamp time.Time
	APIInfo   *thermostat.APIInfo
	QueryInfo *thermostat.QueryInfo
	Sensors   []*thermostat.Sensor
	Runtime   []*thermostat.Runtime
	Alerts    []*thermostat.Alert
}
