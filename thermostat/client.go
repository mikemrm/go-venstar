package thermostat

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
	"time"

	"github.com/pkg/errors"
)

var (
	defaultTimeout = 5 * time.Second
	userAgent      = "github.com/mikemrm/go-venstar:0.1"
)

type thermostatClient interface {
	Do(*http.Request) (*http.Response, error)
}

type updateObject interface {
	BuildRequest(*http.Request) error
}

// Thermostat manages communication with Venstar API.
type Thermostat struct {
	client  thermostatClient
	baseURL url.URL
	pin     string
}

// SetPin sets the unlock pin when device has a pin set.
func (t *Thermostat) SetPin(pin string) {
	t.pin = pin
}

func (t *Thermostat) url(parts ...interface{}) string {
	pathParts := make([]string, len(parts))
	for _, part := range parts {
		pathParts = append(pathParts, strings.TrimLeft(fmt.Sprintf("%s", part), "/"))
	}
	nurl := t.baseURL
	nurl.Path = strings.Join(pathParts, "/")
	return nurl.String()
}

func (t *Thermostat) buildRequest(method, path string, body io.Reader) (*http.Request, error) {
	req, err := http.NewRequest(method, path, body)
	if err != nil {
		return nil, err
	}
	req.Header.Add("User-Agent", userAgent)
	return req, err
}

func (t *Thermostat) getJSON(path string, data interface{}) (*http.Response, error) {
	req, err := t.buildRequest("GET", path, nil)
	if err != nil {
		return nil, errors.Wrap(err, "building "+path+" request")
	}
	resp, err := t.client.Do(req)
	if err != nil {
		return resp, errors.Wrap(err, "requesting "+path)
	}
	err = DecodeBody(resp, data)
	if err != nil {
		return resp, errors.Wrap(err, "decoding "+path+" response")
	}
	return resp, nil
}

func (t *Thermostat) postJSON(path string, update updateObject, data interface{}) (*http.Response, error) {
	req, err := t.buildRequest("POST", path, nil)
	if err != nil {
		return nil, errors.Wrap(err, "building "+path+" request")
	}
	err = update.BuildRequest(req)
	if err != nil {
		return nil, errors.Wrap(err, "building "+path+" update request")
	}
	resp, err := t.client.Do(req)
	if err != nil {
		return resp, errors.Wrap(err, "requesting "+path)
	}
	err = DecodeBody(resp, data)
	if err != nil {
		return resp, errors.Wrap(err, "decoding "+path+" response")
	}
	return resp, nil
}

// GetAPIInfo retreives the general API information from the thermostat.
func (t *Thermostat) GetAPIInfo() (*APIInfo, error) {
	var info APIInfo
	_, err := t.getJSON(t.url("/"), &info)
	if err != nil {
		return nil, errors.Wrap(err, "processing api info request")
	}
	return &info, nil
}

// GetQueryInfo retreives overall stats from the thermostat.
func (t *Thermostat) GetQueryInfo() (*QueryInfo, error) {
	var info QueryInfo
	_, err := t.getJSON(t.url("/query/info"), &info)
	if err != nil {
		return nil, errors.Wrap(err, "processing query info request")
	}
	return &info, nil
}

// GetQuerySensors retreives the sensor readings from the thermostat.
func (t *Thermostat) GetQuerySensors() ([]*Sensor, error) {
	var info QueryResponse
	_, err := t.getJSON(t.url("/query/sensors"), &info)
	if err != nil {
		return nil, errors.Wrap(err, "processing query sensors request")
	}
	return info.Sensors, nil
}

// GetQueryRuntimes retreives the active system duration for each day.
// The results for each timestamp is for 24 hours prior.
func (t *Thermostat) GetQueryRuntimes() ([]*Runtime, error) {
	var info QueryResponse
	_, err := t.getJSON(t.url("/query/runtimes"), &info)
	if err != nil {
		return nil, errors.Wrap(err, "processing query runtime request")
	}
	return info.Runtimes, nil
}

// GetQueryAlerts retreives a list of alerts and whether they are triggered
// or not.
func (t *Thermostat) GetQueryAlerts() ([]*Alert, error) {
	var info QueryResponse
	_, err := t.getJSON(t.url("/query/alerts"), &info)
	if err != nil {
		return nil, errors.Wrap(err, "processing query alerts request")
	}
	return info.Alerts, nil
}

// UpdateControls submits the provided update request returning an error if the
// update failed.
func (t *Thermostat) UpdateControls(cr *ControlRequest) error {
	var updateResponse UpdateResponse
	_, err := t.postJSON(t.url("/control"), cr, &updateResponse)
	if err != nil {
		return errors.Wrap(err, "processing update control request")
	}
	if updateResponse.Error {
		return errors.New("Control Request update error: " + updateResponse.Reason)
	}
	if !updateResponse.Success {
		return errors.New("Control Request unknown error")
	}
	return nil
}

// UpdateSettings submits the provided update request returning an error if the
// update failed.
func (t *Thermostat) UpdateSettings(sr *SettingsRequest) error {
	var updateResponse UpdateResponse
	_, err := t.postJSON(t.url("/settings"), sr, &updateResponse)
	if err != nil {
		return errors.Wrap(err, "processing update settings request")
	}
	if updateResponse.Error {
		return errors.New("Settings Request update error: " + updateResponse.Reason)
	}
	if !updateResponse.Success {
		return errors.New("Settings Request unknown error")
	}
	return nil
}

// DecodeBody decodes the json http response body into the provided interface.
func DecodeBody(resp *http.Response, out interface{}) error {
	if resp.Body != nil {
		defer resp.Body.Close()

		var buf bytes.Buffer
		tee := io.TeeReader(resp.Body, &buf)
		resp.Body = ioutil.NopCloser(&buf)
		err := json.NewDecoder(tee).Decode(out)
		if err != nil {
			return errors.Wrap(err, "decoding json")
		}
	}
	return nil
}

// New creates a new Thermostat instance
func New(host string) *Thermostat {
	client := &http.Client{
		Timeout: defaultTimeout,
	}
	return &Thermostat{
		client:  client,
		baseURL: url.URL{Scheme: "http", Host: host},
	}
}
