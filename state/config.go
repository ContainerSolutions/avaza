package state

import (
	"encoding/json"
	"github.com/ContainerSolutions/avaza/auth"
	xdg_ "github.com/OpenPeeDeeP/xdg"
	"io/ioutil"
	"log"
	"os"
)

var xdg *xdg_.XDG

func init() {
	xdg = xdg_.New("", "avaza-cli")
}

type Config struct {
	BearerToken *auth.BearerToken
}

func LoadConfig() *Config {
	if _, err := os.Stat(ConfigFile()); os.IsNotExist(err) {
		// Nothing configured yet; return empty configuration
		return &Config{}
	}

	file, err := os.Open(ConfigFile())
	if err != nil {
		log.Fatal("Could not open config file")
	}
	defer file.Close()

	decoder := json.NewDecoder(file)
	configuration := Config{}
	err = decoder.Decode(&configuration)
	if err != nil {
		log.Fatalf("Could not parse config file; %s", err)
	}
	return &configuration
}

func SaveConfig(c *Config) {
	jsonConfig, _ := json.MarshalIndent(c, "", "\t")
	err := ioutil.WriteFile(ConfigFile(), jsonConfig, 0644)
	if err != nil {
		log.Fatalf("Could not write config file: %s", err)
	}
}

func ConfigFile() string {
	return xdg.ConfigHome() + ".json"
}

func CacheHome() string {
	return xdg.CacheHome()
}
