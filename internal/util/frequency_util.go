package util

import (
	"bufio"
	"os"
	"strings"
)

func CountFrequency(file *os.File) (map[string]uint64, error) {
	result := make(map[string]uint64)
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		result[strings.ToLower(scanner.Text())] += 1
	}
	if err := scanner.Err(); err != nil {
		return nil, err
	}
	return result, nil
}
