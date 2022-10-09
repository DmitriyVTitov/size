package size

import (
	"testing"

	"github.com/golang/groupcache/lru"
)

func TestOf(t *testing.T) {
	ss := make([][]string, 100, 100) // 100 * 24 + 24
	s := make([]string, 1)           // 24
	s[0] = "1234"                    // 16 + 4
	for i := range ss {
		ss[i] = s
	} // 24 + 16 + 4 + 99 * 24 + 24 = 2444
	tests := []struct {
		name string
		v    interface{}
		want int
	}{
		{
			name: "Array",
			v:    [3]int32{1, 2, 3}, // 3 * 4  = 12
			want: 12,
		},
		{
			name: "Slice",
			v:    make([]int64, 2, 5), // 5 * 8 + 24 = 64
			want: 64,
		},
		{
			name: "String",
			v:    "ABCdef", // 6 + 16 = 22
			want: 22,
		},
		{
			name: "Map",
			// (8 + 3 + 16) + (8 + 4 + 16) = 55
			// 55 + 8 + 10.79 * 2 = 84
			v:    map[int64]string{0: "ABC", 1: "DEFG"},
			want: 84,
		},
		{
			name: "Struct",
			v: struct {
				slice     []int64
				array     [2]bool
				structure struct {
					i int8
					s string
				}
			}{
				slice: []int64{12345, 67890}, // 2 * 8 + 24 = 40
				array: [2]bool{true, false},  // 2 * 1 = 2
				structure: struct {
					i int8
					s string
				}{
					i: 5,     // 1
					s: "abc", // 3 * 1 + 16 = 19
				}, // 20 + 7 (padding) = 27
			}, // 40 + 2 + 27 = 69 + 6 (padding) = 75
			want: 75,
		},
		{
			name: "Struct With Func",
			v: lru.Cache{
				MaxEntries: 0,   // 8
				OnEvicted:  nil, // 0
			}, // + 16 (two more pointers) = 24
			want: 24,
		},
		{
			name: "Slice of strings slices (slice of cloned slices)",
			v:    ss,
			want: 2444,
		},
		{
			name: "Struct with the same slice value in two fields",
			v: struct {
				s1 []string // 24
				s2 []string // 24
			}{
				s1: s, // + 16 + 4
				s2: s, // + 0 (the same)
			}, // 2 * 24 + 16 + 4 = 68
			want: 68,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Of(tt.v); got != tt.want {
				t.Errorf("Of() = %v, want %v", got, tt.want)
			}
		})
	}
}
