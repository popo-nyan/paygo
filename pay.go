package paygo

import (
	"net/http"
	"time"
)

const (
	VERSION          = "0.0.1"
	PAYPAYAPPVERSION = "4.31.0"
)

func New(token string) (s *Session, err error) {
	d := GenerationDeviceInfo()
	s = &Session{
		Client:      &http.Client{Timeout: 30 * time.Second},
		AccessToken: token,
		UserAgent:   "PaypayApp/" + PAYPAYAPPVERSION + " Android" + d.OsReleaseVersion,
	}
	return
}
