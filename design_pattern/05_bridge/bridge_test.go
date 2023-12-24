package bridge

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestExampleCommonSMS(t *testing.T) {
	m := NewCommonMessage(ViaSMS())
	res := m.SendMessage("have a drink?", "bob")
	assert.Equal(t, "send have a drink? to bob via SMS", res)
}

func TestExampleCommonEmail(t *testing.T) {
	m := NewCommonMessage(ViaEmail())
	res := m.SendMessage("have a drink?", "bob")
	assert.Equal(t, "send have a drink? to bob via Email", res)
}

func TestExampleUrgencySMS(t *testing.T) {
	m := NewUrgencyMessage(ViaSMS())
	res := m.SendMessage("have a drink?", "bob")
	assert.Equal(t, "send [Urgency] have a drink? to bob via SMS", res)
}

func TestExampleUrgencyEmail(t *testing.T) {
	m := NewUrgencyMessage(ViaEmail())
	res := m.SendMessage("have a drink?", "bob")
	assert.Equal(t, "send [Urgency] have a drink? to bob via Email", res)
}
