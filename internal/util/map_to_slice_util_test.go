package util

import (
	"frequency_counter/internal/structure"
	"testing"
)

func TestMapToSlice(t *testing.T) {
	testCases := []struct {
		name     string
		input    map[string]uint64
		expected []structure.Pair
	}{
		{
			name:     "Empty map",
			input:    make(map[string]uint64),
			expected: []structure.Pair{},
		},
		{
			name:  "Single element in map",
			input: map[string]uint64{"ольга": 4},
			expected: []structure.Pair{
				{
					Key:   "ольга",
					Value: 4,
				},
			},
		},
		{
			name:  "Several elements in map",
			input: map[string]uint64{"света": 1, "миша": 3, "ольга": 2},
			expected: []structure.Pair{
				{
					Key:   "света",
					Value: 1,
				},
				{
					Key:   "миша",
					Value: 3,
				},
				{
					Key:   "ольга",
					Value: 2,
				},
			},
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			got := MapToSlice(testCase.input)

			if !compareMapWithSlice(testCase.input, got) {
				t.Errorf("Map to slice converting is not correct. Expected %v, got %v", testCase.expected, got)
			}
		})
	}
}

func compareMapWithSlice(inputMap map[string]uint64, slice []structure.Pair) bool {
	if len(inputMap) != len(slice) {
		return false
	}
	for _, x := range slice {
		if mapValue, ok := inputMap[x.Key]; !ok || x.Value != mapValue {
			return false
		}
	}
	return true
}
