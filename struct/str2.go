package main

import (
	"fmt"

)

type person struct {
    firstName string
    lastName  string
}

type book struct {
    title  string
    author person
}


func main() {
	b := book{
    title: "The Majik Gopher",
    author: person{
        firstName: "Christopher",
        lastName:  "Swanson",
    },
}
fmt.Println(b.author.firstName)

}
