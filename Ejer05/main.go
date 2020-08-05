package main

import "fmt"

func main() {
	/*
		fmt.Println(uno(3))

		numero, estado := dos(1)

		fmt.Println(numero)
		fmt.Println(estado)
	*/
	fmt.Println(Calculo(5, 5, 5, 5, 2, 2, 2, 2, 2))
	fmt.Println(Calculo(3))
	fmt.Println(Calculo(27, 73, 1))
	fmt.Println(Calculo(1, 1, 1, 1, 1, 1, 1))
}

func uno(numero int) (z int) {
	z = numero * 2
	return
}

func dos(numero int) (int, bool) {
	if numero == 1 {
		return 5, false
	} else {
		return 10, true
	}
}

func Calculo(numero ...int) int {
	total := 0
	for item, num := range numero {
		total = total + num
		fmt.Printf("\nItem %d: \ny suma: %d\t", item, total)
	}
	return total
}
