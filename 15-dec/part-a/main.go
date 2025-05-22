package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

type Tile struct {
	row, col int
	c        string
}

func main() {

	filename := os.Args[1]

	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err.Error())
	}
	defer file.Close()

	// Map
	pRow := 0
	pCol := 0
	tiles := [][]*Tile{}

	scanner := bufio.NewScanner(file)
	ri := 0

	for scanner.Scan() {
		line := scanner.Text()
		if strings.Trim(line, " ") == "" {
			break
		}

		row := []*Tile{}
		for ci, c := range line {
			if c == '@' {
				pRow = ri
				pCol = ci
			}
			row = append(row, &Tile{
				row: ri,
				col: ci,
				c:   string([]rune{c}),
			})
		}

		tiles = append(tiles, row)
		ri++
	}

	// Instructions
	ops := []string{}
	for scanner.Scan() {
		line := scanner.Text()
		for _, v := range strings.Split(line, "") {
			ops = append(ops, v)
		}
	}

	// Perform instructions.
	for _, o := range ops {
		canMove := doStep(tiles, o, pRow, pCol)
		if canMove {
			switch o {
			case "^":
				pRow--
			case ">":
				pCol++
			case "v":
				pRow++
			case "<":
				pCol--
			}
		}
	}

	score := 0
	for i := 0; i < len(tiles); i++ {
		for j := 0; j < len(tiles[i]); j++ {
			fmt.Printf("%s", tiles[i][j].c)
			c := tiles[i][j].c
			if c == "O" {
				score += 100*i + j
			}
		}
		fmt.Println("")
	}

	fmt.Println("Score:", score)

}

func doStep(tiles [][]*Tile, op string, row, col int) bool {
	c := tiles[row][col].c
	if c == "." {
		return true
	}

	if c == "#" {
		return false
	}

	switch op {
	case "^":
		ret := doStep(tiles, op, row-1, col)
		if ret {
			tiles[row-1][col].c, tiles[row][col].c = tiles[row][col].c, tiles[row-1][col].c
		}
		return ret
	case ">":
		ret := doStep(tiles, op, row, col+1)
		if ret {
			tiles[row][col+1].c, tiles[row][col].c = tiles[row][col].c, tiles[row][col+1].c
		}
		return ret
	case "v":
		ret := doStep(tiles, op, row+1, col)
		if ret {
			tiles[row+1][col].c, tiles[row][col].c = tiles[row][col].c, tiles[row+1][col].c
		}
		return ret
	case "<":
		ret := doStep(tiles, op, row, col-1)
		if ret {
			tiles[row][col-1].c, tiles[row][col].c = tiles[row][col].c, tiles[row][col-1].c
		}
		return ret
	}

	log.Fatal("Unknown op")
	return false
}
