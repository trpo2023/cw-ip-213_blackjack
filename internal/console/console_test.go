package console

import (
	"bufio"
	"github.com/stretchr/testify/require"
	"os"
	"sync"
	"testing"
	"time"
)

func TestNewConsole(t *testing.T) {
	testCases := []struct {
		name  string
		check func(c *Console, err error)
	}{
		{
			name: "Ok",
			check: func(c *Console, err error) {
				require.NoError(t, err)
				require.NotNil(t, c)
			},
		},
	}

	for i := range testCases {
		tc := testCases[i]

		t.Run(tc.name, func(t *testing.T) {
			c, err := NewConsole()
			tc.check(c, err)
		})
	}
}

func TestConsole_Input(t *testing.T) {
	testCases := []struct {
		name  string
		check func(c *Console)
	}{
		{
			name: "Ok",
			check: func(c *Console) {
				timer := time.AfterFunc(time.Second, func() {
					t.Fatal("timeout")
				})
				defer timer.Stop()

				var wg sync.WaitGroup

				wg.Add(1)
				go func() {
					defer wg.Done()
					c.Input()
				}()

				writer := bufio.NewWriter(os.Stdout)
				_, err := writer.WriteString("a")
				require.NoError(t, err)

				wg.Wait()
			},
		},
		{
			name: "Timeout",
			check: func(c *Console) {
				timer := time.AfterFunc(time.Second, func() {
					t.Log("timeout")
				})
				defer timer.Stop()

				c.Input()
			},
		},
	}

	for i := range testCases {
		tc := testCases[i]

		t.Run(tc.name, func(t *testing.T) {
			c, err := NewConsole()
			require.NoError(t, err)

			tc.check(c)
		})
	}
}
