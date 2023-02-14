package main

import (
	"fmt"
)

func main() {
	var cumle string = "   Salam  Araz    "

	fmt.Println(Trim(cumle))

}
func Trim(sentence string) string {
	_first := First(sentence)
	_end := End(_first)
	_middle := Middle(_end)
	return _middle

}
func First(sentence string) string {
	index := 0
	for i, value := range sentence {
		if string(value) != " " {
			index = i
			break
		}
	}
	return sentence[index:]
}
func Middle(sentence string) string {
	startIndex := len(sentence)
	for i, value := range sentence {
		if string(sentence[i]) != " " {
			sentence += string(value)
		} else if string(sentence[i-1]) != " " {
			sentence += " "
		}
	}

	return sentence[startIndex:]
}
func End(sentence string) string {
	index := 0
	for i := len(sentence) - 1; i >= 0; i-- {
		if string(sentence[i]) != " " {
			index = i
			break
		}
	}
	return sentence[:index+1]
}
