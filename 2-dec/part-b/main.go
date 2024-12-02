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
		log.Fatal(err)
	}
	defer file.Close()

	count := 0

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		numberStrings := strings.Split(scanner.Text(), " ")

		numbers := []int64{}
		for _, v := range numberStrings {
			a, _ := strconv.ParseInt(v, 10, 64)
			numbers = append(numbers, a)
		}

		if valid := testLevels(numbers, 0); valid {
			count++
		}
	}

	fmt.Println(count)
}

func testLevels(numbers []int64, level int) bool {
	sign := 0
	first := numbers[1] - numbers[0]

	if first >= 1 && first <= 3 {
		sign = 1
	} else if first <= -1 && first >= -3 {
		sign = -1
	} else {
		if level == 0 {
			return false ||
				testLevels(numbers[1:], 1) ||
				testLevels(append(append([]int64{}, numbers[:1]...), numbers[2:]...), 1)
		} else {
			return false
		}
	}

	for i := range numbers[1 : len(numbers)-1] {
		j := i + 1
		v := numbers[j+1] - numbers[j]
		if sign == 1 && v >= 1 && v <= 3 {
			continue
		} else if sign == -1 && v <= -1 && v >= -3 {
			continue
		} else {
			if level == 0 {
				return false ||
					testLevels(append(append([]int64{}, numbers[:j-1]...), numbers[j+0:]...), 1) ||
					testLevels(append(append([]int64{}, numbers[:j+0]...), numbers[j+1:]...), 1) ||
					testLevels(append(append([]int64{}, numbers[:j+1]...), numbers[j+2:]...), 1)
			} else {
				return false
			}
		}
	}

	return true
}
