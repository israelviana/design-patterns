package main

import (
	singleton "design-patterns/singleton"
	"fmt"
)

func main() {
	//Usando a variável global
	instance := singleton.GetInstance()
	fmt.Print(instance.Text)
}
