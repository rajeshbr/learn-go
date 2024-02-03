package main

import (
	"errors"
	"log"
)

func main() {

	var a int = 1
	b := 0
	log.Println("a + b = ", add(a, b))

	var c int
	var err error
	c, err = divide(a, b)

	if err != nil {
		log.Println(err)
	} else {
		log.Println("a / b = ", c)
	}

	var name string = "Rajesh"

	log.Println("Hello ", name)
}

func add(a int, b int) int {
	return a + b
}

func divide(a int, b int) (int, error) {
	if b == 0 {
		return 0, errors.New("division by zero")
	}

	return a / b, nil
}
