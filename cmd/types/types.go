package main

import (
	"log"
	"time"
)

type User struct {
	FirstName   string
	LastName    string
	PhoneNumber string
	Age         int
	BirthDate   time.Time
}

func (m *User) greetMe() {
	log.Println("Hello ", m.FirstName)
}

func main() {
	var user User = User{
		FirstName:   "Rajesh Reddy",
		LastName:    "Bathula",
		PhoneNumber: "+91 9731133881",
	}

	user2 := User{
		FirstName: "Sushrith Reddy",
		LastName:  "Bathula",
	}

	log.Println(user)
	log.Println(user2)
	user.greetMe()
	user2.greetMe()
}
