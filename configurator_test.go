package cspec

import (
	"fmt"
	"testing"
)

func TestRun(t *testing.T) {

	// take a chainspec in, eg from a client config file
	chainSpec := `"params": {
		"eip2028Transition": "{{.EIP2028.BlockNumber}}",
	}`

	// build the arguments, eg from Hive, Puppeth
	chainSpecArguments := NewChainSpecArguments()
	eip := EIP{
		Name:        2028,                    // the argument we expect to see applied
		BlockNumber: Number{Value: "8a61c8"}, //this would typically come in as marshalled json, need json marshalling tags on the various structs
	}
	chainSpecArguments.SetEIP(eip)

	// let's say we are dealing with Parity, which prefixes hex with 0x
	// set up a configurator, or pull one from a cache
	settings := Settings{NumberPrefix: "0x"}
	cfg, _ := NewConfigurator(settings)

	// apply the arguments to the chainspec:
	output, err := cfg.GenerateConfig(chainSpec, chainSpecArguments)
	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println(output)
	}

}
