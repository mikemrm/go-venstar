package thermostat

type APIInfo struct {
	Version  int    `json:"api_ver"`
	Model    string `json:"model"`
	Firmware string `json:"firmware"`
	Type     string `json:"type"`
}

type UpdateResponse struct {
	Success bool   `json:"success,omitempty"`
	Error   bool   `json:"error"`
	Reason  string `json:"reason,omitempty"`
}
