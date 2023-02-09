package main

import (
	"bufio"
	"fmt" //standard package for using println
	"os"
	"strings"
	"time" //standard package for using time
)

func main() {

	currentTime := time.Now().Hour() //gets the current time in hours
	// Hard-coded list of license plates
	allowedPlates := []string{"ABC123", "DEF456", "GHI789"}

	// Read the license plate from the user
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter your license plate number: ")
	input, _ := reader.ReadString('\n')
	input = strings.TrimSpace(input)

	// Check if the license plate is in the allowed list
	isAllowed := false
	for _, plate := range allowedPlates {
		if input == plate {
			isAllowed = true
			break
		}
	}
	if isAllowed == true {
		if currentTime >= 7 && currentTime < 12 { //check the current time and if between set arguments prints the underlying line
			fmt.Println("Goedemorgen! Welkom bij Fonteyn vakantieparken")
		} else if currentTime >= 12 && currentTime < 18 { //check the current time and if between set arguments prints the underlying line
			fmt.Println("Goedemiddag! Welkom bij Fonteyn vakantieparken")
		} else if currentTime >= 18 && currentTime < 23 { //check the current time and if between set arguments prints the underlying line
			fmt.Println("Goedenavond! Welkom bij Fonteyn vakantieparken")
		} else if currentTime >= 23 && currentTime < 7 { //check the current time and if between set arguments prints the underlying line
			fmt.Println("Sorry, de parkeerplaats is gesloten")
		}
		// Show the appropriate message based on the result
	} else {
		fmt.Println("U heeft helaas geen toegang tot het parkeerterrein")
	}
}
