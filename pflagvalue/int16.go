package pflagvalue

import (
	"fmt"
	"strconv"
)

// int16PtrValue is a flag.Value which stores the value in a *int16 if it
// can be parsed with strconv.ParseInt/ParseUint. If the value was not set the pointer
// is nil.
type int16PtrValue struct {
	v **int16
	b bool
}

func NewInt16PtrValue(p **int16, v *int16) *int16PtrValue {
	*p = v
	return &int16PtrValue{p, v != nil}
}

func (s *int16PtrValue) Set(val string) error {
	n, err := strconv.ParseInt(val, 10, 16)
	if err != nil {
		return err
	}
	nn:=int16(n)
	*s.v, s.b = &nn, true
	return nil
}

func (s *int16PtrValue) Type() string {
	return "int16"
}

func (s *int16PtrValue) String() string {
	if s.b {
		return fmt.Sprint(**s.v)
	}
	return ""
}