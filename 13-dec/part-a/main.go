package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"regexp"
	"strconv"
)

type Entry struct {
	ax, ay int
	bx, by int
	px, py int
}

func main() {

	filename := os.Args[1]

	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err.Error())
	}
	defer file.Close()

	a := regexp.MustCompile("Button A: X\\+(?<x>\\d+), Y\\+(?<y>\\d+)")
	b := regexp.MustCompile("Button B: X\\+(?<x>\\d+), Y\\+(?<y>\\d+)")
	p := regexp.MustCompile("Prize: X=(?<x>\\d+), Y=(?<y>\\d+)")

	section := 0
	entries := []Entry{{}}

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()

		switch section % 4 {
		case 0:
			{
				matches := a.FindStringSubmatch(line)
				if len(matches) <= 0 {
					log.Fatal("No match")
				}

				for i, name := range a.SubexpNames() {
					switch name {
					case "x":
						x := matches[i]
						xi, err := strconv.ParseInt(x, 10, 64)
						if err != nil {
							log.Fatal(err.Error())
						}
						entries[len(entries)-1].ax = int(xi)
					case "y":
						y := matches[i]
						yi, err := strconv.ParseInt(y, 10, 64)
						if err != nil {
							log.Fatal(err.Error())
						}
						entries[len(entries)-1].ay = int(yi)
					}
				}
			}

		case 1:
			{
				matches := b.FindStringSubmatch(line)
				if len(matches) <= 0 {
					log.Fatal("No match")
				}

				for i, name := range b.SubexpNames() {
					switch name {
					case "x":
						x := matches[i]
						xi, err := strconv.ParseInt(x, 10, 64)
						if err != nil {
							log.Fatal(err.Error())
						}
						entries[len(entries)-1].bx = int(xi)
					case "y":
						y := matches[i]
						yi, err := strconv.ParseInt(y, 10, 64)
						if err != nil {
							log.Fatal(err.Error())
						}
						entries[len(entries)-1].by = int(yi)
					}
				}
			}

		case 2:
			{
				matches := p.FindStringSubmatch(line)
				for i, name := range p.SubexpNames() {
					switch name {
					case "x":
						x := matches[i]
						xi, err := strconv.ParseInt(x, 10, 64)
						if err != nil {
							log.Fatal(err.Error())
						}
						entries[len(entries)-1].px = int(xi)
					case "y":
						y := matches[i]
						yi, err := strconv.ParseInt(y, 10, 64)
						if err != nil {
							log.Fatal(err.Error())
						}
						entries[len(entries)-1].py = int(yi)
					}
				}
			}
		case 3:
			entries = append(entries, Entry{})
		}

		section++
	}

	cost := 0
	wins := 0

	for _, entry := range entries {
		costMap := map[string]int{}

		for i := 0; i <= 100; i++ {
			for j := 0; j <= 100; j++ {
				x := entry.ax*i + entry.bx*j
				y := entry.ay*i + entry.by*j
				if x == entry.px && y == entry.py {
					costMap[fmt.Sprintf("%d,%d", i, j)] = i*3 + j*1
				}
			}
		}

		lc := math.MaxInt
		for _, v := range costMap {
			if v < lc {
				lc = v
			}
		}

		if lc < math.MaxInt {
			cost += lc
		}
	}

	fmt.Println("==========================================")
	fmt.Println("Wins:", wins)
	fmt.Println("Cost:", cost)

}
