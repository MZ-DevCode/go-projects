package main

import "fmt"

type Auto interface {
	StepOnGas()
}

type BMW struct{}

func (b BMW) StepOnGas() {
	fmt.Println("Это БМВ")
}

type Lambo struct{}

func (l Lambo) StepOnGas() {
	fmt.Println("Это Lambo")
}

func ride(auto Auto) {
	fmt.Println("Машина")
	auto.StepOnGas()
}

func main() {
	b := BMW{}
	l := Lambo{}
	ride(b)
	ride(l)
}
