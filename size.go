// Package size implements run-time calculation of size of the variable.
// Source code is based on "binary.Size()" function from Go standard library.
package size

import (
	"fmt"
	"reflect"
)

// Of returns the size of 'v' in bytes.
// If there is an error during calculation, Of returns -1.
func Of(v interface{}) int {
	cache := make(map[reflect.Value]bool) // cache with every visited addressable object for recursion detection
	return sizeOf(reflect.Indirect(reflect.ValueOf(v)), cache)
}

// sizeOf returns the number of bytes the actual data represented by v occupies in memory.
// If there is an error, sizeOf returns -1.
func sizeOf(v reflect.Value, cache map[reflect.Value]bool) int {

	// If Value is in cache then it's been already visited - hence it's infinite recursion.
	if v.CanAddr() && cache[v] {
		return 0
	}

	// Every addressable value stored in cache to avoid infinite recursion.
	if v.CanAddr() {
		cache[v] = true
	}

	switch v.Kind() {

	case reflect.Array:
		fallthrough
	case reflect.Slice:
		sum := 0
		for i := 0; i < v.Len(); i++ {
			s := sizeOf(v.Index(i), cache)
			if s < 0 {
				return -1
			}
			sum += s
		}

		return sum

	case reflect.Struct:
		sum := 0
		for i, n := 0, v.NumField(); i < n; i++ {
			s := sizeOf(v.Field(i), cache)
			if s < 0 {
				return -1
			}
			sum += s
		}

		return sum

	case reflect.String:
		str := fmt.Sprintf("%v", v)
		return len(str)

	case reflect.Ptr:
		if v.IsNil() {
			return 0
		}
		s := sizeOf(reflect.Indirect(v), cache)
		if s < 0 {
			return -1
		}
		return s

	case reflect.Bool,
		reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64,
		reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64,
		reflect.Int,
		reflect.Float32, reflect.Float64, reflect.Complex64, reflect.Complex128:
		return int(v.Type().Size())

	// stub for some types for now
	case reflect.Interface:
		fallthrough
	case reflect.Map:
		fallthrough
	case reflect.Chan:
		return 0
	}

	return -1
}
