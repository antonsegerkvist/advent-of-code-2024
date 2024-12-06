package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

const (
	TOP = iota
	RIGHT
	BOTTOM
	LEFT
	END
)

type Guard struct {
	Direction int
	Row       int
	Col       int
	Finished  bool
}

type Node struct {
	HasObstruction bool
}

func main() {

	filename := os.Args[1]

	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	m := [][]Node{}
	g := Guard{}

	scanner := bufio.NewScanner(file)
	row := 0

	for ; scanner.Scan(); row++ {
		line := scanner.Text()

		values := []Node{}
		for col, c := range line {
			switch c {
			case '#':
				values = append(values, Node{HasObstruction: true})
			case '^', '>', '<', 'v':
				values = append(values, Node{HasObstruction: false})

				g.Col = col
				g.Row = row
				switch c {
				case '^':
					g.Direction = TOP
				case '>':
					g.Direction = RIGHT
				case '<':
					g.Direction = LEFT
				case 'v':
					g.Direction = BOTTOM
				}
			default:
				values = append(values, Node{HasObstruction: false})
			}
		}

		m = append(m, values)
	}

	count := 0
	for i, row := range m {
		for j, node := range row {
			if node.HasObstruction == false {
				g0 := g
				s := map[string]bool{}

				m[i][j].HasObstruction = true
				for !s[GuardHash(&g0)] && !g0.Finished {
					s[GuardHash(&g0)] = true

					UpdateGuardState(m, &g0)
				}
				m[i][j].HasObstruction = false

				if !g0.Finished {
					count++
				}
			}
		}
	}

	fmt.Printf("%v\n", g)
	for _, row := range m {
		for _, n := range row {
			if n.HasObstruction {
				fmt.Printf("#")
			} else {
				fmt.Printf(".")
			}
		}
		fmt.Println("")
	}

	fmt.Println("Obstruction posibilities:", count)
}

func UpdateGuardState(m [][]Node, g *Guard) {
	switch g.Direction {
	case TOP:
		if g.Row <= 0 {
			g.Finished = true
		} else if m[g.Row-1][g.Col].HasObstruction {
			g.Direction = RIGHT
		} else {
			g.Row--
		}
	case LEFT:
		if g.Col <= 0 {
			g.Finished = true
		} else if m[g.Row][g.Col-1].HasObstruction {
			g.Direction = TOP
		} else {
			g.Col--
		}
	case BOTTOM:
		if g.Row >= len(m)-1 {
			g.Finished = true
		} else if m[g.Row+1][g.Col].HasObstruction {
			g.Direction = LEFT
		} else {
			g.Row++
		}
	case RIGHT:
		if g.Col >= len(m[0])-1 {
			g.Finished = true
		} else if m[g.Row][g.Col+1].HasObstruction {
			g.Direction = BOTTOM
		} else {
			g.Col++
		}
	}
}

func GuardHash(g *Guard) string {
	return "" +
		strconv.FormatInt(int64(g.Row), 10) +
		"," +
		strconv.FormatInt(int64(g.Col), 10) +
		"," +
		strconv.FormatInt(int64(g.Direction), 10)
}
