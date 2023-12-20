package main

import (
	"fmt"
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

func TestEnforce(t *testing.T) {
	e := NewEnforcer()

	ok, err := e.Enforce("USER:1416962189826199552", "EAM011:v1:ip:prod", "*")
	assert.Nil(t, err)
	assert.True(t, ok)

}

func EnforceWithMatcher(e *casbin.Enforcer) {
	//If the 'act' is 'w', authentication is granted.
	matcher := "r.act == 'w'"
	ok, err := e.EnforceWithMatcher(matcher, "USER:1416962189826199552", "EAM011:v1:ip:prod", "*")
	if err != nil {
		panic(err)
	}
	fmt.Println(ok)

	ok, err = e.EnforceWithMatcher(matcher, "unknown", "unknown", "w")
	if err != nil {
		panic(err)
	}

	fmt.Println(ok)
}

// func main() {
// 	e, err := casbin.NewEnforcer("./model.conf", "./policy.csv")
// 	if err != nil {
// 		panic(err)
// 	}

// 	Enforce(e)
// 	EnforceWithMatcher(e)
// }
