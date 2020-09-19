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
		fields["fan_state"] = results.QueryInfo.FanState
		fields["temp_units"] = results.QueryInfo.TempUnits
		fields["schedule"] = results.QueryInfo.Schedule
		fields["schedule_part"] = results.QueryInfo.SchedulePart
		fields["away"] = results.QueryInfo.Away
		fields["holiday"] = results.QueryInfo.Holiday
		fields["override"] = results.QueryInfo.Override
		fields["override_remaining"] = results.QueryInfo.OverrideRemaining
		fields["force_unoccupied"] = results.QueryInfo.ForceUnoccupied
		fields["space_temp"] = results.QueryInfo.SpaceTemp
		fields["heat_temp"] = results.QueryInfo.HeatTemp
		fields["cool_temp"] = results.QueryInfo.CoolTemp
		fields["cool_temp_min"] = results.QueryInfo.CoolTempMin
		fields["cool_temp_max"] = results.QueryInfo.CoolTempMax
		fields["heat_temp_min"] = results.QueryInfo.HeatTempMin
		fields["heat_temp_max"] = results.QueryInfo.HeatTempMax
		fields["active_stage"] = results.QueryInfo.ActiveStage
		fields["humidity_enabled"] = results.QueryInfo.HumidityEnabled
		fields["humidity"] = results.QueryInfo.Humidity
		fields["humidity_setpoint"] = results.QueryInfo.HumidifySetPoint
		fields["dehumidify_setpoint"] = results.QueryInfo.DehumidifySetPoint
		fields["setpoint_delta"] = results.QueryInfo.SetPointDelta
		fields["available_modes"] = results.QueryInfo.AvailableModes
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
