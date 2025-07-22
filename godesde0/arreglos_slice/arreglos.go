package arreglos_slice

import "fmt"

var tabla [10]int 
var matrix [3][3]int

func MuestroArreglos() {
	tabla[7] = 33
	tabla[8] = 44

	tabla2 := [10]string{"Hola", "Mundo", "Go", "Es", "Genial", "!"}
	fmt.Println(tabla2)
	fmt.Println(tabla)
}
