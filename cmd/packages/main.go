package main

import (
	"log"

	"github.com/rajeshbr/learn-go/helpers"
)

func main() {
	var myUser helpers.User

	myUser = helpers.User{
		FirstName: "Rajesh",
		LastName:  "Bathula",
	}

	log.Println(myUser)
}
