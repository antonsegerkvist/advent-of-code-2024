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
	tcheck    bool
	rcheck    bool
	bcheck    bool
	lcheck    bool
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
				tcheck:    false,
				rcheck:    false,
				bcheck:    false,
				lcheck:    false,
				isVisited: false,
			})
		}

		row++
		tiles = append(tiles, cols)
	}

	cost := 0
	costMap := map[string][]int{}

	for row := 0; row < len(tiles); row++ {
		for col := 0; col < len(tiles[row]); col++ {
			tile := tiles[row][col]
			stack := []*Tile{tile}

			area := 0
			perimiter := 0

			for len(stack) > 0 {
				var curr *Tile
				stack, curr = stack[1:len(stack)], stack[0]

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
					if !tiles[i][j].tcheck {
						fmt.Println("t", i, j)
						perimiter++
					}
					if j > 0 && tiles[i][j-1].char == c {
						tiles[i][j-1].tcheck = true
					}
					if j < len(tiles[i])-1 && tiles[i][j+1].char == c {
						tiles[i][j+1].tcheck = true
					}
				}

				if i < len(tiles)-1 && tiles[i+1][j].char == c && !tiles[i+1][j].isVisited {
					stack = append(stack, tiles[i+1][j])
				}
				if i == len(tiles)-1 || i < len(tiles)-1 && tiles[i+1][j].char != c {
					if !tiles[i][j].bcheck {
						fmt.Println("b", i, j)
						perimiter++
					}
					if j > 0 && tiles[i][j-1].char == c {
						tiles[i][j-1].bcheck = true
					}
					if j < len(tiles[i])-1 && tiles[i][j+1].char == c {
						tiles[i][j+1].bcheck = true
					}
				}

				if j > 0 && tiles[i][j-1].char == c && !tiles[i][j-1].isVisited {
					stack = append(stack, tiles[i][j-1])
				}
				if j == 0 || j > 0 && tiles[i][j-1].char != c {
					if !tiles[i][j].lcheck {
						fmt.Println("l", i, j)
						perimiter++
					}
					if i > 0 && tiles[i-1][j].char == c {
						tiles[i-1][j].lcheck = true
					}
					if i < len(tiles)-1 && tiles[i+1][j].char == c {
						tiles[i+1][j].lcheck = true
					}
				}

				if j < len(tiles[i])-1 && tiles[i][j+1].char == c && !tiles[i][j+1].isVisited {
					stack = append(stack, tiles[i][j+1])
				}
				if j == len(tiles[i])-1 || j < len(tiles[i])-1 && tiles[i][j+1].char != c {
					if !tiles[i][j].rcheck {
						fmt.Println("r", i, j)
						perimiter++
					}
					if i > 0 && tiles[i-1][j].char == c {
						tiles[i-1][j].rcheck = true
					}
					if i < len(tiles)-1 && tiles[i+1][j].char == c {
						tiles[i+1][j].rcheck = true
					}
				}
			}

			if area*perimiter > 0 {
				costMap[fmt.Sprintf("%d,%d", row, col)] = []int{area, perimiter, area * perimiter}
			}
			cost += area * perimiter
		}
	}

	fmt.Println("Cost:", cost)

}
