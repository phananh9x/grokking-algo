package warmup

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_ValidPalindrome(t *testing.T) {
	testCases := []struct {
		desc  string
		input string
		want  bool
	}{
		{
			desc:  "test_case_1",
			input: "A man, a plan, a canal, Panama!",
			want:  true,
		},
		{
			desc:  "test_case_2",
			input: "Was it a car or a cat I saw?",
			want:  true,
		},
		{
			desc:  "test_case_3",
			input: "Was it a car or a cat I 123",
			want:  false,
		},
	}

	for _, v := range testCases {
		t.Run(v.desc, func(t *testing.T) {
			result := isValidPalindrome(v.input)
			if !assert.Equal(t, v.want, result) {
				t.Errorf("test fail, got %v, expected %v", result, v.want)
			}
		})

	}
}
