package IOModel

import (
	"fmt"
	"math"
)

// IOFloat is a 64-bit floating point number. It can be assigned to Value interface.
type IOFloat float64

// Type returns TypeID of IOFloat. It's always TypeFloat.
func (f IOFloat) Type() TypeID {
	return TypeFloat
}

func (f IOFloat) asBool() (bool, error) {
	return false, castError(f.Type(), TypeBool)
}

func (f IOFloat) asInt() (int64, error) {
	return 0, castError(f.Type(), TypeInt)
}

func (f IOFloat) asFloat() (float64, error) {
	return float64(f), nil
}

func (f IOFloat) asString() (string, error) {
	return "", castError(f.Type(), TypeString)
}

func (f IOFloat) asBlob() ([]byte, error) {
	return nil, castError(f.Type(), TypeFloat)
}

// MarshalJSON marshals a IOFloat to JSON. NaN and Inf will be encoded as null.
func (f IOFloat) MarshalJSON() ([]byte, error) {
	// the JSON serialization is defined via the IOString()
	// return value as defined below
	return []byte(f.String()), nil
}

// IOString returns JSON representation of a IOFloat. NaN and Inf will be encoded as null.
func (f IOFloat) String() string {
	fl := float64(f)
	// "NaN and Infinity regardless of sign are represented
	// as the IOString null." (ECMA-262)
	// (The default JSON serializer will return an error instead,
	// cf. <https://github.com/golang/go/issues/3480>)
	if math.IsNaN(fl) {
		return "null"
	} else if math.IsInf(fl, 0) {
		return "null"
	}
	return fmt.Sprintf("%#v", f)
}
