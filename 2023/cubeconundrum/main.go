package main

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

const (
	RED   = 12
	GREEN = 13
	BLUE  = 14
)

// Very unoptimized
func underLimit(shows []string) bool {

	for show := range shows {
		colors := strings.Split(shows[show], ",")
		for _, color := range colors {
			if strings.Contains(color, "red") {
				s := strings.Replace(color, " red", "", 1)
				s = strings.Replace(s, " ", "", -1)
				nb, _ := strconv.Atoi(s)
				println(s)
				if nb > RED {
					return false
				}
				continue
			}
			if strings.Contains(color, "green") {
				s := strings.Replace(color, " green", "", 1)
				s = strings.Replace(s, " ", "", -1)
				nb, _ := strconv.Atoi(s)

				println(s)
				if nb > GREEN {
					return false
				}
				continue
			}
			if strings.Contains(color, "blue") {
				s := strings.Replace(color, " blue", "", 1)
				s = strings.Replace(s, " ", "", -1)
				nb, _ := strconv.Atoi(s)
				println(s)
				if nb > BLUE {
					return false
				}
				continue
			}
		}

	}
	return true
}
func findPower(shows []string) int {
	red, blue, green := 0, 0, 0
	for show := range shows {
		colors := strings.Split(shows[show], ",")
		for _, color := range colors {
			if strings.Contains(color, "red") {
				s := strings.Replace(color, " red", "", 1)
				s = strings.Replace(s, " ", "", -1)
				nb, _ := strconv.Atoi(s)
				if nb > red {
					red = nb
				}
				continue
			}
			if strings.Contains(color, "green") {
				s := strings.Replace(color, " green", "", 1)
				s = strings.Replace(s, " ", "", -1)
				nb, _ := strconv.Atoi(s)

				if nb > green {
					green = nb
				}
				continue
			}
			if strings.Contains(color, "blue") {
				s := strings.Replace(color, " blue", "", 1)
				s = strings.Replace(s, " ", "", -1)
				nb, _ := strconv.Atoi(s)
				if nb > blue {
					blue = nb
				}
				continue
			}
		}

	}
	return red * blue * green
}
func main() {

	file, _ := os.Open("./input.txt")
	scanner := bufio.NewScanner(file)
	sum := 0
	sumPower := 0
	// Lines are not longer than 65 000 character so :
	for scanner.Scan() {
		line := scanner.Text()

		game := strings.Split(line, ":")
		input := game[1]
		number := strings.Replace(game[0], "Game ", "", 1)
		id, _ := strconv.Atoi(number)
		shows := strings.Split(input, ";")
		if underLimit(shows) {
			sum = sum + id
		}
		p := findPower(shows)
		sumPower = sumPower + p
	}
	println("Total sum : ", sum)
	println("Total power sum : ", sumPower)
}
