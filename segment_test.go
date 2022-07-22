package golevel7

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestSegParse(t *testing.T) {
	val := []rune(`PID|||12001||Jones^John^^^Mr.||19670824|M|||123 West St.^^Denver^CO^80020^USA~520 51st Street^^Denver^CO^80020^USA|||||||
`)
	seps := NewDelimeters()
	seg := &Segment{Value: val}
	err := seg.parse(seps)
	require.NoError(t, err)
	assert.Len(t, seg.Fields, 20)
}

func TestSegSet(t *testing.T) {
	seps := NewDelimeters()
	loc := "ZZZ.10"
	l := NewLocation(loc)
	seg := &Segment{}

	err := seg.Set(l, "TEST", seps)
	require.NoError(t, err)

	str, err := seg.Get(l)
	assert.NoError(t, err)
	assert.Equal(t, "TEST", str)
}
