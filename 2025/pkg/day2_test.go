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
	require.Equal(t, 1227775554, invalidSum)
}
