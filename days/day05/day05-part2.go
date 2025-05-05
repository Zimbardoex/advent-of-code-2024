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
			// println("page to find, page", strings.TrimRight(pageToFind, "\r\n"), strings.TrimRight(page, "\r\n"))
			return index
		}
	}

	return -1
}

func reorderUpdatesUntilAllValid(updates [][]string, pageOrderingRules [][]string) [][]string {
	var allValid = false
	var reorderings = 0
	for allValid != true {
		allValid = true
		for _, update := range updates {
			for index, pageNumInUpdate := range update {
				println("number index, ", index)
				var pageOrderRules = getApplicableRules(pageNumInUpdate, pageOrderingRules)

				if len(pageOrderRules) > 0 {
					for _, rule := range pageOrderRules {
						var rulesPageIndex = findPageIndex(rule[0], update)
						var rulesPageIndex1 = findPageIndex(rule[1], update)

						if rulesPageIndex1 >= 0 && rulesPageIndex1 < rulesPageIndex {
							allValid = false
							update[rulesPageIndex] = rule[1]
							update[rulesPageIndex1] = rule[0]
						}
					}
				}
			}

		}
		reorderings++
		println("reorderings: ", reorderings, allValid)
	}

	return updates
}

func filterUpdates(pageOrderingRules [][]string, updates [][]string) [][]string {
	var invalidUpdates [][]string

	for _, update := range updates {
		var valid = true
		for index, pageNumInUpdate := range update {
			println("number index, ", index)
			var pageOrderRules = getApplicableRules(pageNumInUpdate, pageOrderingRules)

			if len(pageOrderRules) > 0 {
				for _, rule := range pageOrderRules {
					var rulesPageIndex = findPageIndex(rule[0], update)
					var rulesPageIndex1 = findPageIndex(rule[1], update)

					if rulesPageIndex1 >= 0 && rulesPageIndex1 < rulesPageIndex {
						println("index, rule's index", index, rulesPageIndex)
						println("rule to be applied: ", rule[0], "|", rule[1])
						println("update before application: ")
						for _, char := range update {
							println(char)
						}
						valid = false
						update[rulesPageIndex] = rule[1]
						update[rulesPageIndex1] = rule[0]
						println("update after application: ")
						for _, char := range update {
							println(char)
						}
					}
				}
			}
		}

		if !valid {
			invalidUpdates = append(invalidUpdates, update)
		}
	}

	invalidUpdates = reorderUpdatesUntilAllValid(invalidUpdates, pageOrderingRules)
	return invalidUpdates
}

func sumMiddles(updates [][]string) int {

	var sum int
	for _, update := range updates {
		sum += convertToInteger(strings.TrimRight(update[len(update)/2], "\r\n"))
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

	var invalidUpdates = filterUpdates(pageOrdering, updates)

	println("invalid lines: ", len(invalidUpdates))

	var checkingAllValid = filterUpdates(pageOrdering, invalidUpdates)

	println("checking all are valid: ", len(checkingAllValid))

	var sum = sumMiddles(invalidUpdates)
	print("Total of middle number of invalid updates: ", sum)

}
