package uuid

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestGeneror(t *testing.T) {
	generator := new()
	uuidStr := generator.Generate()
	err := generator.Parse(uuidStr)

	require.NoError(t, err)
}
