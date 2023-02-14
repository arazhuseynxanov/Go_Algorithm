// Min methodu yaradılır. Daxil olunan arrayın ən kiçik elementini tapır.
package main

import (
	"errors"
	"fmt"
)

func main() {
	myArr := []int{3, 2, 1}
	app, err := Min(myArr)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(app)
	}
}
func Min(arr []int) (int, error) {
	if len(arr) == 0 {
		return 0, errors.New("Array Bosdur...")
	}
	var min = arr[0]
	for _, v := range arr {
		if v < min {
			min = v
		}
	}
	return min, nil
}
