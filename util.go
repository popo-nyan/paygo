package paygo

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"log"
	"math/rand"
	"os"
)

func GenerateDeviceInfo() DeviceInfo {
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

func HashCalculation(msg string) string {
	mac := hmac.New(sha256.New, []byte(HmacKey))
	mac.Write([]byte(msg))
	return hex.EncodeToString(mac.Sum(nil))
}
