package main

import "fmt"

type Product struct {
	ID   int64
	Name string
}

func main() {
product := Product{
	ID:   1,
	Name: "iPhone",
}

func rename(p Product) {
	p.Name = "Samsung"
}

rename(product)

fmt.Println(product)
}
