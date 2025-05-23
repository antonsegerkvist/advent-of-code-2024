package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
)

type Score struct {
	path  map[string]bool
	score int
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
	tiles := [][]rune{}
	ri := 0

	for scanner.Scan() {
		line := scanner.Text()
		row := []rune{}

		for ci, c := range line {
			if c == 'S' {
				startRow = ri
				startCol = ci
			}
			if c == 'E' {
				endRow = ri
				endCol = ci
			}

			row = append(row, c)
		}

		tiles = append(tiles, row)
		ri++
	}

	fmt.Println("Start:", startRow, startCol)
	fmt.Println("End:  ", endRow, endCol)

	for ri := 0; ri < len(tiles); ri++ {
		for ci := 0; ci < len(tiles[ri]); ci++ {
			fmt.Printf("%c", tiles[ri][ci])
		}
		fmt.Println("")
	}

	m := &map[string]int{}
	cost := recurse(tiles, m, startRow, startCol, 1, 0)

	fmt.Println("Score:", cost.score)
	totalTiles := 0
	for range cost.path {
		totalTiles++
	}
	fmt.Println("Tiles:", totalTiles)
}

func recurse(tiles [][]rune, m *map[string]int, row, col, dir, cost int) *Score {
	if tiles[row][col] == 'E' {
		return &Score{
			path: map[string]bool{
				fmt.Sprintf("%d,%d", row, col): true,
			},
			score: cost,
		}
	}

	key := fmt.Sprintf("%d,%d,%d", dir, row, col)
	if v, ok := (*m)[key]; ok && v < cost {
		return &Score{
			path:  map[string]bool{},
			score: math.MaxInt,
		}
	} else {
		(*m)[key] = cost
	}

	scores := []*Score{}
	if tiles[row-1][col] != '#' && dir != 2 {
		numRot := minRotations(dir, 0)
		scores = append(scores, recurse(tiles, m, row-1, col, 0, cost+1+numRot*1000))
	}
	if tiles[row+1][col] != '#' && dir != 0 {
		numRot := minRotations(dir, 2)
		scores = append(scores, recurse(tiles, m, row+1, col, 2, cost+1+numRot*1000))
	}
	if tiles[row][col+1] != '#' && dir != 3 {
		numRot := minRotations(dir, 1)
		scores = append(scores, recurse(tiles, m, row, col+1, 1, cost+1+numRot*1000))
	}
	if tiles[row][col-1] != '#' && dir != 1 {
		numRot := minRotations(dir, 3)
		scores = append(scores, recurse(tiles, m, row, col-1, 3, cost+1+numRot*1000))
	}

	minScore := math.MaxInt
	for _, score := range scores {
		if score.score < minScore {
			minScore = score.score
		}
	}

	ret := &Score{
		path: map[string]bool{
			fmt.Sprintf("%d,%d", row, col): true,
		},
		score: minScore,
	}
	for _, score := range scores {
		if score.score == minScore {
			for k, v := range score.path {
				ret.path[k] = v
			}
		}
	}

	return ret

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
