package niler

import (
	"testing"

	"github.com/davecgh/go-spew/spew"
	"github.com/stretchr/testify/assert"
)

func TestIsNil(t *testing.T) {
	type testCase struct {
		val      any
		expected bool
	}
	ch := make(chan int)
	var emptyCh chan int

	var emptyMap map[int]int
	valueMap := map[int]int{}
	interfaceNil := func() interface{} {
		return nil
	}
	interfaceValue := func() interface{} {
		return "test"
	}
	typedInterfaceNil := func() interface{} {
		var val *string
		return val
	}

	var nilSlice []int
	emptySlice := []int{}

	cases := []testCase{
		{0, false},
		{0., false},
		{"", false},
		{"test", false},
		{nil, true},
		{interfaceNil(), true},
		{interfaceValue(), false},
		{typedInterfaceNil(), true},
		{ch, false},
		{emptyCh, true},
		{emptyMap, true},
		{valueMap, false},
		{nilSlice, true},
		{emptySlice, false},
	}

	for _, v := range cases {
		assert.Equal(t, v.expected, IsNil(v.val), "%v", spew.Sdump(v))
	}
}

func BenchmarkIsNil(b *testing.B) {
	b.Run("untyped nil", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			_ = IsNil(nil)
		}
	})
	b.Run("pointer nil", func(b *testing.B) {
		p := new(string)
		for i := 0; i < b.N; i++ {
			_ = IsNil(p)
		}
	})

}
