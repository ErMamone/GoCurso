package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	leoArchivo()
	leoArchivo2()
	graboArchivo()
	graboArchivo2()
}

func leoArchivo() {
	archivo, err := ioutil.ReadFile("./simple.txt")

	if err != nil {
		fmt.Println("Hubo un error.")
	} else {
		fmt.Println("\n" + string(archivo))
	}
}

func leoArchivo2() {
	archivo, err := os.Open("./simple.txt")

	if err != nil {
		fmt.Println("Hubo un error.")
	} else {
		scanner := bufio.NewScanner(archivo)
		for scanner.Scan() {
			registro := scanner.Text()
			fmt.Printf("Linea > " + registro + "\n")

		}
	}
	archivo.Close()
}

func graboArchivo() {
	archivo, err := os.Create("./nuevoArchivo.txt")
	if err != nil {
		fmt.Println("Hubo un error.")
		return
	}
	fmt.Fprintln(archivo, "QUE ONDA PUTOOOOOOOOOS")
	archivo.Close()
}

func graboArchivo2() {
	fileName := "./nuevoArchivo.txt"
	if Append(fileName, "\nQUE ONDA BROOOOOOOOOOOOOOOOOOOOOOOOOOO NO TELA CREO") == false {
		fmt.Println("Eror en la segunda linea gay")
	}

}

func Append(archivo string, texto string) bool {
	arch, err := os.OpenFile(archivo, os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil {
		fmt.Println("Hubo un error.")
		return false
	}

	_, err = arch.WriteString(texto)
	if err != nil {
		fmt.Println("Hubo un error.")
		return false
	}

	return true
}
