package main

import (
	"errors"
	"fmt"
	"math/rand"
	"time"
)

func main() {
	var numberInput int
	fmt.Print("Ededi Daxil Et:")
	fmt.Scan(&numberInput)
	app, err := Random(numberInput)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(app)
	}

}
func Random(number int) (any, error) {
	arr := []int{number}
	var arr2 []int
	result := number
	if number <= 0 {
		return 0, errors.New("Sifir ve ya Menfi Eded Ola Bilmez")
	} else {
		for i := 1; i < number; i++ {
			result--
			arr2 = append(arr2, result)
		}
		arr = append(arr, arr2...)

		rand.Seed(time.Now().UnixNano())
		for i := len(arr) - 1; i > 0; i-- {
			j := rand.Intn(i + 1)
			arr[i], arr[j] = arr[j], arr[i]
		}
		return arr, nil
	}

}
