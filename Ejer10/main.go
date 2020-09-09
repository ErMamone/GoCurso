package main

import "fmt"

type serVivo interface {
	estaVivo() bool
}
type humano interface {
	respirar()
	pensar()
	comer()
	sexo() string
}
type animal interface {
	respirar()
	comer()
	EsCarnivoro() bool
}
type vegetal interface {
	ClasificacionVegetal() string
}

/*Genero humanos*/
type hombre struct {
	edad       int
	altura     float32
	peso       float32
	respirando bool
	pensando   bool
	comiendo   bool
	esHombre   bool
	vivo       bool
}

type mujer struct {
	hombre
}

//Solo implementa la interfaz cuando se implementan todos los metodos (TODOS, SIN EXCEPCION)
func (h *hombre) respirar() { h.respirando = true }
func (h *hombre) comer()    { h.comiendo = true }
func (h *hombre) pensar()   { h.pensando = true }
func (h *hombre) sexo() string {
	if h.esHombre == true {
		return "Hombre"
	} else {
		return "Mujer"
	}
}
func (h *hombre) estaVivo() bool { return h.vivo }

//Entiende al humano como mujer y hombre en memoria, debido a que implementaron sus metodos
func HumanosRespirando(hu humano) {
	hu.respirar()
	fmt.Printf("Soy un/a %s, y ya estoy respirando \n", hu.sexo())
}

/*------------------------------------------------------------------------------------------*/

//Animales
type perro struct {
	respirando bool
	comiendo   bool
	carnivoro  bool
	vivo       bool
}

//Funciones de animal
func (p *perro) respirar()         { p.respirando = true }
func (p *perro) comer()            { p.comiendo = true }
func (p *perro) EsCarnivoro() bool { return p.carnivoro }
func (p *perro) estaVivo() bool    { return p.vivo }

func AnimalRespirar(an animal) {
	an.respirar()
	fmt.Printf("Soy un animal y estoy respirando\n")
}

func AnimalCarnivoro(an animal) int {
	if an.EsCarnivoro() == true {
		return 1
	}
	return 0
}

/* Ser Vivo*/

func estoyVivo(v serVivo) bool {
	return v.estaVivo()
}

func main() {
	//Se crea un hombre y una mujer, puede aplicar los metodos de la interfaz
	//debido a los metodos implementados mas arriba, se entiende que hombre y mujer implementan humano
	/*
		Pedro := new(hombre)
		Pedro.esHombre = true
		HumanosRespirando(Pedro)

		Maria := new(mujer)
		HumanosRespirando(Maria)
	*/

	totalCarnivoros := 0
	Dogo := new(perro)
	Dogo.vivo = true
	Dogo.carnivoro = true
	AnimalRespirar(Dogo)
	totalCarnivoros = +AnimalCarnivoro(Dogo)

	fmt.Printf("Total carnivoros %d vivos \n", totalCarnivoros)

	fmt.Printf("Estoy vivo = %t", estoyVivo(Dogo))

}
