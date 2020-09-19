# Venstar Go library

## Example Output

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