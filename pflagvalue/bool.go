package pflagvalue

import "strconv"

// boolPtrValue is a flag.Value which stores the value in a *bool if it
// can be parsed with strconv.ParseBool. If the value was not set the
// pointer is nil.
type boolPtrValue struct {
	v **bool
	b bool
}

func NewBoolPtrValue(p **bool, v *bool) *boolPtrValue {
	*p = v
	return &boolPtrValue{p, v != nil}
}

func (s *boolPtrValue) IsBoolFlag() bool { return true }

func (s *boolPtrValue) Set(val string) error {
	b, err := strconv.ParseBool(val)
	if err != nil {
		return err
	}
	*s.v, s.b = &b, true
	return nil
}

func (s *boolPtrValue) Type() string {
	return "bool"
}

func (s *boolPtrValue) String() string {
	if s.b {
		return strconv.FormatBool(**s.v)
	}
	return "false"
}
