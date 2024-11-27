package main

import "fmt"

func main() {
	//Delcaracion de contstantes
	const message string = "Hola Mundo, Go pal mundo...!"
	const pi float64 = 3.14
	const pi2 = 3.15
	fmt.Println(message)
	fmt.Println("pi:", pi)
	fmt.Println("pi2:", pi2)

	// Declaracion de variables enteras
	base := 12
	var altura int = 14
	var area int

	fmt.Println(base, altura, area)

	// Zero values
	var a int
	var b float64
	var c string
	var d bool

	fmt.Println(a, b, c, d)

	//Area cuadrado
	const baseCuadrado = 10
	areaCuadrado := baseCuadrado * baseCuadrado

	fmt.Println("Area cuadrado:", areaCuadrado)

	// rectacgulo

	baseRectangulo := 20
	alturaRectangulo := 10

	areaRectangulo := baseRectangulo * alturaRectangulo

	fmt.Println("Area rectangulo:", areaRectangulo)

	// Circulo: area circulo = pi por radio al cuadrado
	const PI_LEO float64 = 3.14
	var radioCirculo float64 = 10

	areaCiruclo := PI_LEO * (radioCirculo * radioCirculo)

	fmt.Println("Area ciruclo:", areaCiruclo)

	// trapecio

	var base1 float64 = 6
	var base2 float64 = 15
	var alturaTrapecio float64 = 25

	areaTrapecio := ((base1 + base2) * alturaTrapecio) / 2

	fmt.Println("Area trapecio:", areaTrapecio)
}
