package main

import (
	"fmt"

	"github.com/MrEngineer/godesde0/variables"
)

func main() {
	estado, texto := variables.ConvierteATexto(1588)
	fmt.Println(estado)
	fmt.Println(texto)
	
}	