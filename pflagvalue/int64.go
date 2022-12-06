package pflagvalue

import (
	"fmt"
	"strconv"
)

// int64PtrValue is a flag.Value which stores the value in a *int64 if it
// can be parsed with strconv.ParseInt/ParseUint. If the value was not set the pointer
// is nil.
type int64PtrValue struct {
	v **int64
	b bool
}

func NewInt64PtrValue(p **int64, v *int64) *int64PtrValue {
	*p = v
	return &int64PtrValue{p, v != nil}
}

func (s *int64PtrValue) Set(val string) error {
	n, err := strconv.ParseInt(val, 10, 64)
	if err != nil {
		return err
	}
	nn:=int64(n)
	*s.v, s.b = &nn, true
	return nil
}

func (s *int64PtrValue) Type() string {
	return "int64"
}

func (s *int64PtrValue) String() string {
	if s.b {
		return fmt.Sprint(**s.v)
	}
	return ""
}