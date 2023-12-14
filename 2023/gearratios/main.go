package main

import (
	"bufio"
	"bytes"
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
	part1()
	part2()
}

const ANSWER_TO_THE_UNIVERSE = 42

func part2() {
	gearList := []gear{}
	sum := 0
	file, _ := os.ReadFile("./input.txt")
	rows := bytes.Split(file, []byte("\n"))
	rowMaxId := len(rows[0]) - 1
	columnMaxId := len(rows) - 1
	for rowId := range rows {
		for columnId := range rows[rowId] {
			if rows[rowId][columnId] == ANSWER_TO_THE_UNIVERSE {
				g := gear{
					row:      rowId,
					column:   columnId,
					adjacent: 0,
					values:   []Value{},
				}
				g.findAdjacent(rows, rowMaxId, columnMaxId)
				if g.adjacent == 2 {
					gearList = append(gearList, g)
				}

			}
		}
	}
	for gId := range gearList {
		v, _ := strconv.Atoi(string(gearList[gId].values[0].value))

		v2, _ := strconv.Atoi(string(gearList[gId].values[1].value))
		sum = sum + (v * v2)
	}
	println("Sum of gears : ", sum)
}
func part1() {
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
			for idx := range vals[lineNum] {
				for i := vals[lineNum][idx].positionStart - 1; i <= vals[lineNum][idx].positionEnd+1; i++ {
					if i == -1 {

						continue
					}
					if i > len(previousLine[ind])-1 {
						continue
					}
					// My worst mistake, and a dumb one of course :
					// I didn't realise 47 was /... so the line below obviously didn't work
					// if previousLine[i] < 46 || previousLine[i] > 57 {
					if previousLine[ind][i] != 46 && (previousLine[ind][i] < 48 || previousLine[ind][i] > 57) {
						vals[lineNum][idx].exists = true

					}
				}
			}
		}

		lineNum++
		ind++
		// GOD I HATE THE SCANNER
		// This line NEEDS to be a clone, else it will not work. The scanner sometimes override the value. (every 3920 character or 28 lines,in my case)
		previousLine = append(previousLine, bytes.Clone(line))
		// What I still don't understand :
		// Why 28 lines ? We're in a loop, so it's kind of weird that it'd overwrite it with a line 28 lines later, not just the line below. Maybe Scan actually initially scans the
		// entire file, then processes the line individually and concurrently ?
		// Also, does that mean that scan has already the next tokens ready even though the for() loop is not finished yet ?
	}
	for o := range vals {
		for o2 := range vals[o] {
			if vals[o][o2].exists {
				val, err := strconv.Atoi(string(vals[o][o2].value))
				if err != nil {

					println("Error processing values :", err.Error())
				}
				sum = sum + val
			}
		}
	}
	println("sum : ", sum)
}

type gear struct {
	row      int
	column   int
	adjacent int
	values   []Value
}

func isDigit(b byte) bool {
	if b > 47 && b < 58 {
		return true
	}
	return false
}

