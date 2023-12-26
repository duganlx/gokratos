package mysqlop

import (
	"testing"

	"github.com/casbin/casbin/v2"
	"github.com/casbin/casbin/v2/model"
	xormadapter "github.com/casbin/xorm-adapter/v2"
	_ "github.com/go-sql-driver/mysql"
	"github.com/stretchr/testify/assert"
)

// Mysql data preparation
// INSERT INTO jhl_uc.casbin_rbac VALUES('p', 'USER:lvx', 'EAM01:v1:test', '*', '', '', '');
func NewEnforcer() *casbin.Enforcer {
	driver := "mysql"
	source := "root:root@tcp(192.168.15.42:3306)/jhl_uc"
	tbname := "rbac"
	tbprefix := "casbin_"

	rbacAdapter, err := xormadapter.NewAdapterWithTableName(driver, source, tbname, tbprefix, true)
	if err != nil {
		panic(err)
	}

	casbinModel, err := model.NewModelFromFile("../model.conf")
	if err != nil {
		panic(err)
	}

	e, err := casbin.NewEnforcer(casbinModel, rbacAdapter)
	if err != nil {
		panic(err)
	}

	e.LoadPolicy()
	return e
}

func TestEnforce(t *testing.T) {
	e := NewEnforcer()

	var ok bool
	var err error

	ok, err = e.Enforce("USER:lvx", "EAM01:v1:test", "*")
	assert.Nil(t, err)
	assert.True(t, ok)
}

func TestInsert(t *testing.T) {
	e := NewEnforcer()

	var ok bool
	var err error

	ok, err = e.AddPolicy("USER:tim", "EAM01:v1:test", "*", "allow")
	assert.Nil(t, err)
	assert.True(t, ok)

	ok, err = e.AddGroupingPolicy("USER:tim", "SRCGROUP:happy")
	assert.Nil(t, err)
	assert.True(t, ok)
}

func TestDelete(t *testing.T) {
	e := NewEnforcer()

	var ok bool
	var err error

	ok, err = e.RemovePolicy("USER:tim", "EAM01:v1:test", "*", "allow")
	assert.Nil(t, err)
	assert.True(t, ok)

	ok, err = e.RemoveGroupingPolicy("USER:tim", "SRCGROUP:happy")
	assert.Nil(t, err)
	assert.True(t, ok)
}

func TestUpdate(t *testing.T) {
	e := NewEnforcer()

	var ok bool
	var err error

	ok, err = e.UpdatePolicy([]string{"USER:tim", "EAM01:v1:test", "*", "allow"}, []string{"USER:tim", "EAM01:v1:test", "r", "deny"})
	assert.Nil(t, err)
	assert.True(t, ok)

	ok, err = e.UpdateGroupingPolicy([]string{"USER:tim", "SRCGROUP:happy"}, []string{"USER:tim", "SRCGROUP:sad"})
	assert.Nil(t, err)
	assert.True(t, ok)
}
