package ejer_interfaces

import (
	"fmt"

	"github.com/MrEngineer/godesde0/interfaces"
)

func HumanoRespira(h interfaces.Humano) {
	h.Respirar()
	fmt.Printf("Sy un@ %s y estoy respirando\n", h.Sexo())
}