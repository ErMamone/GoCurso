package main

import "fmt"

/*
Son interceptores que permiten ejecutar instrucciones comunes
a varias funciones que reciben y devuelven
los mismos tipos variables

Los middleware se usan para: agregarle a todas las rutinas, funciones de desarrollos grandes
funcionalidad, seguridad, encriptacion, hace todos los capas
*/
var result int

func main() {
	fmt.Println("Inicio")

	result = operacionesMidd(sumar)(2, 3)
	fmt.Println(result)

	result = operacionesMidd(restar)(10, 6)
	fmt.Println(result)

	result = operacionesMidd(multiplicar)(2, 4)
	fmt.Println(result)
}

func sumar(a, b int) int {
	return a + b
}

func restar(a, b int) int {
	return a - b
}

func multiplicar(a, b int) int {
	return a * b
}

func operacionesMidd(f func(int, int) int) func(int, int) int {
	//Funcion anonima que retorna la funcion que se le pasa por parametros
	return func(a, b int) int {
		fmt.Println("Inicio de operacion:")
		return f(a, b)
	}
}
