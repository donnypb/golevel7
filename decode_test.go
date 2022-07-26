package golevel7

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

type Address struct {
	Street string `hl7:"PID.11.1"`
	City   string `hl7:"PID.11.3"`
}

type my7 struct {
	FirstName        string    `hl7:"PID.5.2"`
	LastName         string    `hl7:"PID.5.1"`
	PatientAddresses []Address `hl7:"PID.11"`
	Countries        []string  `hl7:"PID.11.6"`
}

func TestDecode(t *testing.T) {
	fname := "./testdata/msg.hl7"
	file, err := os.Open(fname)
	require.NoError(t, err)
	defer file.Close()

	st := my7{}
	msgs, err := NewDecoder(file).Messages()
	require.NoError(t, err)
	require.Len(t, msgs, 1)

	err = msgs[0].Unmarshal(&st)
	require.NoError(t, err)

	assert.Equal(t, "John", st.FirstName)
	assert.Equal(t, "Jones", st.LastName)

	assert.Len(t, st.PatientAddresses, 2)
	assert.Equal(t, "123 West St.", st.PatientAddresses[0].Street)
	assert.Equal(t, "Denver", st.PatientAddresses[0].City)
	assert.Equal(t, "520 51st Street", st.PatientAddresses[1].Street)
	assert.Equal(t, "Denver", st.PatientAddresses[1].City)

	assert.Len(t, st.Countries, 2)
	assert.Equal(t, "USA", st.Countries[0])
	assert.Equal(t, "USA", st.Countries[1])
}
