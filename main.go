package main

import "github.com/pjgg/Go-Option/option"
import "fmt"
import "strings"

//Example of usage
func main() {
	userA := new(user)
	userA.age = 29
	userA.name = "pablo"

	userAOption := option.Of(userA)
	fmt.Println(userAOption.IsPresent())
	fmt.Println(userAOption.Get())

	userB := new(user)
	userB.age = 49
	userB.name = "pablo"

	userBOption := option.Of(userB)
	fmt.Println(userBOption.IsPresent())
	fmt.Println(userBOption.Get())
	fmt.Println(userBOption.Filter(exampleUserPredicateEqualUser(*userA)))
}

type user struct {
	name string
	age  int
}

func exampleUserPredicateEqualUser(userA user) option.Predicate {
	return func(value interface{}) bool {
		var result bool
		if userB, ok := value.(*user); ok == true {
			fmt.Println(strings.Compare(userA.name, userB.name))
			result = userA.name == userB.name
		} else {
			result = false
		}
		return result
	}
}
