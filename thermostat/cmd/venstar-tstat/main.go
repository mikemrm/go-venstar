package main

import (
	"fmt"
	"os"

	"github.com/mikemrm/go-venstar/thermostat"
)

func main() {
	t := thermostat.New(os.Args[1])

	fmt.Println("API Info:")
	info, err := t.GetAPIInfo()
	if err != nil {
		panic(err)
	}
	fmt.Println("  Type     :", info.Type)
	fmt.Println("  Model    :", info.Model)
	fmt.Println("  Version  :", info.Version)
	fmt.Println("  Firmware :", info.Firmware)

	fmt.Println("Query Info:")
	qinfo, err := t.GetQueryInfo()
	if err != nil {
		panic(err)
	}
	fmt.Println("  Name               :", qinfo.Name)
	fmt.Println("  Mode               :", qinfo.Mode)
	fmt.Println("  State              :", qinfo.State)
	fmt.Println("  Fan                :", qinfo.Fan)
	fmt.Println("  Fanstate           :", qinfo.Fanstate)
	fmt.Println("  Tempunits          :", qinfo.Tempunits)
	fmt.Println("  Schedule           :", qinfo.Schedule)
	fmt.Println("  Schedulepart       :", qinfo.Schedulepart)
	fmt.Println("  Away               :", qinfo.Away)
	fmt.Println("  Holiday            :", qinfo.Holiday)
	fmt.Println("  Override           :", qinfo.Override)
	fmt.Println("  OverrideTime       :", qinfo.OverrideTime)
	fmt.Println("  Forceunocc         :", qinfo.Forceunocc)
	fmt.Println("  SpaceTemp          :", qinfo.SpaceTemp)
	fmt.Println("  HeatTemp           :", qinfo.HeatTemp)
	fmt.Println("  CoolTemp           :", qinfo.CoolTemp)
	fmt.Println("  CoolTempMin        :", qinfo.CoolTempMin)
	fmt.Println("  CoolTempMax        :", qinfo.CoolTempMax)
	fmt.Println("  HeatTempMin        :", qinfo.HeatTempMin)
	fmt.Println("  HeatTempMax        :", qinfo.HeatTempMax)
	fmt.Println("  ActiveStage        :", qinfo.ActiveStage)
	fmt.Println("  HumidityEnabled    :", qinfo.HumidityEnabled)
	fmt.Println("  Humidity           :", qinfo.Humidity)
	fmt.Println("  HumidifySetPoint   :", qinfo.HumidifySetPoint)
	fmt.Println("  DehumidifySetPoint :", qinfo.DehumidifySetPoint)
	fmt.Println("  SetPointDelta      :", qinfo.SetPointDelta)
	fmt.Println("  AvailableModes     :", qinfo.AvailableModes)

	fmt.Println("Query Sensors:")
	sensors, err := t.GetQuerySensors()
	if err != nil {
		panic(err)
	}
	for _, sensor := range sensors {
		fmt.Printf("  %s = %.1f\n", sensor.Name, sensor.Temp)
	}

	fmt.Println("Query Runtimes:")
	runtimes, err := t.GetQueryRuntimes()
	if err != nil {
		panic(err)
	}
	for i, runtime := range runtimes {
		if i == 0 {
			fmt.Println("  +-----------------+--------+--------+--------+--------+-------+-------+----+----+")
			fmt.Println("  | Timestamp       | Heat 1 | Heat 2 | Cool 1 | Cool 2 | Aux 1 | Aux 2 | FC | OV |")
			fmt.Println("  +-----------------+--------+--------+--------+--------+-------+-------+----+----+")
		}
		fmt.Printf("  | %15d | %6d | %6d | %6d | %6d | %5d | %5d | %2d | %2d |\n", runtime.Time, runtime.Heat1, runtime.Heat2, runtime.Cool1, runtime.Cool2, runtime.Aux1, runtime.Aux2, runtime.FC, runtime.OV)
		if i == len(runtimes)-1 {
			fmt.Println("  +-----------------+--------+--------+--------+--------+-------+-------+----+----+")
		}
	}

	fmt.Println("Query Alerts:")
	alerts, err := t.GetQueryAlerts()
	if err != nil {
		panic(err)
	}
	for _, alert := range alerts {
		fmt.Printf("  %10s: Active = %v\n", alert.Name, alert.Active)
	}
}
