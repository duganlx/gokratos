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

func TestEnforceEx(t *testing.T) {
	// EnforceEx 多一个 reason 参数看出最终命中的规则
	e := NewEnforcer()

	var ok bool
	var err error
	var reason []string

	ok, reason, err = e.EnforceEx("unknown", "unknown", "w")
	assert.Nil(t, err)
	assert.False(t, ok)
	assert.Equal(t, reason, []string{})

	ok, reason, err = e.EnforceEx("USER:lvx", "DRW001:v1:test", "*")
	assert.Nil(t, err)
	assert.True(t, ok)
	assert.Equal(t, reason, []string{"USER:lvx", "DRW001:v1:test", "*", "allow"})

	ok, reason, err = e.EnforceEx("USER:lvx", "DRW011:v1:test", "*")
	assert.Nil(t, err)
	assert.True(t, ok)
	assert.Equal(t, reason, []string{"USERGROUP:it", "DRW011:v1:test", "*", "allow"})

	ok, reason, err = e.EnforceEx("USER:lvx", "DRW101:v1:test", "*")
	assert.Nil(t, err)
	assert.True(t, ok)
	assert.Equal(t, reason, []string{"SRCGROUP:drw", "DRW101:v1:test", "*", "allow"})

	ok, reason, err = e.EnforceEx("USER:lvx", "EAM101:v1:test", "*")
	assert.Nil(t, err)
	assert.True(t, ok)
	assert.Equal(t, reason, []string{"SRCGROUP:eam", "EAM101:v1:test", "*", "allow"})
}

func TestEnforceExWithMatcher(t *testing.T) {
	e := NewEnforcer()

	var ok bool
	var err error
	var reason []string

	ok, reason, err = e.EnforceEx("USER:lvx", "x:v1:test", "*")
	assert.Nil(t, err)
	assert.False(t, ok)
	assert.Equal(t, reason, []string{})

	// If the 'act' is 'w', authentication is granted. 所以第一条策略就匹配成功
	matcher := "r.act == 'w'"
	ok, reason, err = e.EnforceExWithMatcher(matcher, "USER:lvx", "x:v1:test", "w")
	assert.Nil(t, err)
	assert.True(t, ok)
	assert.Equal(t, reason, []string{"USER:lvx", "DRW001:v1:test", "*", "allow"})

	ok, reason, err = e.EnforceExWithMatcher(matcher, "unknown", "unknown", "w")
	assert.Nil(t, err)
	assert.True(t, ok)
	assert.Equal(t, reason, []string{"USER:lvx", "DRW001:v1:test", "*", "allow"})
}

func TestBatchEnforce(t *testing.T) {
	e := NewEnforcer()

	var boolArray []bool
	var err error

	requests := [][]interface{}{
		{"USER:lvx", "DRW001:v1:test", "*"},
		{"USER:lvx", "DRW002:v1:test", "r"},
		{"USER:lvx", "DRW002:v1:test", "w"},
		{"USER:lvx", "DRW011:v1:test", "w"},
		{"USER:lvx", "DRW014:v1:test", "w"},
		{"USER:lvx", "EAM101:v1:test", "w"},
		{"USER:lvx", "EAM102:v1:test", "w"},
	}
	boolArray, err = e.BatchEnforce(requests)

	res := []bool{true, true, false, true, false, true, false}
	assert.Nil(t, err)
	assert.Equal(t, boolArray, res)
}

func TestGetAllSubjects(t *testing.T) {
	e := NewEnforcer()

	allSubjects := e.GetAllSubjects()

	expected := []string{"USER:lvx", "USERGROUP:it", "SRCGROUP:drw", "SRCGROUP:eam", "SRCGROUP:ev2"}
	for _, sub := range allSubjects {
		assert.Contains(t, expected, sub)
	}
}

func TestGetAllNamedSubjects(t *testing.T) {
	e := NewEnforcer()

	// only policy
	allNamedSubjects := e.GetAllNamedSubjects("p")

	expected := []string{"USER:lvx", "USERGROUP:it", "SRCGROUP:drw", "SRCGROUP:eam", "SRCGROUP:ev2"}
	for _, sub := range allNamedSubjects {
		assert.Contains(t, expected, sub)
	}
}

func TestGetAllObjects(t *testing.T) {
	e := NewEnforcer()

	allObjects := e.GetAllObjects()

	expected := []string{"DRW001:v1:test", "DRW002:v1:test", "DRW003:v1:test", "DRW004:v1:test", "DRW005:v1:test", "DRW011:v1:test", "DRW012:v1:test", "DRW013:v1:test", "DRW014:v1:test", "DRW101:v1:test", "DRW102:v1:test", "DRW103:v1:test", "DRW104:v1:test", "EAM101:v1:test", "EAM102:v1:test", "EAM103:v1:test", "EAM104:v1:test", "EAM101:v2:test", "EAM102:v2:test", "EAM103:v2:test", "EAM104:v2:test", "EAM104:v2:test"}
	for _, sub := range allObjects {
		assert.Contains(t, expected, sub)
	}
}

