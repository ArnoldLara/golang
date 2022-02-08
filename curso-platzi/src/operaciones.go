//Operadores aritmeticos
package main

import "fmt"



func main() {
	
	//Calcular el áera del cuadrado
	const baseCuadrado = 10
	areaCuadrado := baseCuadrado * baseCuadrado

	fmt.Println("El área del cuadrado es:", areaCuadrado)

	x := 50
	y := 10
	
	//Suma
	result := x + y
	fmt.Println("La suma entre X y Y es",result)
	//Resta
	result = x - y
	fmt.Println("La resta entre X y Y es",result)

	//Multiplicación
	result = x * y
	fmt.Println("La  multiplicación entre X y Y es",result)

	//División
	result = x / y
	fmt.Println("La división entre X y Y es",result)

	//El residuo de la División
	result = x % y
	fmt.Println("La resudio de la división entre X y Y es",result)

	//Incremental
	x++
	fmt.Println("El incremento de X es",x)

	//Decremento
	y--
	fmt.Println("El incremento de Y es",y)
}