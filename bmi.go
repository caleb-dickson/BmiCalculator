package main

import (
	"fmt"
	"github.com/tacDev-io/BmiCalculator/info"
	"strconv"
	"strings"
)

func main() {

	fmt.Println(info.MainTitle)
	fmt.Println(info.Separator)
	fmt.Println(info.WeightPrompt)

	// READS A STRING ENTERED IN THE CLI, UP UNTIL ENTER (\N) IS PRESSED
	weightInput, _ := reader.ReadString('\n')

	fmt.Println(info.HeightPrompt)
	heightInput, _ := reader.ReadString('\n')

	weightInput = strings.Replace(weightInput, "\n", "", -1)
	heightInput = strings.Replace(heightInput, "\n", "", -1)

	weight, _ := strconv.ParseFloat(weightInput, 64)
	height, _ := strconv.ParseFloat(heightInput, 64)

	bmi := (weight / (height * height)) * 703

	fmt.Printf("BMI = %.2f", bmi)

}
