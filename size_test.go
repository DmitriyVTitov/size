package size

import (
	"reflect"
	"testing"
)

func TestOf(t *testing.T) {

	tests := testCases()

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Of(tt.v); got != tt.want {
				t.Errorf("Of() = %v, want %v", got, tt.want)
			}
		})

		// cleaning recursion detection cache
		values = make(map[reflect.Value]bool)
	}
}
