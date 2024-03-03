package models

type Device struct {
	DeviceName string `json:"device_name"`
	Browser    string `json:"browser"`
	IP         string `json:"ip"`
	BrowserVersion    string `json:"browser_version"`
}
