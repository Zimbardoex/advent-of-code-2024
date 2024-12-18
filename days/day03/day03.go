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

	matches := r.FindAllStringSubmatch(str, -1)

	return matches
}

func multiplyAndSum(multiplications [][]string) int {
	var total int

	for _, match := range multiplications {
		var numbers = strings.Split(match[1], ",")
		var num1 = convertToInteger(numbers[0])
		var num2 = convertToInteger(numbers[1])
		total += num1 * num2
	}

	return total
}

func main() {

	//read file
	var fileData = getFileData("./data.txt")

	// extract the mul(x,y)s
	var multiplications = extractMultiplicaions(fileData)

	// split, muliply and sum
	total := multiplyAndSum(multiplications)

	print("Total of multiplications: ", total)
}
