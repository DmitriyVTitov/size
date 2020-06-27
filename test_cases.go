package size

type testCase struct {
	name string
	v    interface{}
	want int
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

func testCases() []testCase {

	var v1 = t1{
		a: 10,              // 8
		b: `1234567890123`, // 13
		c: 20,              // 8
	}

	var v2 = struct {
		a int
		b string
		c t1
	}{
		a: 2,       // 8
		b: "12345", // 5
		c: t1{
			a: 1,     // 8
			b: "123", // 3
			// c = 8
		},
	}

	var v3 = struct {
		a int
		b string
		c t1
		d [3]int
	}{
		a: 2,       // 8
		b: "12345", // 5
		c: t1{
			a: 1,     // 8
			b: "123", // 3
			// c = 8
		},
		d: [3]int{11, 22, 33}, // 8 * 3 = 24
	}

	var v4 = struct {
		a int
		b string
		c t1
		d []int
	}{
		a: 2,       // 8
		b: "12345", // 5
		c: t1{
			a: 1,     // 8
			b: "123", // 3
			// c = 8
		},
		d: []int{10, 20, 30, 40}, // 8 * 4 = 32
	}

	var v5 *t1 = &t1{
		b: "String",
	}

	var v6 = t2{
		a: []int{1, 2, 3},
		b: v5,
	}

	var v7 = t2{
		a: []int{1, 2, 3},
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

	var v9 = make(map[int]string) // 42
	v9[0] = "ABC"                 // 8 + 3
	v9[1] = "CDEFG"               // 8 + 5
	v9[2] = "ABCDEFGHHI"          // 8 + 10

	var v10 interface{}
	v10 = 100

	var v11 interface{}
	v11 = "ABCDEF"

	var v12 = make(chan int) // 8 - size of chan in Go

	var tests = []testCase{
		{
			name: "v1",
			v:    v1,
			want: 29,
		},
		{
			name: "v2",
			v:    v2,
			want: 32,
		},
		{
			name: "v3",
			v:    v3,
			want: 56,
		},
		{
			name: "v4",
			v:    v4,
			want: 64,
		},
		{
			name: "v5",
			v:    v5,
			want: 22,
		},
		{
			name: "v6",
			v:    v6,
			want: 46,
		},
		{
			name: "v7",
			v:    v7,
			want: 24,
		},
		{
			name: "v8",
			v:    v8,
			want: 4,
		},
		{
			name: "v9",
			v:    v9,
			want: 42,
		},
		{
			name: "v10",
			v:    v10,
			want: 8,
		},
		{
			name: "v11",
			v:    v11,
			want: 6,
		},
		{
			name: "v12",
			v:    v12,
			want: 8,
		},
	}
	return tests
}
