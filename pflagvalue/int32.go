package pflagvalue

import (
	"fmt"
	"strconv"
)

// int32PtrValue is a flag.Value which stores the value in a *int32 if it
// can be parsed with strconv.ParseInt/ParseUint. If the value was not set the pointer
// is nil.
type int32PtrValue struct {
	v **int32
	b bool
}

func NewInt32PtrValue(p **int32, v *int32) *int32PtrValue {
	*p = v
	return &int32PtrValue{p, v != nil}
}

func (s *int32PtrValue) Set(val string) error {
	n, err := strconv.ParseInt(val, 10, 32)
	if err != nil {
		return err
	}
	nn:=int32(n)
	*s.v, s.b = &nn, true
	return nil
}

func (s *int32PtrValue) Type() string {
	return "int32"
}

func (s *int32PtrValue) String() string {
	if s.b {
		return fmt.Sprint(**s.v)
	}
	return ""
}