package models

type Device struct {
	Id             string `json:"id"`
	UserId         string `json:"user_id"`
	DeviceName     string `json:"device_name"`
	Browser        string `json:"browser"`
	IP             string `json:"ip"`
	BrowserVersion string `json:"browser_version"`
	CreatedAt      string `json:"created_at"`
}
type DeviceCreate struct {
	Id             string `json:"id"`
	UserId         string `json:"user_id"`
	DeviceName     string `json:"device_name"`
	Browser        string `json:"browser"`
	BrowserVersion string `json:"browser_version"`
	IP             string `json:"ip"`
}
type DevicePrimaryKey struct {
	Id string `json:"id"`
}

type DeviceGetListRequest struct {
	Offset int `json:"offset"`
	Limit  int `json:"limit"`
	UserId string `json:"user_id"`
}

type DeviceGetListResponse struct {
	Count   int       `json:"count"`
	Devices []*Device `json:"datas"`
}
