package IOModel

import (
	"encoding/json"
	"fmt"
	"math"
)

type IOValue interface {
	Type() TypeID //.每个数据类型都有自己唯一的id
	asBool() (bool, error)
	asInt() (int64, error)
	asFloat() (float64, error)
	asString() (string, error)
	asBlob() ([]byte, error)
	String() string
}

func castError(from TypeID, to TypeID) error {
	return fmt.Errorf("unsupported cast %v from %v", to.String(), from.String())
}

func NewValue(v IOValue) interface{} {
	var result interface{}
	switch v.Type() {
	case TypeBool:
		result, _ = v.asBool()
	case TypeInt:
		result, _ = v.asInt()
	case TypeFloat:
		result, _ = v.asFloat()
	case TypeString:
		result, _ = v.asString()
	case TypeBlob:
		result, _ = v.asBlob()
	case TypeNull:
		result = nil
	default:
		//do nothing
	}
	return result
}

// NewValue returns a Value object from interface{}.
// Returns an error when value type is not supported in SensorBee.
func NewIOValue(v interface{}) (IOValue, error) {
	switch vt := v.(type) {
	case json.Number:
		// json.Number is a string and must be checked before the actual string type.
		if i, err := vt.Int64(); err == nil {
			return IOInt(i), nil
		}
		f, err := vt.Float64()
		if err != nil {
			return nil, err
		}
		return IOFloat(f), nil
	case bool:
		return IOBool(vt), nil
	case int:
		return IOInt(vt), nil
	case int8:
		return IOInt(vt), nil
	case int16:
		return IOInt(vt), nil
	case int32:
		return IOInt(vt), nil
	case int64:
		return IOInt(vt), nil
	case uint:
		if vt > uint(math.MaxUint64) {
			return nil, fmt.Errorf("an int value must be less than %v: %v", math.MaxUint64, vt)
		}
		return IOInt(vt), nil
	case uint8:
		return IOInt(vt), nil
	case uint16:
		return IOInt(vt), nil
	case uint32:
		return IOInt(vt), nil
	case uint64:
		if vt > math.MaxInt64 {
			return nil, fmt.Errorf("an int value must be less than 2^63: %v", vt)
		}
		return IOInt(vt), nil
	case float32:
		return IOFloat(vt), nil
	case float64:
		return IOFloat(vt), nil
	case string:
		return IOString(vt), nil
	case []byte:
		return IOBlob(vt), nil
	case nil:
		return IONull{}, nil
		// support some tuple types for convenience
	case IOValue:
		return vt, nil
	default:
		return nil, fmt.Errorf("unsupported type %T", v)
	}
}
