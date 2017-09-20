package util

import (
	"encoding/json"
	"html/template"
	"io/ioutil"
	"log"
	"os"

	redis "gopkg.in/redis.v5"

	s "github.com/dechristopher/sms.party/src/strings"
)

var (
	// Conf Global configuration
	Conf Configuration
	// Templates stores statically compiled HTML templates
	Templates = template.Must(template.ParseFiles("web/index.html"))
	// R a redis client instance
	R = redis.NewClient(&redis.Options{
		Addr:     "",
		Password: "",
		DB:       1,
	})
	// RedisConnected is what it sounds like
	RedisConnected = false
)

// Configuration for csms-backend
type Configuration struct {
	Port      string `json:"port"`
	Prefix    string `json:"prefix"`
	Rate      int    `json:"rate"`
	Version   string `json:"version"`
	RedisConf Redis  `json:"redisconf"`
}

// Redis configuration
type Redis struct {
	Address  string `json:"address"`
	Password string `json:"password"`
	DB       int    `json:"db"`
}

// ReadConfig reads config.json into memory for global configuration
func ReadConfig() Configuration {
	// Check to make sure config.json exists
	raw, err := ioutil.ReadFile("./config.json")
	if err != nil {
		log.Fatal(s.LogPrefix + "Configuration Error! config.json improperly loaded or does not exist :: " + err.Error())
		os.Exit(1)
	}

	// Unmarshal the raw json config into a Configuration object
	var c Configuration
	if errUnm := json.Unmarshal(raw, &c); errUnm != nil {
		log.Fatal(s.LogPrefix + "Configuration Error! config.json improperly unmarshaled :: " + errUnm.Error())
		os.Exit(1)
	}

	// Set the full API prefix slug
	c.Prefix = c.Prefix + "/v" + c.Version

	// Trash the fake init client and actually connect to the datastore
	R = redis.NewClient(&redis.Options{
		Addr:     c.RedisConf.Address,
		Password: c.RedisConf.Password,
		DB:       c.RedisConf.DB,
	})

	if rperr := R.Ping().Err(); rperr != nil {
		LogErr(s.ErrRedisConnection)
		RedisConnected = false
	} else {
		RedisConnected = true
	}

	return c
}
