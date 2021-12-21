/*
Vamos a hacer que nuestro programa sea un poco más complejo y útil.

Desarrolla las funciones necesarias para permitir a la empresa calcular:
1)Salario mensual de un trabajador según la cantidad de horas trabajadas.

a)La función recibirá las horas trabajadas en el mes y el valor de la hora como argumento.
Dicha función deberá retornar más de un valor (salario calculado y error).
En caso de que el salario mensual sea igual o superior a $150.000, se le deberá descontar el 10%
en concepto de impuesto.
En caso de que la cantidad de horas mensuales ingresadas sea menor a 80 o un número negativo,
la función debe devolver un error. El mismo deberá indicar “error: el trabajador no puede haber
trabajado menos de 80 hs mensuales”.

b) Calcular el medio aguinaldo correspondiente al trabajador
Fórmula de cálculo de aguinaldo:
[mejor salario del semestre] / 12 * [meses trabajados en el semestre].
La función que realice el cálculo deberá retornar más de un valor,
incluyendo un error en caso de que se ingrese un número negativo.


2)Desarrolla el código necesario para cumplir con las funcionalidades requeridas,
utilizando “errors.New()”, “fmt.Errorf()” y “errors.Unwrap()”.
No olvides realizar las validaciones de los retornos de error en tu función “main()”.

*/

package main

import (
	"errors"
	"fmt"
)

func salarioMensual(horas float64, valor float64) (float64, error) {
	if horas < 80 || horas < 0 {
		return 0, errors.New("error: el trabajador no puede haber trabajado menos de 80 hs mensuales\n")
	}
	salary := float64(horas) * valor

	if salary >= 150000 {
		salary -= salary * 0.1
	}

	return salary, nil

}

func calcularAguinaldo(horasPorMes []float64, valorHora float64) (float64, error) {

	salariosDelSemestre := make([]float64, 6)

	for i, value := range horasPorMes {
		salario, err := salarioMensual(value, valorHora)

		if err != nil {
			e2 := fmt.Errorf("%w", err)
			return 0, e2
		}

		salariosDelSemestre[i] = salario
	}
	fmt.Println(salariosDelSemestre)
	mejorSalario, mesesTrabajados := buscarMejorSalario(salariosDelSemestre)

	return mejorSalario * 12 / float64(mesesTrabajados), nil
}

func buscarMejorSalario(salarios []float64) (float64, int) {
	max := salarios[0]
	meses := 0

	for _, salario := range salarios {
		if salario == 0 {
			break
		}

		if max < salario {
			max = salario
		}

		meses++
	}
	return max, meses
}

func main() {
	salary, err := salarioMensual(90, 200)

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("El salario mensual es: %.2f\n", salary)
	}

	horas := []float64{90, 95, 85, 100}
	aguinaldo, err2 := calcularAguinaldo(horas, 100)

	if err2 != nil {
		fmt.Println(errors.Unwrap(err2))
	} else {
		fmt.Printf("Aguinaldo: %2.f\n", aguinaldo)
	}
}
