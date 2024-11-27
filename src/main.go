package main

import (
	"fmt"
	"math"
)

func main() {

	// fnDeclaracionesVariables()
	// fnOperacionesAritmeticas()
	// fnTiposDatosPrimitivos()
	// fnPaqueteFmt()
	// fnUsoFunciones()
	fnUsoCiclos()

}

func fnDeclaracionesVariables() {
	fmt.Println("DECLARACION VARIABLES, CONSTANTES, ZERO VALUES")
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
}

func fnOperacionesAritmeticas() {
	fmt.Println("")
	fmt.Println("OPERADORES ARITMETICOS")

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

func fnTiposDatosPrimitivos() {
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

func fnPaqueteFmt() {
	fmt.Println("")
	fmt.Println("USO PAQUETE FMT")

	helloMessage := "Hello"
	worldMessage := "world"

	// Println: Salto de Linea Automatico
	fmt.Println(helloMessage, worldMessage)
	fmt.Println(helloMessage, worldMessage)

	// Printf
	nombre := "Platzi"
	cursos := 500
	// Con valores seguros
	fmt.Printf("%s tiene más de %d cursos\n", nombre, cursos)
	// Con valores inseguros
	fmt.Printf("%v tiene más de %v cursos\n", nombre, cursos)

	// Sprintf
	message2 := fmt.Sprintf("%v tiene más de %v cursos\n", nombre, cursos)
	fmt.Println(message2)

	// Tipo de datos:
	fmt.Printf("helloMessage: %T\n", helloMessage)
	fmt.Printf("cursos: %T\n", cursos)
}

func fnUsoFunciones() {
	fmt.Println("")
	fmt.Println("USO FUNCIONES")
	fmt.Printf("Circulo %.2f \n", fnAreaCirculo(2))
	fmt.Printf("Rectangulo %.2f \n", fnAreaRectangulo(5, 10))
	fmt.Printf("Trapezoide %.2f \n", fnAreaTrapezoide(10, 5, 3))
}

func fnAreaCirculo(radio float64) float64 {
	return math.Pi * radio * radio
}
func fnAreaRectangulo(base float64, altura float64) float64 {
	return base * altura
}

func fnAreaTrapezoide(B float64, b float64, h float64) float64 {
	return h * (B + b) / 2
}

func fnUsoCiclos() {
	fmt.Println("")
	fmt.Println("CICLOS")

	fmt.Println("-------------For condicional-------------")
	// For condicional
	for i := 0; i <= 5; i++ {
		fmt.Println(i)
	}

	fmt.Println("-------------For while-------------")
	// For while
	counter := 0
	for counter < 10 {
		fmt.Println(counter)
		counter++
	}

	fmt.Println("-------------For forever-------------")
	//For forever
	counterForever := 0
	for {
		fmt.Println(counterForever)
		counterForever++
		if counterForever > 12 {
			break
		}
	}

	fmt.Println("-------------For Range-------------")
	//For Range
	arreglo := [8]int{0, 1, 4, 6, 10, 9}
	fmt.Println("Arreglo:", arreglo)

	fmt.Println("Primer ejemplo")
	for i, j := range arreglo {
		fmt.Printf("indice i: %d tiene como valor #%d\n", i, j)
	}

	fmt.Println("Segundo ejemplo")
	for i := range arreglo {
		fmt.Printf("Valor de i: %d\n", i)
	}

	fmt.Println("Tercer ejemplo")
	for _, j := range arreglo {
		fmt.Printf("Valor de i: %d\n", j)
	}

	fmt.Println("-------------For con funciones-------------")
	// For con funciones
	for i := preFor(); condicion(i); i = postFor(i) {
		fmt.Printf("Valor de i: %d", i)
		if i == 7 {
			fmt.Printf(" así que saldremos del ciclo...\n")
			break /// este ejemplo es para usar el break
		}
		fmt.Printf("\n")
	}

	fmt.Println("-------------For con goto tags-------------")
	var i int
CICLO:
	fmt.Println("estamos fuera del for")
	for i < 10 {
		if i == 6 {
			i = i + 3
			fmt.Println("Saltando a etiqueta CICLO con i = i + 3")
			goto CICLO2
		}
		fmt.Printf("Valor de i: %d\n", i)
		i++
	}
CICLO2:
	fmt.Printf("ciclo 2 Valor de i: %d\n", i)
	if i == 9 {
		fmt.Printf("Valor de i: %d\n", i)
		i = i + 3
		fmt.Println("Saltando a etiqueta CICLO con i = i + 3")
		goto CICLO
	}
	fmt.Printf("terminamos\n")

	fmt.Println("-------------Reto for decreciente-------------")
	// Reto for decreciente
	for i := 10; i >= 0; i-- {
		fmt.Println(i)
	}
}

func preFor() int {
	fmt.Println("prefor i")
	return 0
}
func postFor(i int) int {
	fmt.Println("postFor sumemos i")
	i++
	return i
}
func condicion(i int) bool {
	fmt.Println("condicion i")

	return (i < 10)
}
