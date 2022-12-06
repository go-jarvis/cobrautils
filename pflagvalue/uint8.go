package pflagvalue

import (
	"fmt"
	"strconv"
)

// uint8PtrValue is a flag.Value which stores the value in a *uint8 if it
// can be parsed with strconv.ParseInt/ParseUint. If the value was not set the pointer
// is nil.
type uint8PtrValue struct {
	v **uint8
	b bool
}

func NewUint8PtrValue(p **uint8, v *uint8) *uint8PtrValue {
	*p = v
	return &uint8PtrValue{p, v != nil}
}

func (s *uint8PtrValue) Set(val string) error {
	n, err := strconv.ParseUint(val, 10, 8)
	if err != nil {
		return err
	}
	nn:=uint8(n)
	*s.v, s.b = &nn, true
	return nil
}

func (s *uint8PtrValue) Type() string {
	return "uint8"
}

func (s *uint8PtrValue) String() string {
	if s.b {
		return fmt.Sprint(**s.v)
	}
	return ""
}