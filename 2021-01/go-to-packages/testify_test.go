package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestWithTestify(t *testing.T) {
	assert.True(t, false, "an extra message for context")
	require.False(t, true)
	assert.Falsef(t, true, "this will run %d times", 0)
}
