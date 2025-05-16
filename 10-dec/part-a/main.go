package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

type Tile struct {
	c         rune
	row       int
	col       int
	isVisited bool
}

type Node struct {
	row int
	col int
	c   rune
}

func main() {

	filename := os.Args[1]

	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err.Error())
	}
	defer file.Close()

	rowCount := 0
	tiles := [][]Tile{}
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()
		row := []Tile{}

		for col, c := range line {
			row = append(row, Tile{
				c:         c,
				row:       rowCount,
				col:       col,
				isVisited: false,
			})
		}

		tiles = append(tiles, row)
		rowCount++
	}

	stack := []*Tile{}
	total := 0

	for row := 0; row < len(tiles); row++ {
		for col := 0; col < len(tiles[row]); col++ {
			t := &tiles[row][col]
			score := 0
			t.isVisited = true

			if t.c != '0' {
				continue
			}

			for i := 0; i < len(tiles); i++ {
				for j := 0; j < len(tiles[i]); j++ {
					tiles[i][j].isVisited = false
				}
			}

			stack = append(stack, t)
			for len(stack) > 0 {
				var t0 *Tile
				t0, stack = stack[len(stack)-1], stack[:len(stack)-1]

				if t0.c == '9' {
					score += 1
					continue
				}

				if t0.col > 0 && tiles[t0.row][t0.col-1].isVisited == false && tiles[t0.row][t0.col-1].c == (t0.c+1) {
					tiles[t0.row][t0.col-1].isVisited = true
					stack = append(stack, &tiles[t0.row][t0.col-1])
				}
				if t0.col < len(tiles[t0.row])-1 && tiles[t0.row][t0.col+1].isVisited == false && tiles[t0.row][t0.col+1].c == (t0.c+1) {
					tiles[t0.row][t0.col+1].isVisited = true
					stack = append(stack, &tiles[t0.row][t0.col+1])
				}
				if t0.row > 0 && tiles[t0.row-1][t0.col].isVisited == false && tiles[t0.row-1][t0.col].c == (t0.c+1) {
					tiles[t0.row-1][t0.col].isVisited = true
					stack = append(stack, &tiles[t0.row-1][t0.col])
				}
				if t0.row < len(tiles)-1 && tiles[t0.row+1][t0.col].isVisited == false && tiles[t0.row+1][t0.col].c == (t0.c+1) {
					tiles[t0.row+1][t0.col].isVisited = true
					stack = append(stack, &tiles[t0.row+1][t0.col])
				}
			}

			total += score
		}
	}

	fmt.Println("The total score is:", total)

}
