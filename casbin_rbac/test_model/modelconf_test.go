package testmodel

import (
	"testing"

	"github.com/casbin/casbin/v2"
	"github.com/stretchr/testify/assert"
)

func NewEnforcer(policies [][]string) *casbin.Enforcer {

	e, err := casbin.NewEnforcer("../model.conf", "./empty.csv")
	if err != nil {
		panic(err)
	}

	// Load Policy Configuration
	for _, policy := range policies {
		tye := policy[0]

		switch tye {
		case "p":
			e.AddPolicy(policy[1:])
		case "g":
			e.AddGroupingPolicy(policy[1:])
		default:
		}
	}
	return e
}

func TestModelcnf(t *testing.T) {
	// user -> src
	policies := [][]string{
		{"p", "USER:lvx", "EAM01:v1:test", "*", "allow"},

		{"p", "USER:lvx", "EAM01:v2:test", "*", "allow"},
		{"p", "USER:lvx", "EAM01:v2:test", "*", "deny"},

		{"p", "USER:lvx", "EAM02:v2:test", "r", "allow"},
		{"p", "USER:lvx", "EAM03:v2:test", "w", "allow"},
		{"p", "USER:lvx", "EAM04:v2:test", "r", "allow"},
		{"p", "USER:lvx", "EAM04:v2:test", "w", "allow"},
	}
	e := NewEnforcer(policies)

	var ok bool
	var err error

	// scene 1: 通配符权限
	ok, err = e.Enforce("USER:lvx", "EAM01:v1:test", "*")
	assert.Nil(t, err)
	assert.True(t, ok)
	ok, err = e.Enforce("USER:lvx", "EAM01:v1:test", "r")
	assert.Nil(t, err)
	assert.True(t, ok)
	ok, err = e.Enforce("USER:lvx", "EAM01:v1:test", "w")
	assert.Nil(t, err)
	assert.True(t, ok)
	ok, err = e.Enforce("USER:lvx", "EAM01:v1:test", "x")
	assert.Nil(t, err)
	assert.True(t, ok)

	// scene 2: 同时存在 允许&拒接 策略
	ok, err = e.Enforce("USER:lvx", "EAM01:v2:test", "*")
	assert.Nil(t, err)
	assert.False(t, ok)

	// scene 3: 读写权限
	ok, err = e.Enforce("USER:lvx", "EAM02:v2:test", "r")
	assert.Nil(t, err)
	assert.True(t, ok)
	ok, err = e.Enforce("USER:lvx", "EAM02:v2:test", "w")
	assert.Nil(t, err)
	assert.False(t, ok)
	ok, err = e.Enforce("USER:lvx", "EAM03:v2:test", "r")
	assert.Nil(t, err)
	assert.False(t, ok)
	ok, err = e.Enforce("USER:lvx", "EAM03:v2:test", "w")
	assert.Nil(t, err)
	assert.True(t, ok)
	ok, err = e.Enforce("USER:lvx", "EAM04:v2:test", "r")
	assert.Nil(t, err)
	assert.True(t, ok)
	ok, err = e.Enforce("USER:lvx", "EAM04:v2:test", "w")
	assert.Nil(t, err)
	assert.True(t, ok)

	// scene 4: 无权限用户 & 无权限资源
	ok, err = e.Enforce("USER:lvx", "EAM01:v1:prod", "r")
	assert.Nil(t, err)
	assert.False(t, ok)
	ok, err = e.Enforce("USER:xxx", "EAM01:v1:test", "r")
	assert.Nil(t, err)
	assert.False(t, ok)
}

