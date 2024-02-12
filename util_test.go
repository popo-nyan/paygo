package paygo

import (
	"fmt"
	"testing"
)

func TestGenerateDeviceInfo(t *testing.T) {
	d := GenerateDeviceInfo()
	fmt.Println(d)
}

func TestHashCalculation(t *testing.T) {
	h := HashCalculation("v1/config", "GET")
	fmt.Println(h)
}
