package main

import (
	"fmt"
	pk "introduction-go/src/pkstructs"
	"math"
	"reflect"
	"strings"
)

func main() {

	// fnDeclaracionesVariables()
	// fnOperacionesAritmeticas()
	// fnTiposDatosPrimitivos()
	// fnPaqueteFmt()
	// fnUsoFunciones()
	// fnUsoCiclos()
	// fnCondicionalIf()
	// fnUsoSwitch()
	// fnUsoKeyWords()
	// fnUsoArraysAndSlices()
	// fnUsoRecorridoSlicesConRange()
	// fnLlaveValorConMaps()
	// fnUsoStructs()
	// fnUsoModificadoresAcceso()
	// fnStructAndPunteros()
	// fnUsoStringers()
	fnUsoInterfacesYlistas()

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
	fmt.Printf("Circulo %.2f \n", areaCirculo(2))
	fmt.Printf("Rectangulo %.2f \n", areaRectangulo(5, 10))
	fmt.Printf("Trapezoide %.2f \n", areaTrapezoide(10, 5, 3))
}

func areaCirculo(radio float64) float64 {
	return math.Pi * radio * radio
}
func areaRectangulo(base float64, altura float64) float64 {
	return base * altura
}

func areaTrapezoide(B float64, b float64, h float64) float64 {
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

func fnCondicionalIf() {
	fmt.Println("")
	fmt.Println("CONDICIONAL IF")
	parOImpar(13)
	parOImpar(10)

	validateUser("admin", "admin123")
	validateUser("leo", "admin123")
}

func parOImpar(numero int) {
	if numero%2 == 0 {
		fmt.Printf("El número %d es par \n", numero)
	} else {
		fmt.Printf("El número %d es impar\n", numero)
	}
}

func validateUser(user string, password string) {
	if user == "admin" && password == "admin123" {
		fmt.Printf("Acceso permitido \n")
	} else {
		fmt.Printf("Acceso denegado \n")
	}
}

func fnUsoSwitch() {
	fmt.Println("")
	fmt.Println("USO SWITCH")

	palabra := "en platzi nunca paramos de aprender"
	a, e, i, o, u := contadorVocales(palabra)
	fmt.Printf("la frase '%s' tiene: \n", palabra)
	fmt.Printf("%d vocales a \n", a)
	fmt.Printf("%d vocales e \n", e)
	fmt.Printf("%d vocales i \n", i)
	fmt.Printf("%d vocales o \n", o)
	fmt.Printf("%d vocales u \n", u)
}

func contadorVocales(palabra string) (int, int, int, int, int) {
	conta := 0
	conte := 0
	conti := 0
	conto := 0
	contu := 0
	for _, valor := range palabra {
		variable := string(valor)
		switch variable {
		case "a":
			conta++
		case "e":
			conte++
		case "i":
			conti++
		case "o":
			conto++
		case "u":
			contu++
		}
	}
	return conta, conte, conti, conto, contu
}

func fnUsoKeyWords() {
	// Continue y break
	for i := 0; i < 10; i++ {
		fmt.Println(i)
		if i == 4 {
			fmt.Println("Es 4")
			continue
		}

		// Break
		if i == 4 {
			fmt.Println("Break")
			break
		}
	}
}

func fnUsoArraysAndSlices() {
	// Array  (son inmutables)
	var array [4]int
	array[0] = 1
	array[1] = 2
	fmt.Println(array, len(array), cap(array))

	// Slice (son dinamicos)
	slice := []int{0, 1, 2, 3, 4, 5, 6, 7, 8}
	fmt.Println(slice, len(slice), cap(slice))

	// Métodos en el slice
	fmt.Println(slice[0])
	fmt.Println(slice[:3])
	fmt.Println(slice[2:4])
	fmt.Println(slice[4:])

	// Append
	slice = append(slice, 11)
	fmt.Println(slice)

	// Append
	newSlice := []int{12, 13, 14}
	slice = append(slice, newSlice...)
	fmt.Println(slice)
}

func fnUsoRecorridoSlicesConRange() {

	slice := []string{"hola", "que", "haces"}
	for i := range slice {
		fmt.Println(i)
	}

	var palabra string
	fmt.Scan(&palabra)
	minus := strings.ToLower(palabra)
	esPalindromo(minus)
}

func esPalindromo(text string) {
	var textReverse string
	for i := len(text) - 1; i >= 0; i-- {
		textReverse += string(text[i])
	}

	if text == textReverse {
		fmt.Println("Es palindromo")
	} else {
		fmt.Println("No es palindromo")
	}
}

func fnLlaveValorConMaps() {
	fmt.Println()
	fmt.Println("====================================================")
	fmt.Println("Llave valor con Maps (también llamados diccionarios)")
	fmt.Println("====================================================")
	fmt.Println()

	maps := make(map[string]int)

	maps["Jose"] = 14
	maps["Luis"] = 8

	fmt.Println(maps)

	//Recorrer map
	recorrerMap(maps)

	//Encontrar un valor
	fmt.Println("Encontrando valor")
	v1 := maps["Jose"]
	v2 := maps["Josep"]
	// valor, existe
	v3, ok3 := maps["Jose"]
	v4, ok4 := maps["Josep"]
	fmt.Println(v1)
	fmt.Println(v2)
	fmt.Println(v3, ok3)
	fmt.Println(v4, ok4)

	//Declaration
	var map_1 = map[string]int{
		"Car":      50000,
		"House":    20000,
		"Computer": 1000,
	}
	recorrerMap(map_1)

}

func recorrerMap(mapDemo map[string]int) {
	for key, value := range mapDemo {
		fmt.Printf("Llave: %s - Valor: %d \n", key, value)
	}
}

type car struct {
	brand   string
	year    int
	seating int
	color   string
	owner   string
}

func fnUsoStructs() {
	fmt.Println()
	fmt.Println("====================================================")
	fmt.Println("Structs: forma de hacer clases en Go")
	fmt.Println("====================================================")
	fmt.Println()

	fmt.Println("Primera forma de instanciación:::")
	fmt.Println("---------------------------------")
	fmt.Println()

	myCar := car{brand: "Renault", year: 2021}
	fmt.Println(myCar)

	myCar2 := car{brand: "Toyota", year: 2018, seating: 10, color: "Rojo", owner: "Eliaz Bobadilla"}
	fmt.Println("Los Datos de mi auto son:", myCar2)

	fmt.Println("Segunda forma de instanciación:::")
	fmt.Println("---------------------------------")
	fmt.Println()

	var otherCar car
	otherCar.brand = "Chevrolet"
	fmt.Println(otherCar)
	//Si un atributo del struct no se inicializa, toma un Zero Value, y por lo tanto se imprime así lo referente al año:
	fmt.Println()

}

func fnUsoModificadoresAcceso() {
	fmt.Println()
	fmt.Println("====================================================")
	fmt.Println("Modificadores de acceso en funciones y Structs")
	fmt.Println("====================================================")
	fmt.Println()

	var myCar pk.Car
	myCar.Brand = "FORD"
	myCar.Year = 2020
	fmt.Println(myCar)

}

func fnStructAndPunteros() {
	fmt.Println()
	fmt.Println("====================================================")
	fmt.Println("Structs y Punteros")
	fmt.Println("====================================================")
	fmt.Println()

	a := 50
	b := &a

	fmt.Println(b)
	fmt.Println(*b)

	// El "&" accede a la dirección del espacio de memoria de la variable.
	//"*" accede al valor que contiene ese espacio de memoria, dado el nombre de una variable o una dirección especifica.
	*b = 100
	fmt.Println(a)

	myPC := pk.NewPc(16, 200, "msi")
	fmt.Println(myPC)

	myPC.DuplicateRAM()
	fmt.Println(myPC)

	myPC.DuplicateRAM()
	fmt.Println(myPC)
}

func fnUsoStringers() {
	fmt.Println()
	fmt.Println("====================================================")
	fmt.Println("Stringers")
	fmt.Println("====================================================")
	fmt.Println()
	myPC := pk.NewPc(16, 200, "msi")
	fmt.Println(myPC)
	fmt.Printf("%+v", myPC) // resultado=> {ram:16 disk:200 brand:msi}%
}

type figure2D interface {
	getArea() float64
}

type square struct {
	base float64
}

type rectangle struct {
	high  float64
	width float64
}

type trapezoid struct {
	baseA float64
	baseB float64
	high  float64
}

type circle struct {
	radio float64
}

func (s square) getArea() float64 {
	return s.base * s.base
}

func (r rectangle) getArea() float64 {
	return r.high * r.width
}

func (t trapezoid) getArea() float64 {
	return ((t.baseA + t.baseB) / 2) * t.high
}

func (c circle) getArea() float64 {
	return math.Pi * math.Pow(c.radio, 2)
}

func calculateArea(f figure2D) {
	fmt.Printf("Area of %s: %.2f\n", reflect.TypeOf(f).Name(), f.getArea())
}

func fnUsoInterfacesYlistas() {
	fmt.Println()
	fmt.Println("====================================================")
	fmt.Println("Interfaces y listas")
	fmt.Println("====================================================")
	fmt.Println()

	mySquare := square{base: 5}
	myRectangle := rectangle{width: 4, high: 5}
	myTrapezoid := trapezoid{baseA: 18, baseB: 10, high: 5}
	myCircle := circle{radio: 4}
	calculateArea(mySquare)
	calculateArea(myRectangle)
	calculateArea(myTrapezoid)
	calculateArea(myCircle)

}
