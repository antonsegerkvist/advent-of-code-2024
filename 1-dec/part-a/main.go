package main

import (
	"bufio"
	"container/heap"
	"errors"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
)

var ErrNoMatch = errors.New("list entry not found")

type Heap []int

func (h Heap) Len() int           { return len(h) }
func (h Heap) Less(i, j int) bool { return h[i] < h[j] }
func (h Heap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *Heap) Push(x any) {
	*h = append(*h, x.(int))
}

func (h *Heap) Pop() any {
	tmp := *h
	n := len(tmp)
	x := tmp[n-1]
	*h = tmp[0 : n-1]
	return x
}

func main() {
	filename := os.Args[1]

	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err.Error())
	}
	defer file.Close()

	aHeap := &Heap{}
	bHeap := &Heap{}

	heap.Init(aHeap)
	heap.Init(bHeap)

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		a, b, err := ExtractNumbers(scanner.Text())
		if err != nil {
			log.Fatal(err)
		}

		heap.Push(aHeap, a)
		heap.Push(bHeap, b)
	}

	a, b, sum := 0, 0, 0

	for aHeap.Len() > 0 {
		a = heap.Pop(aHeap).(int)
		b = heap.Pop(bHeap).(int)

		if a > b {
			sum += a - b
		} else {
			sum += b - a
		}
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
