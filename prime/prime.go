package main

import "fmt"

func main() {
	num := 4
	isPrime := true
	for i := 2; i < num; i++ {
		if num%i == 0 {
			fmt.Println("Number is not prime")
			isPrime = false
			break
		}

	}

	if isPrime {
		fmt.Println("Number is prime")
	}

}
