package search

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSearchMultipleKeysFromStringDataset(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		name          string
		keys          []string
		dataset       []string
		caseSensitive bool
		expect        []ResultFromStringDataset
	}{
		{
			name:          "found one for each key (case insensitive)",
			keys:          []string{"x", "y", "z"},
			dataset:       []string{"1", "X", "8", "17", "Y", "Z", "78", "113"},
			caseSensitive: false,
			expect: []ResultFromStringDataset{
				{
					Key:       "x",
					Positions: []int{1},
				},
				{
					Key:       "y",
					Positions: []int{4},
				},
				{
					Key:       "z",
					Positions: []int{5},
				},
			},
		},
		{
			name:          "found more than one for each key (case insensitive)",
			keys:          []string{"x", "y", "z"},
			dataset:       []string{"1", "X", "8", "17", "Y", "Z", "78", "113", "X", "Z", "Z", "Y", "y", "y"},
			caseSensitive: false,
			expect: []ResultFromStringDataset{
				{
					Key:       "x",
					Positions: []int{1, 8},
				},
				{
					Key:       "y",
					Positions: []int{4, 11, 12, 13},
				},
				{
					Key:       "z",
					Positions: []int{5, 9, 10},
				},
			},
		},
		{
			name:          "not found (case insensitive)",
			keys:          []string{"x", "y", "z"},
			dataset:       []string{"1", "8", "17", "78", "113"},
			caseSensitive: false,
			expect: []ResultFromStringDataset{
				{
					Key:       "x",
					Positions: nil,
				},
				{
					Key:       "y",
					Positions: nil,
				},
				{
					Key:       "z",
					Positions: nil,
				},
			},
		},
		{
			name:          "found one for each key (case sensitive)",
			keys:          []string{"x", "y", "z"},
			dataset:       []string{"1", "x", "8", "17", "y", "z", "78", "113"},
			caseSensitive: true,
			expect: []ResultFromStringDataset{
				{
					Key:       "x",
					Positions: []int{1},
				},
				{
					Key:       "y",
					Positions: []int{4},
				},
				{
					Key:       "z",
					Positions: []int{5},
				},
			},
		},
		{
			name:          "found more than one for each key (case sensitive)",
			keys:          []string{"x", "y", "z"},
			dataset:       []string{"1", "X", "8", "17", "Y", "Z", "78", "113", "x", "z", "Z", "Y", "y", "y"},
			caseSensitive: true,
			expect: []ResultFromStringDataset{
				{
					Key:       "x",
					Positions: []int{8},
				},
				{
					Key:       "y",
					Positions: []int{12, 13},
				},
				{
					Key:       "z",
					Positions: []int{9},
				},
			},
		},
		{
			name:          "not found (case sensitive)",
			keys:          []string{"x", "y", "z"},
			dataset:       []string{"1", "8", "17", "78", "113", "X", "Y", "Z"},
			caseSensitive: true,
			expect: []ResultFromStringDataset{
				{
					Key:       "x",
					Positions: nil,
				},
				{
					Key:       "y",
					Positions: nil,
				},
				{
					Key:       "z",
					Positions: nil,
				},
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			s := New()

			r := s.SearchMultipleKeysFromStringDataset(tc.keys, tc.dataset, tc.caseSensitive)

			assert.EqualValues(t, tc.expect, r)
		})
	}

}
