package cspec

import "fmt"

// Number is usually a hex value
type Number struct {
	Value string // should this be an actual numeric?
}

func (n *Number) String() string {
	return fmt.Sprintf("{{.FormatNumber \"%s\" }}", n.Value)
}