func TestGetAllNamedObjects(t *testing.T) {
	e := NewEnforcer()

	allNamedObjects := e.GetAllNamedObjects("p")

	expected := []string{"DRW001:v1:test", "DRW002:v1:test", "DRW003:v1:test", "DRW004:v1:test", "DRW005:v1:test", "DRW011:v1:test", "DRW012:v1:test", "DRW013:v1:test", "DRW014:v1:test", "DRW101:v1:test", "DRW102:v1:test", "DRW103:v1:test", "DRW104:v1:test", "EAM101:v1:test", "EAM102:v1:test", "EAM103:v1:test", "EAM104:v1:test", "EAM101:v2:test", "EAM102:v2:test", "EAM103:v2:test", "EAM104:v2:test", "EAM104:v2:test"}
	for _, sub := range allNamedObjects {
		assert.Contains(t, expected, sub)
	}
}

func TestGetAllActions(t *testing.T) {
	e := NewEnforcer()

	allActions := e.GetAllActions()

	expected := []string{"*", "r", "w"}
	for _, act := range allActions {
		assert.Contains(t, expected, act)
	}
}

func TestGetAllNamedActions(t *testing.T) {
	e := NewEnforcer()

	allActions := e.GetAllNamedActions("p")

	expected := []string{"*", "r", "w"}
	for _, act := range allActions {
		assert.Contains(t, expected, act)
	}
}

func TestGetAllRoles(t *testing.T) {
	e := NewEnforcer()

	allRoles := e.GetAllRoles()

	expected := []string{"USERGROUP:it", "SRCGROUP:drw", "USERGROUP:eam", "SRCGROUP:eam"}
	for _, role := range allRoles {
		assert.Contains(t, expected, role)
	}
}

func TestGetAllNamedRoles(t *testing.T) {
	e := NewEnforcer()

	allRoles := e.GetAllNamedRoles("g")

	expected := []string{"USERGROUP:it", "SRCGROUP:drw", "USERGROUP:eam", "SRCGROUP:eam"}
	for _, role := range allRoles {
		assert.Contains(t, expected, role)
	}
}

func TestGetPolicy(t *testing.T) {
	e := NewEnforcer()

	policies := e.GetPolicy()

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
		{"SRCGROUP:ev2", "EAM101:v2:test", "*", "allow"},
		{"SRCGROUP:ev2", "EAM102:v2:test", "r", "allow"},
		{"SRCGROUP:ev2", "EAM103:v2:test", "w", "allow"},
		{"SRCGROUP:ev2", "EAM104:v2:test", "*", "allow"},
		{"SRCGROUP:ev2", "EAM104:v2:test", "w", "deny"},
	}

	for _, policy := range policies {
		assert.Contains(t, expected, policy)
	}
}

func TestGetFilteredPolicy(t *testing.T) {
	e := NewEnforcer()

	policies := e.GetFilteredPolicy(0, "USER:lvx")

	expected := [][]string{
		{"USER:lvx", "DRW001:v1:test", "*", "allow"},
		{"USER:lvx", "DRW002:v1:test", "r", "allow"},
		{"USER:lvx", "DRW003:v1:test", "w", "allow"},
		{"USER:lvx", "DRW004:v1:test", "*", "allow"},
		{"USER:lvx", "DRW004:v1:test", "w", "deny"},
		{"USER:lvx", "DRW005:v1:test", "r", "allow"},
		{"USER:lvx", "DRW005:v1:test", "w", "allow"},
	}

	for _, policy := range policies {
		assert.Contains(t, expected, policy)
	}
}

func TestGetNamedPolicy(t *testing.T) {
	e := NewEnforcer()

	namedPolicy := e.GetNamedPolicy("p")

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
		{"SRCGROUP:ev2", "EAM101:v2:test", "*", "allow"},
		{"SRCGROUP:ev2", "EAM102:v2:test", "r", "allow"},
		{"SRCGROUP:ev2", "EAM103:v2:test", "w", "allow"},
		{"SRCGROUP:ev2", "EAM104:v2:test", "*", "allow"},
		{"SRCGROUP:ev2", "EAM104:v2:test", "w", "deny"},
	}

	for _, policy := range namedPolicy {
		assert.Contains(t, expected, policy)
	}
}

func TestGetFilteredNamedPolicy(t *testing.T) {
	e := NewEnforcer()

	policies := e.GetFilteredNamedPolicy("p", 0, "USER:lvx")

	expected := [][]string{
		{"USER:lvx", "DRW001:v1:test", "*", "allow"},
		{"USER:lvx", "DRW002:v1:test", "r", "allow"},
		{"USER:lvx", "DRW003:v1:test", "w", "allow"},
		{"USER:lvx", "DRW004:v1:test", "*", "allow"},
		{"USER:lvx", "DRW004:v1:test", "w", "deny"},
		{"USER:lvx", "DRW005:v1:test", "r", "allow"},
		{"USER:lvx", "DRW005:v1:test", "w", "allow"},
	}

	for _, policy := range policies {
		assert.Contains(t, expected, policy)
	}
}

