package variables

import "fmt"

func ShowInteger() {
	//Variable declarada
	var intcomun int
	//Variable por asignacion
	//Devuelve un entero de 32 bytes
	intde32 := int32(10)
	intde64 := int64(100)

	fmt.Println("Entero comun =", intcomun)
	fmt.Println("Entero de 32 bytes =", intde32)
	fmt.Println("Entero de 64 bytes =", intde64)

}

