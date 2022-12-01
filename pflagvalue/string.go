package pflagvalue

// stringPtrValue is a flag.Value which stores the value in a *string.
// If the value was not set the pointer is nil.
type stringPtrValue struct {
	v **string
	b bool
}

func NewStringPtrValue(p **string, v *string) *stringPtrValue {
	*p = v
	return &stringPtrValue{p, v != nil}
}

func (s *stringPtrValue) Set(val string) error {
	*s.v, s.b = &val, true
	return nil
}

func (s *stringPtrValue) Type() string {
	return "string"
}

func (s *stringPtrValue) String() string {
	if s.b {
		return **s.v
	}
	return ""
}