func TestGetGroupingPolicy(t *testing.T) {
	e := NewEnforcer()

	groupingPolicy := e.GetGroupingPolicy()

	expected := [][]string{
		{"USER:lvx", "USERGROUP:it"},
		{"USER:lvx", "SRCGROUP:drw"},
		{"USER:lvx", "USERGROUP:eam"},
		{"USERGROUP:eam", "SRCGROUP:eam"},
	}

	for _, policy := range groupingPolicy {
		assert.Contains(t, expected, policy)
	}
}

func TestGetFilteredGroupingPolicy(t *testing.T) {
	e := NewEnforcer()

	filteredGroupingPolicy := e.GetFilteredGroupingPolicy(0, "USER:lvx")

	expected := [][]string{
		{"USER:lvx", "USERGROUP:it"},
		{"USER:lvx", "SRCGROUP:drw"},
		{"USER:lvx", "USERGROUP:eam"},
	}

	for _, policy := range filteredGroupingPolicy {
		assert.Contains(t, expected, policy)
	}
}

func TestGetNamedGroupingPolicy(t *testing.T) {
	e := NewEnforcer()

	namedGroupingPolicy := e.GetNamedGroupingPolicy("g")

	expected := [][]string{
		{"USER:lvx", "USERGROUP:it"},
		{"USER:lvx", "SRCGROUP:drw"},
		{"USER:lvx", "USERGROUP:eam"},
		{"USERGROUP:eam", "SRCGROUP:eam"},
	}

	for _, policy := range namedGroupingPolicy {
		assert.Contains(t, expected, policy)
	}
}

func TestGetFilteredNamedGroupingPolicy(t *testing.T) {
	e := NewEnforcer()

	filteredGroupingPolicy := e.GetFilteredNamedGroupingPolicy("g", 0, "USER:lvx")

	expected := [][]string{
		{"USER:lvx", "USERGROUP:it"},
		{"USER:lvx", "SRCGROUP:drw"},
		{"USER:lvx", "USERGROUP:eam"},
	}

	for _, policy := range filteredGroupingPolicy {
		assert.Contains(t, expected, policy)
	}
}

func TestHasPolicy(t *testing.T) {
	e := NewEnforcer()

	var hasPolicy bool

	hasPolicy = e.HasPolicy("USER:lvx", "DRW001:v1:test", "*", "allow")
	assert.True(t, hasPolicy)
	hasPolicy = e.HasPolicy("USER:lvx", "DRW002:v1:test", "*", "allow")
	assert.False(t, hasPolicy)
	hasPolicy = e.HasPolicy("USER:lvx", "DRW005:v1:test", "*", "allow")
	assert.False(t, hasPolicy)
	hasPolicy = e.HasPolicy("USER:lvx", "DRW004:v1:test", "w", "deny")
	assert.True(t, hasPolicy)
}

func TestHasNamedPolicy(t *testing.T) {
	e := NewEnforcer()

	var hasPolicy bool

	hasPolicy = e.HasNamedPolicy("p", "USER:lvx", "DRW001:v1:test", "*", "allow")
	assert.True(t, hasPolicy)
	hasPolicy = e.HasNamedPolicy("p", "USER:lvx", "DRW002:v1:test", "*", "allow")
	assert.False(t, hasPolicy)
	hasPolicy = e.HasNamedPolicy("p", "USER:lvx", "DRW005:v1:test", "*", "allow")
	assert.False(t, hasPolicy)
	hasPolicy = e.HasNamedPolicy("p", "USER:lvx", "DRW004:v1:test", "w", "deny")
	assert.True(t, hasPolicy)
}

func TestAddPolicy(t *testing.T) {
	e := NewEnforcer()

	var ok bool
	var err error

	ok, err = e.AddPolicy("USER:lvx", "DRW001:v1:test", "*", "allow")
	assert.Nil(t, err)
	assert.False(t, ok)

	ok, err = e.Enforce("USER:lvx", "DRW011:v1:tmp", "*")
	assert.Nil(t, err)
	assert.False(t, ok)

	ok, err = e.AddPolicy("USER:lvx", "DRW011:v1:tmp", "*", "allow")
	assert.Nil(t, err)
	assert.True(t, ok)

	ok, err = e.Enforce("USER:lvx", "DRW011:v1:tmp", "*")
	assert.Nil(t, err)
	assert.True(t, ok)
}

