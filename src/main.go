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

	//Numeros enteros
	//int = Depende del OS (32 o 64 bits)
	//int8 = 8bits = -128 a 127
	//int16 = 16bits = -2^15 a 2^15-1
	//int32 = 32bits = -2^31 a 2^31-1
	//int64 = 64bits = -2^63 a 2^63-1

	//Optimizar memoria cuando sabemos que el dato simpre va ser positivo
	//uint = Depende del OS (32 o 64 bits)
	//uint8 = 8bits = 0 a 127
	//uint16 = 16bits = 0 a 2^15-1
	//uint32 = 32bits = 0 a 2^31-1
	//uint64 = 64bits = 0 a 2^63-1

	//numeros decimales
	// float32 = 32 bits = +/- 1.18e^-38 +/- -3.4e^38
	// float64 = 64 bits = +/- 2.23e^-308 +/- -1.8e^308

	//textos y booleanos
	//string = ""
	//bool = true or false

	//numeros complejos
	//Complex64 = Real e Imaginario float32
	//Complex128 = Real e Imaginario float64
	//Ejemplo : c:=10 + 8i
}
