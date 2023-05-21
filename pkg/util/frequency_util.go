package util

import (
	"bufio"
	"os"
	"strings"
)

func CountFrequency(file *os.File) map[string]int {
	result := make(map[string]int)
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		result[strings.ToLower(scanner.Text())] += 1
	}
	return result
}
