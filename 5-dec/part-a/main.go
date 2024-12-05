package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {

	filename := os.Args[1]

	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err.Error())
	}
	defer file.Close()

	ruleTable := map[int64]map[int64]bool{}
	invRuleTable := map[int64]map[int64]bool{}
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()
		numbers := strings.Split(line, "|")

		if len(numbers) <= 1 {
			break
		}

		a, _ := strconv.ParseInt(numbers[0], 10, 64)
		b, _ := strconv.ParseInt(numbers[1], 10, 64)

		if _, ok := ruleTable[a]; ok {
			ruleTable[a][b] = true
		} else {
			ruleTable[a] = map[int64]bool{b: true}
		}

		if _, ok := invRuleTable[b]; ok {
			invRuleTable[b][a] = true
		} else {
			invRuleTable[b] = map[int64]bool{a: true}
		}
	}

	fmt.Println(ruleTable)
	fmt.Println(invRuleTable)

	sum := int64(0)

	for scanner.Scan() {
		line := scanner.Text()
		list := strings.Split(line, ",")
		numbers := []int64{}

		for _, v := range list {
			n, _ := strconv.ParseInt(v, 10, 64)
			numbers = append(numbers, n)
		}

		issueFound := false
		for i := 0; i < len(numbers); i++ {
			for j := i + 1; j < len(numbers); j++ {
				a := numbers[i]
				b := numbers[j]

				if valids, ok := invRuleTable[a]; ok && valids[b] == true {
					issueFound = true
				}
			}
		}

		if issueFound {
			continue
		}

		fmt.Println("Numbers:", numbers)
		fmt.Println("Valid")
		sum += numbers[len(numbers)>>1]
	}

	fmt.Println("Middle sum:", sum)

}
