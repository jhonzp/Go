// Package recoverexample is a example Package
package recoverexample

import "fmt"

// func recoverexample() {
// 	first()
// 	fmt.Println("Returned normally from f.")
// }

//Esta función presenta el funcionamiento de recover, defer y panic
func first() {
	//Esta función literal se ejecuta al final
	defer func() {
		//Cuando se lanza el Panic es capturado por recover (los defer se continuan ejecutando)
		if r := recover(); r != nil {
			fmt.Println("Recovered in first", r)
		}
	}()
	fmt.Println("Calling g.")
	Gotest(0)
	fmt.Println("Returned normally from g.")
}

//funcion de test
func Gotest(i int) {
	if i > 3 {
		fmt.Println("Panicking!")
		//El valor de i es recuperado en el recover (los defer se siguen ejecutando)
		panic(fmt.Sprintf("%v", i))
	}
	//Este Print se hace al final
	defer fmt.Println("Defer in g", i)
	fmt.Println("Printing in g", i)
	//llamado recursivo
	Gotest(i + 1)
}
