package random

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestNew(t *testing.T) {
	testCases := []struct {
		name   string
		config Config
		check  func(r *Random)
	}{
		{
			name: "With empty config",
			config: Config{
				Chars: "",
			},
			check: func(r *Random) {
				require.NotNil(t, r)
			},
		},
		{
			name: "With CharsDefault config",
			config: Config{
				Chars: CharsDefault,
			},
			check: func(r *Random) {
				require.NotNil(t, r)
			},
		},
		{
			name: "With ArabicNumerals config",
			config: Config{
				Chars: CharsArabicNumerals,
			},
			check: func(r *Random) {
				require.NotNil(t, r)
			},
		},
	}

	for i := range testCases {
		tc := testCases[i]

		t.Run(tc.name, func(t *testing.T) {
			r := New(tc.config)
			tc.check(r)
		})
	}
}

func TestRunes(t *testing.T) {
	testCases := []struct {
		name  string
		size  int
		check func(r []rune, size int)
	}{
		{
			name: "Len 0",
			size: 0,
			check: func(r []rune, size int) {
				require.NotNil(t, r)
				require.Len(t, r, size)
			},
		},
		{
			name: "Len 10",
			size: 10,
			check: func(r []rune, size int) {
				require.NotNil(t, r)
				require.Len(t, r, size)
			},
		},
		{
			name: "Len 10000",
			size: 10000,
			check: func(r []rune, size int) {
				require.NotNil(t, r)
				require.Len(t, r, size)
			},
		},
	}

	for i := range testCases {
		tc := testCases[i]

		t.Run(tc.name, func(t *testing.T) {
			r := Runes(tc.size)
			tc.check(r, tc.size)
		})
	}
}

func TestRandString(t *testing.T) {
	testCases := []struct {
		name  string
		size  int
		check func(s string, size int)
	}{
		{
			name: "Len 0",
			size: 0,
			check: func(s string, size int) {
				require.Len(t, s, size)
			},
		},
		{
			name: "Len 10",
			size: 10,
			check: func(s string, size int) {
				require.Len(t, s, size)
			},
		},
		{
			name: "Len 10000",
			size: 10000,
			check: func(s string, size int) {
				require.Len(t, s, size)
			},
		},
		{
			name: "Len 0",
			size: 0,
			check: func(s string, size int) {
				require.Len(t, s, size)
			},
		},
	}

	for i := range testCases {
		tc := testCases[i]

		t.Run(tc.name, func(t *testing.T) {
			s := RandString(tc.size)
			tc.check(s, tc.size)
		})
	}
}

func TestInt(t *testing.T) {
	testCases := []struct {
		name  string
		check func(n int)
	}{
		{
			name: "Try 1",
			check: func(n int) {
				require.NotEmpty(t, n)
			},
		},
		{
			name: "Try 2",
			check: func(n int) {
				require.NotEmpty(t, n)
			},
		},
		{
			name: "Try 3",
			check: func(n int) {
				require.NotEmpty(t, n)
			},
		},
	}

	for i := range testCases {
		tc := testCases[i]

		t.Run(tc.name, func(t *testing.T) {
			n := Int()
			tc.check(n)
		})
	}
}

func TestRandInt(t *testing.T) {
	testCases := []struct {
		name    string
		min     int
		max     int
		recover func(err any)
		check   func(n, min, max int)
	}{
		{
			name: "Min<Max",
			min:  10,
			max:  1000,
			check: func(n, min, max int) {
				require.NotEmpty(t, n)
				require.LessOrEqual(t, n, max)
				require.GreaterOrEqual(t, n, min)
			},
		},
		{
			name: "Min==Max",
			min:  1000,
			max:  1000,
			recover: func(err any) {
				require.NotEmpty(t, err)
			},
		},
		{
			name: "Min>Max",
			min:  1000,
			max:  10,
			recover: func(err any) {
				require.NotEmpty(t, err)
			},
		},
	}

	for i := range testCases {
		tc := testCases[i]

		t.Run(tc.name, func(t *testing.T) {
			defer func() {
				if err := recover(); err != nil {
					tc.recover(err)
				}
			}()

			n := RandInt(tc.min, tc.max)
			tc.check(n, tc.min, tc.max)
		})
	}
}
