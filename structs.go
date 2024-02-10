package paygo

import (
	"net/http"
)

type Session struct {
	Client *http.Client

	AccessToken string

	Debug    bool
	logLevel int

	AppVersion string
	ClientName string

	DeviceUuid string
	ClientUuid string
	UserAgent  string
}
type Devices struct {
	Samsung []DeviceInfo `json:"samsung"`
}

type DeviceInfo struct {
	ModelName        string `json:"modelName"`
	OsVersion        string `json:"osVersion"`
	OsReleaseVersion string `json:"osReleaseVersion"`
	HardwareName     string `json:"hardwareName"`
}