func TestAddPolicies(t *testing.T) {
	// AddPolicy 向当前策略添加授权许多规则。该操作本质上是原子的，因此，如果授权规则由不符合现行政策的规则组成，
	// 函数返回false，当前政策中没有添加任何政策规则。如果所有授权规则都符合政策规则，则函数返回true，每项政策规则都被添加到目前的政策中。
	e := NewEnforcer()

	var ok bool
	var err error

	ok, err = e.Enforce("USER:lvx", "DRW001:v1:tmp", "*")
	assert.Nil(t, err)
	assert.False(t, ok)
	ok, err = e.Enforce("USER:lvx", "DRW002:v1:tmp", "*")
	assert.Nil(t, err)
	assert.False(t, ok)

	rules := [][]string{
		{"USER:lvx", "DRW001:v1:tmp", "*", "allow"},
		{"USER:lvx", "DRW002:v1:tmp", "*", "allow"},
	}
	ok, err = e.AddPolicies(rules)
	assert.Nil(t, err)
	assert.True(t, ok)

	ok, err = e.Enforce("USER:lvx", "DRW001:v1:tmp", "*")
	assert.Nil(t, err)
	assert.True(t, ok)
	ok, err = e.Enforce("USER:lvx", "DRW002:v1:tmp", "*")
	assert.Nil(t, err)
	assert.True(t, ok)

	// test atomicity
	rules = [][]string{
		{"USER:lvx", "DRW001:v1:tmp", "*", "allow"},
		{"USER:lvx", "DRW002:v1:tmp", "*", "allow"},
		{"USER:lvx", "DRW003:v1:tmp", "*", "allow"},
	}
	ok, err = e.AddPolicies(rules)
	assert.Nil(t, err)
	assert.False(t, ok)
	ok, err = e.Enforce("USER:lvx", "DRW003:v1:tmp", "*")
	assert.Nil(t, err)
	assert.False(t, ok)
}

func TestAddPoliciesEx(t *testing.T) {
	// AddPoliciesEx adds authorization rules to the current policy.
	// If the rule already exists, the rule will not be added.
	// But unlike AddPolicies, other non-existent rules are added instead of returning false directly
	e := NewEnforcer()

	var ok bool
	var err error

	ok, err = e.Enforce("USER:lvx", "DRW001:v1:tmp", "*")
	assert.Nil(t, err)
	assert.False(t, ok)
	ok, err = e.Enforce("USER:lvx", "DRW002:v1:tmp", "*")
	assert.Nil(t, err)
	assert.False(t, ok)

	rules := [][]string{
		{"USER:lvx", "DRW001:v1:tmp", "*", "allow"},
		{"USER:lvx", "DRW002:v1:tmp", "*", "allow"},
	}
	ok, err = e.AddPoliciesEx(rules)
	assert.Nil(t, err)
	assert.True(t, ok)

	ok, err = e.Enforce("USER:lvx", "DRW001:v1:tmp", "*")
	assert.Nil(t, err)
	assert.True(t, ok)
	ok, err = e.Enforce("USER:lvx", "DRW002:v1:tmp", "*")
	assert.Nil(t, err)
	assert.True(t, ok)

	// test atomicity
	rules = [][]string{
		{"USER:lvx", "DRW001:v1:tmp", "*", "allow"},
		{"USER:lvx", "DRW002:v1:tmp", "*", "allow"},
		{"USER:lvx", "DRW003:v1:tmp", "*", "allow"},
	}
	ok, err = e.AddPoliciesEx(rules)
	assert.Nil(t, err)
	assert.True(t, ok)
	ok, err = e.Enforce("USER:lvx", "DRW003:v1:tmp", "*")
	assert.Nil(t, err)
	assert.True(t, ok)
}

func TestAddNamedPolicy(t *testing.T) {
	e := NewEnforcer()

	var ok bool
	var err error

	ok, err = e.AddNamedPolicy("p", "USER:lvx", "DRW001:v1:test", "*", "allow")
	assert.Nil(t, err)
	assert.False(t, ok)

	ok, err = e.Enforce("USER:lvx", "DRW011:v1:tmp", "*")
	assert.Nil(t, err)
	assert.False(t, ok)

	ok, err = e.AddNamedPolicy("p", "USER:lvx", "DRW011:v1:tmp", "*", "allow")
	assert.Nil(t, err)
	assert.True(t, ok)

	ok, err = e.Enforce("USER:lvx", "DRW011:v1:tmp", "*")
	assert.Nil(t, err)
	assert.True(t, ok)
}

