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

type MessageType struct {
	MessageCode      string `hl7:"MSH.9.1"`
	TriggerEvent     string `hl7:"MSH.9.2"`
	MessageStructure string `hl7:"MSH.9.3"`
}

type PID struct {
	PatientIdentifierList string   `hl7:"PID.3"`
	DoB                   string   `hl7:"PID.7"`
	Sex                   string   `hl7:"PID.8"`
	Addresses             []string `hl7:"PID.11"`
}

type my7 struct {
	AcceptAcknowledgmentType string    `hl7:"MSH.15"`
	FirstName                string    `hl7:"PID.5.2"`
	LastName                 string    `hl7:"PID.5.1"`
	PatientAddresses         []Address `hl7:"PID.11"`
	Countries                []string  `hl7:"PID.11.6"`
	PID                      PID       `hl7:"PID"`
	MessageType              `hl7:"MSH.9"`
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

	assert.Equal(t, "", st.AcceptAcknowledgmentType)
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

	assert.Equal(t, "12001", st.PID.PatientIdentifierList)
	assert.Equal(t, "19670824", st.PID.DoB)
	assert.Equal(t, "M", st.PID.Sex)
	assert.Len(t, st.PID.Addresses, 2)

	assert.Equal(t, "ORM", st.MessageType.MessageCode)
	assert.Equal(t, "O01", st.MessageType.TriggerEvent)
	assert.Equal(t, "", st.MessageType.MessageStructure)
}

type ORC struct {
	OrderControl      string `hl7:"ORC.1"`
	PlacerOrderNumber string `hl7:"ORC.2"`
	OrderStatus       string `hl7:"ORC.5"`
}

type OBR struct {
	SetID              string `hl7:"OBR.1"`
	PlacerOrderNumber  string `hl7:"OBR.2"`
	UniversalServiceID string `hl7:"OBR.4.1"`
}

type OBX struct {
	SetID            string `hl7:"OBX.1"`
	Type             string `hl7:"OBX.2"`
	ObservationID    string `hl7:"OBX.3.1"`
	ObeservationText string `hl7:"OBX.3.2"`
	ObservationValue string `hl7:"OBX.5.1"`
}

type OMLMsg struct {
	ORCs []ORC `hl7:"ORC"`
	OBRs []OBR `hl7:"OBR"`
	OBXs []OBX `hl7:"OBX"`
}

func TestDecodeOML(t *testing.T) {
	fname := "./testdata/msg_oml_O21.hl7"
	file, err := os.Open(fname)
	require.NoError(t, err)
	defer file.Close()

	st := OMLMsg{}
	msgs, err := NewDecoder(file).Messages()
	require.NoError(t, err)
	require.Len(t, msgs, 1)

	err = msgs[0].Unmarshal(&st)
	require.NoError(t, err)

	assert.Len(t, st.ORCs, 2)
	assert.Equal(t, ORC{OrderControl: "NW", PlacerOrderNumber: "2021P0067674", OrderStatus: ""}, st.ORCs[0])
	assert.Equal(t, ORC{OrderControl: "NW", PlacerOrderNumber: "2021P0067674", OrderStatus: ""}, st.ORCs[1])

	assert.Len(t, st.OBRs, 2)
	assert.Equal(t, OBR{SetID: "1", PlacerOrderNumber: "41510399", UniversalServiceID: "A000172"}, st.OBRs[0])
	assert.Equal(t, OBR{SetID: "1", PlacerOrderNumber: "41510399", UniversalServiceID: "A001312"}, st.OBRs[1])

	assert.Len(t, st.OBXs, 2)
	assert.Equal(t, OBX{SetID: "1", Type: "ST", ObservationID: "A000172", ObeservationText: "", ObservationValue: "34"}, st.OBXs[0])
	assert.Equal(t, OBX{SetID: "1", Type: "ST", ObservationID: "A001312", ObeservationText: "", ObservationValue: "238"}, st.OBXs[1])
}

func TestFailedDecode(t *testing.T) {
	fname := "./testdata/msg.hl7"
	file, err := os.Open(fname)
	require.NoError(t, err)
	defer file.Close()

	msgs, err := NewDecoder(file).Messages()
	require.NoError(t, err)
	require.Len(t, msgs, 1)

	err = msgs[0].Unmarshal(nil)
	assert.EqualError(t, err, "hl7: cannot unmarshal to non-pointer <nil>")

	var st *my7
	err = msgs[0].Unmarshal(st)
	assert.EqualError(t, err, "hl7: cannot unmarshal to nil value of \"*golevel7.my7\"")
}
