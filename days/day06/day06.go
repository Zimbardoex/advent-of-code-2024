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

func processFileData(data string) []string {

	var warehouseGrid = strings.Split(data, "\n")

	for i, line := range warehouseGrid {
		warehouseGrid[i] = strings.TrimRight(line, "\r\n")
	}
	return warehouseGrid
}

func getStartingLocation(warehouseData []string) (int, int) {

	for yIndex, line := range warehouseData {
		for xIndex, char := range line {
			if char == '^' {
				return xIndex, yIndex
			}
		}
	}

	panic("starting location couldn't be found!")
}

func getCountOfUniqueVisits(startingX int, startingY int, warehouseGrid []string) int {
	xVel := 0
	yVel := -1

	var x = startingX
	var y = startingY
	var nextX, nextY = nextPosition(x, y, xVel, yVel)

	var visits = 0

	var vistedLocations = make(map[string]bool)

	for checkPostitionValid(nextX, nextY, warehouseGrid) {
		// 35 = '#'
		if warehouseGrid[nextY][nextX] == 35 {
			if yVel == -1 {
				// up -> right
				xVel = 1
				yVel = 0
			} else if xVel == 1 {
				// right -> down
				xVel = 0
				yVel = 1
			} else if yVel == 1 {
				// down -> left
				xVel = -1
				yVel = 0
			} else if xVel == -1 {
				// left -> up
				xVel = 0
				yVel = -1
			}
		}
		vistedLocations[strconv.Itoa(x)+"."+strconv.Itoa(y)] = true
		visits++
		x, y = nextPosition(x, y, xVel, yVel)
		nextX, nextY = nextPosition(x, y, xVel, yVel)
	}

	return len(vistedLocations) + 1
}

func nextPosition(x int, y int, xVel int, yVel int) (int, int) {
	return x + xVel, y + yVel
}

func checkPostitionValid(x int, y int, warehouseGrid []string) bool {
	if y < 0 || y > len(warehouseGrid)-1 {
		return false
	}
	return x >= 0 && x < len(warehouseGrid[y])
}

func main() {
	//read files
	var warehouseData = getFileData("./data.txt")

	var warehouseGrid = processFileData(warehouseData)

	var x, y = getStartingLocation(warehouseGrid)

	var countOfVisits = getCountOfUniqueVisits(x, y, warehouseGrid)

	print("locations visted: ", countOfVisits)
}
