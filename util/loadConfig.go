package util

import (
	"Geo-DNS/structs"
	"encoding/json"
	"log"
	"os"
)

func LoadConfig() structs.Data {
	c, e := os.ReadFile("./config.json")
	if e != nil {
		log.Fatal(e)
	}

	var payload structs.Data
	e = json.Unmarshal(c, &payload)
	if e != nil {
		log.Fatal(e)
	}
	return payload
}
