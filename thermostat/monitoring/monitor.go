package monitoring

import (
	"time"

	"github.com/mikemrm/go-venstar/thermostat"
	"github.com/pkg/errors"
)

var (
	defaultDuration = 5 * time.Second
)

type Monitor struct {
	devices   []*thermostat.Thermostat
	frequency time.Duration
}

func (m *Monitor) GetDeviceResults(device *thermostat.Thermostat) (*Results, error) {
	results := &Results{
		Timestamp: time.Now(),
	}
	apiinfo, err := device.GetAPIInfo()
	if err != nil {
		return nil, errors.Wrap(err, "loading api info")
	}
	queryinfo, err := device.GetQueryInfo()
	if err != nil {
		return nil, errors.Wrap(err, "loading query info")
	}
	results.APIInfo = apiinfo
	results.QueryInfo = queryinfo
	return results, nil
}

func (m *Monitor) run(resultsChan chan *Results, errorsChan chan error) {
	var results *Results
	var err error
	for _, device := range m.devices {
		results, err = m.GetDeviceResults(device)
		if err != nil {
			errorsChan <- err
			continue
		}
		resultsChan <- results
	}
}

func (m *Monitor) monitor(resultsChan chan *Results, errorsChan chan error, stopChan chan bool) {
	ticker := time.NewTicker(m.frequency)
	defer ticker.Stop()
	for {
		select {
		case <-stopChan:
			return
		case <-ticker.C:
			m.run(resultsChan, errorsChan)
		}
	}
}

func (m *Monitor) Monitor(stopChan chan bool) (chan *Results, chan error) {
	resultsChan := make(chan *Results, 1)
	errorsChan := make(chan error, 1)
	go func() {
		defer close(resultsChan)
		defer close(errorsChan)
		m.monitor(resultsChan, errorsChan, stopChan)
	}()
	return resultsChan, errorsChan
}

func New(hosts ...string) *Monitor {
	devices := make([]*thermostat.Thermostat, len(hosts))
	for i, host := range hosts {
		devices[i] = thermostat.New(host)
	}
	return &Monitor{
		devices:   devices,
		frequency: defaultDuration,
	}
}
