package util

import (
	"frequency_counter/internal/structure"
)

func MapToSlice(inputMap map[string]uint64) []structure.Pair {
	result := make([]structure.Pair, 0, len(inputMap))
	for key, value := range inputMap {
		result = append(result, structure.Pair{
			Key:   key,
			Value: value,
		})
	}
	return result
}
