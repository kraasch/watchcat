// Package gocfg reads a config file and parses it to TOML.
package gocfg

import (
	"fmt"
	"os"

	toml "github.com/pelletier/go-toml"
)

type Config struct {
	Filename string
	Path     string
}

var osSep = string(os.PathSeparator)

func (c *Config) ReadRawText() (string, error) {
	path := c.Path + osSep + c.Filename
	data, err := os.ReadFile(path)
	return string(data), err
}

type Target struct {
	Name  string
	Dir   string
	Rules string // multiline string
}

type SerializedConfig struct {
	Targets []Target `toml:"targets"`
}

func (c *Config) ParseToml() (SerializedConfig, error) {
	raw, err0 := c.ReadRawText()
	if err0 != nil {
		return SerializedConfig{}, err0
	}
	var cfg SerializedConfig
	err := toml.Unmarshal([]byte(raw), &cfg)
	return cfg, err
}

const (
	value = "Toast: "
)

func Toast(in string) string {
	return fmt.Sprintf("%#v", value+in)
}
