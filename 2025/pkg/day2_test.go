package pkg

import (
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestDay2(t *testing.T) {
	inputFilePath := filepath.Join(INPUTFILES_DIR, "day2-test-input.txt")
	invalidSum, err := Day2(inputFilePath)
	require.NoError(t, err)
	require.Equal(t, 4174379265, invalidSum)
}

func TestGetRepeatedString(t *testing.T) {
	repeatedString, err := GetRepeatedValue("1", 5)
	require.NoError(t, err)
	require.Equal(t, 11111, repeatedString)
}

func TestAddRepeatedValuesInRange(t *testing.T) {
	var repeatedValues []int
	AddRepeatedValuesInRange("1", 5, 12, &repeatedValues)
	require.Contains(t, repeatedValues, 11)
	repeatedValues = repeatedValues[:0]
	AddRepeatedValuesInRange("82482", 824824821, 824824827, &repeatedValues)
	require.Contains(t, repeatedValues, 824824824)
}

func TestAddRepeatedValuesInRange2(t *testing.T) {
	var repeatedValues []int
	AddRepeatedValuesInRange("10", 998, 1012, &repeatedValues)
	require.Contains(t, repeatedValues, 999)
	require.Contains(t, repeatedValues, 1010)
}
