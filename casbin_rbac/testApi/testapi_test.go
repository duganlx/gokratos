package testapi

import (
	"testing"

	"github.com/casbin/casbin/v2"
	"github.com/stretchr/testify/assert"
)

func NewEnforcer() *casbin.Enforcer {

	e, err := casbin.NewEnforcer("../model.conf", "../policy.csv")
	if err != nil {
		panic(err)
	}

	return e
}

func TestEnforce(t *testing.T) {
	e := NewEnforcer()

	var ok bool
	var err error

	ok, err = e.Enforce("USER:lvx", "DRW011:v1:test", "*")
	assert.Nil(t, err)
	assert.True(t, ok)
}

func TestEnforceWithMatcher(t *testing.T) {
	e := NewEnforcer()

	var ok bool
	var err error

	ok, err = e.Enforce("USER:lvx", "x:v1:test", "*")
	assert.Nil(t, err)
	assert.False(t, ok)

	//If the 'act' is 'w', authentication is granted.
	matcher := "r.act == 'w'"
	ok, err = e.EnforceWithMatcher(matcher, "USER:lvx", "x:v1:test", "w")
	assert.Nil(t, err)
	assert.True(t, ok)
	ok, err = e.EnforceWithMatcher(matcher, "unknown", "unknown", "w")
	assert.Nil(t, err)
	assert.True(t, ok)
}
