package main

import (
	"bufio"
	"fmt"
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

func main() {
	dat, err := os.Open("./data.txt")
	check(err)

	sc := bufio.NewScanner(dat)

	var list1, list2 []int
	for sc.Scan() {
		row := strings.Split(sc.Text(), "   ")
		list1 = append(list1, convertToInteger(row[0]))
		list2 = append(list2, convertToInteger(row[1]))
	}

	coordinCount := make(map[int]int)

	for _, value := range list2 {
		current := coordinCount[value]
		coordinCount[value] = current + 1
	}

	fmt.Printf("coordinCount: %v\n", coordinCount)

	var similarity int
	for _, value := range list1 {
		similarity += value * coordinCount[value]
	}

	print("Total similarity: ", similarity)
}
