package main

import (
	"fmt"
)

type person struct {
	firstName string
	lastName  string
}

//func (p person) String() string {
//	return fmt.Sprintf("The person name is: %v %v", p.firstName, p.lastName)
//}

func (p person) printIt() {
	fmt.Printf("Person: %+v\n", p)
}

func (p *person) updateIt(firstName string)  {
	(*p).firstName = firstName
}

func main() {
	person1 := person{"Alex", "Anderson"}
	person1.updateIt("Mariana")
	person1.printIt()
}
