package main

import "fmt"

func main() {
	/*
			for i := 1; i <= 10; i++ {
				fmt.Println(i)
			}

			numero := 0
			for {
				fmt.Println("Continuo?")
				fmt.Println("Ingrese el numero secreto")
				fmt.Scanln(&numero)
				if numero == 100 {
					break
				}
			}


		var i = 0
		for i < 10 {
			fmt.Printf("\nValor de I: %d", i)
			if i == 5 {
				fmt.Printf("\n multiplicamos por 2\n")
				i = i * 2
				//el continue envia de nuevo a checkear la condicion sin ejecutar lo de abajo
				continue
			}
			fmt.Printf("	Paso por el if")
			i++
		}
	*/

	var i int = 0

RUTINA:
	fmt.Println("Aca empieza rutina")
	for i < 10 {
		if i == 4 {
			i = i + 2
			fmt.Println("Voy a rutina sumando 2 a i")
			goto RUTINA
		}
		fmt.Printf("Valor de i: %d\n", i)
		i++
	}
}
