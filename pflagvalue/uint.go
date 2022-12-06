package pflagvalue

import (
	"fmt"
	"strconv"
)

// uintPtrValue is a flag.Value which stores the value in a *uint if it
// can be parsed with strconv.ParseInt/ParseUint. If the value was not set the pointer
// is nil.
type uintPtrValue struct {
	v **uint
	b bool
}

func NewUintPtrValue(p **uint, v *uint) *uintPtrValue {
	*p = v
	return &uintPtrValue{p, v != nil}
}

func (s *uintPtrValue) Set(val string) error {
	n, err := strconv.ParseUint(val, 10, 0)
	if err != nil {
		return err
	}
	nn:=uint(n)
	*s.v, s.b = &nn, true
	return nil
}

func (s *uintPtrValue) Type() string {
	return "uint"
}

func (s *uintPtrValue) String() string {
	if s.b {
		return fmt.Sprint(**s.v)
	}
	return ""
}