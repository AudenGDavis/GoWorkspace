package main

import "fmt"

type Animal interface {
	makeSound()
	setName(string)
}

type Dog struct {
	name   string
	age    int
	weight float32
}

func (dog Dog) makeSound() {
	fmt.Println("sound made")
}

func (dog Dog) setName(newName string) {
	dog.name = newName
}

func (dog Dog) doTrick() {
	fmt.Println("Trick completed")
}

func main() {
	var marigold Dog = Dog{"Marigold", 8, 35.0}

	marigold.doTrick()

	fmt.Println(marigold)

}
