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

func checkForNormalXmas(fileData string) int {
	var count int
	var lines = strings.Split(fileData, "\n")

	for _, line := range lines {
		for charIndex := range lines {
			if charIndex < len(line)-3 {
				var substring = line[charIndex : charIndex+4]
				println(substring)
				if substring == "XMAS" || substring == "SAMX" {
					count += 1
				}
			}
		}
	}

	println("total normal XMAS: ", count)

	return count
}

func checkForReverseXmas(fileData string) int {
	var count int
	var lines = strings.Split(fileData, "\n")

	for _, line := range lines {
		for charIndex := range lines {
			if charIndex > 3 {
				var substring = line[charIndex-4 : charIndex]
				println(substring)
				if substring == "SAMX" {
					count += 1
				}
			}
		}
	}

	println("total reverse XMAS: ", count)

	return count
}

func extractRightDownDiagonalSubstring(x int, y int, subStringLength int, data []string) string {

	var substring string
	for count := 0; count < subStringLength; count += 1 {
		currentX := x + count
		currentY := y + count

		substring += string(data[currentY][currentX])
	}

	return substring
}

func extractLeftDownDiagonalSubstring(x int, y int, subStringLength int, data []string) string {

	var substring string
	for count := 0; count < subStringLength; count += 1 {
		currentX := x - count
		currentY := y + count

		character := string(data[currentY][currentX])
		substring += character
	}

	return substring
}

func checkForDiagonalNormalXmas(fileData string) int {
	var subStringLength = 4
	var count int
	var lines = strings.Split(fileData, "\n")

	for lineIndex, line := range lines {
		for charIndex := range line {
			if charIndex < len(line)-subStringLength && lineIndex <= len(lines)-subStringLength {
				var substring = extractRightDiagonalSubstring(charIndex, lineIndex, subStringLength, lines)
				if substring == "XMAS" || substring == "SAMX" {
					count += 1
				}
			}

			if charIndex > 3 && charIndex < len(line) && lineIndex <= len(lines)-subStringLength {

				var substring = extractLeftDiagonalSubstring(charIndex-1, lineIndex, subStringLength, lines)

				if substring == "XMAS" || substring == "SAMX" {
					count += 1
				}
			}
		}
	}

	println("total diagonal XMAS: ", count)

	return count
}

func checkForVerticalXmas(fileData string) int {
	var count int
	var lines = strings.Split(fileData, "\n")

	for index := range lines {
		for charIndex := range lines {
			if index < len(lines)-3 {
				var firstChar = string(lines[index][charIndex])
				var secondChar = string(lines[index+1][charIndex])
				var thirdChar = string(lines[index+2][charIndex])
				var forthChar = string(lines[index+3][charIndex])
				var substring = firstChar + secondChar + thirdChar + forthChar
				if substring == "XMAS" || substring == "SAMX" {
					count += 1
				}
			}
		}
	}

	println("total vertical XMAS: ", count)

	return count
}

func main() {

	var total int
	//read file
	var fileData = getFileData("./data.txt")

	total += checkForNormalXmas(fileData)

	total += checkForReverseXmas(fileData)

	total += checkForDiagonalNormalXmas(fileData)

	total += checkForVerticalXmas(fileData)

	print("Total of all xmas's: ", total)

}
