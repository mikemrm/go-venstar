[![GoDev](https://img.shields.io/static/v1?label=godev&message=reference&color=00add8)](https://pkg.go.dev/github.com/mikemrm/go-venstar)
[![BuildStatus](https://github.com/mikemrm/go-venstar/workflows/Test/badge.svg)](https://github.com/mikemrm/go-venstar/actions?workflow=Test)
[![codecov](https://codecov.io/gh/mikemrm/go-venstar/branch/main/graph/badge.svg)](https://codecov.io/gh/mikemrm/go-venstar)

# Venstar Go library

Venstar is a Thermostat manufacturer which provides a snappy http API that can
be used to interact and retreive stats about your thermostat.

go-venstar aims to provide a clean library for interacting with this api
through Go.

*Note:* This library has only been tested on the _ColorTouch T8900_ however it
should work on all Venstar thermostats which support the Local API and follows
the restful docs.

## venstar-tstat

`venstar-tstat`s requires the ip of the thermostat to be provided.

The default action of venstar-tstat is to print all information available over
the venstar api.

If you pass any additional arguments, the command will process the updates,
followed by printing all information available over the api.

```shell
$ venstar-tstat -help
Usage of venstar-tstat:
  -controls.cool int
      Update Cool to temp (default -1)
  -controls.fan string
      Update Fan auto/on
  -controls.heat int
      Update Heat to temp (default -1)
  -controls.mode string
      Update Mode off/heat/cool/auto
  -settings.away string
      Update Away yes/no
  -settings.dehumidify-setpoint int
      Update Dehumidify SetPoint (25-99) (default -1)
  -settings.humidify-setpoint int
      Update Humidify SetPoint (0-60) (default -1)
  -settings.schedule string
      Update Schedule off/on
  -settings.tempunits string
      Update temperature units f/c fahrenheit/celsius
```

```shell
$ venstar-tstat 192.168.1.105
API Info:
  Type     : commercial
  Model    : COLORTOUCH
  Version  : 7
  Firmware : 5.10
Query Info:
  Name               : Thermostat
  Mode               : auto (3)
  State              : idle (0)
  Fan                : on (1)
  FanState           : off (0)
  ActiveStage        : 0
  TempUnits          : fahrenheit (0)
  Schedule           : inactive (0)
  SchedulePart       : inactive (255)
  Away               : home (0)
  Holiday            : not observing (0)
  Override           : off (0)
  OverrideRemaining  : 0s (0)
  ForceUnoccupied    : off (0)
  SpaceTemp          : 74.0
  HeatTemp           : 70.0
  CoolTemp           : 74.0
  CoolTempMin        : 35.0
  CoolTempMax        : 99.0
  HeatTempMin        : 35.0
  HeatTempMax        : 99.0
  HumidityEnabled    : disabled (0)
  Humidity           : 48%
  HumidifySetPoint   : 0%
  DehumidifySetPoint : 99%
  SetPointDelta      : 4.0
  AvailableModes     : all (0)
Query Sensors:
  Thermostat = 74.0
  Space Temp = 74.0
Query Runtimes:
  +-------------------------+--------+--------+----------+--------+-------+-------+--------------+----------+
  | Timestamp               | Heat 1 | Heat 2 | Cool 1   | Cool 2 | Aux 1 | Aux 2 | Free Cooling | Override |
  +-------------------------+--------+--------+----------+--------+-------+-------+--------------+----------+
  | 2015-09-15 00:00:00 UTC | 0s     | 0s     | 3h43m0s  | 1m0s   | 0s    | 0s    | 0s           | 0s       |
  | 2015-09-16 00:00:00 UTC | 0s     | 0s     | 7h6m0s   | 0s     | 0s    | 0s    | 0s           | 0s       |
  | 2015-09-17 00:00:00 UTC | 0s     | 0s     | 9h44m0s  | 0s     | 0s    | 0s    | 0s           | 0s       |
  | 2015-09-18 00:00:00 UTC | 0s     | 0s     | 9h19m0s  | 0s     | 0s    | 0s    | 0s           | 0s       |
  | 2015-09-19 00:00:00 UTC | 0s     | 0s     | 10h38m0s | 0s     | 0s    | 0s    | 0s           | 0s       |
  | 2015-09-19 10:16:10 UTC | 0s     | 0s     | 2h31m0s  | 0s     | 0s    | 0s    | 0s           | 0s       |
  +-------------------------+--------+--------+----------+--------+-------+-------+--------------+----------+
Query Alerts:
  Air Filter: Active = false
     UV Lamp: Active = false
     Service: Active = false
```

```shell
$ venstar-tstat/main.go -controls.mode auto -controls.heat 72 -controls.cool 76 192.168.1.105
Controls updated!
API Info:
  Type     : commercial
  Model    : COLORTOUCH
  Version  : 7
  Firmware : 5.10
Query Info:
...
  SpaceTemp          : 76.0
  HeatTemp           : 72.0
  CoolTemp           : 76.0
...
```