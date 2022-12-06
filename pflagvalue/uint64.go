package pflagvalue

import (
	"fmt"
	"strconv"
)

// uint64PtrValue is a flag.Value which stores the value in a *uint64 if it
// can be parsed with strconv.ParseInt/ParseUint. If the value was not set the pointer
// is nil.
type uint64PtrValue struct {
	v **uint64
	b bool
}

func NewUint64PtrValue(p **uint64, v *uint64) *uint64PtrValue {
	*p = v
	return &uint64PtrValue{p, v != nil}
}

func (s *uint64PtrValue) Set(val string) error {
	n, err := strconv.ParseUint(val, 10, 64)
	if err != nil {
		return err
	}
	nn:=uint64(n)
	*s.v, s.b = &nn, true
	return nil
}

func (s *uint64PtrValue) Type() string {
	return "uint64"
}

func (s *uint64PtrValue) String() string {
	if s.b {
		return fmt.Sprint(**s.v)
	}
	return ""
}