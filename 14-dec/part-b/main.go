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

const (
	WIDTH  = 101
	HEIGHT = 103
)

type Robot struct {
	X, Y   int
	VX, VY int
}

func main() {

	robots := []*Robot{}
	filename := os.Args[1]

	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err.Error())
	}
	defer file.Close()

	r := regexp.MustCompile("p=(?<x>-?\\d+),(?<y>-?\\d+) v=(?<vx>-?\\d+),(?<vy>-?\\d+)")

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		matches := r.FindStringSubmatch(line)

		robot := Robot{}
		for i, name := range r.SubexpNames() {
			switch name {
			case "x":
				robot.X, err = strconv.Atoi(matches[i])
				if err != nil {
					log.Fatal(err.Error())
				}
			case "y":
				robot.Y, err = strconv.Atoi(matches[i])
				if err != nil {
					log.Fatal(err.Error())
				}
			case "vx":
				robot.VX, err = strconv.Atoi(matches[i])
				if err != nil {
					log.Fatal(err.Error())
				}
			case "vy":
				robot.VY, err = strconv.Atoi(matches[i])
				if err != nil {
					log.Fatal(err.Error())
				}
			}
		}

		robots = append(robots, &robot)
	}

	for i := 0; i < math.MaxInt; i++ {
		for _, robot := range robots {
			robot.X = (robot.X + robot.VX + WIDTH) % WIDTH
			robot.Y = (robot.Y + robot.VY + HEIGHT) % HEIGHT
		}

		metric := 0.
		for j := 0; j < len(robots); j++ {
			for k := j + 1; k < len(robots); k++ {
				r1 := robots[j]
				r2 := robots[k]

				metric += math.Abs(float64(r1.X-r2.X)) + math.Abs(float64(r1.Y-r2.Y))
			}
		}

		if metric < 6800000 {
			fmt.Println("============", i+1)
			printR(&robots)
		}
	}

}

func printR(robots *[]*Robot) {
	m := map[string]int{}
	for _, robot := range *robots {
		m[fmt.Sprintf("%d,%d", robot.X, robot.Y)]++
	}

	for i := 0; i < WIDTH; i++ {
		for j := 0; j < HEIGHT; j++ {
			if v, ok := m[fmt.Sprintf("%d,%d", i, j)]; ok {
				fmt.Print(v)
			} else {
				fmt.Print(".")
			}
		}
		fmt.Println()
	}
	bufio.NewReader(os.Stdin).ReadBytes('\n')
}
