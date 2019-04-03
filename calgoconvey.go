package main

import (
	"errors"
	"fmt"
)

func Add(x int, y int) (z int) {
	z = x + y
	return z
}

func Minus(x int, y int) (z int) {
	z = x - y
	return z
}
func Multiplies(x int, y int) (z int) {
	z = x * y
	return z
}

func DividedBy(x int, y int) (z float64, err error) {
	z = float64(x) / float64(y)
	if y == 0 {
		return 0, errors.New("** y != 0")
	}

	return z, nil
}

func main() {
	add := Add(5, 2)
	fmt.Println("add(+) = ", add)
	minus := Minus(5, 2)
	fmt.Println("minus(-) = ", minus)
	multiplies := Multiplies(5, 2)
	fmt.Println("multiplies(*) = ", multiplies)
	dividedBy, _ := DividedBy(5, 3)
	fmt.Println("dividedBy(/) = ", dividedBy)
	_, err := DividedBy(5, 0)
	fmt.Println("dividedBy(/) = ", err)
}
