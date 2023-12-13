package main

import (
	"bufio"
	"os"
	"strconv"
)

type Value struct {
	positionStart int
	positionEnd   int
	value         []byte
	exists        bool
	started       bool
}

func main() {
	file, _ := os.Open("./input.txt")
	scanner := bufio.NewScanner(file)
	// Lines are not longer than 65 000 character so :
	var previousLine [][]byte
	sum := 0
	lineNum := 0
	ind := -1
	vals := [][]Value{}
	for scanner.Scan() {

		line := scanner.Bytes()
		lastLineSymbolIdx := -3
		latest := Value{}
		for pos, characterByte := range line {
			vals = append(vals, []Value{})
			if characterByte == 46 {
				if latest.started && latest.positionEnd == pos-1 {
					vals[lineNum] = append(vals[lineNum], latest)
					latest = Value{}
				}
				continue
			}
			//
			if characterByte > 47 && characterByte < 58 {
				latest.value = append(latest.value, characterByte)
				if !latest.started {
					latest.started = true
					latest.positionStart = pos
				}
				if latest.positionStart == lastLineSymbolIdx+1 {
					latest.exists = true
				}
				latest.positionEnd = pos

				//This one was a bit tricky. I forgot that I was only adding numbers when we hit an identifier...
				// Which obviously doesn't work when the number is at the end of the line.
				if pos == len(line)-1 {
					vals[lineNum] = append(vals[lineNum], latest)
					latest = Value{}
				}
				continue
			}
			//if character is a symbol :
			if latest.started && latest.positionEnd == pos-1 {
				latest.exists = true

				vals[lineNum] = append(vals[lineNum], latest)
				latest = Value{}
			}
			lastLineSymbolIdx = pos
			if ind != -1 {
				for idx := range vals[ind] {
					if pos >= vals[ind][idx].positionStart-1 && pos <= vals[ind][idx].positionEnd+1 {
						vals[ind][idx].exists = true
					}
				}
			}
		}
		if ind != -1 {
			if lineNum == 28 {
				println("first faulty line")
			}
			for idx := range vals[lineNum] {
				// if lineNum == 28 && string(vals[lineNum][idx].value) == "807" {
				// 	println("faulty value")
				// }
				for i := vals[lineNum][idx].positionStart - 1; i <= vals[lineNum][idx].positionEnd+1; i++ {
					if i == -1 {

						continue
					}
					if i > len(previousLine[ind])-1 {
						continue
					}
					if lineNum == 28 && string(vals[lineNum][idx].value) == "807" {
						println("len ", len(previousLine))
						println("accessing previousline ", ind, "==>", string(previousLine[ind-1]))
					}
					// My worst mistake, and a dumb one of course :
					// I didn't realie 47 was /...
					// if previousLine[i] < 46 || previousLine[i] > 57 {
					if previousLine[ind][i] != 46 && (previousLine[ind][i] < 48 || previousLine[ind][i] > 57) {
						vals[lineNum][idx].exists = true

					}
				}
			}
		}
		lineNum++
		ind++
		// println(lineNum, string(previousLine))
		// println("appending line", lineNum)
		// println(len(previousLine))
		previousLine = append(previousLine, line)
	}
	for o := range vals {
		for o2 := range vals[o] {
			if vals[o][o2].exists {
				val, err := strconv.Atoi(string(vals[o][o2].value))
				if err != nil {

					println("ERROR")
				}
				sum = sum + val
			} else {
				// println("doesn't exist : L :", o, "=>", string(vals[o][o2].value))
			}
		}
	}

	println("27 ==>", string(previousLine[27]))
	println("###########################################")

	println("sum : ", sum)
}