func TestAddNamedPolicies(t *testing.T) {
	// AddNamedPolicies 向当前命名策略中添加授权规则。该操作本质上是原子的 因此，如果授权规则由不符合现行政策的规则组成，函数返回false，当前政策中没有添加任何政策规则。
	// 如果所有授权规则都符合政策规则，则函数返回true，每项政策规则都被添加到目前的政策中。
	e := NewEnforcer()

	var ok bool
	var err error

	ok, err = e.Enforce("USER:lvx", "DRW001:v1:tmp", "*")
	assert.Nil(t, err)
	assert.False(t, ok)
	ok, err = e.Enforce("USER:lvx", "DRW002:v1:tmp", "*")
	assert.Nil(t, err)
	assert.False(t, ok)

	rules := [][]string{
		{"USER:lvx", "DRW001:v1:tmp", "*", "allow"},
		{"USER:lvx", "DRW002:v1:tmp", "*", "allow"},
	}
	ok, err = e.AddNamedPolicies("p", rules)
	assert.Nil(t, err)
	assert.True(t, ok)

	ok, err = e.Enforce("USER:lvx", "DRW001:v1:tmp", "*")
	assert.Nil(t, err)
	assert.True(t, ok)
	ok, err = e.Enforce("USER:lvx", "DRW002:v1:tmp", "*")
	assert.Nil(t, err)
	assert.True(t, ok)

	// test atomicity
	rules = [][]string{
		{"USER:lvx", "DRW001:v1:tmp", "*", "allow"},
		{"USER:lvx", "DRW002:v1:tmp", "*", "allow"},
		{"USER:lvx", "DRW003:v1:tmp", "*", "allow"},
	}
	ok, err = e.AddNamedPolicies("p", rules)
	assert.Nil(t, err)
	assert.False(t, ok)
	ok, err = e.Enforce("USER:lvx", "DRW003:v1:tmp", "*")
	assert.Nil(t, err)
	assert.False(t, ok)
}

func TestAddNamedPoliciesEx(t *testing.T) {
	// AddNamedPoliciesEx adds authorization rules to the current named policy.
	// If the rule already exists, the rule will not be added.
	// But unlike AddNamedPolicies, other non-existent rules are added instead of returning false directly
	e := NewEnforcer()

	var ok bool
	var err error
	var rules [][]string

	ok, err = e.Enforce("USER:lvx", "DRW001:v1:tmp", "*")
	assert.Nil(t, err)
	assert.False(t, ok)
	ok, err = e.Enforce("USER:lvx", "DRW002:v1:tmp", "*")
	assert.Nil(t, err)
	assert.False(t, ok)

	rules = [][]string{
		{"USER:lvx", "DRW001:v1:tmp", "*", "allow"},
		{"USER:lvx", "DRW002:v1:tmp", "*", "allow"},
	}
	ok, err = e.AddNamedPoliciesEx("p", rules)
	assert.Nil(t, err)
	assert.True(t, ok)

	ok, err = e.Enforce("USER:lvx", "DRW001:v1:tmp", "*")
	assert.Nil(t, err)
	assert.True(t, ok)
	ok, err = e.Enforce("USER:lvx", "DRW002:v1:tmp", "*")
	assert.Nil(t, err)
	assert.True(t, ok)

	// test atomicity
	rules = [][]string{
		{"USER:lvx", "DRW001:v1:tmp", "*", "allow"},
		{"USER:lvx", "DRW002:v1:tmp", "*", "allow"},
		{"USER:lvx", "DRW003:v1:tmp", "*", "allow"},
	}
	ok, err = e.AddNamedPoliciesEx("p", rules)
	assert.Nil(t, err)
	assert.True(t, ok)
	ok, err = e.Enforce("USER:lvx", "DRW003:v1:tmp", "*")
	assert.Nil(t, err)
	assert.True(t, ok)
}

func TestSelfAddPoliciesEx(t *testing.T) {
	// SelfAddPoliciesEx adds authorization rules to the current named policy with autoNotifyWatcher disabled.
	// If the rule already exists, the rule will not be added.
	// But unlike SelfAddPolicies, other non-existent rules are added instead of returning false directly
	e := NewEnforcer()

	var ok bool
	var err error

	ok, err = e.Enforce("USER:lvx", "DRW001:v1:tmp", "*")
	assert.Nil(t, err)
	assert.False(t, ok)

	rules := [][]string{
		{"USER:lvx", "DRW001:v1:tmp", "*", "allow"},
	}

	ok, err = e.SelfAddPoliciesEx("p", "p", rules)
	assert.Nil(t, err)
	assert.True(t, ok)

	ok, err = e.Enforce("USER:lvx", "DRW001:v1:tmp", "*")
	assert.Nil(t, err)
	assert.True(t, ok)
}

func TestRemovePolicy(t *testing.T) {
	e := NewEnforcer()

	var ok bool
	var err error

	ok, err = e.Enforce("USER:lvx", "DRW001:v1:test", "*")
	assert.Nil(t, err)
	assert.True(t, ok)

	ok, err = e.RemovePolicy("USER:lvx", "DRW001:v1:test", "*", "allow")
	assert.Nil(t, err)
	assert.True(t, ok)

	ok, err = e.Enforce("USER:lvx", "DRW001:v1:test", "*")
	assert.Nil(t, err)
	assert.False(t, ok)

	ok, err = e.RemovePolicy("USER:lvx", "DRW001:v1:test", "*", "allow")
	assert.Nil(t, err)
	assert.False(t, ok)
}

