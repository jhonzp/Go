package main

import (
	"errors"
	"fmt"
	"log"
)

type raizError struct {
	lat  string
	long string
	err  error
}

func (re raizError) Error() string {
	return fmt.Sprintf("error matemático: %v %v \n %v", re.lat, re.long, re.err)
}

func main() {
	v, err := raiz(-9)
	if err != nil {
		log.Println(err)
	}
	fmt.Println(v)
}

func raiz(f float64) (float64, error) {
	if f < 0 {
		// Escribe tu código aquí
		e := errors.New(fmt.Sprintf("%v Must be grather than 0", f))
		return 0, raizError{lat: "50.2289 N", long: "99.4656 W", err: e}
	}
	return 42, nil
}
