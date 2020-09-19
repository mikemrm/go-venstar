package main

import (
	"fmt"
	"os"
	"sort"
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
	fmt.Printf("  Name               : %s\n", qinfo.Name)
	fmt.Printf("  Mode               : %s (%d)\n", qinfo.Mode.String(), qinfo.Mode)
	fmt.Printf("  State              : %s (%d)\n", qinfo.State.String(), qinfo.State)
	fmt.Printf("  Fan                : %s (%d)\n", qinfo.Fan.String(), qinfo.Fan)
	fmt.Printf("  FanState           : %s (%d)\n", qinfo.FanState.String(), qinfo.FanState)
	fmt.Printf("  ActiveStage        : %d\n", qinfo.ActiveStage)
	fmt.Printf("  TempUnits          : %s (%d)\n", qinfo.TempUnits.String(), qinfo.TempUnits)
	fmt.Printf("  Schedule           : %s (%d)\n", qinfo.Schedule.String(), qinfo.Schedule)
	fmt.Printf("  SchedulePart       : %s (%d)\n", qinfo.SchedulePart.String(), qinfo.SchedulePart)
	fmt.Printf("  Away               : %s (%d)\n", qinfo.Away.String(), qinfo.Away)
	fmt.Printf("  Holiday            : %s (%d)\n", qinfo.Holiday.String(), qinfo.Holiday)
	fmt.Printf("  Override           : %s (%d)\n", qinfo.Override.String(), qinfo.Override)
	fmt.Printf("  OverrideRemaining  : %s (%d)\n", qinfo.OverrideRemaining.String(), qinfo.OverrideRemaining)
	fmt.Printf("  ForceUnoccupied    : %s (%d)\n", qinfo.ForceUnoccupied.String(), qinfo.ForceUnoccupied)
	fmt.Printf("  SpaceTemp          : %.1f\n", qinfo.SpaceTemp)
	fmt.Printf("  HeatTemp           : %.1f\n", qinfo.HeatTemp)
	fmt.Printf("  CoolTemp           : %.1f\n", qinfo.CoolTemp)
	fmt.Printf("  CoolTempMin        : %.1f\n", qinfo.CoolTempMin)
	fmt.Printf("  CoolTempMax        : %.1f\n", qinfo.CoolTempMax)
	fmt.Printf("  HeatTempMin        : %.1f\n", qinfo.HeatTempMin)
	fmt.Printf("  HeatTempMax        : %.1f\n", qinfo.HeatTempMax)
	fmt.Printf("  HumidityEnabled    : %s (%d)\n", qinfo.HumidityEnabled.String(), qinfo.HumidityEnabled)
	fmt.Printf("  Humidity           : %d%%\n", qinfo.Humidity)
	fmt.Printf("  HumidifySetPoint   : %d%%\n", qinfo.HumidifySetPoint)
	fmt.Printf("  DehumidifySetPoint : %d%%\n", qinfo.DehumidifySetPoint)
	fmt.Printf("  SetPointDelta      : %.1f\n", qinfo.SetPointDelta)
	fmt.Printf("  AvailableModes     : %s (%d)\n", qinfo.AvailableModes.String(), qinfo.AvailableModes)

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
		values["Free Cooling"] = runtime.FreeCooling.String()
		values["Override"] = runtime.Override.String()
		for k, v := range runtime.Heaters {
			idx := "Heat " + k
			values[idx] = v.String()
		}
		for k, v := range runtime.Coolers {
			idx := "Cool " + k
			values[idx] = v.String()
		}
		for k, v := range runtime.Aux {
			idx := "Aux " + k
			values[idx] = v.String()
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
	columns := make([]string, len(colWidths))
	colNext := 0
	for k := range colWidths {
		columns[colNext] = k
		colNext++
	}
	sort.Strings(columns)

	colOrder := make([]string, len(columns))
	colOrder[0] = "Timestamp"
	colOrder[len(columns)-2] = "Free Cooling"
	colOrder[len(columns)-1] = "Override"
	colNext = 1
	for _, k := range columns {
		if strings.HasPrefix(k, "Heat") {
			colOrder[colNext] = k
			colNext++
		}
	}
	for _, k := range columns {
		if strings.HasPrefix(k, "Cool") {
			colOrder[colNext] = k
			colNext++
		}
	}
	for _, k := range columns {
		if strings.HasPrefix(k, "Aux") {
			colOrder[colNext] = k
			colNext++
		}
	}
	divBits := make([]string, len(columns))
	values := make([]string, len(columns))
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
			values[i] = fmt.Sprintf("%-*s", width, row[k])
		}
		fmt.Printf("  | %s |\n", strings.Join(values, " | "))
	}
	fmt.Println(divider)
}
