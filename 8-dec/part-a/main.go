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

				antiNode1Row := row2 - row1 + row2
				antiNode1Col := col2 - col1 + col2
				antiNode2Row := row1 - row2 + row1
				antiNode2Col := col1 - col2 + col1

				if antiNode1Row >= 0 && antiNode1Row < rows && antiNode1Col >= 0 && antiNode1Col < cols {
					m[antiNode1Row][antiNode1Col].HasAntinode = true
				}
				if antiNode2Row >= 0 && antiNode2Row < rows && antiNode2Col >= 0 && antiNode2Col < cols {
					m[antiNode2Row][antiNode2Col].HasAntinode = true
				}
			}
		}
	}

	count := 0
	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			if !m[i][j].HasAntinode {
				fmt.Printf("%s", string(m[i][j].Rune))
			} else {
				count++
				fmt.Printf("%s", "#")
			}
		}
		fmt.Print("\n")
	}

	fmt.Println("Unique antinodes:", count)

}
