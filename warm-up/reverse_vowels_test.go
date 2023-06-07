package warmup

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_ReverseVowels(t *testing.T) {
	testCases := []struct {
		desc  string
		input string
		want  string
	}{
		{
			desc:  "test_case_1",
			input: "hello",
			want:  "holle",
		},
		{
			desc:  "test_case_2",
			input: "AEIOU",
			want:  "UOIEA",
		},
		{
			desc:  "test_case_3",
			input: "DesignGUrus",
			want:  "DusUgnGires",
		},
	}
	for _, v := range testCases {
		t.Run(v.desc, func(t *testing.T) {
			result := reverseVowels(v.input)
			if !assert.Equal(t, v.want, result) {
				t.Errorf("test fail, got %v, expected %v", result, v.want)
			}
		})

	}
}
