package config

import (
	"fmt"
	"github.com/xeipuuv/gojsonschema"
	"gopkg.in/sensorbee/sensorbee.v0/data"
)

const (
	// DefaultPort is the default port number used by the SensorBee server.
	DefaultPort = 15601
)

// Network has configuration parameters related to the network.
type Network struct {
	// ListenOn has binding information in "host:port" format.
	ListenOn string `json:"listen_on" yaml:"listen_on"`
}

var (
	networkSchemaString = `{
	"type": "object",
	"properties": {
		"listen_on": {
			"type": "string",
			"pattern": "^.*:[0-9]+$"
		}
	},
	"additionalProperties": false
}`
	networkSchema *gojsonschema.Schema
)

func init() {
	s, err := gojsonschema.NewSchema(gojsonschema.NewStringLoader(networkSchemaString))
	if err != nil {
		panic(err)
	}
	networkSchema = s
}

// NewNetwork creates a Newtork config parameters from a given map.
func NewNetwork(m data.Map) (*Network, error) {
	if err := validate(networkSchema, m); err != nil {
		return nil, err
	}
	return newNetwork(m), nil
}

func newNetwork(m data.Map) *Network {
	return &Network{
		ListenOn: mustAsString(getWithDefault(m, "listen_on", data.String(fmt.Sprintf(":%d", DefaultPort)))),
	}
}

// ToMap returns network config information as data.Map.
func (n *Network) ToMap() data.Map {
	return data.Map{
		"listen_on": data.String(n.ListenOn),
	}
}
