package main

import (
	"log"

	"github.com/rajeshbr/learn-go/helpers"
)

const numPool = 100

func CalculateValue(intChan chan int) {
	randomNumber := helpers.RandomNumber(numPool)
	intChan <- randomNumber
}

func main() {
	log.Println("This is about channels")

	intChan := make(chan int)
	defer close(intChan)

	go CalculateValue(intChan)

	num := <-intChan
	log.Println(num)
}
