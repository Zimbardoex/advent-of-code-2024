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

func isRowSafe(row []int) bool {
	var safe bool = true
	var increasing bool
	var decreasing bool

	for i, dp := range row {
		if i == 0 {
			continue
		}

		if row[i-1] > dp {
			decreasing = true
		}

		if row[i-1] < dp {
			increasing = true
		}

		var absoluteDiff = diff(dp, row[i-1])

		if absoluteDiff < 1 || absoluteDiff > 3 {
			safe = false
			break
		}

		if increasing {
			if decreasing {
				safe = false
				break
			}
		}

		if decreasing {
			if increasing {
				safe = false
				break
			}
		}
	}
	return safe
}

func removeIndex(s []int, index int) []int {
	return append(s[:index], s[index+1:]...)
}

func main() {

	var fileData = getFileData("./data.txt")

	var safeRowCount = len(fileData)

	for _, fileRow := range fileData {
		var isOrginalRowSafe = isRowSafe(fileRow)

		if !isOrginalRowSafe {
			var isDampendRowSafe bool

			for index := range fileRow {
				rowToDampen := make([]int, len(fileRow))
				copy(rowToDampen, fileRow)

				var dampendRow = removeIndex(rowToDampen, index)

				isDampendRowSafe = isRowSafe(dampendRow)
				if isDampendRowSafe {
					isDampendRowSafe = true
					break
				}
			}

			if !isDampendRowSafe {
				safeRowCount -= 1
			}
		}
	}

	print("Total safe lines: ", safeRowCount)
}
