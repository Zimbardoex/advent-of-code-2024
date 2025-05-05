package main

import (
	"os"
	"strings"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func getFileData(filepath string) string {
	dat, err := os.ReadFile(filepath)
	check(err)

	return string(dat)
}

func extractRightDiagonalSubstring(x int, y int, data []string) string {
	println("x, y: ", x, y)

	var substring string
	for count := 0; count < 3; count += 1 {
		currentY := y - 1 + count
		currentX := x - 1 + count

		substring += string(data[currentY][currentX])
	}

	return substring
}

func extractLeftDiagonalSubstring(x int, y int, data []string) string {
	// println("x, y: ", x, y)

	var substring string
	for count := 0; count < 3; count += 1 {
		currentX := x + 1 - count
		currentY := y - 1 + count

		character := string(data[currentY][currentX])
		substring += character
	}

	return substring
}

func checkForDiagonalMas(fileData string) int {
	var subStringLength = 2
	var count int
	var lines = strings.Split(fileData, "\n")

	for lineIndex, line := range lines {
		for charIndex := range line {
			if charIndex < len(line)-subStringLength && charIndex > 0 && lineIndex <= len(lines)-subStringLength && lineIndex > 0 {
				var substring = extractRightDiagonalSubstring(charIndex, lineIndex, lines)
				println("substring 1: ", substring)
				if substring == "MAS" || substring == "SAM" {
					if charIndex > 0 && charIndex < len(line) && lineIndex <= len(lines)-subStringLength && lineIndex > 0 {

						var substring1 = extractLeftDiagonalSubstring(charIndex, lineIndex, lines)
						println("substring 2: ", substring1)

						if substring1 == "MAS" || substring1 == "SAM" {
							count += 1
						}
					}
				}

			}

		}
	}

	println("total MAS cross: ", count)

	return count
}

func main() {

	var total int
	//read file
	var fileData = getFileData("./data.txt")

	total += checkForDiagonalMas(fileData)

	print("Total of all mas x's: ", total)

}
