package thermostat

import (
	"errors"
	"io"
	"net/http"
	"strings"
	"testing"
)

func TestSetPin(t *testing.T) {
	tstat := &Thermostat{}
	want := "1597"
	tstat.SetPin(want)
	if tstat.pin != want {
		t.Error("pin, got:", tstat.pin, "want:", want)
	}
}

func TestBuildRequest(t *testing.T) {
	tstat := &Thermostat{}

	t.Run("invalid request returns error", func(t *testing.T) {
		_, err := tstat.buildRequest("GET", "://bad", nil)
		if err == nil {
			t.Error("error expected, but none returned")
		}
	})

	testMethods := []string{"GET", "POST"}

	for _, method := range testMethods {
		t.Run(method+" method is set", func(t *testing.T) {
			req, err := tstat.buildRequest(method, "http://localhost", nil)
			if err != nil {
				t.Fatal("returned error when none was expected:", err)
			}
			if req.Method != method {
				t.Error("method doesn't match, got:", req.Method, "want:", method)
			}
		})
	}

	wantUserAgent := "go.mrm.dev/venstar:0.1"
	for _, method := range testMethods {
		t.Run(method+" User-Agent header is set", func(t *testing.T) {
			req, err := tstat.buildRequest(method, "http://localhost", nil)
			if err != nil {
				t.Fatal("returned error when none was expected:", err)
			}
			gotUserAgent := req.Header.Get("User-Agent")
			if gotUserAgent != wantUserAgent {
				t.Error("User-Agent header doesn't match,",
					"got:", gotUserAgent, "want:", wantUserAgent)
			}
		})
	}

	t.Run("nil request body stays nil", func(t *testing.T) {
		req, err := tstat.buildRequest("POST", "http://localhost", nil)
		if err != nil {
			t.Fatal("returned error when none was expected:", err)
		}
		if req.Body != nil {
			t.Error("body not nil, expected nil")
		}
	})

	t.Run("request body stays in tact", func(t *testing.T) {
		want := "this is test text"
		body := strings.NewReader(want)
		req, err := tstat.buildRequest("POST", "http://localhost", body)
		if err != nil {
			t.Fatal("returned error when none was expected:", err)
		}
		if req.Body == nil {
			t.Fatal("body is nil, want:", want)
		}
		got, err := io.ReadAll(req.Body)
		if err != nil {
			t.Fatal("error produced reading body:", err)
		}
		if string(got) != want {
			t.Error("body got:", got, "want:", want)
		}
	})
}

type fakeThermostatClient struct {
	body  string
	error bool
}

func (c *fakeThermostatClient) Do(_ *http.Request) (*http.Response, error) {
	var body io.ReadCloser
	if c.body != "" {
		body = io.NopCloser(strings.NewReader(c.body))
	}
	resp := &http.Response{
		Body: body,
	}
	var err error
	if c.error {
		err = errors.New("this is an error")
	}
	return resp, err
}

func TestGetJSON(t *testing.T) {
	tests := []struct {
		name    string
		url     string
		reqBody string
		expBody string
		expErr  string
	}{
		{"invalid request returns error", "://bad", "", "", `building ://bad request: parse "://bad": missing protocol scheme`},
		{"bad response returns error", "req/path", "", "", `requesting req/path: this is an error`},
		{"invalid json returns error", "req/path", `invalid"`, "", `decoding req/path response: decoding json: invalid character 'i' looking for beginning of value`},
		{"valid json gets decoded", "req/path", `"valid"`, "valid", ""},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			tstatClient := &fakeThermostatClient{
				body: test.reqBody,
			}
			if test.reqBody == "" && test.expErr != "" {
				tstatClient.error = true
			}
			tstat := &Thermostat{
				client: tstatClient,
			}

			var out string
			_, err := tstat.getJSON(test.url, &out)

			if test.expErr != "" && err == nil {
				t.Fatal("error", test.expErr, "wanted, got nil")
			}
			if test.expErr != "" && err != nil && test.expErr != err.Error() {
				t.Fatal("error expected, got:", err.Error(), "want:", test.expErr)
			}
		})
	}
}

type fakeUpdate struct {
	error bool
}

func (fu *fakeUpdate) BuildRequest(_ *http.Request) error {
	if fu.error {
		return errors.New("this is an error")
	}
	return nil
}

func TestPostJSON(t *testing.T) {
	tests := []struct {
		name      string
		url       string
		reqBody   string
		expBody   string
		updateErr bool
		expErr    string
	}{
		{"invalid request returns error", "://bad", "", "", false, `building ://bad request: parse "://bad": missing protocol scheme`},
		{"invalid update returns error", "req/path", "", "", true, `building req/path update request: this is an error`},
		{"bad response returns error", "req/path", "", "", false, `requesting req/path: this is an error`},
		{"invalid json returns error", "req/path", `invalid"`, "", false, `decoding req/path response: decoding json: invalid character 'i' looking for beginning of value`},
		{"valid json gets decoded", "req/path", `"valid"`, "valid", false, ""},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			tstatClient := &fakeThermostatClient{
				body: test.reqBody,
			}
			if test.reqBody == "" && test.expErr != "" {
				tstatClient.error = true
			}
			tstat := &Thermostat{
				client: tstatClient,
			}

			update := &fakeUpdate{test.updateErr}

			var out string
			_, err := tstat.postJSON(test.url, update, &out)

			if test.expErr != "" && err == nil {
				t.Fatal("error", test.expErr, "wanted, got nil")
			}
			if test.expErr != "" && err != nil && test.expErr != err.Error() {
				t.Fatal("error expected, got:", err.Error(), "want:", test.expErr)
			}
		})
	}
}

func TestNew(t *testing.T) {
	tstat := New("127.0.0.1")
	t.Run("empty on creation", func(t *testing.T) {
		if tstat.pin != "" {
			t.Error("pin: got", tstat.pin, "want ''")
		}
	})
}
