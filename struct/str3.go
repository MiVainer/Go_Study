package main

import (
	"fmt"

)

type user struct {
    name  string
    karma struct {
        value int
        title string
    }
}

type comment struct {
    text   string
    author *user
}


func main() {
	
	u := user{
    name: "Chris",
    karma: struct {
        value int
        title string
    }{
        value: 100,
        title: "^-^",
    },
}
fmt.Printf("%+v\n", u.karma.value)

//-----------------------------------------------

	chris := user{
		name: "Chris",
	}
	c := comment{
		text:   "Gophers are awesome!",
		author: &chris,
	}
	fmt.Printf("%+v\n", c)

}