func TestRemovePolicies(t *testing.T) {
	e := NewEnforcer()

	var ok bool
	var err error
	var boolArray []bool

	rules := [][]string{
		{"USER:lvx", "DRW001:v1:test", "*", "allow"},
		{"USER:lvx", "DRW002:v1:test", "r", "allow"},
		{"USER:lvx", "DRW004:v1:test", "w", "deny"},
	}

	ok, err = e.RemovePolicies(rules)
	assert.Nil(t, err)
	assert.True(t, ok)

	requests := [][]interface{}{
		{"USER:lvx", "DRW001:v1:test", "*"},
		{"USER:lvx", "DRW002:v1:test", "r"},
		{"USER:lvx", "DRW004:v1:test", "w"},
	}
	boolArray, err = e.BatchEnforce(requests)

	res := []bool{false, false, true}
	assert.Nil(t, err)
	assert.Equal(t, boolArray, res)
}

func TestRemoveFilteredPolicy(t *testing.T) {
	e := NewEnforcer()

	var ok bool
	var err error
	var boolArray []bool

	ok, err = e.RemoveFilteredPolicy(0, "USERGROUP:it")
	assert.Nil(t, err)
	assert.True(t, ok)

	requests := [][]interface{}{
		{"USER:lvx", "DRW011:v1:test", "*"},
		{"USER:lvx", "DRW012:v1:test", "r"},
		{"USER:lvx", "DRW013:v1:test", "w"},
		{"USER:lvx", "DRW001:v1:test", "w"},
	}
	boolArray, err = e.BatchEnforce(requests)

	res := []bool{false, false, false, true}
	assert.Nil(t, err)
	assert.Equal(t, boolArray, res)
}

func TestRemoveNamedPolicy(t *testing.T) {
	e := NewEnforcer()

	var ok bool
	var err error

	ok, err = e.RemoveNamedPolicy("p", "USER:lvx", "DRW001:v1:test", "*", "allow")
	assert.Nil(t, err)
	assert.True(t, ok)

	ok, err = e.Enforce("USER:lvx", "DRW001:v1:test", "*")
	assert.Nil(t, err)
	assert.False(t, ok)
}

func TestRemoveNamedPolicies(t *testing.T) {
	// RemoveNamedPolicies 从当前命名策略中删除授权规则。该操作本质上是原子的，因此，
	// 如果授权规则由不符合现行政策的规则组成，函数返回 false ，当前政策中没有任何政策规则被删除。
	// 如果所有授权规则都符合政策规则，则函数返回true，每项政策规则都从现行政策中删除。
	e := NewEnforcer()

	var ok bool
	var err error
	var boolArray []bool

	rules := [][]string{
		{"USER:lvx", "DRW001:v1:test", "*", "allow"},
		{"USER:lvx", "DRW002:v1:test", "r", "allow"},
		{"USER:lvx", "DRW003:v1:test", "w", "allow"},
	}

	ok, err = e.RemoveNamedPolicies("p", rules)
	assert.Nil(t, err)
	assert.True(t, ok)

	requests := [][]interface{}{
		{"USER:lvx", "DRW001:v1:test", "*"},
		{"USER:lvx", "DRW002:v1:test", "r"},
		{"USER:lvx", "DRW003:v1:test", "w"},
		{"USER:lvx", "DRW001:v1:test", "w"},
	}
	boolArray, err = e.BatchEnforce(requests)

	res := []bool{false, false, false, false}
	assert.Nil(t, err)
	assert.Equal(t, boolArray, res)
}

func TestRemoveFilteredNamedPolicy(t *testing.T) {
	e := NewEnforcer()

	var ok bool
	var err error

	ok, err = e.RemoveFilteredNamedPolicy("p", 0, "USERGROUP:it", "DRW014:v1:test")
	assert.True(t, ok)
	assert.Nil(t, err)

	ok, err = e.Enforce("USER:lvx", "DRW014:v1:test", "*")
	assert.False(t, ok)
	assert.Nil(t, err)

	ok, err = e.Enforce("USER:lvx", "DRW013:v1:test", "w")
	assert.True(t, ok)
	assert.Nil(t, err)
}

func TestHasGroupingPolicy(t *testing.T) {
	// HasGroupingPolicy 确定是否存在角色继承规则。
	e := NewEnforcer()

	var ok bool
	ok = e.HasGroupingPolicy("USER:lvx", "USERGROUP:eam")
	assert.True(t, ok)

	// 不支持链式传导
	ok = e.HasGroupingPolicy("USER:lvx", "SRCGROUP:eam")
	assert.False(t, ok)
}

func TestHasNamedGroupingPolicy(t *testing.T) {
	e := NewEnforcer()

	var ok bool
	ok = e.HasNamedGroupingPolicy("g", "USER:lvx", "USERGROUP:eam")
	assert.True(t, ok)

	ok = e.HasNamedGroupingPolicy("g", "USER:lvx", "SRCGROUP:eam")
	assert.False(t, ok)
}

