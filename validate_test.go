package golevel7

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestValid(t *testing.T) {
	fname := "./testdata/msg.hl7"
	file, err := os.Open(fname)
	require.NoError(t, err)

	defer file.Close()
	msgs, err := NewDecoder(file).Messages()
	require.NoError(t, err)

	valid, failures := msgs[0].IsValid(NewValidMSH24())
	assert.Truef(t, valid, "failures: %+v", failures)

	valid, failures = msgs[0].IsValid(NewValidPID24())
	assert.Truef(t, valid, "failures: %+v", failures)
}
