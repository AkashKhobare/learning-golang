package main

import "fmt"

type user struct {
	username, email string
}

func (user *user) updateEmail(email string) {
	user.email = email
}

func (user user) printUserDetails() {
	fmt.Printf("\nusername: %v, email: %v\n", user.username, user.email)
}

func main() {
	user := user{
		username: "akashk",
		email:    "akashk@gmail.com",
	}
	user.printUserDetails()

	// update email of user
	user.updateEmail("akashkv2@gmail.com")
	user.printUserDetails()
}


