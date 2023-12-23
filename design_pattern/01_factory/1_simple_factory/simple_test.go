package simplefactory

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// TestType1 test get hiapi with factory
func TestType1(t *testing.T) {
	api := NewAPI(1)
	s := api.Say("Tom")
	assert.Equal(t, "Hi, Tom", s)
}

func TestType2(t *testing.T) {
	api := NewAPI(2)
	s := api.Say("Tom")
	assert.Equal(t, "Hello, Tom", s)
}
