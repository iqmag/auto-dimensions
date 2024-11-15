package main

import (
	"auto/cars"
	"auto/dimensions"
	"log"
)

func main() {
bmw := cars.BMW{}
	log.Println("BMW Dimensions:")
	log.Printf("Length: %.2f CM\n", bmw.Dimensions().Length().Get(dimensions.CM))
	log.Printf("Width: %.2f CM\n", bmw.Dimensions().Width().Get(dimensions.CM))
	log.Printf("Height: %.2f CM\n", bmw.Dimensions().Height().Get(dimensions.CM))

	mercedes := cars.Mercedes{}
	log.Println("\nMercedes Dimensions:")
	log.Printf("Length: %.2f CM\n", mercedes.Dimensions().Length().Get(dimensions.CM))
	log.Printf("Width: %.2f CM\n", mercedes.Dimensions().Width().Get(dimensions.CM))
	log.Printf("Height: %.2f CM\n", mercedes.Dimensions().Height().Get(dimensions.CM))

	dodge := cars.Dodge{}
	log.Println("\nDodge Dimensions:")
	log.Printf("Length: %.2f inches\n", dodge.Dimensions().Length().Get(dimensions.Inch))
	log.Printf("Width: %.2f inches\n", dodge.Dimensions().Width().Get(dimensions.Inch))
	log.Printf("Height: %.2f inches\n", dodge.Dimensions().Height().Get(dimensions.Inch))
}