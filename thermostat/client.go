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

type Thermostat struct {
	client  *http.Client
	baseURL url.URL
	pin     string
}

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

func (t *Thermostat) request(method, path string, body interface{}) (*http.Request, error) {
	req, err := http.NewRequest(method, path, body)
	if err != nil {
		return nil, err
	}
	req.Header.Add("User-Agent", userAgent)
	return req, err
}

func (t *Thermostat) getJSON(path string, data interface{}) (*http.Response, error) {
	req, err := t.request("GET", path, nil)
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

func (t *Thermostat) GetAPIInfo() (*APIInfo, error) {

	var info APIInfo
	_, err := t.getJSON(t.url("/"), &info)
	if err != nil {
		return nil, errors.Wrap(err, "processing api info request")
	}
	return &info, nil
}

func (t *Thermostat) GetQueryInfo() (*QueryInfo, error) {
	var info QueryInfo
	_, err := t.getJSON(t.url("/query/info"), &info)
	if err != nil {
		return nil, errors.Wrap(err, "processing query info request")
	}
	return &info, nil
}

func (t *Thermostat) GetQuerySensors() ([]*Sensor, error) {
	var info QueryResponse
	_, err := t.getJSON(t.url("/query/sensors"), &info)
	if err != nil {
		return nil, errors.Wrap(err, "processing query sensors request")
	}
	return info.Sensors, nil
}

func (t *Thermostat) GetQueryRuntimes() ([]*Runtime, error) {
	var info QueryResponse
	_, err := t.getJSON(t.url("/query/runtimes"), &info)
	if err != nil {
		return nil, errors.Wrap(err, "processing query runtime request")
	}
	return info.Runtimes, nil
}

func (t *Thermostat) GetQueryAlerts() ([]*Alert, error) {
	var info QueryResponse
	_, err := t.getJSON(t.url("/query/alerts"), &info)
	if err != nil {
		return nil, errors.Wrap(err, "processing query alerts request")
	}
	return info.Alerts, nil
}

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

func New(host string) *Thermostat {
	client := &http.Client{
		Timeout: defaultTimeout,
	}
	return &Thermostat{
		client:  client,
		baseURL: url.URL{Scheme: "http", Host: host},
	}
}
