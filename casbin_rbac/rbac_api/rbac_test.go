package rbacapi

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

func TestGetRolesForUser(t *testing.T) {
	e := NewEnforcer()

	expected := []string{"USERGROUP:it", "SRCGROUP:drw", "USERGROUP:eam"}
	res, err := e.GetRolesForUser("USER:lvx")
	assert.Nil(t, err)
	for _, item := range res {
		assert.Contains(t, expected, item)
	}
}

func TestGetUsersForRole(t *testing.T) {
	e := NewEnforcer()

	expected := []string{"USERGROUP:eam"}
	res, err := e.GetUsersForRole("SRCGROUP:eam")
	assert.Nil(t, err)
	assert.Equal(t, expected, res)
}

func TestHasRoleForUser(t *testing.T) {
	e := NewEnforcer()

	var has bool
	var err error

	has, err = e.HasRoleForUser("USER:lvx", "USERGROUP:it")
	assert.Nil(t, err)
	assert.True(t, has)

	has, err = e.HasRoleForUser("USER:lvx", "SRCGROUP:eam")
	assert.Nil(t, err)
	assert.False(t, has)
}

func TestAddRoleForUser(t *testing.T) {
	e := NewEnforcer()

	var ok bool
	var err error

	ok, err = e.AddRoleForUser("USER:lvx", "SRCGROUP:ev2")
	assert.Nil(t, err)
	assert.True(t, ok)

	ok, err = e.AddRoleForUser("USER:lvx", "USERGROUP:it")
	assert.Nil(t, err)
	assert.False(t, ok)

	ok, err = e.Enforce("USER:lvx", "DRW011:v1:test", "*")
	assert.Nil(t, err)
	assert.True(t, ok)
}
