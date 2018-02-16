package main

import "fmt"

func main() {
	fmt.Println(weatherCelsius(25, "Bangkok few cloud"))
	fmt.Println(weatherCelsius(34, "Tak sunny"))
	fmt.Println(weatherCelsius(17, "Phuket rainy"))
	fmt.Println(weatherCelsius(9, "Chiang-mai cold"))
}

func weatherCelsius(number int, description string) (display string) {
	var numdisplay string
	numdisplay = fmt.Sprintf("_  _\n_||_\n")
	secondDigit := number%10
	firstDigit := number - secondDigit
	fmt.Printf("======%v\n", firstDigit)
	fmt.Printf("======%v\n", secondDigit)
	display = fmt.Sprintf("%s\n%v c\n%s\n", numdisplay, number, description)
	return display
}

func numdisplay(number int, line int) (display string) {
	if number == 2 {
		if line == 1 {
			display = " _ "
		} else if line == 2 {
			display = " _|"
		} else {
			display = "|_ "
		}
	}
	return display
}