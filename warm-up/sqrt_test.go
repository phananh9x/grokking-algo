package warmup

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Sqrt(t *testing.T) {
	testCases := []struct {
		desc  string
		input int
		want  int
	}{
		{
			desc:  "test_case_1",
			input: 8,
			want:  2,
		},
		{
			desc:  "test_case_2",
			input: 4,
			want:  2,
		},
		{
			desc:  "test_case_3",
			input: 2,
			want:  1,
		},
		{
			desc:  "test_case_4",
			input: 1,
			want:  1,
		},
		{
			desc:  "test_case_5",
			input: 15,
			want:  3,
		},
	}

	for _, v := range testCases {
		t.Run(v.desc, func(t *testing.T) {
			result := sqrt(v.input)
			if !assert.Equal(t, v.want, result) {
				t.Errorf("test fail, got %v, expected %v", result, v.want)
			}
		})
	}

}

func Test_SqrtBinarySearch(t *testing.T) {
	testCases := []struct {
		desc  string
		input int
		want  int
	}{
		{
			desc:  "test_case_1",
			input: 8,
			want:  2,
		},
		{
			desc:  "test_case_2",
			input: 4,
			want:  2,
		},
		{
			desc:  "test_case_3",
			input: 2,
			want:  1,
		},
		{
			desc:  "test_case_4",
			input: 1,
			want:  1,
		},
		{
			desc:  "test_case_5",
			input: 15,
			want:  3,
		},
	}

	for _, v := range testCases {
		t.Run(v.desc, func(t *testing.T) {
			result := sqrtBinarySearch(v.input)
			if !assert.Equal(t, v.want, result) {
				t.Errorf("test fail, got %v, expected %v", result, v.want)
			}
		})
	}

}
