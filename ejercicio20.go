package main

import "fmt"

func conversionEnPies(altura *float32) float32 {
    *altura = *altura - 0.10
    return *altura

}

var altura float32 = 1.70

func equivalenciaEnPies(altura float32) float32 {
    altura = altura * 3.28
    return altura
}

var altura float32 = 1.70


func main() {
    //fmt.Println("La altura es:", altura, "mts")
    //fmt.Println("La altura es:", conversionEnPies(&altura), " pies")
    //fmt.Println("Despues de envejecer:", altura, "mts")

    fmt.Println("La altura es:", altura, "mts")
    fmt.Println("La altura es:", conversionEnPies(&altura), " pies")
    fmt.Println("Despues de envejecer:", altura, "mts")
}