// This is a bit of an horror because I treated each case separately, and I copypasted some logic out of laziness. But it's nothing very complicated to solve for the machine !
func (t *gear) findAdjacent(rows [][]byte, maxRow int, maxColumn int) {
	top := t.row - 1
	bottom := t.row + 1
	left := t.column - 1
	right := t.column + 1
	vals := []Value{}
	topVals := []Value{}
	bottomVals := []Value{}
	if left > 0 {
		if isDigit(rows[t.row][left]) {
			t.adjacent++

			leftVal := Value{
				value: []byte{rows[t.row][left]},
			}
			i := left
			for {
				i = i - 1
				if i < 0 || !isDigit(rows[t.row][i]) {
					break
				}
				leftVal.value = append(leftVal.value, rows[t.row][i])
			}

			for i2, j := 0, len(leftVal.value)-1; i2 < j; i2, j = i2+1, j-1 {
				leftVal.value[i2], leftVal.value[j] = leftVal.value[j], leftVal.value[i2]
			}
			vals = append(vals, leftVal)

		}
	}
	if right <= maxColumn {
		if isDigit(rows[t.row][right]) {
			t.adjacent++

			rightVal := Value{
				value: []byte{rows[t.row][right]},
			}
			i := right
			for {
				i = i + 1
				if i > maxColumn || !isDigit(rows[t.row][i]) {
					break
				}
				rightVal.positionEnd = i
				rightVal.value = append(rightVal.value, rows[t.row][i])
			}

			vals = append(vals, rightVal)

		}
	}

	if top >= 0 {
		i := t.column
		j := left
		leftVal := Value{}
		rightVal := Value{}
		if isDigit(rows[top][left]) {
			t.adjacent++
			leftVal.value = append(leftVal.value, rows[top][left])
			for {
				j = j - 1
				if j < 0 || !isDigit(rows[top][j]) {
					break
				}
				leftVal.value = append(leftVal.value, rows[top][j])
			}

			for i2, j2 := 0, len(leftVal.value)-1; i2 < j2; i2, j2 = i2+1, j2-1 {
				leftVal.value[i2], leftVal.value[j2] = leftVal.value[j2], leftVal.value[i2]
			}

			topVals = append(topVals, leftVal)
		}

		if isDigit(rows[top][i]) {
			rightVal.value = append(rightVal.value, rows[top][i])
			k := i
			for {
				k = k + 1
				if k > maxColumn || !isDigit(rows[top][k]) {
					break
				}
				rightVal.value = append(rightVal.value, rows[top][k])
			}

			if len(topVals) > 0 {
				topVals[0].value = append(topVals[0].value, rightVal.value...)
			} else {

				t.adjacent++
				topVals = append(topVals, rightVal)
			}
		} else if isDigit(rows[top][right]) {
			t.adjacent++
			rightVal.value = append(rightVal.value, rows[top][right])
			k := right
			for {
				k = k + 1
				if k > maxColumn || !isDigit(rows[top][k]) {
					break
				}
				rightVal.value = append(rightVal.value, rows[top][k])
			}

			topVals = append(topVals, rightVal)
		}
	}

	if bottom <= maxRow {
		i := t.column
		j := left
		leftVal := Value{}
		rightVal := Value{}
		if isDigit(rows[bottom][left]) {
			t.adjacent++
			leftVal.value = append(leftVal.value, rows[bottom][left])
			for {
				j = j - 1
				if j < 0 || !isDigit(rows[bottom][j]) {
					break
				}
				leftVal.value = append(leftVal.value, rows[bottom][j])
			}

			for i2, j2 := 0, len(leftVal.value)-1; i2 < j2; i2, j2 = i2+1, j2-1 {
				leftVal.value[i2], leftVal.value[j2] = leftVal.value[j2], leftVal.value[i2]
			}

			bottomVals = append(bottomVals, leftVal)
		}

		if isDigit(rows[bottom][i]) {
			rightVal.value = append(rightVal.value, rows[bottom][i])
			k := i
			for {
				k = k + 1
				if k > maxColumn || !isDigit(rows[bottom][k]) {
					break
				}
				rightVal.value = append(rightVal.value, rows[bottom][k])
			}

			if len(bottomVals) > 0 {
				bottomVals[0].value = append(bottomVals[0].value, rightVal.value...)
			} else {

				t.adjacent++
				bottomVals = append(bottomVals, rightVal)
			}
		} else if isDigit(rows[bottom][right]) {
			t.adjacent++
			rightVal.value = append(rightVal.value, rows[bottom][right])
			k := right
			for {
				k = k + 1
				if k > maxColumn || !isDigit(rows[bottom][k]) {
					break
				}
				rightVal.value = append(rightVal.value, rows[bottom][k])
			}

			bottomVals = append(bottomVals, rightVal)
		}
	}
	vals = append(vals, bottomVals...)
	vals = append(vals, topVals...)
	t.values = vals
}
