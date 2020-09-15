package influx

import (
	"fmt"
	"sync"

	"github.com/influxdata/influxdb1-client/v2"
	"github.com/mikemrm/go-venstar/thermostat/monitoring"
)

type Config struct {
	Addr string
	User string
	Pass string

	Database        string
	Measurement     string
	RetentionPolicy string
}

type InfluxWriter struct {
	client          client.Client
	Database        string
	Measurement     string
	RetentionPolicy string

	buffer   []*monitoring.Results
	bufferMu sync.RWMutex
}

func (w *InfluxWriter) WriteResults(results *monitoring.Results) error {
	bp, err := client.NewBatchPoints(client.BatchPointsConfig{
		Database:        w.Database,
		Precision:       "s",
		RetentionPolicy: w.RetentionPolicy,
	})
	if err != nil {
		return err
	}
	point, err := client.NewPoint(w.Measurement, resultTags(results), resultFields(results), results.Timestamp)
	if err != nil {
		return err
	}
	bp.AddPoint(point)
	return w.client.Write(bp)
}

func resultTags(results *monitoring.Results) map[string]string {
	tags := make(map[string]string)
	if results.APIInfo != nil {
		if results.APIInfo.Model != "" {
			tags["system_model"] = fmt.Sprintf("%v", results.APIInfo.Model)
		}
		if results.APIInfo.Type != "" {
			tags["system_type"] = results.APIInfo.Type
		}
	}
	if results.QueryInfo != nil {
		if results.QueryInfo.Name != "" {
			tags["name"] = results.QueryInfo.Name
		}
		tags["state"] = fmt.Sprintf("%d", results.QueryInfo.State)
	}
	return tags
}

func resultFields(results *monitoring.Results) map[string]interface{} {
	fields := make(map[string]interface{})
	if results.APIInfo != nil {
		fields["system_api_version"] = results.APIInfo.Version
		if results.APIInfo.Model != "" {
			fields["system_model"] = results.APIInfo.Model
		}
		if results.APIInfo.Firmware != "" {
			fields["system_firmware"] = results.APIInfo.Firmware
		}
		if results.APIInfo.Type != "" {
			fields["system_type"] = results.APIInfo.Type
		}
	}
	if results.QueryInfo != nil {
		fields["name"] = results.QueryInfo.Name
		fields["mode"] = results.QueryInfo.Mode
		fields["state"] = results.QueryInfo.State
		fields["fan"] = results.QueryInfo.Fan
		fields["fanstate"] = results.QueryInfo.Fanstate
		fields["tempunits"] = results.QueryInfo.Tempunits
		fields["schedule"] = results.QueryInfo.Schedule
		fields["schedulepart"] = results.QueryInfo.Schedulepart
		fields["away"] = results.QueryInfo.Away
		fields["holiday"] = results.QueryInfo.Holiday
		fields["override"] = results.QueryInfo.Override
		fields["overridetime"] = results.QueryInfo.OverrideTime
		fields["forceunocc"] = results.QueryInfo.Forceunocc
		fields["spacetemp"] = results.QueryInfo.SpaceTemp
		fields["heattemp"] = results.QueryInfo.HeatTemp
		fields["cooltemp"] = results.QueryInfo.CoolTemp
		fields["cooltempmin"] = results.QueryInfo.CoolTempMin
		fields["cooltempmax"] = results.QueryInfo.CoolTempMax
		fields["heattempmin"] = results.QueryInfo.HeatTempMin
		fields["heattempmax"] = results.QueryInfo.HeatTempMax
		fields["activestage"] = results.QueryInfo.ActiveStage
		fields["hum_active"] = results.QueryInfo.HumidityEnabled
		fields["hum"] = results.QueryInfo.Humidity
		fields["hum_setpoint"] = results.QueryInfo.HumidifySetPoint
		fields["dehum_setpoint"] = results.QueryInfo.DehumidifySetPoint
		fields["setpointdelta"] = results.QueryInfo.SetPointDelta
		fields["availablemodes"] = results.QueryInfo.AvailableModes
	}
	return fields
}

func NewWriter(config Config) (*InfluxWriter, error) {
	client, err := client.NewHTTPClient(client.HTTPConfig{
		Addr:     config.Addr,
		Username: config.User,
		Password: config.Pass,
	})
	if err != nil {
		return nil, err
	}
	return &InfluxWriter{
		client:          client,
		Database:        config.Database,
		Measurement:     config.Measurement,
		RetentionPolicy: config.RetentionPolicy,
	}, nil
}
