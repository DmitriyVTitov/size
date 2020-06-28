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
		b: `1234567890123`, // 13 + 16
		c: 20,              // 8
	}

	var v2 = struct {
		a int
		b string
		c t1
	}{
		a: 2,       // 8
		b: "12345", // 5 + 16
		c: t1{
			a: 1,     // 8
			b: "123", // 3 + 16
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
		b: "12345", // 5 + 16 = 21
		c: t1{
			a: 1,     // 8
			b: "123", // 3 + 16 = 19
			// c = 8
		},
		d: [3]int{11, 22, 33}, // 8 * 3 = 24 + 24 = 48
	}

	var v4 = struct {
		a int
		b string
		c t1
		d []int
	}{
		a: 2,       // 8
		b: "12345", // 5 + 16 = 21
		c: t1{
			a: 1,     // 8
			b: "123", // 3 + 16 = 19
			// c = 8
		},
		d: []int{10, 20, 30, 40}, // 8 * 4 = 32 + 24 = 56
	}

	var v5 *t1 = &t1{
		b: "String", // 38
	}

	var v6 = t2{
		a: []int{1, 2, 3}, // 32 + 24 = 56
		b: v5,             // 38
	}

	var v7 = t2{
		a: []int{1, 2, 3}, // 24 + 24 = 48
		// ptr = 8
	}

	var v8 = t4{
		data: []t3{ // 24
			{
				text: "c1", // 2 + 16 + 8 = 26
			},
			{
				text: "c2", // 2 + 16 + 8 = 26
			},
		},
	}
	for i := range v8.data {
		v8.data[i].parent = &v8
	}

	var v9 = make(map[int]string) // 90 + 8 = 98 - size of Map is 8
	v9[0] = "ABC"                 // 8 + 3 + 16 = 27
	v9[1] = "CDEFG"               // 8 + 5 + 16 = 29
	v9[2] = "ABCDEFGHHI"          // 8 + 10 + 16 = 34

	var v10 interface{}
	v10 = 100 // 8

	var v11 interface{}
	v11 = "ABCDEF" // 6 + 16 = 22

	var v12 = make(chan int) // 8 - size of chan in Go

	var tests = []testCase{
		{
			name: "v1",
			v:    v1,
			want: 45,
		},
		{
			name: "v2",
			v:    v2,
			want: 64,
		},
		{
			name: "v3",
			v:    v3,
			want: 112,
		},
		{
			name: "v4",
			v:    v4,
			want: 120,
		},
		{
			name: "v5",
			v:    v5,
			want: 38,
		},
		{
			name: "v6",
			v:    v6,
			want: 94,
		},
		{
			name: "v7",
			v:    v7,
			want: 56,
		},
		{
			name: "v8",
			v:    v8,
			want: 76,
		},
		{
			name: "v9",
			v:    v9,
			want: 98,
		},
		{
			name: "v10",
			v:    v10,
			want: 8,
		},
		{
			name: "v11",
			v:    v11,
			want: 22,
		},
		{
			name: "v12",
			v:    v12,
			want: 8,
		},
	}
	return tests
}
