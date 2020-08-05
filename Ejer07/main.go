package main

import "fmt"

/*
/Estan en todo el entorno del archivo
/var tabla [10]int
/tabla := [10]int{5, 1, 2, 6, 4, 8, 22, 10, 9}

var matriz [5][7]int
*/

func main() {
	/*
				tabla[0] = 1
				tabla[5] = 15

			tabla := [10]int{5, 1, 2, 6, 4, 8, 22, 10, 9}

			for i := 0; i < len(tabla); i++ {
				fmt.Println(tabla[i])
			}

		//Matriz
		matriz[3][5] = 1

		fmt.Println(matriz)
	*/

	matriz := []int{2, 4, 5}

	fmt.Println(matriz)

	variante2()
	variante3()
	variante4()
}

func variante2() {
	elementos := [5]int{1, 2, 3, 4, 5}
	porcion := elementos[2:4]

	fmt.Println(elementos)
	fmt.Println(porcion)

}

func variante3() {
	elementos := make([]int, 5, 20)
	fmt.Printf("Largo %d, capacidad %d", len(elementos), cap(elementos))
}

func variante4() {
	nums := make([]int, 0, 0)

	for i := 0; i < 579; i++ {
		nums = append(nums, i)
	}
	fmt.Printf("\nLargo %d y capacidad %d", len(nums), cap(nums))

}
