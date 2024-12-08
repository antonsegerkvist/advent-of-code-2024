package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

type Node struct {
	Rune        rune
	HasAntinode bool
}

type Antenna struct {
	Row int
	Col int
}

func main() {

	filename := os.Args[1]

	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err.Error())
	}
	defer file.Close()

	m := [][]Node{}

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()

		row := []Node{}
		for _, c := range line {
			row = append(row, Node{
				Rune: c,
			})
		}

		m = append(m, row)
	}

	rows, cols := len(m), len(m[0])
	fmt.Println("rows x cols:", rows, cols)

	runeLocations := map[rune][]Antenna{}
	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			r := m[i][j].Rune
			if r != '.' {
				runeLocations[r] = append(runeLocations[r], Antenna{
					Row: i,
					Col: j,
				})
			}
		}
	}

	for _, v := range runeLocations {
		for i := 0; i < len(v); i++ {
			for j := i + 1; j < len(v); j++ {
				row1, col1 := v[i].Row, v[i].Col
				row2, col2 := v[j].Row, v[j].Col

				dRow := row2 - row1
				dCol := col2 - col1
				antiNodeRow := row2
				antiNodeCol := col2

				curRow := antiNodeRow
				curCol := antiNodeCol

				for k := 0; curRow < rows && curCol < cols && curRow >= 0 && curCol >= 0; k++ {
					m[curRow][curCol].HasAntinode = true
					curRow = antiNodeRow + k*dRow
					curCol = antiNodeCol + k*dCol
				}

				curRow = antiNodeRow - dRow
				curCol = antiNodeCol - dCol

				for k := 0; curRow < rows && curCol < cols && curRow >= 0 && curCol >= 0; k-- {
					m[curRow][curCol].HasAntinode = true
					curRow = antiNodeRow + k*dRow
					curCol = antiNodeCol + k*dCol
				}
			}
		}
	}

	count := 0
	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			if m[i][j].HasAntinode && m[i][j].Rune != '.' {
				count++
				fmt.Printf("%s", string(m[i][j].Rune))
			} else if m[i][j].HasAntinode {
				count++
				fmt.Printf("%s", "#")
			} else {
				fmt.Printf("%s", string(m[i][j].Rune))
			}
		}
		fmt.Print("\n")
	}

	fmt.Println("Unique antinodes:", count)

}
