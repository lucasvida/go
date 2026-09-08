package main

import (
	"fmt"
	"net/http"
)

const jsonResponse = `
{
	"message": "Hello, World! Meu primeiro servidor em Go!",
	"status": "success",
	"users": [
		{
			"name": "Lucas Vida",
			"age": 20
		},
		{
			"name": "Anna",
			"age": 18
		},
		{
			"name": "João",
			"age": 22
		},
		{
			"name": "Maria",
			"age": 24
		},	
		{
			"name": "Pedro",
			"age": 26
		},
		{
			"name": "Ana",
			"age": 28
		},
		{
			"name": "Carlos",
			"age": 30
		},
		{
			"name": "Julia",
			"age": 32
		},
		{
			"name": "Bruno",
			"age": 34
		},
		{
			"name": "Larissa",
			"age": 36
		}
	]
}
`

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, jsonResponse)
	})

	fmt.Println("Servidor em http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}