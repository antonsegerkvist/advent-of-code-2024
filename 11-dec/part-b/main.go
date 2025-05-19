package main

import (
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
	"strings"
	"sync"
	"sync/atomic"
	"unicode/utf8"
)

var cache = map[string]int{}
var rwlock = sync.RWMutex{}

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

	count := atomic.Int64{}
	wg := sync.WaitGroup{}
	for _, s := range curr {
		wg.Add(1)
		go func(s string) {
			defer wg.Done()
			count.Add(int64(dac(s, 0, steps)))
		}(s)
	}

	wg.Wait()
	fmt.Println("Count:", count.Load())

}

func calc(curr string) []string {
	next := []string{}
	len := utf8.RuneCountInString(curr)

	if curr == "0" {
		next = append(next, "1")
	} else if len%2 == 0 {
		nl := curr[:len>>1]
		nr := curr[len>>1:]

		intNL, err := strconv.ParseUint(nl, 10, 64)
		if err != nil {
			log.Fatal(err.Error())
		}
		intNR, err := strconv.ParseUint(nr, 10, 64)
		if err != nil {
			log.Fatal(err.Error())
		}

		nl = strconv.FormatUint(intNL, 10)
		nr = strconv.FormatUint(intNR, 10)

		next = append(next, nl, nr)
	} else {
		n, err := strconv.ParseInt(curr, 10, 64)
		if err != nil {
			log.Fatal(err.Error())
		}
		ns := strconv.FormatInt(n*2024, 10)
		next = append(next, ns)
	}

	return next
}

func dac(s string, level, steps int) int {
	if level == steps {
		return 1
	}

	rwlock.RLock()
	count, ok := cache[s+","+strconv.Itoa(level)]
	if ok {
		rwlock.RUnlock()
		return count
	}
	rwlock.RUnlock()

	next := calc(s)
	if len(next) == 2 {
		count += dac(next[0], level+1, steps)
		count += dac(next[1], level+1, steps)
	} else {
		count += dac(next[0], level+1, steps)
	}

	rwlock.Lock()
	cache[s+","+strconv.Itoa(level)] = count
	rwlock.Unlock()

	return count
}
