package funciones

import "fmt"

func Tabla(valor int) func() int {
	numero := valor
	secuencia := 0
	return func() int {
		secuencia++
		resultado := numero * secuencia
		fmt.Printf("%d x %d = %d\n", numero, secuencia, resultado)
		return resultado
	}
}

func LLamarClosure() {
	tabladel1 := 2
	MiTabla := Tabla(tabladel1)
	for i := 1; i<=10; i++ {
		fmt.Println(MiTabla())
	}
}