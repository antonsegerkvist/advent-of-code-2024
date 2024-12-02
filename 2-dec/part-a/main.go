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
		numbers := strings.Split(scanner.Text(), " ")
		fmt.Println(numbers)
		sign := int64(0)

		a0, _ := strconv.ParseInt(numbers[0], 10, 64)
		b0, _ := strconv.ParseInt(numbers[1], 10, 64)

		if b0-a0 >= 1 && b0-a0 <= 3 {
			sign = 1
		} else if b0-a0 <= -1 && b0-a0 >= -3 {
			sign = -1
		} else {
			continue
		}

		foundIssue := false
		for i := range numbers[1 : len(numbers)-1] {
			a, _ := strconv.ParseInt(numbers[i+1], 10, 64)
			b, _ := strconv.ParseInt(numbers[i+2], 10, 64)

			if sign == 1 && b-a >= 1 && b-a <= 3 {
				continue
			} else if sign == -1 && b-a <= -1 && b-a >= -3 {
				continue
			} else {
				foundIssue = true
				break
			}
		}

		if foundIssue == false {
			count += 1
			fmt.Println(true)
		} else {
			fmt.Println(false)
		}
	}

	fmt.Println(count)
}
