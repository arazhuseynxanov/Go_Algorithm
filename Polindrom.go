package main

import "fmt"

func main() {
	fmt.Println(01)
	var num, a, b, result int
	fmt.Print("Enter The Number:")
	fmt.Scan(&num)
	b = num
	for {
		a = num % 10
		result = result*10 + a
		num = num / 10
		if num == 0 {
			break
		}
	}
	if b == result {
		fmt.Println("Polindrom")
	} else {
		fmt.Println("Polindrom Deyil")
	}

}

/*
mes 121
1
2
1

*/
