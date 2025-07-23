package users

import (
	"fmt"

	"github.com/MrEngineer/godesde0/models"
)

func AltaUsuario() {
	u := new(models.User)
	u.AddUser(1, "Jhon Doe", "andresigto", true)

	fmt.Println(u.Name)

}
