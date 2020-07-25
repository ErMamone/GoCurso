package main

import "fmt"

var estado bool

func main() {
	estado = true

	if estado = false; estado == true {
		fmt.Println(estado)
	} else {
		fmt.Println(estado)
	}

	//PARTE DEL SWITCH HEEEEEEEEEEEEEREEEEEEEEEEE

	switch numero := 5; numero {
	case 1:
		fmt.Println(numero)
	case 2:
		fmt.Println(numero)
	case 3:
		fmt.Println(numero)
	case 4:
		fmt.Println(numero)
	case 5:
		fmt.Println(numero)
	default:
		fmt.Println(numero)
	}

}
