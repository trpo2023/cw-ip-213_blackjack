package util

import (
	"github.com/stretchr/testify/require"
	"math"
	"testing"
)

func TestGetRandomString(t *testing.T) {
	str := GetRandomString()

	require.NotEmpty(t, str)
}

func TestTwoRandomStringIsNotEqual(t *testing.T) {
	str1 := GetRandomString()
	str2 := GetRandomString()

	require.NotEqual(t, str1, str2)
}

func TestGetRandomIntn(t *testing.T) {
	num := GetRandomIntn(10)

	require.NotEmpty(t, num)
}

func TestTwoRandomIntnIsNotEqual(t *testing.T) {
	num1 := GetRandomIntn(math.MaxInt)
	num2 := GetRandomIntn(math.MaxInt)

	require.NotEqual(t, num1, num2)
}
