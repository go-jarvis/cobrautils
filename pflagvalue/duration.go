package pflagvalue

import "time"

// durationPtrValue is a flag.Value which stores the value in a
// *time.Duration if it can be parsed with time.ParseDuration. If the
// value was not set the pointer is nil.
type durationPtrValue struct {
	v **time.Duration
	b bool
}

// github.com/AlekSi/pointer does not have it
func ToDuration(d time.Duration) *time.Duration { return &d }

func NewDurationPtrValue(p **time.Duration, v *time.Duration) *durationPtrValue {
	*p = v
	return &durationPtrValue{p, v != nil}
}

func (s *durationPtrValue) Set(val string) error {
	d, err := time.ParseDuration(val)
	if err != nil {
		return err
	}
	*s.v, s.b = &d, true
	return nil
}

func (s *durationPtrValue) Type() string {
	return "duration"
}

func (s *durationPtrValue) String() string {
	if s.b {
		return (*(*s).v).String()
	}
	return ""
}
