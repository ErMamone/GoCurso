package main

import "fmt"


func main()  {
	campeonato := map[string]int {
		"Barcelona": 39,
		"Real madrid": 38,
		"Chivas": 37,
		"Boquita papa" 30}
	

	campeonato["River"] = 25
	campeonato["Chivas"] = 55
	delete(campeonato, "Real madrid")
	fmt.PrintF(campeonato)

	for equipo, puntaje := range campeonato{
		fmt.Printf("El equipo %s, tiene un puntaje de: %d \n", equipo, puntaje)
	}
	puntaje, existe := campeonato["Chivas"]
	fmt.Printf("El puntaje es de %d, y el equipo existe %t \n",puntaje,existe)

	
}