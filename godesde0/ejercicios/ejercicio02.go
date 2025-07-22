package ejercicios

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var numero int
var err error

func InputKey()  string {
	var multiplicacion string
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("Ingresa un numero por teclado: ")
	for {
		if scanner.Scan() {
			numero, err = strconv.Atoi(scanner.Text())
			if err != nil {
				fmt.Println("Error vuelve a ingresar un numero porfavor")
				multiplicacion = "No es un numero!"
			} else {
				multiplicacion = TablaMultiplicar(numero)
				break
			}
		}
	}
	return multiplicacion
}

func TablaMultiplicar(numero int) string {
	var multiplicacion string
	for i := 1; i <= 10; i++ {
		multiplicacion += fmt.Sprintf("%d x %d = %d\n", numero, i, numero*i)
	}
	return multiplicacion

}
