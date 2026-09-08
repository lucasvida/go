package main

import "fmt"

func sayHello(name string) {
	fmt.Println("Hello, " + name + "!")
}

func somar(a int, b int) int {
	return a + b;
}

func modulo(a int, b int) int {
	return a % b;
}

func main() {
	sayHello("Lucas Vida")
	fmt.Println(somar(10, 20))
	fmt.Println(modulo(6,11))
}