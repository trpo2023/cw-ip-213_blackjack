package console

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func newConsole(t *testing.T) *Console {
	cnsl, err := NewConsole()

	require.NoError(t, err)
	require.Empty(t, cnsl)

	return cnsl
}

func TestNewConsole(t *testing.T) {
	newConsole(t)
}
