package main

import (
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
)

func main() {

	filename := os.Args[1]

	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err.Error())
	}
	defer file.Close()

	bytes, err := io.ReadAll(file)
	if err != nil {
		log.Fatal(err.Error())
	}

	input := string(bytes)
	seq := make([]int, 0, 1024)

	for index, r := range input {
		id := index >> 1
		count, err := strconv.ParseUint(string(r), 10, 64)
		if err != nil {
			log.Fatal(err.Error())
		}

		for i := 0; i < int(count); i++ {
			if index%2 == 0 {
				seq = append(seq, id)
			} else {
				seq = append(seq, -1)
			}
		}
	}

	i := 0
	j := len(seq) - 1

	for i < j {
		for seq[i] != -1 && i < j {
			i++
		}
		for seq[j] == -1 && i < j {
			j--
		}
		seq[i] = seq[j]
		seq[j] = -1
	}

	fmt.Println(seq)

	checksum := 0
	for i, v := range seq {
		if v > 0 {
			checksum += i * v
		}
	}
	fmt.Println("checksum:", checksum)
}
