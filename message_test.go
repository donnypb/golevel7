package golevel7

import (
	"io/ioutil"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"golang.org/x/net/html/charset"
)

func readFile(fname string) ([]byte, error) {
	file, err := os.Open(fname)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	reader, err := charset.NewReader(file, "text/plain")
	if err != nil {
		return nil, err
	}

	data, err := ioutil.ReadAll(reader)
	if err != nil {
		return nil, err
	}

	return data, nil
}

func TestMessage(t *testing.T) {
	data, err := readFile("./testdata/msg.hl7")
	require.NoError(t, err)

	msg := &Message{Value: []rune(string(data))}
	err = msg.parse()
	require.NoError(t, err)
	assert.Len(t, msg.Segments, 5)

	data, err = readFile("./testdata/msg.lf.hl7")
	require.NoError(t, err)

	msg = &Message{Value: []rune(string(data))}
	err = msg.parse()
	require.NoError(t, err)
	assert.Len(t, msg.Segments, 5)

	data, err = readFile("./testdata/msg.crlf.hl7")
	require.NoError(t, err)

	msg = &Message{Value: []rune(string(data))}
	err = msg.parse()
	require.NoError(t, err)
	assert.Len(t, msg.Segments, 5)

	data, err = readFile("./testdata/msg2.hl7")
	require.NoError(t, err)

	msg = &Message{Value: []rune(string(data))}
	err = msg.parse()
	require.NoError(t, err)
	assert.Len(t, msg.Segments, 5)

	data, err = readFile("./testdata/msg3.hl7")
	require.NoError(t, err)

	msg = &Message{Value: []rune(string(data))}
	err = msg.parse()
	require.NoError(t, err)
	assert.Len(t, msg.Segments, 9)

	data, err = readFile("./testdata/msg4.hl7")
	require.NoError(t, err)

	msg = &Message{Value: []rune(string(data))}
	err = msg.parse()
	require.NoError(t, err)
	assert.Len(t, msg.Segments, 9)
}

func TestMsgUnmarshal(t *testing.T) {
	fname := "./testdata/msg.hl7"
	file, err := os.Open(fname)
	require.NoError(t, err)
	defer file.Close()

	msgs, err := NewDecoder(file).Messages()
	require.NoError(t, err)

	st := my7{}
	err = msgs[0].Unmarshal(&st)
	require.NoError(t, err)

	assert.Equal(t, "John", st.FirstName)
	assert.Equal(t, "Jones", st.LastName)
}
