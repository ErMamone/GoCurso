package main

import "log"

func main() { /*
		archivo := "prueba.txt"
		f, err := os.Open(archivo)

		defer f.Close()

		if err != nil {
			fmt.Println("Error abriendo el archivo")
			os.Exit(1)
		}
	*/
	ejemploPanic()

}

func ejemploPanic() {
	//Defer ejecuta cualquier cosa como por ejemplo instruccion o funcion anonima al final del metodo
	defer func() {
		//Recoven recupera lo que tiro el panic
		reco := recover()
		if reco != nil {
			//Se puede verbosear
			log.Fatalf("ocurrio un error que genero Panic \n %v", reco)
		}
	}()

	a := 1

	if a == 1 {
		//Paniquea y te tira el dataso
		panic("Encontro el valor de 1")
	}
}
