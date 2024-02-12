package paygo

import (
	"fmt"
	"testing"
)

func NewCode(c string) PayPayResultCode {
	return PayPayResultCode(c)
}

func TestCode(t *testing.T) {
	var c PayPayResultCode
	c = NewCode("S0000")
	fmt.Println(c)
}
