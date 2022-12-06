package pflagvalue

import (
	"fmt"
	"strconv"
)

// int8PtrValue is a flag.Value which stores the value in a *int8 if it
// can be parsed with strconv.ParseInt/ParseUint. If the value was not set the pointer
// is nil.
type int8PtrValue struct {
	v **int8
	b bool
}

func NewInt8PtrValue(p **int8, v *int8) *int8PtrValue {
	*p = v
	return &int8PtrValue{p, v != nil}
}

func (s *int8PtrValue) Set(val string) error {
	n, err := strconv.ParseInt(val, 10, 8)
	if err != nil {
		return err
	}
	nn:=int8(n)
	*s.v, s.b = &nn, true
	return nil
}

func (s *int8PtrValue) Type() string {
	return "int8"
}

func (s *int8PtrValue) String() string {
	if s.b {
		return fmt.Sprint(**s.v)
	}
	return ""
}