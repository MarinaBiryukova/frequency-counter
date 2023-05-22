package util

import (
	"os"
	"testing"
)

func TestCountFrequency(t *testing.T) {
	testCases := []struct {
		name     string
		input    string
		expected map[string]uint64
	}{
		{
			name:     "Empty file",
			input:    "",
			expected: make(map[string]uint64),
		},
		{
			name:     "Single name",
			input:    "Дима\nДима",
			expected: map[string]uint64{"дима": 2},
		},
		{
			name:     "Single name, case insensitivity",
			input:    "Дима\nДИМА\nДиМа",
			expected: map[string]uint64{"дима": 3},
		},
		{
			name:     "Several names",
			input:    "Дима\nСвета\nОльга\nМиша\nСвета\nМиша\nСвета",
			expected: map[string]uint64{"дима": 1, "света": 3, "ольга": 1, "миша": 2},
		},
		{
			name:     "Several names, case insensitivity",
			input:    "СвЕтА\nМИША\nСвета\nМиша\nСВЕТА\nОльга\nОЛЬгА",
			expected: map[string]uint64{"света": 3, "миша": 2, "ольга": 2},
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			tempFile := createTempFile(testCase.input, t)
			defer func(name string) {
				err := os.Remove(name)
				if err != nil {
					t.Log(err)
				}
			}(tempFile.Name())
			file, err := os.Open(tempFile.Name())
			if err != nil {
				t.Fatal(err)
			}
			defer func(file *os.File) {
				err := file.Close()
				if err != nil {
					t.Log(err)
				}
			}(file)

			got, err := CountFrequency(file)
			if err != nil {
				t.Fatal(err)
			}

			if !compareMaps(got, testCase.expected) {
				t.Errorf("Frequency count is not correct. Expected %v, got %v", testCase.expected, got)
			}
		})
	}
}

func createTempFile(input string, t *testing.T) *os.File {
	tempFile, err := os.CreateTemp("", "test")
	if err != nil {
		t.Fatal(err)
	}
	if _, err := tempFile.Write([]byte(input)); err != nil {
		t.Fatal(err)
	}
	if err := tempFile.Close(); err != nil {
		t.Log(err)
	}
	return tempFile
}

func compareMaps(map1, map2 map[string]uint64) bool {
	if len(map1) != len(map2) {
		return false
	}
	for key, value := range map1 {
		if value2, ok := map2[key]; !ok || value2 != value {
			return false
		}
	}
	return true
}
