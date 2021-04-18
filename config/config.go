package config

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"

	logger "github.com/dm1trypon/easy-logger"
	"github.com/dm1trypon/engine-srv/jsonvalidator"
)

// LC - Logging category
const LC = "CONFIG"

const (
	// ConfigPath - path to config file
	ConfigPath = "./config.json"
	// ConfigSchemaPath - path to config's schema file
	ConfigSchemaPath = "./config.schema.json"
)

type Config struct {
	configData configData
}

// ConfigInst - config instance
var ConfigInst *Config

func (c *Config) Create() *Config {
	logger.Info(LC, "Creating Config module")

	c.readConfig()

	return c
}

func (c *Config) GetConfigData() configData {
	return c.configData
}

// readConfig - a method that checks the config for validation.
func (c *Config) readConfig() {
	cfgByte, err := ioutil.ReadFile(ConfigPath)
	if err != nil {
		logger.Error(LC, fmt.Sprint("Config read error: ", err.Error()))
		os.Exit(-1)
	}

	if !jsonvalidator.Check(cfgByte, ConfigSchemaPath) {
		os.Exit(-1)
	}

	if err := json.Unmarshal(cfgByte, &c.configData); err != nil {
		logger.Error(LC, fmt.Sprint("Error writing config to structure: ", err.Error()))
		os.Exit(-1)
	}
}
