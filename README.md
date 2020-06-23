# size
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
	}{
		a: 10,
		b: "Text",
		d: 25,
	}

	fmt.Println(size.Of(a))
}

// Output: 17
```
