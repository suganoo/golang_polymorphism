package main

import "fmt"

type Cat struct {
}

func (c *Cat)callCat(a *Animal) {
	fmt.Println("--- callCat ---")
	fmt.Println(a.FirstName + " " + a.LastName)
	a.hello()
}

type Dog struct {
}

func (d *Dog)callDog(a *Animal) {
	fmt.Println("--- callDog ---")
	a.callCat(a)
}

type Animal struct {
	Cat
	Dog
	FirstName string
	LastName  string
}

func (a *Animal)hello() {
	fmt.Println("animal hello")
}

func main() {
	var cat Cat = Cat{}
	var dog Dog = Dog{}

	var animal = Animal{
		Cat: cat,
		Dog: dog,
		FirstName: "lady",
		LastName:  "gaga",
	}
	animal.hello()
	animal.callCat(&animal)
	animal.callDog(&animal)
}
