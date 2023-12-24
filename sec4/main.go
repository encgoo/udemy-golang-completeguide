package main

import (
	"fmt"
	"udemy/sec4/person"
)

func main() {
	person := person.MakePerson("Alex", "Anderson", "aa@example.com", 91210)
	//	pP := &person
	//	(*pP).SetEmail("aa@hotmail.com")

	// pointer shortcut
	person.SetEmail("aa@gmail.com")

	fmt.Printf("%+v", person)
}
