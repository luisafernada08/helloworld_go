package main

import "fmt"

func pensionar() {
	var nombre, apellido string = "julian", "gonzales"
	var edad int = 29
	var pensionado bool = false

	fmt.Println(nombre, apellido)
	fmt.Println("Edad: ", edad)
	fmt.Println("Pensionado: ", pensionado)
}

func peso_persona() {
	var estudinte bool
	var peso float64
	var edad int
	fmt.Println("estudinte:", estudinte)
	fmt.Println("Edad: ", edad)
	fmt.Println("peso: ", peso)
}

func profecion() {
	var ingeniero = "sistemas "
	var sueldo = 1000000
	fmt.Println("ingeniero ", ingeniero)
	fmt.Println("Sueldo: ", sueldo)
}

func verificar_nivel() {
	var var1 = "es igua nivel1"
	var var2 = "es igual nivel2"
	var var3 = "es mayor nivel1"
	fmt.Println(var3)
	fmt.Println(var1)
	fmt.Println(var2)
}

func colores() {
	var color1 = "rojo"
	var color2 = "azul"
	var color3 = "verde"
	fmt.Println(color3)
	fmt.Println(color1)
	fmt.Println(color2)
}

const Pi = 3.1416

func area_de_un_triangulo(radio float64) float64 {
	return Pi * radio * radio
}

func main() {
	//pensionar()
	//peso_persona()
	//profecion()
	//verificar_nivel()
	//colores()
	//fmt.Println("Area de un triangulo: ", area_de_un_triangulo(5))

}