func TestAddGroupingPolicy(t *testing.T) {
	// AddGroupingPolicy 向当前策略添加角色继承规则。如果规则已经存在，
	// 函数返回false，并且不会添加规则。否则，函数通过添加新规则并返回true。
	e := NewEnforcer()

	var ok bool
	var err error

	ok, err = e.AddGroupingPolicy("USER:lvx", "SRCGROUP:ev2")
	assert.Nil(t, err)
	assert.True(t, ok)

	ok, err = e.Enforce("USER:lvx", "EAM101:v2:test", "*")
	assert.Nil(t, err)
	assert.True(t, ok)
}

func TestAddGroupingPolicies(t *testing.T) {
	// AddGroupingPolicies adds role inheritance rules to the current policy.
	// If the rule already exists, the function returns false for the corresponding policy rule and the rule will not be added.
	// Otherwise the function returns true for the corresponding policy rule by adding the new rule.
	e := NewEnforcer()

	var ok bool
	var err error

	rules := [][]string{
		{"USER:lvx", "SRCGROUP:ev2"},
	}
	ok, err = e.AddGroupingPolicies(rules)
	assert.Nil(t, err)
	assert.True(t, ok)

	ok, err = e.Enforce("USER:lvx", "EAM101:v2:test", "*")
	assert.Nil(t, err)
	assert.True(t, ok)
}

func TestAddGroupingPoliciesEx(t *testing.T) {
	// AddGroupingPoliciesEx adds role inheritance rules to the current policy.
	// If the rule already exists, the rule will not be added.
	// But unlike AddGroupingPolicies, other non-existent rules are added instead of returning false directly.
	e := NewEnforcer()

	var ok bool
	var err error

	rules := [][]string{
		{"USER:lvx", "SRCGROUP:ev2"},
	}
	ok, err = e.AddGroupingPoliciesEx(rules)
	assert.Nil(t, err)
	assert.True(t, ok)

	ok, err = e.Enforce("USER:lvx", "EAM101:v2:test", "*")
	assert.Nil(t, err)
	assert.True(t, ok)
}

func TestAddNamedGroupingPolicy(t *testing.T) {
	e := NewEnforcer()

	var ok bool
	var err error

	ok, err = e.AddNamedGroupingPolicy("g", "USER:lvx", "SRCGROUP:ev2")
	assert.Nil(t, err)
	assert.True(t, ok)

	ok, err = e.Enforce("USER:lvx", "EAM101:v2:test", "*")
	assert.Nil(t, err)
	assert.True(t, ok)
}

func TestAddNamedGroupingPolicies(t *testing.T) {
	e := NewEnforcer()

	var ok bool
	var err error

	rules := [][]string{
		{"USER:lvx", "SRCGROUP:ev2"},
	}
	ok, err = e.AddNamedGroupingPolicies("g", rules)
	assert.Nil(t, err)
	assert.True(t, ok)

	ok, err = e.Enforce("USER:lvx", "EAM101:v2:test", "*")
	assert.Nil(t, err)
	assert.True(t, ok)
}

func TestAddNamedGroupingPoliciesEx(t *testing.T) {
	e := NewEnforcer()

	var ok bool
	var err error

	rules := [][]string{
		{"USER:lvx", "SRCGROUP:ev2"},
	}
	ok, err = e.AddNamedGroupingPoliciesEx("g", rules)
	assert.Nil(t, err)
	assert.True(t, ok)

	ok, err = e.Enforce("USER:lvx", "EAM101:v2:test", "*")
	assert.Nil(t, err)
	assert.True(t, ok)
}

func TestRemoveGroupingPolicy(t *testing.T) {
	e := NewEnforcer()

	var ok bool
	var err error

	ok, err = e.RemoveGroupingPolicy("USER:lvx", "USERGROUP:it")
	assert.Nil(t, err)
	assert.True(t, ok)

	ok, err = e.Enforce("USER:lvx", "DRW011:v1:test", "*")
	assert.Nil(t, err)
	assert.False(t, ok)
}

func TestRemoveGroupingPolicies(t *testing.T) {
	e := NewEnforcer()

	var ok bool
	var err error
	rules := [][]string{
		{"USER:lvx", "USERGROUP:it"},
	}

	ok, err = e.RemoveGroupingPolicies(rules)
	assert.Nil(t, err)
	assert.True(t, ok)

	ok, err = e.Enforce("USER:lvx", "DRW011:v1:test", "*")
	assert.Nil(t, err)
	assert.False(t, ok)
}

func TestRemoveFilteredGroupingPolicy(t *testing.T) {
	e := NewEnforcer()

	var ok bool
	var err error
	ok, err = e.RemoveFilteredGroupingPolicy(0, "USER:lvx")
	assert.Nil(t, err)
	assert.True(t, ok)

	ok, err = e.Enforce("USER:lvx", "DRW101:v1:test", "*")
	assert.Nil(t, err)
	assert.False(t, ok)

	ok, err = e.Enforce("USER:lvx", "DRW001:v1:test", "*")
	assert.Nil(t, err)
	assert.True(t, ok)
}

