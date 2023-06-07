package warmup

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_ContainsDupplicateBruthForce(t *testing.T) {
	testCases := []struct {
		input []int
		desc  string
		want  bool
	}{
		{
			input: []int{1, 2, 3, 4},
			desc:  "test_case 1: false",
			want:  false,
		},
		{
			input: []int{1, 2, 3, 1},
			desc:  "test_case 2: true",
			want:  true,
		},
	}

	for _, test := range testCases {
		t.Run(test.desc, func(t *testing.T) {
			result := containsDupplicateBruthForce(test.input)
			if !assert.Equal(t, test.want, result) {
				t.Errorf("test fail, got %v, expected %v", result, test.want)
			}
		})
	}
}

func Test_ContainsDupplicate(t *testing.T) {
	testCases := []struct {
		input []int
		desc  string
		want  bool
	}{
		{
			input: []int{1, 2, 3, 4},
			desc:  "test_case 1: false",
			want:  false,
		},
		{
			input: []int{1, 2, 3, 1},
			desc:  "test_case 2: true",
			want:  true,
		},
	}

	for _, test := range testCases {
		t.Run(test.desc, func(t *testing.T) {
			result := containsDupplicate(test.input)
			if !assert.Equal(t, test.want, result) {
				t.Errorf("test fail, got %v, expected %v", result, test.want)
			}
		})
	}
}
