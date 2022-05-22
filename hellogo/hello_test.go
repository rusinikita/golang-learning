package hellogo_test

import (
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/golang-learning/hellogo"
)

func TestHello(t *testing.T) {
	t.Parallel()

	require.Equal(t, hellogo.Hello(), "Hello, GoLang")
}
