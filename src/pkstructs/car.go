package pkstructs

import "fmt"

// Car
type Car struct {
	Brand string
	Year  int
}

type carPrivate struct {
	brand string
	year  int
}

func PrintMessage() {
	fmt.Println("Hola")
}
