package main

import (
	"log"
	"sort"
)

func main() {
	var mySlice []string

	mySlice = append(mySlice, "Rajesh")
	mySlice = append(mySlice, "Sushrith")
	mySlice = append(mySlice, "Anjali")

	log.Println(mySlice)
	sort.Strings(mySlice)
	log.Println(mySlice)

	myNumbers := []int{1, 2, 3, 4, 6, 5, 7, 8, 9, 10}

	log.Println(myNumbers)

	sort.Ints(myNumbers)

	log.Println(myNumbers)

}
