package main

import (
	"fmt"
	"reflect"
	"strconv"
)

func imprimir() {
	fmt.Println("texto desde la uncion imprimir")
}

func areglo() {
	var aprendices []string
	aprendices[0] = "leonardo"
	aprendices[1] = "juan sebastian"
	aprendices[2] = "frank"
	aprendices[3] = "jaider"
	fmt.Println(aprendices[0])

}
func comparar_bool() {
	var a bool
	b := true
	a = false
	if a == b {
		fmt.Println("a y b Son iguales")
	} else {
		fmt.Println("a y b no son iguales")
	}

}

func tipos_datos() {
	var texto string = "luisa"
	var numero int = 10000
	var decimal float32 = 0.00001
	fmt.Println(reflect.TypeOf(texto))
	fmt.Println(reflect.TypeOf(numero))
	fmt.Println(reflect.TypeOf(decimal))
}
func convertir() {
	var palabra string = "true"
	boolean, err := strconv.ParseBool("palabra")
	fmt.Println(boolean, reflect.TypeOf(boolean))
	fmt.Println("este es el erros,err")

}

func convertitbooltostrin() {
	var palabrabool bool = true
	strbool := strconv.FormatBool(palabrabool)
	fmt.Println(strbool, reflect.TypeOf(strbool))
}

func main() {
	//fmt.Println("texto desde la funcion main")
	//imprimir()
	tipos_datos()
	//convertir()
	//areglo()
	//comparar_bool()
	//fmt.Println(hola_string("luisa"))

}
