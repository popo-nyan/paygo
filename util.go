package paygo

import (
	"encoding/json"
	"log"
	"math/rand"
	"os"
)

func GenerationDeviceInfo() DeviceInfo {
	file, err := os.Open("devices.json")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	var d Devices
	decoder := json.NewDecoder(file)
	if err := decoder.Decode(&d); err != nil {
		log.Fatal(err)
	}
	return d.Samsung[rand.Intn(3-0)+0]
}
