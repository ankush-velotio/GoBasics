package main

import "fmt"

type Car struct {
	name  string
	brand string
}

func (c Car) getCarName() string {
	return c.name
}

func main() {
	db11 := Car{"db11", "Aston Martin"}
	carName := db11.getCarName()
	fmt.Println(carName)
}
