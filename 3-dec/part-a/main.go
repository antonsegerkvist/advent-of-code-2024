package main

import (
	"fmt"
	"io"
	"log"
	"os"
	"regexp"
	"strconv"
)

func main() {

	filename := os.Args[1]

	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	b, err := io.ReadAll(file)
	if err != nil {
		log.Fatal(err.Error())
	}

	input := string(b)
	r := regexp.MustCompile(`mul\((?P<a>\d+),(?P<b>\d+)\)`)
	matches := r.FindAllStringSubmatch(input, -1)
	names := r.SubexpNames()

	fmt.Println(names)
	fmt.Println(matches)

	var sum int64 = 0

	for _, match := range matches {
		a, b := int64(0), int64(0)
		for i, v := range r.SubexpNames() {
			switch v {
			case "a":
				a, _ = strconv.ParseInt(match[i], 10, 64)
			case "b":
				b, _ = strconv.ParseInt(match[i], 10, 64)
			}
		}

		sum += a * b
	}

	fmt.Println(sum)

}
