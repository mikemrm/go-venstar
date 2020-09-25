package thermostat

// APIInfo is the struct used for `/` results.
type APIInfo struct {
	Version  int    `json:"api_ver"`
	Model    string `json:"model"`
	Firmware string `json:"firmware"`
	Type     string `json:"type"`
}

// UpdateResponse is the struct used for update request responses.
type UpdateResponse struct {
	Success bool   `json:"success,omitempty"`
	Error   bool   `json:"error"`
	Reason  string `json:"reason,omitempty"`
}
