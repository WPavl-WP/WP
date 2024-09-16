package main

import "fmt"

func main() {
	romanNumber := "VIIL"

	//res := 0

	m := make(map[string]int)
	m["I"] = 1
	m["V"] = 5
	m["X"] = 10
	m["L"] = 50
	m["C"] = 100
	m["D"] = 500
	m["M"] = 1000

	output := 0

	for a := len(romanNumber); a > 0; a = a-1 {

		//if romanNumber[a] == romanNumber[a-1] {
		//	output = output + int(romanNumber[a+1]) + int(romanNumber[a+1])
		//}

		//if romanNumber[a]+romanNumber[a-1] < romanNumber[a-2] {
		//	output = output - (int(romanNumber[a])+int(romanNumber[a+1]))*2
		//}
		//if romanNumber[a]+romanNumber[a+1]+romanNumber[a-2] < romanNumber[a-3] {
		//	output = output - (int(romanNumber[a])+int(romanNumber[a+1])+int(romanNumber[a+2]))*2
		//}

		if romanNumber[a] < romanNumber[a-1] && a>0 && a-1 != 0{
			output = output + int(romanNumber[a])
		} else if romanNumber[a] > romanNumber[a-1] && a>0{
			output = output + int(romanNumber[a])
		} else if romanNumber[a] < romanNumber[a-1] && a>0 && a-1=0{
			output = output + int(romanNumber[a-1])	+ int(romanNumber[a])
		} else if a=0 && ...{
				output = output + int(romanNumber[a])
		} else if a=0 && ...{
			output = output - int(romanNumber[a])
		} else {
			fmt.Println("error")
		}

	}

	fmt.Println(output)
}
