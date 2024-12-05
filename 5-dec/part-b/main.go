package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

var lessTable = map[int64]map[int64]bool{}
var moreTable = map[int64]map[int64]bool{}

func main() {

	filename := os.Args[1]

	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err.Error())
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		numbers := strings.Split(line, "|")

		if len(numbers) <= 1 {
			break
		}

		a, _ := strconv.ParseInt(numbers[0], 10, 64)
		b, _ := strconv.ParseInt(numbers[1], 10, 64)

		if _, ok := lessTable[a]; ok {
			lessTable[a][b] = true
		} else {
			lessTable[a] = map[int64]bool{b: true}
		}

		if _, ok := moreTable[b]; ok {
			moreTable[b][a] = true
		} else {
			moreTable[b] = map[int64]bool{a: true}
		}
	}

	fmt.Println(lessTable)
	fmt.Println(moreTable)

	sum := int64(0)

	for scanner.Scan() {
		line := scanner.Text()
		list := strings.Split(line, ",")
		numbers := []int64{}

		for _, v := range list {
			n, _ := strconv.ParseInt(v, 10, 64)
			numbers = append(numbers, n)
		}

		changes := SortNumbers(numbers)
		fmt.Println("numbers:", numbers)
		fmt.Println("changes:", changes)

		if changes > 0 {
			sum += numbers[len(numbers)>>1]
		}
	}

	fmt.Println("Middle sum:", sum)
}

func SortNumbers(arr []int64) int {
	changes := 0

	for i := 0; i < len(arr); i++ {
		for j := i + 1; j < len(arr); j++ {
			a := arr[i]
			b := arr[j]

			if m, ok := lessTable[a]; ok && m[b] == true {
				continue
			} else if m, ok := moreTable[a]; ok && m[b] == true {
				arr[i], arr[j] = arr[j], arr[i]
				changes++
			}
		}
	}

	return changes
}
