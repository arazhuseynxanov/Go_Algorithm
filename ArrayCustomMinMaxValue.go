package main

import (
	"errors"
	"fmt"
)

func main() {
	array := []int{5, 7, -1, 1, 8, 111111}
	fmt.Scan(&array)
	fmt.Print("Minimum:")
	eded, err := Min(array)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(eded)
	}
	fmt.Print("Max:")
	eded, errr := Max(array)
	if errr != nil {
		fmt.Println(errr)
	} else {
		fmt.Println(eded)
	}

}
func Min(arr []int) (int, error) {
	if len(arr) == 0 {
		return 0, errors.New("Array Bosdur")
	}
	var min = arr[0]
	for _, value := range arr {
		if value < min {
			min = value
		}
	}
	return min, nil
}
func Max(arr []int) (int, error) {
	if len(arr) == 0 {
		return 0, errors.New("Array Bosdur")
	}
	var max = arr[0]
	for _, value := range arr {
		if max < value {
			max = value
		}
	}
	return max, nil
}
