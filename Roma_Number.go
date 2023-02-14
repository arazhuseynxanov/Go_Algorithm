package main

import (
	"fmt"
	"strings"
)

func main() {

	var NumOne string
	var NumTwo string
	fmt.Print("Birinci Ededi Daxil Et:")
	fmt.Scan(&NumOne)
	fmt.Print("Ikinci Ededi Daxil Et:")
	fmt.Scan(&NumTwo)
	sum := RomaNumStringToInt(NumOne) + RomaNumStringToInt(NumTwo)
	fmt.Println(RomaNumIntToString(sum))

}

var rumNum = map[string]int{
	"I":  1,
	"IV": 4,
	"V":  5,
	"IX": 9,
	"X":  10,
	"XL": 40,
	"L":  50,
	"XC": 90,
	"C":  100,
	"CD": 400,
	"D":  500,
	"CM": 900,
	"M":  1000,
}

func RomaNumStringToInt(s string) int {
	sum := 0
	result := 0
	for i := len(s) - 1; i >= 0; i-- {
		letter := s[i]
		num := rumNum[string(letter)]
		if num >= result {
			result = num
			sum = sum + num
			continue
		}
		sum = sum - num
	}
	return sum
}
func RomaNumIntToString(number int) string {
	myStruct := []struct {
		key   string
		value int
	}{
		{"M", 1000},
		{"CM", 900},
		{"D", 500},
		{"CD", 400},
		{"C", 100},
		{"XC", 90},
		{"L", 50},
		{"XL", 40},
		{"X", 10},
		{"IX", 9},
		{"V", 5},
		{"IV", 4},
		{"I", 1},
	}
	var roma strings.Builder
	for _, str := range myStruct {
		for number >= str.value {
			roma.WriteString(str.key)
			number = number - str.value
		}
	}
	return roma.String()
}
