package main

import "fmt"

func main() {

	fmt.Println(Reciprocal("Svool Dliow! Gsrh Rmkfg HZNKOV! 10, 15"))
}

func Reciprocal(input string) string {
	i := []rune(input)        //Using rune for fetch chars in input
	o := make([]rune, len(i)) //We create slice for our output with capacity input

	for x := range i {
		if IsLower(i[x]) { //Lowercase letters correspond to byte values between 97 and 122
			o = append(o, 219-i[x])
		} else if IsUpper(i[x]) { //Uppercase letters correspond to byte values between 65 and 90
			o = append(o, 155-i[x])
		} else {
			o = append(o, i[x]) //If its not letter
		}

	}

	return string(o)
}

func IsUpper(input rune) bool {
	var u bool
	if input <= 90 && input >= 65 {
		u = true
	} else {
		u = false
	}
	return u
}

func IsLower(input rune) bool {
	var l bool
	if input <= 122 && input >= 97 {
		l = true
	} else {
		l = false
	}
	return l
}
