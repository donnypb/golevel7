package golevel7

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestFind(t *testing.T) {
	data, err := readFile("./testdata/msg3.hl7")
	require.NoError(t, err)

	msg := &Message{Value: []rune(string(data))}
	err = msg.parse()
	require.NoError(t, err)

	val, err := msg.Find("OBR.4.3")
	assert.NoError(t, err)
	assert.Equal(t, "CPT-4", val)

	val, err = msg.Find("OBX.3.3")
	assert.NoError(t, err)
	assert.Equal(t, "LOINC", val)
}

func TestFindAll(t *testing.T) {
	data, err := readFile("./testdata/msg3.hl7")
	require.NoError(t, err)
	msg := &Message{Value: []rune(string(data))}
	msg.parse()
	require.NoError(t, err)

	vals, err := msg.FindAll("OBX.1")
	assert.NoError(t, err)

	assert.Len(t, vals, 4)
	assert.Equal(t, "1", vals[0])
	assert.Equal(t, "2", vals[1])
	assert.Equal(t, "3", vals[2])
	assert.Equal(t, "4", vals[3])
}

func TestRepFields(t *testing.T) {
	data, err := readFile("./testdata/msg.hl7")
	require.NoError(t, err)
	msg := &Message{Value: []rune(string(data))}
	msg.parse()
	require.NoError(t, err)

	vals, err := msg.FindAll("PID.11.3")
	assert.NoError(t, err)
	assert.Len(t, vals, 2)
}