func TestModelcnf2(t *testing.T) {
	// user -> userGroup/srcGroup -> src
	policies := [][]string{
		{"g", "USER:lvx", "USERGROUP:it"},

		{"p", "USERGROUP:it", "EAM01:v1:test", "*", "allow"},
		{"p", "USERGROUP:it", "EAM02:v1:test", "r", "allow"},
		{"p", "USERGROUP:it", "EAM03:v1:test", "w", "allow"},
		{"p", "USERGROUP:it", "EAM04:v1:test", "r", "allow"},
		{"p", "USERGROUP:it", "EAM04:v1:test", "w", "allow"},

		{"p", "USER:lvx", "EAM04:v1:test", "r", "deny"},
		{"p", "USER:lvx", "EAM02:v1:test", "w", "allow"},
		{"p", "USER:lvx", "EAM01:v1:test", "w", "deny"},
	}
	e := NewEnforcer(policies)

	var ok bool
	var err error

	ok, err = e.Enforce("USER:lvx", "EAM01:v1:test", "*")
	assert.Nil(t, err)
	assert.True(t, ok)
	ok, err = e.Enforce("USER:lvx", "EAM01:v1:test", "r")
	assert.Nil(t, err)
	assert.True(t, ok)
	ok, err = e.Enforce("USER:lvx", "EAM01:v1:test", "w")
	assert.Nil(t, err)
	assert.False(t, ok)

	ok, err = e.Enforce("USER:lvx", "EAM02:v1:test", "*")
	assert.Nil(t, err)
	assert.False(t, ok)
	ok, err = e.Enforce("USER:lvx", "EAM02:v1:test", "r")
	assert.Nil(t, err)
	assert.True(t, ok)
	ok, err = e.Enforce("USER:lvx", "EAM02:v1:test", "w")
	assert.Nil(t, err)
	assert.True(t, ok)

	ok, err = e.Enforce("USER:lvx", "EAM03:v1:test", "*")
	assert.Nil(t, err)
	assert.False(t, ok)
	ok, err = e.Enforce("USER:lvx", "EAM03:v1:test", "r")
	assert.Nil(t, err)
	assert.False(t, ok)
	ok, err = e.Enforce("USER:lvx", "EAM03:v1:test", "w")
	assert.Nil(t, err)
	assert.True(t, ok)

	ok, err = e.Enforce("USER:lvx", "EAM04:v1:test", "*")
	assert.Nil(t, err)
	assert.False(t, ok)
	ok, err = e.Enforce("USER:lvx", "EAM04:v1:test", "r")
	assert.Nil(t, err)
	assert.False(t, ok)
	ok, err = e.Enforce("USER:lvx", "EAM04:v1:test", "w")
	assert.Nil(t, err)
	assert.True(t, ok)

	ok, err = e.Enforce("USER:hacker", "EAM01:v1:test", "*")
	assert.Nil(t, err)
	assert.False(t, ok)
}

func TestModelcnf3(t *testing.T) {
	// user -> userGroup -> srcGroup -> src
	policies := [][]string{
		{"g", "USER:lvx", "USERGROUP:it"},
		{"g", "USERGROUP:it", "SRCGROUP:EAM"},

		{"p", "SRCGROUP:EAM", "EAM01:v1:test", "*", "allow"},
		{"p", "SRCGROUP:EAM", "EAM02:v1:test", "r", "allow"},
		{"p", "SRCGROUP:EAM", "EAM03:v1:test", "w", "allow"},
		{"p", "SRCGROUP:EAM", "EAM04:v1:test", "*", "allow"},
		{"p", "SRCGROUP:EAM", "EAM04:v1:test", "w", "deny"},
	}
	e := NewEnforcer(policies)

	var ok bool
	var err error
	ok, err = e.Enforce("USER:lvx", "EAM01:v1:test", "*")
	assert.Nil(t, err)
	assert.True(t, ok)
	ok, err = e.Enforce("USER:lvx", "EAM01:v1:test", "r")
	assert.Nil(t, err)
	assert.True(t, ok)
	ok, err = e.Enforce("USER:lvx", "EAM01:v1:test", "w")
	assert.Nil(t, err)
	assert.True(t, ok)

	ok, err = e.Enforce("USER:lvx", "EAM02:v1:test", "*")
	assert.Nil(t, err)
	assert.False(t, ok)
	ok, err = e.Enforce("USER:lvx", "EAM02:v1:test", "r")
	assert.Nil(t, err)
	assert.True(t, ok)
	ok, err = e.Enforce("USER:lvx", "EAM02:v1:test", "w")
	assert.Nil(t, err)
	assert.False(t, ok)

	ok, err = e.Enforce("USER:lvx", "EAM03:v1:test", "*")
	assert.Nil(t, err)
	assert.False(t, ok)
	ok, err = e.Enforce("USER:lvx", "EAM03:v1:test", "r")
	assert.Nil(t, err)
	assert.False(t, ok)
	ok, err = e.Enforce("USER:lvx", "EAM03:v1:test", "w")
	assert.Nil(t, err)
	assert.True(t, ok)

	ok, err = e.Enforce("USER:lvx", "EAM04:v1:test", "*")
	assert.Nil(t, err)
	assert.True(t, ok)
	ok, err = e.Enforce("USER:lvx", "EAM04:v1:test", "r")
	assert.Nil(t, err)
	assert.True(t, ok)
	ok, err = e.Enforce("USER:lvx", "EAM04:v1:test", "w")
	assert.Nil(t, err)
	assert.False(t, ok)
}
