package cspec

import (
	"fmt"
	"reflect"
	"sync"
)

// A chainspec is considered here as a
// parameterised config file template.
// This structure represents the arguments
// to such a template.
type chainSpecArguments struct {
	mutex       sync.Mutex
	dynamicType reflect.Value //the type used by the template engine
}

// ChainSpecArguments is a set of chainspec arguments
// to a parameterised config template expressed as a
// set of EIPs.
type ChainSpecArguments interface {
	SetEIP(eip EIP) //pass by value?
	getInternal() reflect.Value
}

// SetEIP passes an EIP value into the chainspecarguments
func (c *chainSpecArguments) SetEIP(eip EIP) {
	c.mutex.Lock()

	//If the name is "0" we should attach the field directly to the arguments
	//so that we get things like maxCodeSize directly available
	//instead of having to indirect via an EIP0 field in the
	//chainspec

	eipName := fmt.Sprintf("EIP%d", eip.Name)
	c.dynamicType.Elem().FieldByName(eipName).Set(reflect.ValueOf(eip))
	c.mutex.Unlock()
}

func (c *chainSpecArguments) getInternal() reflect.Value {
	return c.dynamicType
}

func newChainSpecArguments() ChainSpecArguments {

	// the template engine requires that properties be
	// directly available. a *map* of eip names to
	// config arguments would require an overcomplex 'index'
	// syntax, which is no good for maintenance or
	// adoption. for this reason we synthesise a type
	// dynamically:
	eipFields := make([]reflect.StructField, maxEIPs)
	for i := 0; i < maxEIPs; i++ {
		eipFields[i] = reflect.StructField{Name: fmt.Sprintf("EIP%d", i), Type: eipType}
	}
	return &chainSpecArguments{
		dynamicType: reflect.New(reflect.StructOf(eipFields)),
	}
}

// NewChainSpecArguments is a constructor for the empty chainspec arguments
func NewChainSpecArguments() ChainSpecArguments {
	return newChainSpecArguments()
}

// TODO - add more default constructors here, eg: Istanbul Mainnet etc...
