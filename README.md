# size - calculates variable's memory consumption at runtime

### Part of the [Transflow Project](http://transflow.ru/)

Sometimes you may need a tool to measure the size of object in your Go program at runtime. This package makes an attempt to do so. Package based on `binary.Size()` from Go standard library.

Features:
- supports non-fixed size variables and struct fields: `struct`, `int`, `slice`, `string`, `map`;
- supports complex types including structs with non-fixed size fields;
- supports all basic types (numbers, bool);
- supports `chan` and `interface`;
- supports pointers;
- implements infinite recursion detection (i.e. pointer inside struct field references to parent struct).

### Usage example

```
package main

import (
	"fmt"

	"github.com/DmitriyVTitov/size"
)

func main() {
	a := struct {
		a int
		b string
		c bool
		d int32
		e []byte
		f [3]int64
	}{
		a: 10,                    // 8 bytes
		b: "Text",                // 4 bytes
		c: true,                  // 1 byte
		d: 25,                    // 4 bytes
		e: []byte{'c', 'd', 'e'}, // 3 bytes
		f: [3]int64{1, 2, 3},     // 24 bytes
	}

	fmt.Println(size.Of(a))
}

// Output: 44
```
