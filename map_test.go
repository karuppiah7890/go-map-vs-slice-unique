package mapvsslice

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUniqueWithMap(t *testing.T) {
	for _, testCase := range uniqueTestCases {
		output := UniqueWithMap(testCase.input)
		ok := assert.ElementsMatch(t, output, testCase.expectedOutput)

		if !ok {
			fmt.Println("Expected: ", testCase.expectedOutput, ". Got: ", output)
		}
	}
}

func BenchmarkUniqueWithMap(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, testCase := range uniqueTestCases {
			UniqueWithMap(testCase.input)
		}
	}
}

func TestUniqueWithSlice(t *testing.T) {
	for _, testCase := range uniqueTestCases {
		output := UniqueWithSlice(testCase.input)
		ok := assert.ElementsMatch(t, output, testCase.expectedOutput)

		if !ok {
			fmt.Println("Expected: ", testCase.expectedOutput, ". Got: ", output)
		}
	}
}

func BenchmarkUniqueWithSlice(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, testCase := range uniqueTestCases {
			UniqueWithSlice(testCase.input)
		}
	}
}
