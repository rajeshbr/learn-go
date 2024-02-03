package main

import "log"

func main() {

	for i := 0; i <= 10; i++ {
		log.Println("i = ", i)
	}

	animals := []string{"cow", "horse", "cat"}

	for i, animal := range animals {
		log.Println(i, " ", animal)
	}

	for _, animal := range animals {
		log.Println(animal)
	}

	animals2 := make(map[string]string)
	animals2["dog"] = "fido"
	animals2["cat"] = "garfield"
	animals2["horse"] = "dexter"

	for animalType, animal := range animals2 {
		log.Println(animalType, animal)
	}

	var myline = "Once upon a time there was a guy"

	for i, l := range myline {
		log.Println(i, l)
	}

	type User struct {
		FirstName   string
		LastName    string
		PhoneNumber string
		Age         int
	}

	var myUsers []User

	myUsers = append(myUsers, User{"Rajesh", "Bathula", "9999999999", 42})
	myUsers = append(myUsers, User{"Sushrith", "Bathula", "8888888888", 8})
	myUsers = append(myUsers, User{"Anjali", "Mulamreddy", "7777777777", 32})

	for _, user := range myUsers {
		log.Println(user.FirstName, user.LastName, user.PhoneNumber, user.Age)
	}

}
