package golevel7

import (
	"errors"
	"fmt"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestAcknowledge(t *testing.T) {
	fname := "./testdata/msg.hl7"
	file, err := os.Open(fname)
	require.NoError(t, err)
	defer file.Close()

	msgs, err := NewDecoder(file).Messages()
	require.NoError(t, err)

	mi, err := msgs[0].Info()
	require.NoError(t, err)

	ack := Acknowledge(mi, nil)
	assert.NotNil(t, ack)

	for _, s := range ack.Segments {
		for _, f := range s.Fields {
			fmt.Println(string(f.Value))
		}
	}

	ack = Acknowledge(mi, errors.New("This is a test error"))
	assert.NotNil(t, ack)

	m := NewMsgInfo()
	m.ReceivingApp = "ORG_REC_APP"
	m.ReceivingFacility = "ORG_REC_FAC"
	m.SendingApp = "ORG_SEND_APP"
	m.SendingFacility = "ORG_SEND_FAC"
	ack = Acknowledge(*m, errors.New("Fatal error"))
	assert.NotNil(t, ack)
}
