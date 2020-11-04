package main

import (
	"fmt"

	"github.com/jhonzp/Go/Errors/recoverexample"
)

func main() {
	recoverexample.First()
	fmt.Println("Returned normally from f.")
}
