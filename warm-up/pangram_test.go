package warmup

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Pangram(t *testing.T) {
	testCases := []struct {
		desc  string
		input string
		want  bool
	}{
		{
			input: "TheQuickBrownFoxJumpsOverTheLazyDog",
			desc:  "test_case 1: true",
			want:  true,
		},
		{
			input: "This is not a pangram",
			desc:  "test_case 2: false",
			want:  false,
		},
	}
	for _, v := range testCases {
		t.Run(v.desc, func(t *testing.T) {
			result := checkIfPagram(v.input)
			if !assert.Equal(t, v.want, result) {
				t.Errorf("test fail, got %v, expected %v", result, v.want)
			}
		})
	}

}
