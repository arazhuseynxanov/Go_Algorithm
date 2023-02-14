package main

import (
	"fmt"
	"regexp"
	"strings"
)

func main() {
	var gmail string
	var password string
	fmt.Print("Gmaili Daxil Edin:")
	fmt.Scan(&gmail)
	fmt.Print("Password Daxil Edin:")
	fmt.Scan(&password)
	adminGmail := "admin@gmail.com"
	const adminPassword = "ADMIN123"

	var output string
	for _, r := range adminPassword {
		if r >= 'A' && r <= 'Z' {
			r = caesarCipherShiftRune(r, 13)
		}
		output += string(r)
	}
	password = output

	if isEmailValid(gmail) == true {
		if gmail == adminGmail && strings.ToUpper(password) == output {
			fmt.Println("Xos Gelmisiniz")
		} else if gmail != adminGmail && password == adminPassword {
			fmt.Println("Istifadeci Ve Ya Password Sehvdir")
		} else {
			fmt.Println("Istifadeci Ve Ya Password Sehvdir")
		}
	} else {
		fmt.Println("Validation Sehvdir....")
	}
}
func isEmailValid(e string) bool {
	emailRegex := regexp.MustCompile("^[a-zA-Z0-9.!#$%&'*+/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$")
	return emailRegex.MatchString(e)
}
func caesarCipherShiftRune(r rune, shift uint) rune {
	s := rune(shift % 26)

	if s == 0 {
		return r
	}
	return (((r - 'A') + s) % 26) + 'A'
}
