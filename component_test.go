package golevel7

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestCompParse(t *testing.T) {
	val := []rune("v1&v2&v3&&v5")
	seps := NewDelimeters()

	cmp := &Component{Value: val}
	err := cmp.parse(seps)
	require.NoError(t, err)

	assert.Len(t, cmp.SubComponents, 5)
}

func TestCompSet(t *testing.T) {
	seps := NewDelimeters()
	loc := "ZZZ.1.0.5"
	l := NewLocation(loc)

	cmp := &Component{}
	err := cmp.Set(l, "TEST", seps)
	require.NoError(t, err)

	assert.Len(t, cmp.SubComponents, 6)
	assert.Equal(t, []rune("TEST"), cmp.SubComponents[5].Value)
}
