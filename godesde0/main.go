package main

import (
	"fmt"

	"github.com/MrEngineer/godesde0/ejercicios"
)

func main() {
	/**estado, texto := variables.ConvierteATexto(1588)
	fmt.Println(estado)
	fmt.Println(texto)
	if os := runtime.GOOS; os == "windows" {
		fmt.Println("Ejecutando en Windows")
	}

	switch os := runtime.GOOS; os {
	case "linux":
		println("Ejecutando en Linux")
	case "darwin":
		println("Ejecutando en MacOS")
	default:
		fmt.Printf("%s \n", os)

	}*/
	numero, texto := ejercicios.ConvertStirng("Andresito")

	fmt.Println("Número convertido:", numero)
	fmt.Println("Texto original:", texto)

}
