package main

import (
	"fmt"

	"github.com/casbin/casbin/v2"
)

func Enforce(e *casbin.Enforcer) {
	ok, err := e.Enforce("USER:1416962189826199552", "EAM011:v1:ip:prod", "*")
	if err != nil {
		panic(err)
	}

	fmt.Println(ok)
}

// func EnforceWithMatcher(e *casbin.Enforcer) {
// 	ok, err := e.EnforceWithMatcher(matcher, request)
// }

func main() {
	e, err := casbin.NewEnforcer("./model.conf", "./policy.csv")
	if err != nil {
		panic(err)
	}

	Enforce(e)
}
