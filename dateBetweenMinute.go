package main

import (
	"fmt"
	"time"
)

func main() {
	var dateString string
	var dateString2 string

	fmt.Print("Tarix Daxil Edin:")
	fmt.Scan(&dateString)
	fmt.Print("Tarix Daxil Edin:")
	fmt.Scan(&dateString2)
	date, error := time.Parse("_2-1-2006", dateString)

	if error != nil {
		fmt.Println(error)
		return
	}
	date2, err := time.Parse("_2-1-2006", dateString2)

	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("Value of date: %v", date)
	fmt.Println()
	fmt.Printf("Value of date: %v", date2)
	fmt.Println()
	fmt.Printf("The Minute difference is: %f", date2.Sub(date).Minutes())

}
