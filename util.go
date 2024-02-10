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
func MakeHMAC(msg, key string) string {
	mac := hmac.New(sha256.New, []byte(key))
	mac.Write([]byte(msg))
	return hex.EncodeToString(mac.Sum(nil))

}
func CalculationHash(hpm string, ep string, by string) string {
	mac := hmac.New(sha256.New, []byte(HmacKey))
	mac.Write(([]))

}
