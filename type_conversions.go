package IOModel

import (
	"encoding/base64"
	"fmt"
	"math"
	"strconv"
	"strings"
)

const (
	// MaxConvFloat64 is the largest float64 that can be converted to int64.
	MaxConvFloat64 = float64(math.MaxInt64)

	// MinConvFloat64 is the smallest float64 that can be converted to int64
	MinConvFloat64 = float64(math.MinInt64)

	errNilConversionFormat = "cannot convert nil to %v"
)

// AsBool returns a bool value only when the type of Value is TypeBool,
// otherwise it returns error.
func AsBool(v IOValue) (bool, error) {
	if v == nil {
		return false, fmt.Errorf(errNilConversionFormat, TypeBool)
	}
	return v.asBool()
}

// AsInt returns an integer value only when the type of Value is TypeInt,
// otherwise it returns error.
func AsInt(v IOValue) (int64, error) {
	if v == nil {
		return 0, fmt.Errorf(errNilConversionFormat, TypeInt)
	}
	return v.asInt()
}

// AsFloat returns a float value only when the type of Value is TypeFloat,
// otherwise it returns error.
func AsFloat(v IOValue) (float64, error) {
	if v == nil {
		return 0, fmt.Errorf(errNilConversionFormat, TypeFloat)
	}
	return v.asFloat()
}

// AsString returns a string only when the type of Value is TypeString,
// otherwise it returns error.
func AsString(v IOValue) (string, error) {
	if v == nil {
		return "", fmt.Errorf(errNilConversionFormat, TypeString)
	}
	return v.asString()
}

// AsBlob returns an array of bytes only when the type of Value is TypeBlob,
// otherwise it returns error.
func AsBlob(v IOValue) ([]byte, error) {
	if v == nil {
		return nil, fmt.Errorf(errNilConversionFormat, TypeBlob)
	}
	return v.asBlob()
}

// ToBool converts a given Value to a bool, if possible. The conversion
// rules are similar to those in Python:
//
//  * Null: false
//  * Bool: actual boolean value
//  * Int: true if non-zero
//  * Float: true if non-zero and not NaN
//  * String: true if non-empty
//  * Blob: true if non-empty
//  * Timestamp: true if IsZero() is false
//  * Array: true if non-empty
//  * Map: true if non-empty
func ToBool(v IOValue) (bool, error) {
	defaultValue := false
	switch v.Type() {
	case TypeNull:
		return defaultValue, nil
	case TypeBool:
		return v.asBool()
	case TypeInt:
		val, _ := v.asInt()
		return val != 0, nil
	case TypeFloat:
		val, _ := v.asFloat()
		return val != 0.0 && !math.IsNaN(val), nil
	case TypeString:
		val, _ := v.asString()
		val = strings.TrimSpace(val) // keep this for error reporting
		switch strings.ToLower(val) {
		case "t", "true", "y", "yes", "on", "1":
			return true, nil
		case "f", "false", "n", "no", "off", "0":
			return false, nil
		}
		return false, fmt.Errorf("invalid string as a bool literal: %v", val)
	case TypeBlob:
		val, _ := v.asBlob()
		return len(val) > 0, nil
	default:
		return defaultValue,
			fmt.Errorf("cannot convert %T to bool", v)
	}
}

// ToInt converts a given Value to an int64, if possible. The conversion
// rules are as follows:
//
//  * Null: 0
//  * Bool: 0 if false, 1 if true
//  * Int: actual value
//  * Float: conversion as done by int64(value)
//    (values outside of valid int64 bounds will lead to an error)
//  * String: parsed integer with base 0 as per strconv.ParseInt
//    (values outside of valid int64 bounds will lead to an error)
//  * Blob: (error)
//  * Timestamp: the number of second elapsed since January 1, 1970 UTC.
//  * Array: (error)
//  * Map: (error)
func ToInt(v IOValue) (int64, error) {
	defaultValue := int64(0)
	switch v.Type() {
	case TypeNull:
		return defaultValue, nil
	case TypeBool:
		val, _ := v.asBool()
		if val {
			return 1, nil
		}
		return 0, nil
	case TypeInt:
		return v.asInt()
	case TypeFloat:
		val, _ := v.asFloat()
		if val >= MinConvFloat64 && val <= MaxConvFloat64 {
			return int64(val), nil
		}
		return defaultValue,
			fmt.Errorf("%v is out of bounds for int64 conversion", val)
	case TypeString:
		val, _ := v.asString()
		return strconv.ParseInt(val, 0, 64)
	default:
		return defaultValue,
			fmt.Errorf("cannot convert %T to int64", v)
	}
}

// ToFloat converts a given Value to a float64, if possible. The conversion
// rules are as follows:
//
//  * Null: 0.0
//  * Bool: 0.0 if false, 1.0 if true
//  * Int: conversion as done by float64(value)
//  * Float: actual value
//  * String: parsed float as per strconv.ParseFloat
//    (values outside of valid float64 bounds will lead to an error)
//  * Blob: (error)
//  * Timestamp: the number of seconds (not microseconds!) elapsed since
//    January 1, 1970 UTC, with a decimal part
//  * Array: (error)
//  * Map: (error)
func ToFloat(v IOValue) (float64, error) {
	defaultValue := float64(0)
	switch v.Type() {
	case TypeNull:
		return defaultValue, nil
	case TypeBool:
		val, _ := v.asBool()
		if val {
			return 1.0, nil
		}
		return 0.0, nil
	case TypeInt:
		val, _ := v.asInt()
		return float64(val), nil
	case TypeFloat:
		return v.asFloat()
	case TypeString:
		val, _ := v.asString()
		return strconv.ParseFloat(val, 64)
	default:
		return defaultValue,
			fmt.Errorf("cannot convert %T to float64", v)
	}
}

// ToString converts a given Value to a string. The conversion
// rules are as follows:
//
//  * Null: ""
//  * String: the actual string
//  * Blob: base64-encoded string
//  * Timestamp: ISO 8601 representation, see time.RFC3339
//  * other: Go's "%#v" representation
func ToString(v IOValue) (string, error) {
	switch v.Type() {
	case TypeNull:
		return "", nil
	case TypeString:
		// if we used "%#v", we will get a quoted string; if
		// we used "%v", we will get the result of String()
		// (which is JSON, i.e., also quoted)
		return v.asString()
	case TypeBlob:
		val, _ := v.asBlob()
		return base64.StdEncoding.EncodeToString(val), nil
	default:
		return fmt.Sprintf("%#v", v), nil
	}
}

// ToBlob converts a given Value to []byte, if possible.
// The conversion rules are as follows:
//
//  * Null: nil
//  * String: []byte just copied from string
//  * Blob: actual value
//  * other: (error)
func ToBlob(v IOValue) ([]byte, error) {
	switch v.Type() {
	case TypeNull:
		return nil, nil
	case TypeString:
		val, _ := v.asString()
		return base64.StdEncoding.DecodeString(val)
	case TypeBlob:
		return v.asBlob()
	default:
		return nil, fmt.Errorf("cannot convert %T to Blob", v)
	}
}
