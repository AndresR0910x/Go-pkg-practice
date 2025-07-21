package variables

import (
	"fmt"
	"time"
	"strconv"
)

//SCOPE
var Nombre string
var Estado bool
var Sueldo float32
var Fecha time.Time



func RestoVariables() {
	Nombre = "Pedro"
	Estado = true
	Sueldo = 1577.66
	Fecha = time.Now()
	fmt.Println(Nombre)
	fmt.Println(Estado)
	fmt.Println(Sueldo)
	fmt.Println(Fecha)

}

func ConvierteATexto(number int) (bool, string) {
	text := strconv.Itoa(number)
	return true, text
}