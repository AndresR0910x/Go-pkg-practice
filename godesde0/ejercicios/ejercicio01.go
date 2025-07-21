package ejercicios

import (
	"fmt"
	"strconv"
)

func ConvertString (text string) (int, string) {
	entero, error := strconv.Atoi(text)
	if error != nil {
		fmt.Println("Error al convertir el texto a entero:", error)
	} else if entero > 100 {
		fmt.Println("El número es mayor que 100")
	} else {
		fmt.Println("El número es menor o igual a 100")
	}

	return entero, text
}
