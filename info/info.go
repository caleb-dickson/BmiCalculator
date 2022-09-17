package info

import "fmt"

const (
	mainTitle    = "BMI Calculator App"
	separator    = "==================================="
	WeightPrompt = "Please enter weight in lbs: "
	HeightPrompt = "Enter height in inches: "
)

func PrintWelcome() {
	fmt.Println("\n\n", mainTitle)
	fmt.Println(separator)
}
