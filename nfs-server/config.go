package nfsserver

import (
	"errors"
	"os"
	"strings"

	"github.com/allocamelus/allocamelus/pkg/logger"
	jsoniter "github.com/json-iterator/go"
	"k8s.io/klog/v2"
)

var json = jsoniter.ConfigCompatibleWithStandardLibrary

type Config struct {
	Exports []Export `json:"exports"`
}

type Export struct {
	Path  string  `json:"path"`
	Rules []Rules `json:"rules"`
}

type Rules struct {
	IP      string   `json:"ip"`
	Options []string `json:"options"`
}

// ReadConfig initializes Config
func ReadConfig(path string) (*Config, error) {
	configration := new(Config)
	file, err := os.Open(path)
	if err != nil {
		return &Config{}, errors.New("Error reading config @ " + path)
	}
	defer file.Close()

	decoder := json.NewDecoder(file)
	if err := decoder.Decode(&configration); err != nil {
		return &Config{}, err
	}
	return configration, nil
}

func (c *Config) GenExports() string {
	var exportsFile strings.Builder
	// Build
	// Path ip(options) ip?(options?)
	for _, e := range c.Exports {
		exportsFile.WriteString(e.Path)

		// Make dir if not exist
		if err := os.MkdirAll(e.Path, os.ModeSticky|os.ModePerm); err != nil {
			klog.Fatal("Error: Creating "+e.Path+" Failed: ", err)
		}

		// Check rules
		if len(e.Rules) == 0 {
			logger.Fatal(errors.New("Error: " + e.Path + " has no rules"))
		}
		// Iterate over rules
		for _, r := range e.Rules {
			exportsFile.WriteString(" " + r.IP + "(")
			oLen := len(r.Options)
			// Check Options
			if oLen == 0 {
				logger.Fatal(errors.New("Error: " + e.Path + " " + r.IP + " has no options"))
			}
			// Iterate over options
			for oi, o := range r.Options {
				exportsFile.WriteString(o)
				// Add comma if not last option
				if oLen-1 != oi {
					exportsFile.WriteRune(',')
				}
			}
			// Close (
			exportsFile.WriteRune(')')
		}
		// Add new line
		exportsFile.WriteRune('\n')
	}
	return exportsFile.String()
}
