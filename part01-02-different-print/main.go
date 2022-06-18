package main

import "fmt"

func main() {
	fmt.Print("Using the normal Print func.\n")
	fmt.Println("With Println it have the same behavior as Print except this func will add break line at the of the text.")
	fmt.Printf("You could specifi any format that you want, for exmaple \ndecimal %d", 42384)
	fmt.Printf("\nfloat %f", 42.234)
	fmt.Printf("\nfloat with 2 percision %.2f", 42.234)
}
