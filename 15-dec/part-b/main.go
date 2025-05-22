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
				pCol = ci * 2
			}

			switch c {
			case '.':
				row = append(row,
					&Tile{
						row: ri,
						col: ci * 2,
						c:   ".",
					},
					&Tile{
						row: ri,
						col: ci*2 + 1,
						c:   ".",
					},
				)
			case '#':
				row = append(row,
					&Tile{
						row: ri,
						col: ci * 2,
						c:   "#",
					},
					&Tile{
						row: ri,
						col: ci*2 + 1,
						c:   "#",
					},
				)
			case 'O':
				row = append(row,
					&Tile{
						row: ri,
						col: ci * 2,
						c:   "[",
					},
					&Tile{
						row: ri,
						col: ci*2 + 1,
						c:   "]",
					},
				)
			case '@':
				row = append(row,
					&Tile{
						row: ri,
						col: ci * 2,
						c:   "@",
					},
					&Tile{
						row: ri,
						col: ci*2 + 1,
						c:   ".",
					},
				)
			}
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
		canMove := canMove(tiles, o, pRow, pCol)
		if canMove {
			doStep(tiles, o, pRow, pCol)
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
			if c == "[" {
				score += 100*i + j
			}
		}
		fmt.Println("")
	}

	fmt.Println("Score:", score)

}

func canMove(tiles [][]*Tile, op string, row, col int) bool {
	c := tiles[row][col].c
	if c == "." {
		return true
	}

	if c == "#" {
		return false
	}

	if c == "@" {
		switch op {
		case "^":
			return canMove(tiles, op, row-1, col)
		case ">":
			return canMove(tiles, op, row, col+1)
		case "v":
			return canMove(tiles, op, row+1, col)
		case "<":
			return canMove(tiles, op, row, col-1)
		}
	}

	switch op {
	case "^":
		if c == "[" {
			return canMove(tiles, op, row-1, col) && canMove(tiles, op, row-1, col+1)
		} else if c == "]" {
			return canMove(tiles, op, row-1, col) && canMove(tiles, op, row-1, col-1)
		}
	case ">":
		return canMove(tiles, op, row, col+2)
	case "<":
		return canMove(tiles, op, row, col-2)
	case "v":
		if c == "[" {
			return canMove(tiles, op, row+1, col) && canMove(tiles, op, row+1, col+1)
		} else if c == "]" {
			return canMove(tiles, op, row+1, col) && canMove(tiles, op, row+1, col-1)
		}
	}

	log.Fatal("Invalid op: ", op)
	return false
}

func doStep(tiles [][]*Tile, op string, row, col int) {
	c := tiles[row][col].c
	if c == "." {
		return
	}

	if c == "#" {
		return
	}

	if c == "@" {
		switch op {
		case "^":
			doStep(tiles, op, row-1, col)
			tiles[row-1][col].c, tiles[row][col].c = tiles[row][col].c, tiles[row-1][col].c
			return
		case ">":
			doStep(tiles, op, row, col+1)
			tiles[row][col+1].c, tiles[row][col].c = tiles[row][col].c, tiles[row][col+1].c
			return
		case "v":
			doStep(tiles, op, row+1, col)
			tiles[row+1][col].c, tiles[row][col].c = tiles[row][col].c, tiles[row+1][col].c
			return
		case "<":
			doStep(tiles, op, row, col-1)
			tiles[row][col-1].c, tiles[row][col].c = tiles[row][col].c, tiles[row][col-1].c
			return
		}
	}

	switch op {
	case "^":
		if c == "[" {
			doStep(tiles, op, row-1, col)
			doStep(tiles, op, row-1, col+1)
			tiles[row-1][col+0].c, tiles[row][col+0].c = tiles[row][col+0].c, tiles[row-1][col+0].c
			tiles[row-1][col+1].c, tiles[row][col+1].c = tiles[row][col+1].c, tiles[row-1][col+1].c
		} else if c == "]" {
			doStep(tiles, op, row-1, col)
			doStep(tiles, op, row-1, col-1)
			tiles[row-1][col+0].c, tiles[row][col+0].c = tiles[row][col+0].c, tiles[row-1][col+0].c
			tiles[row-1][col-1].c, tiles[row][col-1].c = tiles[row][col-1].c, tiles[row-1][col-1].c
		}
		return
	case ">":
		doStep(tiles, op, row, col+1)
		tiles[row][col+1].c, tiles[row][col].c = tiles[row][col].c, tiles[row][col+1].c
		return
	case "v":
		if c == "[" {
			doStep(tiles, op, row+1, col)
			doStep(tiles, op, row+1, col+1)
			tiles[row+1][col+0].c, tiles[row][col+0].c = tiles[row][col+0].c, tiles[row+1][col+0].c
			tiles[row+1][col+1].c, tiles[row][col+1].c = tiles[row][col+1].c, tiles[row+1][col+1].c
		} else if c == "]" {
			doStep(tiles, op, row+1, col)
			doStep(tiles, op, row+1, col-1)
			tiles[row+1][col+0].c, tiles[row][col+0].c = tiles[row][col+0].c, tiles[row+1][col+0].c
			tiles[row+1][col-1].c, tiles[row][col-1].c = tiles[row][col-1].c, tiles[row+1][col-1].c
		}
		return
	case "<":
		doStep(tiles, op, row, col-1)
		tiles[row][col-1].c, tiles[row][col].c = tiles[row][col].c, tiles[row][col-1].c
		return
	}

	log.Fatal("Unknown op")
}
