package main

import (
	"fmt"
	"io"
	"log"
	"os"
	"slices"
	"strconv"
)

type File struct {
	ID    int
	Index int
	Size  int
}

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
	files := make([]File, 0, 1024)
	size := 0

	for i, r := range input {
		id := i >> 1

		count, err := strconv.ParseInt(string(r), 10, 64)
		if err != nil {
			log.Fatal(err.Error())
		}

		if i%2 == 0 {
			files = append(files, File{
				ID:    id,
				Index: size,
				Size:  int(count),
			})
		} else {
			files = append(files, File{
				ID:    -1,
				Index: size,
				Size:  int(count),
			})
		}

		size += int(count)
	}

	// PrintFiles(files)

	for j := len(files) - 1; j > 0; j-- {
		if files[j].ID == -1 {
			continue
		}

		for i := 1; i < j; i++ {
			if files[i].ID != -1 {
				continue
			}

			if files[i].Size == files[j].Size {
				files[i].ID = files[j].ID
				files[j].ID = -1
				break
			} else if files[i].Size > files[j].Size {
				newFile := File{
					ID:    -1,
					Index: files[i].Index + files[j].Size,
					Size:  files[i].Size - files[j].Size,
				}

				files[i].ID = files[j].ID
				files[i].Size = files[j].Size

				files[j].ID = -1

				files = slices.Insert(files, i+1, newFile)
				j += 1
				break
			}
		}

		// PrintFiles(files)
	}

	checksum := 0
	for _, file := range files {
		if file.ID == -1 {
			continue
		}

		for i := file.Index; i < file.Index+file.Size; i++ {
			checksum += file.ID * i
		}
	}

	fmt.Println(checksum)
}

func PrintFiles(files []File) {
	for _, f := range files {
		for i := 0; i < f.Size; i++ {
			if f.ID == -1 {
				fmt.Print(".")
			} else {
				fmt.Print(f.ID)
			}
		}
	}
	fmt.Println("")
}
