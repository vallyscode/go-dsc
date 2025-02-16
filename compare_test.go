package godsc_test

import (
	"fmt"
	"testing"

	godsc "github.com/vallyscode/go-dsc"
)

func TestCompare(t *testing.T) {
	tests := []struct {
		left, right *string
		want        float64
	}{
		{nil, nil, 0},
		{nil, str("foo"), 0},
		{str("foo"), nil, 0},
		{str(""), str(""), 1},
		{str("foo"), str("foo"), 1},
		{str("foo"), str("bar"), 0},
		{str(""), str("foo"), 0},
		{str("foo"), str(""), 0},
		{str("night"), str("nacht"), 0.25},
		{str("healed"), str("sealed"), 0.8},
	}
	for _, test := range tests {
		if got := godsc.Compare(test.left, test.right); got != test.want {
			t.Errorf("Compare(%#v, %#v) = %v, want %v", test.left, test.right, got, test.want)
		}
	}
}

var table = []struct {
	left, right string
	want        float64
}{
	{"", "", 1},
	{"foo", "foo", 1},
	{"foo", "bar", 0},
	{"", "foo", 0},
	{"foo", "", 0},
	{"night", "nacht", 0.25},
	{"healed", "sealed", 0.8},
}

func BenchmarkCompare1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		godsc.Compare(str("night"), str("nacht"))
	}
}

func BenchmarkCompare2(b *testing.B) {
	for _, v := range table {
		b.Run(fmt.Sprintf("input %s %s", v.left, v.right), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				godsc.Compare(str(v.left), str(v.right))
			}
		})
	}
}

func str(s string) *string {
	return &s
}
