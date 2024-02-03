package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type User struct {
	FirstName   string `json:"first_name"`
	LastName    string `json:"last_name"`
	PhoneNumber string `json:"phone_number"`
	Age         int    `json:"age"`
}

func main() {

	myJson := `
	
	[
		{
			"first_name":"rajesh",
			"last_name":"bathula",
			"phone_number":"9999999999",
			"age":42
		},
		{
			"first_name":"sushrith",
			"last_name":"bathula",
			"phone_number":"9999999999",
			"age":8
		}
	]

	`

	var unmarshalled []User

	err := json.Unmarshal([]byte(myJson), &unmarshalled)
	if err != nil {
		log.Println("Error unmarshalling ", err)
	} else {
		log.Printf("unmarshalled: %v", unmarshalled)
	}

	newJson, err := json.MarshalIndent(&unmarshalled, "", "  ")

	if err != nil {
		log.Println("Error marshalling", err)
	} else {
		log.Println(newJson)
		fmt.Println(string(newJson))
	}

}
