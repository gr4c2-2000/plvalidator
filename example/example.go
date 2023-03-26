package main

import (
	"fmt"

	"github.com/gr4c2-2000/plvalidator"
)

func main() {
	pesel := "66111146882"
	peselIsVaLid := plvalidator.Pesel(pesel)
	Pesel, _ := plvalidator.NewPesel(pesel)
	fmt.Printf("________________________________________________\n")
	fmt.Printf("PESEL %v, valid: %t \n", pesel, peselIsVaLid)
	fmt.Printf("Person birthday %v \n", Pesel.Birthday().Format("2006-02-01"))
	fmt.Printf("Person gender %v \n", Pesel.Gender())
	fmt.Printf("________________________________________________\n")
	regon := "590882968"
	regonIsValid := plvalidator.Regon(regon)
	fmt.Printf("REGON %v, valid: %t \n", regon, regonIsValid)
}
