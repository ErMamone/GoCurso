package main

import (
	"fmt"
	"strconv"
)

var numero int
var texto string
var boleano bool
var Status bool

func main() {
	var numero2 int
	numero2 = 1
	numero3 := 4
	numero3 = 15

	num1, num2, num3, txt := 4, 2, 67, "oleputo"
	Status = false
	var moneda int64 = 120

	num2 = +int(moneda)

	//Paquete para convertir tipos de datos STRCONV
	txt = strconv.Itoa(int(moneda))

	fmt.Println(numero2 + numero3)
	fmt.Println(num1)
	fmt.Println(num2)
	fmt.Println(num3)
	fmt.Println(txt)
	MostrarStatus()
}

func MostrarStatus() {
	fmt.Println(Status)
}
