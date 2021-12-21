/*


El mismo estudio del ejercicio anterior, solicita una funcionalidad para poder
registrar datos de nuevos clientes. Los datos requeridos para registrar a un cliente son:

Legajo
Nombre y Apellido
DNI
Número de teléfono
Domicilio

Tarea 1:

El número de legajo debe ser asignado o generado por separado y en forma previa a la carga de los
restantes gastos. Desarrolla e implementa una función para generar un ID que luego utilizarás para
asignarlo como valor a “Legajo”. Si por algún motivo esta función retorna valor “nil”,
debe generar un
panic que interrumpa la ejecución y aborte.

Tarea 2:
Antes de registrar a un cliente, debes verificar si el mismo ya existe.
Para ello, necesitas leer los datos de un archivo .txt. En algún lugar de tu código,
implementa la función para leer un archivo llamado “customers.txt” (como en el ejercicio anterior,
este archivo no existe, por lo que la función que intente leerlo devolverá un error).
Debes manipular adecuadamente ese error como hemos visto hasta aquí. Ese error deberá:

1.- generar un panic;
2.- lanzar por consola el mensaje: “error: el archivo indicado no fue encontrado o está dañado”,
y continuar con la ejecución del programa normalmente.

Tarea 3:
Luego de intentar verificar si el cliente a registrar ya existe, desarrolla una función para validar
que todos los datos a registrar de un cliente contienen un valor distinto de cero.
Esta función debe retornar, al menos, dos valores. Uno de los valores retornados deberá ser de tipo
error para el caso de que se ingrese por parámetro algún valor cero (recuerda los valores cero de cada
tipo de dato, ej: 0, “”, nil).

Tarea 4:
Antes de finalizar la ejecución, incluso si surgen panics, se deberán imprimir por consola
los siguientes mensajes: “Fin de la ejecución”, “Se detectaron varios errores en tiempo de ejecución”
y “No han quedado archivos abiertos” (en ese orden).
Utiliza defer para cumplir con este requerimiento.


*/

package main

import (
	"errors"
	"fmt"
	"os"
)

type Cliente struct {
	legajo    int
	nya       string
	dni       string
	telefono  int
	domicilio string
}

var ultimoLegajo int = 1000

func generarLegajo() int {
	ultimoLegajo++
	return ultimoLegajo
}

func leerArchivo(path string) {
	defer func() {
		err := recover()

		if err != nil {
			fmt.Println(err)
		}
	}()
	_, err := os.ReadFile(path)

	if err != nil {
		panic("el archivo indicado no fue encontrado o esta dañado")
	}

}

func validarDatosClientes(c Cliente) (bool, error) {

	leerArchivo("./customers.txt")

	if c.legajo == 0 || c.nya == "" || c.dni == "" || c.telefono == 0 || c.domicilio == "" {
		return false, errors.New("Faltan datos")
	}

	return true, nil
}

func main() {
	defer func() {
		err := recover()

		if err != nil {
			fmt.Println(err)
			fmt.Println("----------------------------------------------------")
			fmt.Println("Fin de la ejecución")
			fmt.Println("Se detectaron varios errores en tiempo de ejecución")
			fmt.Println("No han quedado archivos abiertos")

		}

	}()

	legajo := generarLegajo()

	if legajo == 0 {
		panic("Aborto la aplicación")
	}
	cliente := Cliente{}

	_, err := validarDatosClientes(cliente)

	if err != nil {
		panic(err)
	}

}
