package main

import (
	"fmt"  //standard package for using println
	"time" //standard package for using time
)

func main() {
	currentTime := time.Now().Hour() //gets the current time in hours

	if currentTime >= 7 && currentTime <= 12 { //check the current time and if between set arguments prints the underlying line
		fmt.Println("Goedemorgen! Welkom bij Fonteyn vakantieparken")
	} else if currentTime >= 12 && currentTime <= 18 { //check the current time and if between set arguments prints the underlying line
		fmt.Println("Goedemiddag! Welkom bij Fonteyn vakantieparken")
	} else if currentTime >= 18 && currentTime <= 23 { //check the current time and if between set arguments prints the underlying line
		fmt.Println("Goedenavond! Welkom bij Fonteyn vakantieparken")
	}
}
