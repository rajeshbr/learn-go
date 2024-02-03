package main

import "log"

func main() {

	var isTrue bool
	isTrue = false

	if isTrue {
		log.Println("isTrue is ", isTrue)
	} else {
		log.Println("isTrue is ", isTrue)
	}

	cat := "cat"

	if cat == "cat" {
		log.Println("cat  is cat")
	} else {
		log.Println("cat  is not cat")
	}

	myVar := "dog"

	switch myVar {
	case "cat":
		log.Println("myVar is a cat")
	case "dog":
		log.Println("myVar is a dog")
	case "rooster":
		log.Println("myVar is a rooster")
	default:
		log.Println("myVar is something else")
	}
}
