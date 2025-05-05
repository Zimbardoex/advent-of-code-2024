package main

import (
	"os"
	"strconv"
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

func convertToInteger(str string) int {
	i, err := strconv.Atoi(str)
	if err != nil {
		print("string %s couldn't be parsed as an integer", str)
		panic(err)
	}

	return i
}

func getApplicableRules(pageNum string, pageOrdering [][]string) [][]string {
	var rules [][]string
	for _, pageOrder := range pageOrdering {
		if strings.TrimRight(pageNum, "\r\n") == strings.TrimRight(pageOrder[0], "\r\n") {
			rules = append(rules, pageOrder)
		}
	}

	return rules
}

func findPageIndex(pageToFind string, updateToSearch []string) int {
	for index, page := range updateToSearch {

		if strings.TrimRight(page, "\r\n") == strings.TrimRight(pageToFind, "\r\n") {
			return index
		}
	}

	return -1
}

func filterUpdates(pageOrderingRules [][]string, updates [][]string) [][]string {
	var validUpdates [][]string

	for _, update := range updates {
		var valid = true
		for index, pageNumInUpdate := range update {
			var pageOrderRules = getApplicableRules(pageNumInUpdate, pageOrderingRules)

			if len(pageOrderRules) > 0 {
				for _, rule := range pageOrderRules {
					var rulesPageIndex = findPageIndex(rule[1], update)

					if rulesPageIndex >= 0 && rulesPageIndex < index {
						valid = false
						// TODO: could break here, what will be "broken" back to?
					}
				}
			}
		}

		if valid {
			validUpdates = append(validUpdates, update)
		}
	}
	return validUpdates
}

func sumMiddles(updates [][]string) int {

	var sum int
	for _, update := range updates {
		sum += convertToInteger(update[len(update)/2])
	}

	return sum
}

func extractData(orderings string, delimiter string) [][]string {
	var extractedData [][]string
	var lines = strings.Split(orderings, "\n")
	for _, line := range lines {
		extractedData = append(extractedData, strings.Split(line, delimiter))
	}

	return extractedData
}

func main() {
	//read files
	var rawOrdering = getFileData("./ordering.txt")
	var rawUpdates = getFileData("./updates.txt")

	var pageOrdering = extractData(rawOrdering, "|")

	var updates = extractData(rawUpdates, ",")

	var validUpdates = filterUpdates(pageOrdering, updates)

	println("valid lines: ", len(validUpdates))

	var sum = sumMiddles(validUpdates)
	print("Total of middle number of valid updates: ", sum)

}
