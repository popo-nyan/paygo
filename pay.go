package paygo

import (
	"net/http"
	"time"
)

const VERSION = "0.0.1"

func New(token string) (s *Session, err error) {
	d := GenerateDeviceInfo()
	s = &Session{
		Client:      &http.Client{Timeout: 30 * time.Second},
		AccessToken: token,
		UserAgent:   "PaypayApp/" + PayPayAppVersion + " Android" + d.OsReleaseVersion,
	}
	c := OrderStatusSuccess
	return
}
