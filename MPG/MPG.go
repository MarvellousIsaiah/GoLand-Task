package main

import (
	"fmt"
	"math"
)

func main() {
	milePerGallon := 0.0
	totalMpg := 0.0
	for index := 1; index < math.MaxInt; index++ {
		fmt.Println("enter the miles driven")
		var milesDriven float64
		fmt.Scanln(&milesDriven)
		if milesDriven == -1 {
			break
		}
		fmt.Println("enter the gallon used for this trip ")
		var gallonUsed float64
		fmt.Scanln(&gallonUsed)
		milePerGallon = milesDriven / gallonUsed
		fmt.Println("mpg for journey", index, "is", milePerGallon)
		totalMpg += milePerGallon

	}
	fmt.Println("the total mpg of these journry is", totalMpg)

}
