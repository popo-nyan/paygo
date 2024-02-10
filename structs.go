package paygo

import "net/http"

type Session struct {
	Client *http.Client

	AccessToken string

	Debug    bool
	logLevel int

	AppVersion string
	ClientName string

	PayPayLang             string
	DeviceUuid             string
	ClientUuid             string
	DeviceName             string
	DeviceOsVersion        string
	DeviceOsReleaseVersion string
	Timezone               string

	UserAgent string
}
