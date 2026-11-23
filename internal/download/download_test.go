package download

import (
	"context"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

type otherKey struct{}

func TestIsMarked(t *testing.T) {
	assert.False(t, IsMarked(context.Background()))

	marked := Mark(context.Background())
	assert.True(t, IsMarked(marked))

	derived, cancel := context.WithTimeout(context.WithValue(marked, otherKey{}, "x"), time.Minute)
	defer cancel()
	assert.True(t, IsMarked(derived), "a context derived from a marked one stays marked")
}
