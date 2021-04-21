package helper

import (
	"testing"

	. "github.com/onsi/gomega"
)

func TestDataSize(t *testing.T) {
	type testCase struct {
		desp string

		v    interface{}
		size int
	}
	boolValue := true
	int8Value := int8(1)
	uint8Value := uint8(1)
	int16Value := int16(1)
	uint16Value := uint16(1)
	testCases := []testCase{
		{
			desp: "bool size",
			v:    boolValue,
			size: 1,
		},
		{
			desp: "int8 size",
			v:    int8Value,
			size: 1,
		},
		{
			desp: "uint8 size",
			v:    uint8Value,
			size: 1,
		},
		{
			desp: "*bool size",
			v:    &boolValue,
			size: 1,
		},
		{
			desp: "*int8 size",
			v:    &int8Value,
			size: 1,
		},
		{
			desp: "*uint8 size",
			v:    &uint8Value,
			size: 1,
		},
		{
			desp: "[]bool 4 element size",
			v:    make([]bool, 4),
			size: 4,
		},
		{
			desp: "[]int8 4 element size",
			v:    make([]int8, 4),
			size: 4,
		},
		{
			desp: "[]uint8 4 element size",
			v:    make([]uint8, 4),
			size: 4,
		},
		{
			desp: "int16 size",
			v:    int16Value,
			size: 2,
		},
		{
			desp: "uint16 size",
			v:    uint16Value,
			size: 2,
		},
		{
			desp: "*int16 size",
			v:    &int16Value,
			size: 2,
		},
		{
			desp: "*uint16 size",
			v:    &uint16Value,
			size: 2,
		},
		{
			desp: "[]int16 4 element size",
			v:    make([]int16, 4),
			size: 8,
		},
		{
			desp: "[]uint16 4 element size",
			v:    make([]uint16, 4),
			size: 8,
		},
	}
	for _, tc := range testCases {
		t.Run(tc.desp, func(t *testing.T) {
			g := NewWithT(t)

			g.Expect(DataSize(tc.v)).To(Equal(tc.size))
		})
	}
}
