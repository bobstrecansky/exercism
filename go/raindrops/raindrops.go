package raindrops

import "strconv"

// Convert - AKA Why did they try and cover up FizzBuzz?
func Convert(in int) (out string) {
	returnString := ""
	if in%3 == 0 {
		returnString += "Pling"
	}
	if in%5 == 0 {
		returnString += "Plang"
	}
	if in%7 == 0 {
		returnString += "Plong"
	}
	if returnString == "" {
		returnString = strconv.Itoa(in)
	}
	return returnString
}
