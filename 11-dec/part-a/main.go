package main

import (
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
	"strings"
	"unicode/utf8"
)

func main() {

	filename := os.Args[1]
	steps, err := strconv.Atoi(os.Args[2])
	if err != nil {
		log.Fatal(err.Error())
	}

	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err.Error())
	}
	defer file.Close()

	b, err := io.ReadAll(file)
	input := string(b)

	curr := strings.Split(input, " ")

	for i := 0; i < steps; i++ {
		next := []string{}

		for _, s := range curr {
			len := utf8.RuneCountInString(s)
			if s == "0" {
				next = append(next, "1")
			} else if len%2 == 0 {
				nl := s[:len>>1]
				nr := s[len>>1:]

				intNL, err := strconv.ParseInt(nl, 10, 64)
				if err != nil {
					log.Fatal(err.Error())
				}
				intNR, err := strconv.ParseInt(nr, 10, 64)
				if err != nil {
					log.Fatal(err.Error())
				}

				nl = strconv.FormatInt(intNL, 10)
				nr = strconv.FormatInt(intNR, 10)

				next = append(next, nl, nr)
			} else {
				n, err := strconv.ParseInt(s, 10, 64)
				if err != nil {
					log.Fatal(err.Error())
				}
				ns := strconv.FormatInt(n*2024, 10)
				next = append(next, ns)
			}
		}

		curr = next
	}

	fmt.Println(len(curr))

}
