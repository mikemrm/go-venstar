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
)

type Thermostat struct {
	client  *http.Client
	baseURL url.URL
	pin     int
}

func (t *Thermostat) SetPin(pin int) {
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

func (t *Thermostat) GetAPIInfo() (*APIInfo, error) {
	req, err := http.NewRequest("GET", t.url("/"), nil)
	if err != nil {
		return nil, errors.Wrap(err, "creating api info request")
	}
	req.Header.Add("User-Agent", "github.com/mikemrm/go-venstar:0.1")
	resp, err := t.client.Do(req)
	if err != nil {
		return nil, errors.Wrap(err, "sending api info request")
	}
	var info APIInfo
	err = DecodeBody(resp, &info)
	if err != nil {
		return nil, errors.Wrap(err, "decoding api info response")
	}
	return &info, nil
}

func (t *Thermostat) GetQueryInfo() (*QueryInfo, error) {
	req, err := http.NewRequest("GET", t.url("/query/info"), nil)
	if err != nil {
		return nil, errors.Wrap(err, "creating query info request")
	}
	req.Header.Add("User-Agent", "github.com/mikemrm/go-venstar:0.1")
	resp, err := t.client.Do(req)
	if err != nil {
		return nil, errors.Wrap(err, "sending query info request")
	}
	var info QueryInfo
	err = DecodeBody(resp, &info)
	if err != nil {
		return nil, errors.Wrap(err, "decoding query info response")
	}
	return &info, nil
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
