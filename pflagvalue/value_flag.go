package pflagvalue

import "github.com/spf13/pflag"

func IntValueFlag(vv interface{}, addr interface{}) pflag.Value {
	switch v := vv.(type) {

	case *int8:
		return NewInt8PtrValue(addr.(**int8), v)

	case *int16:
		return NewInt16PtrValue(addr.(**int16), v)

	case *int32:
		return NewInt32PtrValue(addr.(**int32), v)

	case *uint:
		return NewUintPtrValue(addr.(**uint), v)

	case *uint8:
		return NewUint8PtrValue(addr.(**uint8), v)

	case *uint16:
		return NewUint16PtrValue(addr.(**uint16), v)

	case *uint32:
		return NewUint32PtrValue(addr.(**uint32), v)

	case *uint64:
		return NewUint64PtrValue(addr.(**uint64), v)

	}

	return nil
}
