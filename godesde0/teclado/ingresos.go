package teclado

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var numero1 int
var numero2 int 
var leyenda string 
var err error

func IngresoNumeros() {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("Ingrese numero 1 : ")
	if scanner.Scan() {
		numero1, err = strconv.Atoi(scanner.Text())
		if err != nil {
			panic("El dato no es un numero entero, "+ err.Error())
		}
	}

	fmt.Println("Ingrese numero 2 : ")
	if scanner.Scan() {
		numero2, err = strconv.Atoi(scanner.Text())
		if err != nil {
			panic("El dato no es un numero entero, "+ err.Error())
		}
	}

	fmt.Println("Ingrese una leyenda : ")
	if scanner.Scan() {
		leyenda = scanner.Text()
	}

	fmt.Println(leyenda, numero1 * numero2)
}