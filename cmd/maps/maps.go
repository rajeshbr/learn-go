package main

import "log"

func main() {

	myMap := make(map[string]string)

	myMap["dog"] = "Samson"
	myMap["cat"] = "Cassie"
	myMap["hamster"] = "Derrik"

	log.Println(myMap)

	changeCat(myMap)

	log.Println(myMap)

}

func changeCat(myMap map[string]string) {
	myMap["cat"] = "Kitty"
}