func TestRemoveNamedGroupingPolicy(t *testing.T) {
	e := NewEnforcer()

	var ok bool
	var err error
	ok, err = e.RemoveNamedGroupingPolicy("g", "USER:lvx", "USERGROUP:it")
	assert.Nil(t, err)
	assert.True(t, ok)

	ok, err = e.Enforce("USER:lvx", "DRW011:v1:test", "*")
	assert.Nil(t, err)
	assert.False(t, ok)

	ok, err = e.Enforce("USER:lvx", "DRW001:v1:test", "*")
	assert.Nil(t, err)
	assert.True(t, ok)
}

func TestRemoveNamedGroupingPolicies(t *testing.T) {
	e := NewEnforcer()

	var ok bool
	var err error
	rules := [][]string{
		{"USER:lvx", "USERGROUP:it"},
	}
	ok, err = e.RemoveNamedGroupingPolicies("g", rules)
	assert.Nil(t, err)
	assert.True(t, ok)

	ok, err = e.Enforce("USER:lvx", "DRW011:v1:test", "*")
	assert.Nil(t, err)
	assert.False(t, ok)

	ok, err = e.Enforce("USER:lvx", "DRW001:v1:test", "*")
	assert.Nil(t, err)
	assert.True(t, ok)
}

func TestRemoveFilteredNamedGroupingPolicy(t *testing.T) {
	e := NewEnforcer()

	var ok bool
	var err error

	ok, err = e.RemoveFilteredNamedGroupingPolicy("g", 0, "USER:lvx")
	assert.Nil(t, err)
	assert.True(t, ok)

	ok, err = e.Enforce("USER:lvx", "DRW011:v1:test", "*")
	assert.Nil(t, err)
	assert.False(t, ok)

	ok, err = e.Enforce("USER:lvx", "DRW001:v1:test", "*")
	assert.Nil(t, err)
	assert.True(t, ok)
}

func TestUpdatePolicy(t *testing.T) {
	e := NewEnforcer()

	var ok bool
	var err error

	oldpolicy := []string{"USER:lvx", "DRW001:v1:test", "*", "allow"}
	newpolicy := []string{"USER:lvx", "DRW001:v1:test", "r", "allow"}
	e.UpdatePolicy(oldpolicy, newpolicy)

	ok, err = e.Enforce("USER:lvx", "DRW001:v1:test", "*")
	assert.Nil(t, err)
	assert.False(t, ok)

	ok, err = e.Enforce("USER:lvx", "DRW001:v1:test", "r")
	assert.Nil(t, err)
	assert.True(t, ok)
}

func TestUpdatePolicies(t *testing.T) {
	e := NewEnforcer()

	var ok bool
	var err error

	oldpolicies := [][]string{{"USER:lvx", "DRW001:v1:test", "*", "allow"}}
	newpolicies := [][]string{{"USER:lvx", "DRW001:v1:test", "r", "allow"}}
	e.UpdatePolicies(oldpolicies, newpolicies)

	ok, err = e.Enforce("USER:lvx", "DRW001:v1:test", "*")
	assert.Nil(t, err)
	assert.False(t, ok)

	ok, err = e.Enforce("USER:lvx", "DRW001:v1:test", "r")
	assert.Nil(t, err)
	assert.True(t, ok)
}

func TestAddFunction(t *testing.T) {
	e := NewEnforcer()

	customfunc := func(key1 string, key2 string) bool {
		return key1 == "1" && key2 == "2"
	}

	customfuncwrapper := func(args ...interface{}) (interface{}, error) {
		key1 := args[0].(string)
		key2 := args[1].(string)

		return bool(customfunc(key1, key2)), nil
	}

	e.AddFunction("my_func", customfuncwrapper)
	matcher := "my_func(r.sub, r.obj)"
	ok, err := e.EnforceWithMatcher(matcher, "1", "2", "*")
	assert.Nil(t, err)
	assert.True(t, ok)
}

func TestUpdateNamedGroupingPolicy(t *testing.T) {
	e := NewEnforcer()

	var ok bool
	var err error

	oldRule := []string{"USER:lvx", "SRCGROUP:drw"}
	newRule := []string{"USER:lvx", "SRCGROUP:ev2"}
	ok, err = e.UpdateNamedGroupingPolicy("g", oldRule, newRule)
	assert.Nil(t, err)
	assert.True(t, ok)

	ok, err = e.Enforce("USER:lvx", "DRW101:v1:test", "*")
	assert.Nil(t, err)
	assert.False(t, ok)

	ok, err = e.Enforce("USER:lvx", "EAM101:v2:test", "*")
	assert.Nil(t, err)
	assert.True(t, ok)
}
