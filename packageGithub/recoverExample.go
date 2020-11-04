package main

import (
	"fmt"
	//go get -u github.com/jhonzp/Go/Errors/recoverexample
	"github.com/jhonzp/Go/Errors/recoverexample"
)

func main() {
	recoverexample.First()
	fmt.Println("Returned normally from f.")
}
