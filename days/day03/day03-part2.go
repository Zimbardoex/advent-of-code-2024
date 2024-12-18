package main

import (
	"os"
	"regexp"
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

func getFileData(filepath string) string {
	dat, err := os.ReadFile(filepath)
	check(err)

	return string(dat)
}

func extractMultiplicaions(str string) [][]string {
	r := regexp.MustCompile(`mul\((?P<multi>[0-9]+\,[0-9]+)\)`)

	return r.FindAllStringSubmatch(str, -1)
}

func extractDoAndDonts(str string) []string {
	r := regexp.MustCompile(`do\(\)|don\'t\(\)`)

	return r.FindAllString(str, -1)
}

func splitOnDosAndDonts(str string) []string {
	println(str)
	array := regexp.MustCompile(`(do(n't)?\(\))`).Split(str, -1)
	return array
}

func multiplyAndSum(dataSplitOnDosAndDonts []string, dosAndDonts []string) int {
	var total int

	var multiplications = extractMultiplicaions(dataSplitOnDosAndDonts[0])
	for _, match := range multiplications {
		var numbers = strings.Split(match[1], ",")
		var num1 = convertToInteger(numbers[0])
		var num2 = convertToInteger(numbers[1])
		total += num1 * num2
	}

	for index, doOrDont := range dosAndDonts {
		if doOrDont == "do()" {
			var multiplications = extractMultiplicaions(dataSplitOnDosAndDonts[index+1])

			for _, match := range multiplications {
				var numbers = strings.Split(match[1], ",")
				var num1 = convertToInteger(numbers[0])
				var num2 = convertToInteger(numbers[1])
				total += num1 * num2
			}
		}
	}

	return total
}

func main() {

	//read file
	var fileData = getFileData("./data.txt")

	// get ordered array of dos and don'ts from file data
	var dosAndDonts = extractDoAndDonts(fileData)

	// split data into chunck at "do()" and "don't" delimiters
	var dataSplitOnDosAndDonts = splitOnDosAndDonts(fileData)

	// muliply and sum based on do and don't order
	total := multiplyAndSum(dataSplitOnDosAndDonts, dosAndDonts)

	print("Total of multiplications: ", total)
}
