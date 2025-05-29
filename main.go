package main

import (
	"fmt"
)

func main() {
	var a, b float32
	c := " "
	fmt.Println("Напишите число, операцию и число")
	fmt.Scan(&a, &c, &b)
	if c == "/" && b == 0 {
		fmt.Println("Делить на ноль нельзя!")
	} else {
		fmt.Println("Ответ:")
		switch c {
		case "+":
			fmt.Println(a + b)
		case "-":
			fmt.Println(a - b)
		case "*":
			fmt.Println(a * b)
		case "/":
			fmt.Println(a / b)
		case "%":
			fmt.Println((int)(a) % (int)(b))
		}
	}
}
