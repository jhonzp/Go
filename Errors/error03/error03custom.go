package main

import (
	"fmt"
)

type errorCustom struct {
	message  string
	value    int
	Rollback bool
}

func (ec errorCustom) Error() string {
	serror := fmt.Sprintf("Error  message %v \n value %v \n  Rollback %v", ec.message, ec.value, ec.Rollback)
	return serror
}

func main() {
	ec := errorCustom{
		message:  "Hello",
		value:    500,
		Rollback: false,
	}
	foo(ec)
}

func foo(e error) {
	fmt.Println("foo corrió -", e, "\n")
}
