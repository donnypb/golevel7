package golevel7

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestFieldParse(t *testing.T) {
	val := []rune("520 51st Street^^Denver^CO^80020^USA")
	seps := NewDelimeters()
	fld := &Field{Value: val}

	err := fld.parse(seps)
	require.NoError(t, err)

	assert.Len(t, fld.Components, 6)
}

func TestFieldSet(t *testing.T) {
	seps := NewDelimeters()
	fld := &Field{}
	loc := "ZZZ.1.10"
	l := NewLocation(loc)

	err := fld.Set(l, "TEST", seps)
	require.NoError(t, err)

	require.Len(t, fld.Components, 11)
	assert.Equal(t, "TEST", string(fld.Components[10].SubComponents[0].Value))
}
