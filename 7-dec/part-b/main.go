package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

const (
	ADD = iota
	MULTIPLY
	CONCAT
)

type Node struct {
	Value    int64
	Operator int
}

func main() {

	filename := os.Args[1]

	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err.Error())
	}
	defer file.Close()

	calibration := int64(0)

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()

		values := strings.Split(line, " ")
		target := values[0]
		target = target[:len(target)-1]
		values = values[1:]

		equal, _ := strconv.ParseInt(target, 10, 64)
		nodes := []Node{}

		for _, v := range values {
			p, _ := strconv.ParseInt(v, 10, 64)
			nodes = append(nodes, Node{
				Value: p,
			})
		}

		if CheckCombination(equal, nodes, 0) {
			calibration += equal
		}

		fmt.Println(target, "=", values)
	}

	fmt.Println("Calibration:", calibration)

}

func CheckCombination(target int64, nodes []Node, index int) bool {
	if index == len(nodes) {
		val := int64(0)
		for i := 0; i < len(nodes)-1; i++ {
			switch nodes[i].Operator {
			case ADD:
				if i == 0 {
					val = nodes[i].Value + nodes[i+1].Value
				} else {
					val = val + nodes[i+1].Value
				}
			case MULTIPLY:
				if i == 0 {
					val = nodes[i].Value * nodes[i+1].Value
				} else {
					val = val * nodes[i+1].Value
				}
			case CONCAT:
				if i == 0 {
					val, _ = strconv.ParseInt(
						fmt.Sprintf("%d%d", nodes[i].Value, nodes[i+1].Value),
						10,
						64,
					)
				} else {
					val, _ = strconv.ParseInt(
						fmt.Sprintf("%d%d", val, nodes[i+1].Value),
						10,
						64,
					)
				}
			}
		}
		return val == target
	}

	nodes[index].Operator = ADD
	if CheckCombination(target, nodes, index+1) {
		return true
	}

	nodes[index].Operator = MULTIPLY
	if CheckCombination(target, nodes, index+1) {
		return true
	}

	nodes[index].Operator = CONCAT
	if CheckCombination(target, nodes, index+1) {
		return true
	}
	return false
}
