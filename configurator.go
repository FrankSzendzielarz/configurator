package cspec

import (
	"bytes"
	"fmt"
	"reflect"
	"strings"
	"text/template"
)

var (
	eipType reflect.Type
)

const (
	maxEIPs = 3000
)

func init() {
	var eip EIP
	eipType = reflect.TypeOf(eip)
}

// Configurator applies arguments to
// parameterised chainspec and command line
// strings.
type Configurator interface {
	GenerateConfig(chainspec string, arguments ChainSpecArguments) (string, error)
}

// NewConfigurator creates a new configurator with
// the desired settings
func NewConfigurator(settings Settings) (Configurator, error) {

	configurator := &configurator{
		settings: settings,
	}

	return configurator, nil
}

type configurator struct {
	settings Settings
}

// GenerateConfig accepts the parameterised chain
// spec and fills it using the chainspec arguments
func (c *configurator) GenerateConfig(chainspec string, arguments ChainSpecArguments) (string, error) {

	tmpl, error := template.New("initial").Parse(chainspec)
	if error != nil {
		return "", error
	}

	// generate an intermediate template. the challenge
	// this solves is rendering the template according
	// to the current context, which exists only
	// while the chainspecarguments are being applied
	// to the chainspec. the chainspecarguments are
	// reusable across chainspecs, and the chainspecs
	// are reusable with different arguments. by generating
	// an intermediate template, we get a template that
	// in the desired output format but with functions
	// embedded that call on this configurators formatting
	// methods.

	intermediate := new(bytes.Buffer)
	tmpl.Execute(intermediate, arguments.getInternal())
	intermediateTemplate := intermediate.String()
	tmpl, error = template.New("intermediate").Parse(intermediateTemplate)
	if error != nil {
		return "", error
	}
	output := new(bytes.Buffer)
	tmpl.Execute(output, c) //execute the template

	return output.String(), nil
}

// GenerateConfig accepts the parameterised
// command line and fills it with the flag arguments
func (c *configurator) GenerateFlags(flagspec string) (string, error) {
	//TODO
	return "", nil
}

func (c *configurator) FormatNumber(number string) string {
	if !strings.HasPrefix(number, c.settings.NumberPrefix) {
		return fmt.Sprintf("%s%s", c.settings.NumberPrefix, number)
	}
	return number
}
