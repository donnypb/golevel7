package golevel7

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestBuildMessage(t *testing.T) {
	mi := MsgInfo{
		SendingApp:        "BettrLife",
		SendingFacility:   "UnivIa",
		ReceivingApp:      "Epic",
		ReceivingFacility: "UnivIa",
		MessageType:       "ORM^001",
	}
	msg, err := StartMessage(mi)
	require.NoError(t, err)

	v, err := msg.Find("MSH.3")
	assert.NoError(t, err)
	assert.Equal(t, "BettrLife", v)

	v, err = msg.Find("MSH.4")
	assert.NoError(t, err)
	assert.Equal(t, "UnivIa", v)

	v, err = msg.Find("MSH.5")
	assert.NoError(t, err)
	assert.Equal(t, "Epic", v)

	v, err = msg.Find("MSH.6")
	assert.NoError(t, err)
	assert.Equal(t, "UnivIa", v)

	v, err = msg.Find("MSH.9")
	assert.NoError(t, err)
	assert.Equal(t, "ORM^001", v)

	v, err = msg.Find("MSH.7")
	assert.NoError(t, err)
	assert.NotEmpty(t, v)

	v, err = msg.Find("MSH.10")
	assert.NoError(t, err)
	assert.NotEmpty(t, v)

	v, err = msg.Find("MSH.11")
	assert.NoError(t, err)
	assert.Equal(t, "P", v)

	v, err = msg.Find("MSH.12")
	assert.NoError(t, err)
	assert.Equal(t, "2.4", v)
}

type aMsg struct {
	FirstName string `hl7:"PID.5.1"`
	LastName  string `hl7:"PID.5.0"`
}

func TestMessageBuilding(t *testing.T) {
	mi := MsgInfo{
		SendingApp:        "BettrLife",
		SendingFacility:   "UnivIa",
		ReceivingApp:      "Epic",
		ReceivingFacility: "UnivIa",
		MessageType:       "ORM^001",
	}
	msg, err := StartMessage(mi)
	require.NoError(t, err)

	am := aMsg{FirstName: "Davin", LastName: "Hills"}
	_, err = Marshal(msg, &am)
	assert.NoError(t, err)
}
