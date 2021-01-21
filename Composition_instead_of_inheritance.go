package main

import (
	"fmt"
)

type author struct {
	firstName string
	lastName  string
	bio       string
}

func (x author) fullName() string {
	return fmt.Sprintf("%s %s", x.firstName, x.lastName)
}

type post struct {
	title   string
	content string
	author
}

func (y post) details() {
	fmt.Println("Title: ", y.title)
	fmt.Println("Content: ", y.content)
	fmt.Println("Author: ", y.fullName())
	fmt.Println("Bio: ", y.bio)
}

func main() {
	author1 := author{
		"Sakib",
		"Seaum",
		"Golang Enthusiast",
	}
	post1 := post{
		"Inheritance in Go",
		"Go supports composition instead of inheritance",
		author1,
	}
	post1.details()
}

