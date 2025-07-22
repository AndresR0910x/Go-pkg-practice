package files

import (
	"bufio"
	"fmt"
	"os"

	"github.com/MrEngineer/godesde0/ejercicios"
)

var path string = "./files/archivos/tabla.txt"

func GuardarTabla() {
	var texto string = ejercicios.InputKey()
	archivo, err := os.Create(path)
	if err != nil {
		fmt.Println("Error al crear el archivo:", err.Error())
		return
	}
	fmt.Fprintln(archivo, texto)
	archivo.Close()
}

func SumaTabla() {
	var texto string = ejercicios.InputKey()
	if !Append(path, texto) {
		fmt.Println("Error al concatenar el archivo:")
	}
}

func Append(file string, text string) bool {
	arch, err := os.OpenFile(path, os.O_WRONLY|os.O_APPEND, 0644)

	if err != nil {
		fmt.Println("Error al abrir el archivo:", err.Error())
		return false
	}
	_, err = arch.WriteString(text)

	if err != nil {
		fmt.Println("Error durante el WriteString:", err.Error())
		return false
	}

	arch.Close()
	return true
}

func LeerArchivo() {
	archivo, err := os.Open(path)
	if err != nil {
		fmt.Println("Error al leer el archivo:", err.Error())
		return
	}
	scanner := bufio.NewScanner(archivo)
	for scanner.Scan() {
		registro := scanner.Text()
		fmt.Println("> " + registro)
	}
	archivo.Close()
}
