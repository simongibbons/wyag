package internal_test

import (
	"github.com/simongibbons/wyag/internal"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGreet(t *testing.T) {
	assert.Equal(t, "Hello, World", internal.Greet("World"))
}
