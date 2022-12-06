package pflagvalue

import (
	"fmt"
	"strconv"
)

// uint16PtrValue is a flag.Value which stores the value in a *uint16 if it
// can be parsed with strconv.ParseInt/ParseUint. If the value was not set the pointer
// is nil.
type uint16PtrValue struct {
	v **uint16
	b bool
}

func NewUint16PtrValue(p **uint16, v *uint16) *uint16PtrValue {
	*p = v
	return &uint16PtrValue{p, v != nil}
}

func (s *uint16PtrValue) Set(val string) error {
	n, err := strconv.ParseUint(val, 10, 16)
	if err != nil {
		return err
	}
	nn:=uint16(n)
	*s.v, s.b = &nn, true
	return nil
}

func (s *uint16PtrValue) Type() string {
	return "uint16"
}

func (s *uint16PtrValue) String() string {
	if s.b {
		return fmt.Sprint(**s.v)
	}
	return ""
}