package hellogo_test

import (
	"testing"

	"github.com/golang-learning/hellogo"
	"github.com/stretchr/testify/require"
)

func TestHello(t *testing.T) {
	t.Parallel()

	require.Equal(t, hellogo.Hello(), "Hello, GoLang")
}
