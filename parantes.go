package main

import (
	"fmt"
	"strings"
)

func main() {
	//Birinci Üsul
	//simvol := "[]{}()"

	var result string
	fmt.Println("Daxil Et:")
	fmt.Scan(&result)
	//if len(result) == 6 {
	//	result = strings.Replace(result, "()", "", 1)
	//	result = strings.Replace(result, "[]", "", 1)
	//	result = strings.Replace(result, "{}", "", 1)
	//	if len(result) == 0 {
	//		fmt.Println(true)
	//	}
	//} else {
	//	fmt.Println(false)
	//}
	//Ikinci Üsul
	duz := "[]"
	eyri := "{}"
	adi := "()"
	if len(result) == 6 {
		if strings.Contains(result, duz) && strings.Contains(result, eyri) && strings.Contains(result, adi) {
			fmt.Println(true)
		} else {
			fmt.Println(false)
		}
	} else {
		fmt.Println(false)
	}

}
