package main

import (
	"fmt"

	_ "github.com/lib/pq"
)

func main() {
	fmt.Println("Successfully connected!")

	handleRequests()
}