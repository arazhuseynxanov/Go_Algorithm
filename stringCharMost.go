package main

import (
	"errors"
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
	app, err := Max(Chars)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(app)
	}
}
func Max(arr []Char) (any, error) {
	if len(arr) == 0 {
		return 0, errors.New("Array Bosdur...")
	}
	var max = arr[0]
	for _, v := range arr {
		if v.Value > max.Value {
			max = v
		}
	}
	return max, nil
}
