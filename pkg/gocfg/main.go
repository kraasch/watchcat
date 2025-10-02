// Package gocfg reads a config file and parses it to TOML.
package gocfg

import (
	"fmt"
	"os"
	"strings"

	toml "github.com/pelletier/go-toml"
)

type WatchcatConfig struct {
	Filename string
	Path     string
}

var SEP = string(os.PathSeparator)

func ReadRawText(fullPath string) (string, error) {
	data, err := os.ReadFile(fullPath)
	return string(data), err
}

// InsertSeparator inserts a file path separator symbol, if path does not end in it yet.
func InsertSeparator(path, name string) string {
	var infix string
	if !strings.HasSuffix(path, SEP) {
		infix = SEP
	}
	targetPath := path + infix + name
	return targetPath
}

func (c *WatchcatConfig) GetPath() string {
	return InsertSeparator(c.Path, c.Filename)
}

func (c *WatchcatConfig) ReadConfigRawText() (string, error) {
	data, err := ReadRawText(c.GetPath())
	return string(data), err
}

type Target struct {
	Name          string
	Dir           string
	RulesLocation string
	Rules         string // multiline string
}

type SerializedConfig struct {
	Targets []Target `toml:"targets"`
}

func (c *WatchcatConfig) ParseToml() (SerializedConfig, error) {
	raw, err0 := c.ReadConfigRawText()
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
