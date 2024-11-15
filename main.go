package main

import (
	"auto/cars"
	"auto/dimensions"
	"fmt"
)

func main() {
 bmw := cars.BMW{}
 fmt.Println("BMW Dimensions:")
 fmt.Printf("Length: %.2f CM\n", bmw.Dimensions().Length().Get(dimensions.CM))
 fmt.Printf("Width: %.2f CM\n", bmw.Dimensions().Width().Get(dimensions.CM))
 fmt.Printf("Height: %.2f CM\n", bmw.Dimensions().Height().Get(dimensions.CM))

 mercedes := cars.Mercedes{}
 fmt.Println("\nMercedes Dimensions:")
 fmt.Printf("Length: %.2f CM\n", mercedes.Dimensions().Length().Get(dimensions.CM))
 fmt.Printf("Width: %.2f CM\n", mercedes.Dimensions().Width().Get(dimensions.CM))
 fmt.Printf("Height: %.2f CM\n", mercedes.Dimensions().Height().Get(dimensions.CM))

 dodge := cars.Dodge{}
 fmt.Println("\nDodge Dimensions:")
 fmt.Printf("Length: %.2f inches\n", dodge.Dimensions().Length().Get(dimensions.Inch))
 fmt.Printf("Width: %.2f inches\n", dodge.Dimensions().Width().Get(dimensions.Inch))
 fmt.Printf("Height: %.2f inches\n", dodge.Dimensions().Height().Get(dimensions.Inch))
}