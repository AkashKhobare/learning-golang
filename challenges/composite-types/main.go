package main

import "fmt"

type user struct {
	username, email string
}

func main() {
	user := user{
		username: "akashk",
		email:    "akashk@gmail.com",
	}

	fmt.Printf("username: %v, email: %v\n", user.username, user.email)
}
