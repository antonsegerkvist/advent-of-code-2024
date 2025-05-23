package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
)

type Tile struct {
	char rune
	col  int
	row  int
}

var (
	startRow = 0
	startCol = 0
	endRow   = 0
	endCol   = 0
)

func main() {

	filename := os.Args[1]

	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err.Error())
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	tiles := [][]*Tile{}
	ri := 0

	for scanner.Scan() {
		line := scanner.Text()
		row := []*Tile{}

		for ci, c := range line {
			if c == 'S' {
				startRow = ri
				startCol = ci
			}
			if c == 'E' {
				endRow = ri
				endCol = ci
			}

			row = append(row, &Tile{
				char: c,
				row:  ri,
				col:  ci,
			})
		}

		tiles = append(tiles, row)
		ri++
	}

	fmt.Println("Start:", startRow, startCol)
	fmt.Println("End:  ", endRow, endCol)

	for ri := 0; ri < len(tiles); ri++ {
		for ci := 0; ci < len(tiles[ri]); ci++ {
			fmt.Printf("%c", tiles[ri][ci].char)
		}
		fmt.Println("")
	}

	m := &map[string]int{}
	cost := recurse(tiles, m, startRow, startCol, 1, 0)
	fmt.Println("Score:", cost)
}

func recurse(tiles [][]*Tile, m *map[string]int, row, col, dir, cost int) int {
	if tiles[row][col].char == 'E' {
		return cost
	}

	key := fmt.Sprintf("%d,%d,%d", dir, row, col)
	if v, ok := (*m)[key]; ok && v <= cost {
		return math.MaxInt
	} else {
		(*m)[key] = cost
	}

	scores := []int{}

	if tiles[row-1][col].char != '#' {
		numRot := minRotations(dir, 0)
		scores = append(scores, recurse(tiles, m, row-1, col, 0, cost+1+numRot*1000))
	}
	if tiles[row+1][col].char != '#' {
		numRot := minRotations(dir, 2)
		scores = append(scores, recurse(tiles, m, row+1, col, 2, cost+1+numRot*1000))
	}
	if tiles[row][col+1].char != '#' {
		numRot := minRotations(dir, 1)
		scores = append(scores, recurse(tiles, m, row, col+1, 1, cost+1+numRot*1000))
	}
	if tiles[row][col-1].char != '#' {
		numRot := minRotations(dir, 3)
		scores = append(scores, recurse(tiles, m, row, col-1, 3, cost+1+numRot*1000))
	}

	return min(scores...)
}

func minRotations(old, new int) int {
	n := abs(new - old)
	if n > 2 {
		n = 4 - n
	}
	return n
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func min(x ...int) int {
	if len(x) == 0 {
		return math.MaxInt
	}

	min := x[0]
	for _, v := range x[1:] {
		if v < min {
			min = v
		}
	}
	return min
}
