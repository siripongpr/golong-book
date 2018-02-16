package main

import "fmt"

func main() {
	fmt.Println(weatherCelsius(25, "Bangkok few cloud"))
	fmt.Println(weatherCelsius(34, "Tak sunny"))
	fmt.Println(weatherCelsius(17, "Phuket rainy"))
	fmt.Println(weatherCelsius(9, "Chiang-mai cold"))
}

func weatherCelsius(number int, description string) (display string) {
	var numdisplay1, numdisplay2, numdisplay3 string
	secondDigit := number%10
	firstDigit := (number - secondDigit) / 10
	numdisplay1 = numdisplay(firstDigit, 1) + numdisplay(secondDigit, 1)
	numdisplay2 = numdisplay(firstDigit, 2) + numdisplay(secondDigit, 2)
	numdisplay3 = numdisplay(firstDigit, 3) + numdisplay(secondDigit, 3)
	display = fmt.Sprintf("%s\n%s\n%s c\n%v c\n%s\n", numdisplay1, numdisplay2, numdisplay3, number, description)
	return display
}

func numdisplay(number int, line int) (display string) {
	if number == 1 {
		if line == 1 {
			display = "   "
		} else if line == 2 {
			display = "  |"
		} else {
			display = "  |"
		}
	} else if number == 2 {
		if line == 1 {
			display = " _ "
		} else if line == 2 {
			display = " _|"
		} else {
			display = "|_ "
		}
	} else if number == 3 {
		if line == 1 {
			display = " _ "
		} else if line == 2 {
			display = " _|"
		} else {
			display = " _|"
		}
	} else if number == 4 {
		if line == 1 {
			display = "   "
		} else if line == 2 {
			display = "|_|"
		} else {
			display = "  |"
		}
	} else if number == 5 {
		if line == 1 {
			display = " _ "
		} else if line == 2 {
			display = "|_"
		} else {
			display = " _|"
		}
	} else if number == 7 {
		if line == 1 {
			display = " _ "
		} else if line == 2 {
			display = "  |"
		} else {
			display = "  |"
		}
	} else if number == 9 {
		if line == 1 {
			display = " _ "
		} else if line == 2 {
			display = "|_|"
		} else {
			display = " _|"
		}
	} else {
		display = ""
	}
	return display
}