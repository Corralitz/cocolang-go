package main

import (
	"fmt"
	"os"
	"os/user"

	"github.com/Corralitz/cocolang-go/repl"
)

func main() {
	user, err := user.Current()
	if err != nil {
		panic(err)
	}

	fmt.Printf("Hola %s! Bienvenido a Coco, el lenguaje de programación!\n", user.Username)
	fmt.Printf("Empiece ingresando una expresión\n")
	repl.Start(os.Stdin, os.Stdout)
}
