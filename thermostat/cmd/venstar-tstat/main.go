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
	qinfo, err := t.QueryInfo()
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
}
