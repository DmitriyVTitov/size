package main

import (
	"fmt"
	"reflect"
	"unsafe"
)

type ptrDict map[uintptr]struct{}

func isFixedSize(kind reflect.Kind) bool {
	return kind != reflect.Map && kind != reflect.Array && kind != reflect.Slice && kind != reflect.Struct &&
		kind != reflect.Ptr && kind != reflect.String
}

// got it from unsafe.SizeOf
func sizeOfStruct(kind reflect.Kind) int64 {
	switch kind {
	case reflect.Invalid:
		return 8 // nil interface{} ?
	case reflect.Bool:
		return 1
	case reflect.Int:
		return 4
	case reflect.Int8:
		return 1
	case reflect.Int16:
		return 2
	case reflect.Int32:
		return 4
	case reflect.Int64:
		return 8
	case reflect.Uint:
		return 4
	case reflect.Uint8:
		return 1
	case reflect.Uint16:
		return 2
	case reflect.Uint32:
		return 4
	case reflect.Uint64:
		return 8
	case reflect.Uintptr:
		return 4
	case reflect.Float32:
		return 4
	case reflect.Float64:
		return 8
	case reflect.Complex64:
		return 8
	case reflect.Complex128:
		return 16
	case reflect.Array:
		return 0
	case reflect.Chan:
		return 4
	case reflect.Func:
		return 4
	case reflect.Interface:
		return 4
	case reflect.Map:
		return 4
	case reflect.Ptr:
		return 4
	case reflect.Slice:
		return 24
	case reflect.String:
		return 16
	case reflect.Struct:
		return 0
	case reflect.UnsafePointer:
		return 4
	}

	return 0
}

func sizeOfValue(val reflect.Value, ptrs ptrDict) int64 {
	var size = sizeOfStruct(val.Kind())

	switch val.Kind() {
	case reflect.Array:
		valType := val.Type().Elem().Kind()
		if isFixedSize(valType) {
			size += int64(val.Len()) * sizeOfStruct(valType)
		} else {
			for i := 0; i < val.Len(); i++ {
				size += sizeOfValue(val.Index(i), ptrs)
			}
		}

	case reflect.Map:
		ptr := val.Pointer()
		if _, ok := ptrs[ptr]; !ok {
			ptrs[ptr] = struct{}{}

			keyType := val.Type().Key().Kind()
			valType := val.Type().Elem().Kind()

			if isFixedSize(keyType) && isFixedSize(valType) {
				l := int64(val.Len())
				size += l * sizeOfStruct(keyType)
				size += l * sizeOfStruct(valType)
			} else {
				iter := val.MapRange()
				for iter.Next() {
					size += sizeOfValue(iter.Key(), ptrs)
					size += sizeOfValue(iter.Value(), ptrs)
				}
			}
		}

	case reflect.Ptr:
		ptr := val.Pointer()
		if _, ok := ptrs[ptr]; !ok {
			ptrs[ptr] = struct{}{}
			size += sizeOfValue(val.Elem(), ptrs)
		}

	case reflect.Slice:
		ptr := val.Pointer()
		if _, ok := ptrs[ptr]; !ok {
			ptrs[ptr] = struct{}{}
			valType := val.Type().Elem().Kind()

			if isFixedSize(valType) {
				size += int64(val.Len()) * sizeOfStruct(valType)
			} else {
				for i := 0; i < val.Len(); i++ {
					size += sizeOfValue(val.Index(i), ptrs)
				}
			}
		}

	case reflect.String:
		size += int64(val.Len())

	case reflect.Struct:
		fields := val.NumField()
		for i := 0; i < fields; i++ {
			size += sizeOfValue(val.Field(i), ptrs)
		}
	case reflect.Interface:
		size += sizeOfValue(val.Elem(), ptrs)
	}

	return size
}

func SizeOf(val interface{}) int64 {
	return sizeOfValue(reflect.ValueOf(val), make(ptrDict))
}

type t1 struct {
	a int
	b string
	c int64
}

type t2 = struct {
	a []int
	b *t1
}

type t4 struct {
	data []t3
}

type t3 struct {
	text   string
	parent *t4
}

func main() {

	type t struct {
		s string
	}

	v := struct {
		arr []t
	}{
		arr: []t{
			{
				s: "a",
			},
		},
	}

	var v8 = t4{
		data: []t3{
			{
				text: "c1",
			},
			{
				text: "c2",
			},
		},
	}
	for i := range v8.data {
		v8.data[i].parent = &v8
	}

	fmt.Println(SizeOf(v8))

	fmt.Println(SizeOf(v))

	st := struct{}{}
	fmt.Println(SizeOf(st))

	s := "a"
	fmt.Println(SizeOf(s))

	sl := []string{}
	fmt.Println(SizeOf(sl))

	type t1000 struct {
		a bool
		b int64
		c int32
	}

	var v1000 t1000

	fmt.Println("v1000 size: ", unsafe.Sizeof(v1000))

	m := make(map[int]int)
	fmt.Println(reflect.TypeOf(m).Size())
}
