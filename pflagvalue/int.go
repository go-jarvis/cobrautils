package pflagvalue

import "strconv"

// intPtrValue is a flag.Value which stores the value in a *int if it
// can be parsed with strconv.Atoi. If the value was not set the pointer
// is nil.
type intPtrValue struct {
	v **int
	b bool
}

func NewIntPtrValue(p **int, v *int) *intPtrValue {
	*p = v
	return &intPtrValue{p, v != nil}
}

func (s *intPtrValue) Set(val string) error {
	n, err := strconv.Atoi(val)
	if err != nil {
		return err
	}
	*s.v, s.b = &n, true
	return nil
}

func (s *intPtrValue) Type() string {
	return "int"
}

func (s *intPtrValue) String() string {
	if s.b {
		return strconv.Itoa(**s.v)
	}
	return ""
}
