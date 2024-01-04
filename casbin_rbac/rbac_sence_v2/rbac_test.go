package rbacsencev2

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

func TestUser2Src(t *testing.T) {
	e := NewEnforcer()

	var ok bool
	var err error

	ok, err = e.Enforce("USER:lvx", "DRW001:v1:test", "*", "*")
	assert.True(t, ok)
	assert.Nil(t, err)
	ok, err = e.Enforce("USER:lvx", "DRW001:v1:test", "conf", "*")
	assert.True(t, ok)
	assert.Nil(t, err)
	ok, err = e.Enforce("USER:lvx", "DRW001:v1:test", "conf", "r")
	assert.True(t, ok)
	assert.Nil(t, err)
	ok, err = e.Enforce("USER:lvx", "DRW001:v1:test", "conf", "w")
	assert.True(t, ok)
	assert.Nil(t, err)

	ok, err = e.Enforce("USER:lvx", "DRW002:v1:test", "*", "*")
	assert.False(t, ok)
	assert.Nil(t, err)
	ok, err = e.Enforce("USER:lvx", "DRW002:v1:test", "conf", "*")
	assert.False(t, ok)
	assert.Nil(t, err)
	ok, err = e.Enforce("USER:lvx", "DRW002:v1:test", "conf", "r")
	assert.True(t, ok)
	assert.Nil(t, err)
	ok, err = e.Enforce("USER:lvx", "DRW002:v1:test", "conf", "w")
	assert.False(t, ok)
	assert.Nil(t, err)

	ok, err = e.Enforce("USER:lvx", "DRW003:v1:test", "*", "w")
	assert.True(t, ok)
	assert.Nil(t, err)
	ok, err = e.Enforce("USER:lvx", "DRW003:v1:test", "conf", "w")
	assert.True(t, ok)
	assert.Nil(t, err)
	ok, err = e.Enforce("USER:lvx", "DRW003:v1:test", "conf", "r")
	assert.False(t, ok)
	assert.Nil(t, err)

	ok, err = e.Enforce("USER:lvx", "DRW004:v1:test", "conf", "r")
	assert.True(t, ok)
	assert.Nil(t, err)
	ok, err = e.Enforce("USER:lvx", "DRW004:v1:test", "other", "r")
	assert.False(t, ok)
	assert.Nil(t, err)
	ok, err = e.Enforce("USER:lvx", "DRW004:v1:test", "conf", "w")
	assert.False(t, ok)
	assert.Nil(t, err)
}

func TestUserGroup2Src(t *testing.T) {

}
