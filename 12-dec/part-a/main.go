package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

type Tile struct {
	char      rune
	row       int
	col       int
	isVisited bool
}

func main() {

	filename := os.Args[1]

	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err.Error())
	}
	defer file.Close()

	tiles := [][]*Tile{}

	scanner := bufio.NewScanner(file)
	row := 0

	for scanner.Scan() {
		line := scanner.Text()
		cols := []*Tile{}

		for col, c := range line {
			cols = append(cols, &Tile{
				char:      c,
				row:       row,
				col:       col,
				isVisited: false,
			})
		}

		row++
		tiles = append(tiles, cols)
	}

	cost := 0

	for row := 0; row < len(tiles); row++ {
		for col := 0; col < len(tiles[row]); col++ {
			tile := tiles[row][col]
			stack := []*Tile{tile}

			area := 0
			perimiter := 0

			for len(stack) > 0 {
				var curr *Tile
				stack, curr = stack[:len(stack)-1], stack[len(stack)-1]

				if curr.isVisited {
					continue
				}

				area += 1
				curr.isVisited = true

				i := curr.row
				j := curr.col
				c := curr.char

				if i > 0 && tiles[i-1][j].char == c && !tiles[i-1][j].isVisited {
					stack = append(stack, tiles[i-1][j])
				}
				if i == 0 || i > 0 && tiles[i-1][j].char != c {
					perimiter++
				}

				if i < len(tiles)-1 && tiles[i+1][j].char == c && !tiles[i+1][j].isVisited {
					stack = append(stack, tiles[i+1][j])
				}
				if i == len(tiles)-1 || i < len(tiles)-1 && tiles[i+1][j].char != c {
					perimiter++
				}

				if j > 0 && tiles[i][j-1].char == c && !tiles[i][j-1].isVisited {
					stack = append(stack, tiles[i][j-1])
				}
				if j == 0 || j > 0 && tiles[i][j-1].char != c {
					perimiter++
				}

				if j < len(tiles[i])-1 && tiles[i][j+1].char == c && !tiles[i][j+1].isVisited {
					stack = append(stack, tiles[i][j+1])
				}
				if j == len(tiles[i])-1 || j < len(tiles[i])-1 && tiles[i][j+1].char != c {
					perimiter++
				}
			}

			cost += area * perimiter
		}
	}

	fmt.Println("Cost:", cost)

}
