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

	fmt.Println(width, height)
	for _, row := range m {
		fmt.Println(string(row))
	}

	count := 0
	for i, row := range m {
		for j, c := range row {
			if c == 'A' {
				if Search(i, j) {
					count++
				}
			}
		}
	}

	fmt.Println("XMAS count:", count)

}

func Search(row, col int) bool {
	if row <= 0 || row >= height-1 || col <= 0 || col >= width-1 {
		return false
	}

	diag1 := false
	if m[row-1][col-1] == 'M' && m[row+1][col+1] == 'S' {
		diag1 = true
	}
	if m[row-1][col-1] == 'S' && m[row+1][col+1] == 'M' {
		diag1 = true
	}

	diag2 := false
	if m[row-1][col+1] == 'M' && m[row+1][col-1] == 'S' {
		diag2 = true
	}
	if m[row-1][col+1] == 'S' && m[row+1][col-1] == 'M' {
		diag2 = true
	}

	return diag1 && diag2
}
