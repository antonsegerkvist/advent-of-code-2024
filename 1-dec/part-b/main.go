package main

import (
	"bufio"
	"errors"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
)

var ErrNoMatch = errors.New("list entry not found")

func main() {
	filename := os.Args[1]

	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err.Error())
	}
	defer file.Close()

	aMap := map[int]int{}
	bMap := map[int]int{}

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		a, b, err := ExtractNumbers(scanner.Text())
		if err != nil {
			log.Fatal(err)
		}

		aMap[a] = aMap[a] + 1
		bMap[b] = bMap[b] + 1
	}

	sum := 0
	for k, v := range aMap {
		sum += v * k * bMap[k]
	}
	fmt.Println(sum)
}

func ExtractNumbers(line string) (int, int, error) {
	r := regexp.MustCompile("(?P<a>\\d+)\\s+(?P<b>\\d+)")

	matches := r.FindStringSubmatch(line)
	if matches == nil {
		return 0, 0, ErrNoMatch
	}

	a := int64(0)
	b := int64(0)
	var err error

	for i, name := range r.SubexpNames() {
		switch name {
		case "a":
			a, err = strconv.ParseInt(matches[i], 10, 32)
			if err != nil {
				return 0, 0, err
			}
		case "b":
			b, err = strconv.ParseInt(matches[i], 10, 32)
			if err != nil {
				return 0, 0, err
			}
		}
	}

	return int(a), int(b), nil
}
