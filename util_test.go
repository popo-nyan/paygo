package paygo

import (
	"fmt"
	"testing"
)

func TestGenerationDeviceInfo(t *testing.T) {
	d := GenerationDeviceInfo()
	fmt.Println(d)
}
