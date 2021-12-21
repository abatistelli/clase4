/*


En tu función “main”, define una variable llamada “salary” y asignarle un valor de tipo “int”.
Crea un error personalizado con un struct que implemente
“Error()” con el mensaje “error: el salario ingresado no alcanza el mínimo imponible" y
lánzalo en caso de que “salary” sea menor a 150.000. Caso contrario, imprime por consola
el mensaje “Debe pagar impuesto”.

*/

package main

import "fmt"

type myCustomError struct {
	//msg string
}

func (e *myCustomError) Error() string {
	return fmt.Sprintf("error: el salario ingresado no alcanza el minimo imponible")
}

func main() {
	salary := 100000
	err := &myCustomError{}

	if salary < 150000 {
		fmt.Println(err)
	} else {
		fmt.Println("Debe pagar impuesto")
	}

}
