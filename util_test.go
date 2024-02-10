package paygo

import (
	"fmt"
	"testing"
)

func TestGenerateDeviceInfo(t *testing.T) {
	d := GenerateDeviceInfo()
	fmt.Println(d)
}
