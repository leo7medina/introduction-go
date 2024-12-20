package pkstructs

import "fmt"

type Pc struct {
	ram   int
	disk  int
	brand string
}

func NewPc(rm, dsk int, brn string) Pc {
	myPc := Pc{ram: rm, disk: dsk, brand: brn}
	return myPc
}

func (myPc Pc) FormatPrint() {
	fmt.Printf("Esta pc marca %s cuenta con una ram de %dGB y un disco de %dGB.\n", myPc.brand, myPc.ram, myPc.disk)
}

func (myPc *Pc) DuplicateRAM() {
	myPc.ram = myPc.ram * 2
}

// Stringer
func (juanito Pc) String() string {
	return fmt.Sprintf("PC: %s con %dGB RAM y SSD %dGB", juanito.brand, juanito.ram, juanito.disk)
}

func (myPc Pc) GetRam() int {
	return myPc.ram
}

func (myPc *Pc) SetRam(rm int) {
	myPc.ram = rm
}
