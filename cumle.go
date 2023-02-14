package main

import (
	"fmt"
	"strings"
)

type Char struct {
	Key   string `json:"key"`
	Value int    `json:"value"`
}

func main() {
	var str string
	fmt.Print("Cumle Daxil Et:")
	fmt.Scan(&str)
	ln := len(str)
	var Chars []Char
	for i := 0; i < ln; i++ {
		Chars = append(Chars, Char{
			Key:   string(str[i]),
			Value: strings.Count(str, string(str[i])),
		},
		)
	}
	fmt.Println(Chars)
}
