package jsondepth

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestWithinLimit(t *testing.T) {
	nested := func(depth int) []byte {
		return []byte(strings.Repeat("[", depth) + strings.Repeat("]", depth))
	}
	assert.True(t, WithinLimit(nil))
	assert.True(t, WithinLimit([]byte(`{"a":[1,{"b":"c"}]}`)))
	assert.True(t, WithinLimit(nested(MaxDepth)))
	assert.False(t, WithinLimit(nested(MaxDepth+1)))
	assert.False(t, WithinLimit([]byte(strings.Repeat("{\"a\":", MaxDepth+1))))
	assert.True(t, WithinLimit([]byte(`"`+strings.Repeat("[", MaxDepth+1)+`"`)), "brackets inside a string")
	assert.True(t, WithinLimit([]byte(`"\"`+strings.Repeat("[", MaxDepth+1)+`"`)), "escaped quote keeps the string open")
	assert.False(t, WithinLimit([]byte(`"\\"`+strings.Repeat("[", MaxDepth+1))), "escaped backslash closes the string")
}
