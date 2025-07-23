package mapas

import "fmt"

func MostrarMapas() {
	paises := make(map[string]string, 5)

	//Los mapas poseen claves y valores
	paises["Mexico"] = "CDMX"
	paises["Colombia"] = "Bogotá"
	paises["Argentina"] = "Buenos Aires"
	paises["Chile"] = "Santiago"
	paises["Peru"] = "Lima"
	paises["Venezuela"] = "Caracas"
	paises["Ecuador"] = "Quito"

	fmt.Println(paises)
}
