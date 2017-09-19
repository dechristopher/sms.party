package util

import (
	"encoding/json"
	"html/template"
	"io/ioutil"
	"log"
	"os"

	s "github.com/dechristopher/sms.party/src/strings"
)

var (
	// Conf Global configuration
	Conf Configuration
	// Templates stores statically compiled HTML templates
	Templates = template.Must(template.ParseFiles("web/index.html"))
)

// Configuration for csms-backend
type Configuration struct {
	Port    string `json:"port"`
	Prefix  string `json:"prefix"`
	Rate    int    `json:"rate"`
	Version string `json:"version"`
}

// ReadConfig reads config.json into memory for global configuration
func ReadConfig() Configuration {
	raw, err := ioutil.ReadFile("./config.json")
	if err != nil {
		log.Fatal(s.LogPrefix + "Configuration Error! config.json improperly loaded or does not exist :: " + err.Error())
		os.Exit(1)
	}

	var c Configuration
	if errUnm := json.Unmarshal(raw, &c); errUnm != nil {
		log.Fatal(s.LogPrefix + "Configuration Error! config.json improperly unmarshaled :: " + errUnm.Error())
		os.Exit(1)
	}

	c.Prefix = c.Prefix + "/v" + c.Version

	return c
}
