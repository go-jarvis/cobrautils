package pflagvalue

import (
	"fmt"
	"strconv"
)

// uint32PtrValue is a flag.Value which stores the value in a *uint32 if it
// can be parsed with strconv.ParseInt/ParseUint. If the value was not set the pointer
// is nil.
type uint32PtrValue struct {
	v **uint32
	b bool
}

func NewUint32PtrValue(p **uint32, v *uint32) *uint32PtrValue {
	*p = v
	return &uint32PtrValue{p, v != nil}
}

func (s *uint32PtrValue) Set(val string) error {
	n, err := strconv.ParseUint(val, 10, 32)
	if err != nil {
		return err
	}
	nn:=uint32(n)
	*s.v, s.b = &nn, true
	return nil
}

func (s *uint32PtrValue) Type() string {
	return "uint32"
}

func (s *uint32PtrValue) String() string {
	if s.b {
		return fmt.Sprint(**s.v)
	}
	return ""
}