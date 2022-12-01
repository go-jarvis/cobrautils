package pflagvalue

import "strconv"

// int64PtrValue is a flag.Value which stores the value in a *int64 if it
// can be parsed with strconv.Atoi. If the value was not set the point64er
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
	n, err := strconv.Atoi(val)
	if err != nil {
		return err
	}
	*s.v, s.b = toInt64Ptr(int64(n)), true
	return nil
}

func (s *int64PtrValue) Type() string {
	return "int64"
}

func (s *int64PtrValue) String() string {
	if s.b {
		return strconv.Itoa(int(**s.v))
	}
	return ""
}

func toInt64Ptr(n int64) *int64 {
	return &n
}
