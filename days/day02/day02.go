package main

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func convertToInteger(str string) int {
	i, err := strconv.Atoi(str)
	if err != nil {
		print("string %s couldn't be parsed as an integer", str)
		panic(err)
	}

	return i
}

func diff(a, b int) int {
	if a < b {
		return b - a
	}
	return a - b
}

func getFileData(filepath string) [][]int {
	dat, err := os.Open(filepath)
	check(err)

	sc := bufio.NewScanner(dat)

	var fileData [][]int

	for sc.Scan() {

		row := strings.Split(sc.Text(), " ")

		var intRow []int
		for _, value := range row {
			intRow = append(intRow, convertToInteger(value))
		}

		fileData = append(fileData, intRow)
	}

	return fileData
}

func main() {

	var fileData = getFileData("./data.txt")

	var safeRowCount = len(fileData)

	for _, fileRow := range fileData {
		var increasing bool
		var decreasing bool

		for i, dp := range fileRow {
			if i == 0 {
				continue
			}

			if fileRow[i-1] > dp {
				decreasing = true
			}

			if fileRow[i-1] < dp {
				increasing = true
			}

			var absoluteDiff = diff(dp, fileRow[i-1])

			if absoluteDiff < 1 || absoluteDiff > 3 {
				safeRowCount -= 1
				break
			}

			if increasing {
				if decreasing {
					safeRowCount -= 1
					break
				}
			}

			if decreasing {
				if increasing {
					safeRowCount -= 1
					break
				}
			}
		}
	}

	print("Total safe lines: ", safeRowCount)
}
