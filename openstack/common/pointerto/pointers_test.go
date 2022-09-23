package pointerto

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestBool(t *testing.T) {
	t.Parallel()

	for _, value := range []bool{true, false} {
		require.EqualValues(t, value, *Bool(value))
	}
}

func TestString(t *testing.T) {
	t.Parallel()

	for _, value := range []string{"", "info"} {
		require.EqualValues(t, value, *String(value))
	}
}

func TestInt(t *testing.T) {
	t.Parallel()

	for _, value := range []int{0, 0xff} {
		require.EqualValues(t, value, *Int(value))
	}
}
