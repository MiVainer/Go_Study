package main

import (
	"fmt"

)

type person struct {
    name string
    age  int
}

func newPerson(name string) *person {
    p := person{name: name}
    p.age = 42
    return &p
}

func makePerson(name string) person {
    p := person{name: name}
    p.age = 42
    return p
}

func main() {
	
	bob := person{"Bob", 20}
	fmt.Println(bob)
	//---------------------------------------
	alice := person{name: "Alice", age: 30}
	fmt.Println(alice)
	//---------------------------------------
	fred := person{name: "Fred"}
	fmt.Println(fred)
	//---------------------------------------
	annptr := &person{name: "Ann", age: 40}
	fmt.Println(annptr)

	//---------------------------------------
	john := newPerson("John")
	fmt.Println(john)

	//---------------------------------------
	sean := person{name: "Sean", age: 50}
	fmt.Println(sean.age)
	//---------------------------------------
	sven := &person{name: "Sven", age: 50}
	fmt.Println((*sven).age)
	fmt.Println(sven.age)
	sven.age = 51
	fmt.Println(sven.age)
}
