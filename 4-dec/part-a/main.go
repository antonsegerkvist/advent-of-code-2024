package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"unicode/utf8"
)

var lookupString = []rune("XMAS")
var height = 0
var width = 0
var m = [][]rune{}

const (
	LEFT = iota
	BOTTOM_LEFT
	BOTTOM
	BOTTOM_RIGHT
	RIGHT
	TOP_LEFT
	TOP
	TOP_RIGHT
	END
)

func main() {
	filename := os.Args[1]

	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err.Error())
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	for ; scanner.Scan(); height++ {
		line := scanner.Text()
		width = utf8.RuneCountInString(line)
		m = append(m, []rune(line))
	}

	count := 0
	for i, row := range m {
		for j := range row {
			for k := LEFT; k < END; k++ {
				if Search(i, j, 0, k) {
					count++
				}
			}
		}
	}

	fmt.Println("XMAS count:", count)

}

func Search(row, col, sindex int, direction int) bool {
	if row < 0 || row >= height || col < 0 || col >= width {
		return false
	}

	if m[row][col] != lookupString[sindex] {
		return false
	}

	if sindex == len(lookupString)-1 {
		return true
	}

	switch direction {
	case LEFT:
		return Search(row, col-1, sindex+1, direction)
	case BOTTOM_LEFT:
		return Search(row+1, col-1, sindex+1, direction)
	case BOTTOM:
		return Search(row+1, col, sindex+1, direction)
	case BOTTOM_RIGHT:
		return Search(row+1, col+1, sindex+1, direction)
	case RIGHT:
		return Search(row, col+1, sindex+1, direction)
	case TOP_LEFT:
		return Search(row-1, col-1, sindex+1, direction)
	case TOP:
		return Search(row-1, col, sindex+1, direction)
	case TOP_RIGHT:
		return Search(row-1, col+1, sindex+1, direction)
	}
	return false
}
