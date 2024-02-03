package helpers

import "math/rand"

type User struct {
	FirstName   string
	LastName    string
	PhoneNumber string
	Age         int
}

func RandomNumber(n int) int {
	value := rand.Intn(n)
	return value
}
