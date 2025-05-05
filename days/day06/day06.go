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

func getCountOfVisits(startingX int, startingY int, warehouseGrid []string) int {
	var xVel = 0
	var yVel = -1

	var x = startingX
	var y = startingY
	var nextX, nextY = nextPosition(x, y, xVel, yVel)

	var visits = 0

	for checkPostitionValid(nextX, nextY, warehouseGrid) {
		if warehouseGrid[nextY][nextX] == '#' {
			// up -> right
			if yVel == -1 {
				xVel = 1
				yVel = 0
			}

			// right -> down
			if xVel == 1 {
				xVel = 0
				yVel = 1
			}

			// down -> left
			if yVel == 1 {
				xVel = -1
				yVel = 0
			}

			// left -> up
			if xVel == -1 {
				xVel = 0
				yVel = -1
			}

		}
		visits++
		x = nextX
		y = nextY

		print(" |x and y: ", x, " ", y)
		nextX, nextY = nextPosition(x, y, xVel, yVel)
		print(" |next x and y: ", nextX, " ", nextY)
	}

	return visits
}

func nextPosition(x int, y int, xVel int, yVel int) (int, int) {
	return x + xVel, y + yVel
}

func checkPostitionValid(x int, y int, warehouseGrid []string) bool {
	return x > 0 && x < len(warehouseGrid[y]) && y > 0 && y < len(warehouseGrid)
}

func main() {
	//read files
	var warehouseData = getFileData("./data.txt")

	var warehouseGrid = strings.Split(warehouseData, "\n")

	var x, y = getStartingLocation(warehouseGrid)

	var countOfVisits = getCountOfVisits(x, y, warehouseGrid)

	print("locations visted: ", countOfVisits)

}
