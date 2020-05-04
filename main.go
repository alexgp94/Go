package main

import (
	"fmt"
	"github.com/alexgp94/Go/Tutoriales"
	"github.com/alexgp94/Go/models"
)

// Definimos la funcion pincipal
func main() {
	Tutoriales.Holamundo()
	u := models.User{
		Id:        2,
		Firstname: "Alex",
		Lastname:  "garcia",
	}
	fmt.Println(u)
	u2 := models.User{
		Id:        2,
		Firstname: "Ander",
		Lastname:  "garcia",
	}

	fmt.Println(u2)

}
