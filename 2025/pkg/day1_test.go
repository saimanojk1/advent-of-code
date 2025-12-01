package pkg

import (
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestDay1(t *testing.T) {
	inputFilePath := filepath.Join(INPUTFILES_DIR, "day1-test-input.txt")
	password, err := Day1(inputFilePath)
	require.NoError(t, err)
	require.Equal(t, 6, password)
}
