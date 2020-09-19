package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"

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
	printRuntimes(runtimes)

	fmt.Println("Query Alerts:")
	alerts, err := t.GetQueryAlerts()
	if err != nil {
		panic(err)
	}
	for _, alert := range alerts {
		fmt.Printf("  %10s: Active = %v\n", alert.Name, alert.Active)
	}
}

func printRuntimes(runtimes []*thermostat.Runtime) {
	tsFormat := "2006-01-02 15:04:05 MST"
	colWidths := make(map[string]int)
	rowValues := make([]map[string]string, len(runtimes))
	for i, runtime := range runtimes {
		values := make(map[string]string)
		values["Timestamp"] = runtime.Timestamp.Format(tsFormat)
		values["Free Cooling"] = strconv.Itoa(runtime.FreeCooling)
		values["Override"] = strconv.Itoa(runtime.Override)
		for k, v := range runtime.Heaters {
			idx := "Heat " + k
			values[idx] = strconv.Itoa(v)
		}
		for k, v := range runtime.Coolers {
			idx := "Cool " + k
			values[idx] = strconv.Itoa(v)
		}
		for k, v := range runtime.Aux {
			idx := "Aux " + k
			values[idx] = strconv.Itoa(v)
		}
		for k, v := range values {
			vLen := len(v)
			if l, ok := colWidths[k]; ok {
				if vLen > l {
					colWidths[k] = vLen
				}
			} else {
				colWidths[k] = len(k)
			}
		}
		rowValues[i] = values
	}
	colOrder := make([]string, len(colWidths))
	colOrder[0] = "Timestamp"
	colOrder[len(colWidths)-2] = "Free Cooling"
	colOrder[len(colWidths)-1] = "Override"
	colNext := 1
	for k := range colWidths {
		if strings.HasPrefix(k, "Heat") {
			colOrder[colNext] = k
			colNext++
		}
	}
	for k := range colWidths {
		if strings.HasPrefix(k, "Cool") {
			colOrder[colNext] = k
			colNext++
		}
	}
	for k := range colWidths {
		if strings.HasPrefix(k, "Aux") {
			colOrder[colNext] = k
			colNext++
		}
	}
	divBits := make([]string, len(colWidths))
	values := make([]string, len(colWidths))
	for i, k := range colOrder {
		width := colWidths[k]
		values[i] = fmt.Sprintf("%-*s", width, k)
		divBits[i] = strings.Repeat("-", width)
	}
	divider := fmt.Sprintf("  +-%s-+", strings.Join(divBits, "-+-"))
	fmt.Println(divider)
	fmt.Printf("  | %s |\n", strings.Join(values, " | "))
	fmt.Println(divider)
	for _, row := range rowValues {
		for i, k := range colOrder {
			width := colWidths[k]
			values[i] = fmt.Sprintf("%*s", width, row[k])
		}
		fmt.Printf("  | %s |\n", strings.Join(values, " | "))
	}
	fmt.Println(divider)
}
