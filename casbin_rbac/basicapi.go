package casbin_rbac

import (
	"fmt"

	"github.com/casbin/casbin/v2"
)

func NewEnforcer(policies [][]string) *casbin.Enforcer {

	e, err := casbin.NewEnforcer("./model.conf", "./empty.csv")
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
