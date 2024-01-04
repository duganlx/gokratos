package rbacapi

import (
	"testing"

	"github.com/casbin/casbin/v2"
	"github.com/stretchr/testify/assert"
)

func NewEnforcer() *casbin.Enforcer {

	e, err := casbin.NewEnforcer("./model.conf", "./policy.csv")
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

func TestDeleteRoleForUser(t *testing.T) {
	e := NewEnforcer()

	var ok bool
	var err error

	ok, err = e.DeleteRoleForUser("USER:lvx", "SRCGROUP:drw")
	assert.Nil(t, err)
	assert.True(t, ok)

	ok, err = e.Enforce("USER:lvx", "DRW101:v1:test", "*")
	assert.Nil(t, err)
	assert.False(t, ok)

	ok, err = e.DeleteRoleForUser("USER:lvx", "SRCGROUP:eam")
	assert.Nil(t, err)
	assert.False(t, ok)
}

func TestDeleteRolesForUser(t *testing.T) {
	e := NewEnforcer()

	var ok bool
	var err error

	ok, err = e.DeleteRolesForUser("USER:lvx")
	assert.Nil(t, err)
	assert.True(t, ok)

	ok, err = e.Enforce("USER:lvx", "DRW101:v1:test", "*")
	assert.Nil(t, err)
	assert.False(t, ok)

	ok, err = e.Enforce("USER:lvx", "DRW001:v1:test", "*")
	assert.Nil(t, err)
	assert.True(t, ok)
}

func TestDeleteUser(t *testing.T) {
	e := NewEnforcer()

	var ok bool
	var err error

	ok, err = e.DeleteUser("USER:lvx")
	assert.Nil(t, err)
	assert.True(t, ok)

	ok, err = e.Enforce("USER:lvx", "DRW101:v1:test", "*")
	assert.Nil(t, err)
	assert.False(t, ok)
}

func TestDeleteRole(t *testing.T) {
	e := NewEnforcer()

	var ok bool
	var err error

	ok, err = e.DeleteRole("SRCGROUP:drw")
	assert.Nil(t, err)
	assert.True(t, ok)

	ok, err = e.Enforce("USER:lvx", "DRW101:v1:test", "*")
	assert.Nil(t, err)
	assert.False(t, ok)
}

func TestDeletePermission(t *testing.T) {
	e := NewEnforcer()

	var ok bool
	var err error

	ok, err = e.DeletePermission("DRW001:v1:test")
	assert.Nil(t, err)
	assert.True(t, ok)

	ok, err = e.Enforce("USER:lvx", "DRW001:v1:test", "r")
	assert.Nil(t, err)
	assert.False(t, ok)
}

func TestAddPermissionForUser(t *testing.T) {
	e := NewEnforcer()

	var ok bool
	var err error

	ok, err = e.AddPermissionForUser("USER:lvx", "DRW006:v1:test", "w", "allow")
	assert.Nil(t, err)
	assert.True(t, ok)

	ok, err = e.Enforce("USER:lvx", "DRW006:v1:test", "w")
	assert.Nil(t, err)
	assert.True(t, ok)

	ok, err = e.Enforce("USER:lvx", "DRW006:v1:test", "r")
	assert.Nil(t, err)
	assert.False(t, ok)
}

func TestAddPermissionsForUser(t *testing.T) {
	e := NewEnforcer()

	var ok bool
	var err error

	permissions := [][]string{
		{"DRW006:v1:test", "w", "allow"},
		{"DRW007:v1:test", "r", "allow"},
	}

	ok, err = e.AddPermissionsForUser("USER:lvx", permissions...)
	assert.Nil(t, err)
	assert.True(t, ok)

	ok, err = e.Enforce("USER:lvx", "DRW006:v1:test", "w")
	assert.Nil(t, err)
	assert.True(t, ok)

	ok, err = e.Enforce("USER:lvx", "DRW007:v1:test", "r")
	assert.Nil(t, err)
	assert.True(t, ok)
}

func TestDeletePermissionForUser(t *testing.T) {
	e := NewEnforcer()

	var ok bool
	var err error

	ok, err = e.DeletePermissionForUser("USER:lvx", "DRW001:v1:test", "*", "allow")
	assert.Nil(t, err)
	assert.True(t, ok)

	ok, err = e.Enforce("USER:lvx", "DRW001:v1:test", "*")
	assert.Nil(t, err)
	assert.False(t, ok)
}

func TestDeletePermissionsForUser(t *testing.T) {
	e := NewEnforcer()

	var ok bool
	var err error

	// delete directly relationship
	ok, err = e.DeletePermissionsForUser("USER:lvx")
	assert.Nil(t, err)
	assert.True(t, ok)

	ok, err = e.Enforce("USER:lvx", "DRW001:v1:test", "*")
	assert.Nil(t, err)
	assert.False(t, ok)

	ok, err = e.Enforce("USER:lvx", "DRW101:v1:test", "*")
	assert.Nil(t, err)
	assert.True(t, ok)
}

func TestGetPermissionsForUser(t *testing.T) {
	e := NewEnforcer()

	res := e.GetPermissionsForUser("USER:lvx")

	expected := [][]string{
		{"USER:lvx", "DRW001:v1:test", "*", "allow"},
		{"USER:lvx", "DRW002:v1:test", "r", "allow"},
		{"USER:lvx", "DRW003:v1:test", "w", "allow"},
		{"USER:lvx", "DRW004:v1:test", "*", "allow"},
		{"USER:lvx", "DRW004:v1:test", "w", "deny"},
		{"USER:lvx", "DRW005:v1:test", "r", "allow"},
		{"USER:lvx", "DRW005:v1:test", "w", "allow"},
	}

	for _, item := range res {
		assert.Contains(t, expected, item)
	}
}

func TestHasPermissionForUser(t *testing.T) {
	e := NewEnforcer()

	var ok bool
	ok = e.HasPermissionForUser("USER:lvx", "DRW001:v1:test", "*", "allow")
	assert.True(t, ok)
	ok = e.HasPermissionForUser("USER:lvx", "DRW002:v1:test", "w", "allow")
	assert.False(t, ok)
}

func TestGetImplicitRolesForUser(t *testing.T) {
	e := NewEnforcer()

	res, err := e.GetImplicitRolesForUser("USER:lvx")
	expected := []string{"USERGROUP:it", "SRCGROUP:drw", "USERGROUP:eam", "SRCGROUP:eam"}
	assert.Nil(t, err)
	for _, item := range res {
		assert.Contains(t, expected, item)
	}
}

func TestGetImplicitUsersForRole(t *testing.T) {

	e := NewEnforcer()

	res, err := e.GetImplicitUsersForRole("SRCGROUP:eam")
	assert.Nil(t, err)

	expected := []string{"USERGROUP:eam", "USER:lvx"}
	for _, item := range res {
		assert.Contains(t, expected, item)
	}
}

func TestGetImplicitPermissionsForUser(t *testing.T) {
	e := NewEnforcer()

	res, err := e.GetImplicitPermissionsForUser("USER:lvx")
	assert.Nil(t, err)

	expected := [][]string{
		{"USER:lvx", "DRW001:v1:test", "*", "allow"},
		{"USER:lvx", "DRW002:v1:test", "r", "allow"},
		{"USER:lvx", "DRW003:v1:test", "w", "allow"},
		{"USER:lvx", "DRW004:v1:test", "*", "allow"},
		{"USER:lvx", "DRW004:v1:test", "w", "deny"},
		{"USER:lvx", "DRW005:v1:test", "r", "allow"},
		{"USER:lvx", "DRW005:v1:test", "w", "allow"},
		{"USERGROUP:it", "DRW011:v1:test", "*", "allow"},
		{"USERGROUP:it", "DRW012:v1:test", "r", "allow"},
		{"USERGROUP:it", "DRW013:v1:test", "w", "allow"},
		{"USERGROUP:it", "DRW014:v1:test", "*", "allow"},
		{"USERGROUP:it", "DRW014:v1:test", "w", "deny"},
		{"SRCGROUP:drw", "DRW101:v1:test", "*", "allow"},
		{"SRCGROUP:drw", "DRW102:v1:test", "r", "allow"},
		{"SRCGROUP:drw", "DRW103:v1:test", "w", "allow"},
		{"SRCGROUP:drw", "DRW104:v1:test", "*", "allow"},
		{"SRCGROUP:drw", "DRW104:v1:test", "w", "deny"},
		{"SRCGROUP:eam", "EAM101:v1:test", "*", "allow"},
		{"SRCGROUP:eam", "EAM102:v1:test", "r", "allow"},
		{"SRCGROUP:eam", "EAM103:v1:test", "w", "allow"},
		{"SRCGROUP:eam", "EAM104:v1:test", "*", "allow"},
		{"SRCGROUP:eam", "EAM104:v1:test", "w", "deny"},
	}

	for _, item := range res {
		assert.Contains(t, expected, item)
	}
}

func TestGetNamedImplicitPermissionsForUser(t *testing.T) {
	e := NewEnforcer()

	res, err := e.GetNamedImplicitPermissionsForUser("p", "USER:lvx")
	assert.Nil(t, err)

	expected := [][]string{
		{"USER:lvx", "DRW001:v1:test", "*", "allow"},
		{"USER:lvx", "DRW002:v1:test", "r", "allow"},
		{"USER:lvx", "DRW003:v1:test", "w", "allow"},
		{"USER:lvx", "DRW004:v1:test", "*", "allow"},
		{"USER:lvx", "DRW004:v1:test", "w", "deny"},
		{"USER:lvx", "DRW005:v1:test", "r", "allow"},
		{"USER:lvx", "DRW005:v1:test", "w", "allow"},
		{"USERGROUP:it", "DRW011:v1:test", "*", "allow"},
		{"USERGROUP:it", "DRW012:v1:test", "r", "allow"},
		{"USERGROUP:it", "DRW013:v1:test", "w", "allow"},
		{"USERGROUP:it", "DRW014:v1:test", "*", "allow"},
		{"USERGROUP:it", "DRW014:v1:test", "w", "deny"},
		{"SRCGROUP:drw", "DRW101:v1:test", "*", "allow"},
		{"SRCGROUP:drw", "DRW102:v1:test", "r", "allow"},
		{"SRCGROUP:drw", "DRW103:v1:test", "w", "allow"},
		{"SRCGROUP:drw", "DRW104:v1:test", "*", "allow"},
		{"SRCGROUP:drw", "DRW104:v1:test", "w", "deny"},
		{"SRCGROUP:eam", "EAM101:v1:test", "*", "allow"},
		{"SRCGROUP:eam", "EAM102:v1:test", "r", "allow"},
		{"SRCGROUP:eam", "EAM103:v1:test", "w", "allow"},
		{"SRCGROUP:eam", "EAM104:v1:test", "*", "allow"},
		{"SRCGROUP:eam", "EAM104:v1:test", "w", "deny"},
	}

	for _, item := range res {
		assert.Contains(t, expected, item)
	}
}

func TestGetDomainsForUser(t *testing.T) {
	e := NewEnforcer()

	res, err := e.GetDomainsForUser("USER:lvx")
	assert.Nil(t, err)
	assert.Equal(t, []string{""}, res)
}

func TestGetImplicitResourcesForUser(t *testing.T) {
	e := NewEnforcer()

	res, err := e.GetImplicitResourcesForUser("USER:lvx")
	assert.Nil(t, err)

	expected := [][]string{
		{"USER:lvx", "DRW001:v1:test", "*", "allow"},
		{"USER:lvx", "DRW002:v1:test", "r", "allow"},
		{"USER:lvx", "DRW003:v1:test", "w", "allow"},
		{"USER:lvx", "DRW004:v1:test", "*", "allow"},
		{"USER:lvx", "DRW004:v1:test", "w", "deny"},
		{"USER:lvx", "DRW005:v1:test", "r", "allow"},
		{"USER:lvx", "DRW005:v1:test", "w", "allow"},
		{"USER:lvx", "DRW011:v1:test", "*", "allow"},
		{"USER:lvx", "DRW012:v1:test", "r", "allow"},
		{"USER:lvx", "DRW013:v1:test", "w", "allow"},
		{"USER:lvx", "DRW014:v1:test", "*", "allow"},
		{"USER:lvx", "DRW014:v1:test", "w", "deny"},
		{"USER:lvx", "DRW101:v1:test", "*", "allow"},
		{"USER:lvx", "DRW102:v1:test", "r", "allow"},
		{"USER:lvx", "DRW103:v1:test", "w", "allow"},
		{"USER:lvx", "DRW104:v1:test", "*", "allow"},
		{"USER:lvx", "DRW104:v1:test", "w", "deny"},
		{"USER:lvx", "EAM101:v1:test", "*", "allow"},
		{"USER:lvx", "EAM102:v1:test", "r", "allow"},
		{"USER:lvx", "EAM103:v1:test", "w", "allow"},
		{"USER:lvx", "EAM104:v1:test", "*", "allow"},
		{"USER:lvx", "EAM104:v1:test", "w", "deny"},
	}

	for _, item := range res {
		assert.Contains(t, expected, item)
	}
}

func TestGetImplicitUsersForPermission(t *testing.T) {
	e := NewEnforcer()

	var res []string
	var err error
	res, err = e.GetImplicitUsersForPermission("DRW001:v1:test", "*")
	assert.Nil(t, err)
	assert.Equal(t, []string{"USER:lvx"}, res)

	res, err = e.GetImplicitUsersForPermission("DRW011:v1:test", "*")
	assert.Nil(t, err)
	assert.Equal(t, []string{"USER:lvx"}, res)

	res, err = e.GetImplicitUsersForPermission("EAM101:v1:test", "*")
	assert.Nil(t, err)
	assert.Equal(t, []string{"USER:lvx"}, res)
}

func TestGetAllowedObjectConditions(t *testing.T) {
	// GetAllowedObjectConditions returns a string array of object conditions that the user can access.
}

func TestGetImplicitUsersForResource(t *testing.T) {
	e := NewEnforcer()

	res, err := e.GetImplicitUsersForResource("DRW013:v1:test")
	assert.Nil(t, err)
	assert.Equal(t, [][]string{{"USER:lvx", "DRW013:v1:test", "w", "allow"}}, res)
}
