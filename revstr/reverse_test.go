package revstr

import "testing"

type Cases struct {
	input  string
	output string
}

func TestReverseRunes(t *testing.T) {
	testCases := []Cases{
		{input: "hello world", output: "dlrow olleh"},
		{input: "hi", output: "sih"},
	}

	for _, testCase := range testCases {
		isReversed := ReverseRunes(testCase.input)
		if isReversed != testCase.output {
			t.Errorf("ReverseRunes(%s) = %s; want %s", testCase.input, isReversed, testCase.output)
		}
	}
}
