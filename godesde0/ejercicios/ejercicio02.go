package ejercicios

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var numero int
var err error

func InputKey() {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("Ingresa un numero por teclado: ")
	for {
		if scanner.Scan() {
			numero, err = strconv.Atoi(scanner.Text())
			if err != nil {
				fmt.Println("Error vuelve a ingresar un numero porfavor")
			} else {
				TablaMultiplicar(numero)
				break
			}
		}
	}

}

func TablaMultiplicar(numero int) {
	for i := 1; i <= 10; i++ {
		fmt.Printf("%d x %d = %d\n", numero, i, numero*i)
		fmt.Println("Presiona enter para continuar...")
	}
}