package main

import (
	"bufio"
	"strconv"
	// "bytes"
	// "encoding/binary"
	// "math/big"
	"os"
)

const (
	ZERO  = "zero"
	ONE   = "one"
	TWO   = "two"
	THREE = "three"
	FOUR  = "four"
	FIVE  = "five"
	SIX   = "six"
	SEVEN = "seven"
	EIGHT = "eight"
	NINE  = "nine"
)

var equivSlice = []byte("0123456789")
var strSlice = [][]byte{
	[]byte(ZERO),
	[]byte(ONE),
	[]byte(TWO),
	[]byte(THREE),
	[]byte(FOUR),
	[]byte(FIVE),
	[]byte(SIX), []byte(SEVEN), []byte(EIGHT), []byte(NINE)}

func matchesForward(line []byte, i int, j int) bool {
	for k := range strSlice[j] {
		if line[i+k] != strSlice[j][k] {
			return false
		}
	}
	return true
}
func matchesBackward(line []byte, i int, j int) bool {
	for k := range strSlice[j] {
		if line[len(line)-1-i-k] != strSlice[j][len(strSlice[j])-1-k] {
			return false
		}
	}
	return true
}

func main() {
	file, _ := os.Open("./input.txt")
	scanner := bufio.NewScanner(file)
	// Lines are not longer than 65 000 character so :
	sum := 0
	forwardWorked := 0
	backwardsWorked := 0
	for scanner.Scan() {
		line := scanner.Bytes()
		l := len(line)
		maxIndex := l - 1
		var first byte
		var firstFound bool
		var last byte
		var lastFound bool

		for i := range line {
			for j := range strSlice {
				strLength := len(strSlice[j])
				if strSlice[j][0] == line[i] {
					if strLength > l-i {
						continue
					}
					firstFound = matchesForward(line, i, j)

				}
				if firstFound {
					first = equivSlice[j]
					break
				}
			}
			if firstFound {
				break
			}
			if line[i] <= 57 && 48 <= line[i] {
				first = line[i]

				forwardWorked++
				break
			}
		}
		// Could do a single loop but with more conditions
		for i := range line {
			for j := range strSlice {
				strLength := len(strSlice[j])
				if strSlice[j][strLength-1] == line[len(line)-1-i] {
					if len(line)-strLength+1 < 0 {
						continue
					}
					lastFound = matchesBackward(line, i, j)

				}
				if lastFound {
					backwardsWorked++
					last = equivSlice[j]
					break
				}
			}
			if lastFound {
				break
			}

			if line[maxIndex-i] <= 57 && 48 <= line[maxIndex-i] {
				last = line[maxIndex-i]
				break
			}
		}
		lineNum := []byte{first, last}
		val, _ := strconv.Atoi(string(lineNum))
		sum = sum + val

	}
	println(forwardWorked, backwardsWorked)
	println("Total sum : ", sum)

}
