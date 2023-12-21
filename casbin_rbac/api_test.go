package casbin_rbac

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestModelcnf(t *testing.T) {
	// user 和 src 直接关联
	policies := [][]string{
		{"p", "USER:lvx", "EAM01:v1:test", "*", "allow"},
	}
	e := NewEnforcer(policies)

	var ok bool
	var err error

	ok, err = e.Enforce("USER:lvx", "EAM01:v1:test", "*")
	assert.Nil(t, err)
	assert.True(t, ok)
	ok, err = e.Enforce("USER:lvx", "EAM01:v1:test", "ro")
	assert.Nil(t, err)
	assert.True(t, ok)
	ok, err = e.Enforce("USER:lvx", "EAM01:v1:test", "wo")
	assert.Nil(t, err)
	assert.True(t, ok)
	ok, err = e.Enforce("USER:lvx", "EAM01:v1:test", "rw")
	assert.Nil(t, err)
	assert.True(t, ok)
}

func TestModelcnf1(t *testing.T) {
	// user 和 src 直接关联
	policies := [][]string{
		{"p", "USER:lvx", "EAM01:v1:test", "*", "allow"},
		{"p", "USER:lvx", "EAM01:v1:test", "*", "deny"},
	}
	e := NewEnforcer(policies)

	var ok bool
	var err error

	ok, err = e.Enforce("USER:lvx", "EAM01:v1:test", "*")
	assert.Nil(t, err)
	assert.False(t, ok)
}

// func TestModelcnf1(t *testing.T) {
// 	// user - userGroup - src
// 	policies := [][]string{
// 		{"p", "USERGROUP:it", "EAM01:v1:test", "*", "allow"},
// 		{"g", "USER:lvx", "USERGROUP:it"},
// 	}
// 	e := NewEnforcer(policies)

// 	var ok bool
// 	var err error

// 	ok, err = e.Enforce("USER:lvx", "EAM01:v1:test", "*")
// 	assert.Nil(t, err)
// 	assert.True(t, ok)

// 	ok, err = e.Enforce("USER:hacker", "EAM01:v1:test", "*")
// 	assert.Nil(t, err)
// 	assert.False(t, ok)
// }

// func TestModelcnf2(t *testing.T) {
// 	// user - srcGroup - src
// 	policies := [][]string{
// 		{"p", "SRCGROUP:eam", "EAM01:v1:test", "*", "allow"},
// 		{"g", "USER:lvx", "SRCGROUP:eam"},
// 	}
// 	e := NewEnforcer(policies)

// 	ok, err := e.Enforce("USER:lvx", "EAM01:v1:test", "*")
// 	assert.Nil(t, err)
// 	assert.True(t, ok)
// }
