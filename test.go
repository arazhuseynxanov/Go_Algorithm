package main

import (
	"fmt"
	"strings"
)

func caesarCipherShiftRune(r rune, shift uint) rune {
	s := rune(shift % 26)

	if s == 0 {
		return r
	}

	return (((r - 'A') + s) % 26) + 'A'
}

func caesarCipherEncrypt(value string, shift uint) string {
	var builder strings.Builder
	value = strings.ToUpper(value)

	for _, r := range value {
		if r >= 'A' && r <= 'Z' {
			r = caesarCipherShiftRune(r, shift)
		}

		builder.WriteRune(r)
	}

	return builder.String()
}

func caesarCipherDecrypt(value string, shift uint) string {
	return caesarCipherEncrypt(value, 26-(shift%26))
}

func main() {
	const encryptShift = 18
	const input = "THIS IS A TEST 123"

	encrypted := caesarCipherEncrypt(input, encryptShift)
	decrypted := caesarCipherDecrypt(encrypted, encryptShift)

	fmt.Printf(
		"INPUT: %s\nENCRYPTED: %s\nDECRYPTED: %s\nINPUT EQUALS DECRYPTED: %t\n",
		input, encrypted, decrypted, input == decrypted,
	)
}
