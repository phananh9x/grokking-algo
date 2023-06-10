package warmup

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_IsAnagram(t *testing.T) {
	testCases := []struct {
		desc   string
		input1 string
		input2 string
		want   bool
	}{
		{
			desc:   "test_case_1",
			input1: "listen",
			input2: "silent",
			want:   true,
		},
		{
			desc:   "test_case_2",
			input1: "rat",
			input2: "car",
			want:   false,
		},
		{
			desc:   "test_case_3",
			input1: "hello",
			input2: "world",
			want:   false,
		},
	}

	for _, v := range testCases {
		t.Run(v.desc, func(t *testing.T) {
			result := isValidAnagram(v.input1, v.input2)
			if !assert.Equal(t, v.want, result) {
				t.Errorf("test fail, got %v, expected %v", result, v.want)
			}
		})
	}
}
