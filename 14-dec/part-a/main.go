package main

import (
	"bufio"
	"fmt"
	"log"
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

	for i := 0; i < 100; i++ {
		for _, robot := range robots {
			robot.X = (robot.X + robot.VX + WIDTH) % WIDTH
			robot.Y = (robot.Y + robot.VY + HEIGHT) % HEIGHT
		}
	}

	q1, q2, q3, q4 := 0, 0, 0, 0
	for _, robot := range robots {
		qx := WIDTH >> 1
		qy := HEIGHT >> 1

		if robot.X < qx && robot.Y < qy {
			q1++
		} else if robot.X < qx && robot.Y > (qy) {
			q4++
		} else if robot.X > (qx) && robot.Y > (qy) {
			q3++
		} else if robot.X > (qx) && robot.Y < qy {
			q2++
		}

		fmt.Println(*robot)
	}

	fmt.Println("q1:", q1, "q2:", q2, "q3:", q3, "q4:", q4)
	fmt.Println("score:", q1*q2*q3*q4)

}
