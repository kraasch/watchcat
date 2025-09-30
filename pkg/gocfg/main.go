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

type MyConfig struct { // NOTE: example struct.
	Version int
	Name    string
	Tags    []string
}

// NOTE: example.
var doc = `
version = 2
name = "go-toml"
tags = ["go", "toml"]
`

func (c *Config) ParseToml() (MyConfig, error) {
	var cfg MyConfig
	err := toml.Unmarshal([]byte(doc), &cfg)
	return cfg, err
}

const (
	value = "Toast: "
)

func Toast(in string) string {
	return fmt.Sprintf("%#v", value+in)
}
