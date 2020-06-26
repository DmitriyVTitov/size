# size - Go package for calculation variable size at runtime

### Part of the [Transflow Project](http://transflow.ru/)

Sometimes you may need a tool to measure the size of object in your Go program at runtime. This package makes an attempt to do so.


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
