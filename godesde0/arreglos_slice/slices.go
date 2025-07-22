package arreglos_slice

import "fmt"

var tablaS []int = []int{2,5,4}
var arreglo [10]int = [10]int{1,2,3,4,5,6,7,8,9,10}	
func MuestraSlice() {
	fmt.Println("Muestra Slice")
	porcion := arreglo[3:] // Slice creado desde un vector, desde la posicion 3
	porcion2 := arreglo[:5] //Slice creado desde un vector, desde la posicion 0 hasta la 5 
	porcion3 := arreglo[6:7] // Slice creado desde un vector, desde la posicion 2 hasta la 5

	fmt.Print(porcion)
	fmt.Print(porcion2)
	fmt.Print(porcion3)
}

func Capacidad() {
	elementos := make([]int, 5, 20)
	fmt.Println("Capacidad del Slice: ", len(elementos), " - Capacidad: ", cap(elementos))
}