package blackjack

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestConfig_DefaultPlayerNames(t *testing.T) {
	require.NotEmpty(t, DefaultPlayerNames)
}
