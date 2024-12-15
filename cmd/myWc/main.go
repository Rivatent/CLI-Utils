package main

import (
	"day02/internal/count"
	"flag"
	"fmt"
	"log"
	"sync"
)

func main() {

	var linesFlag, wordsFlag, charsFlag bool

	flag.BoolVar(&linesFlag, "l", false, count.UsageMsg)
	flag.BoolVar(&wordsFlag, "w", false, count.UsageMsg)
	flag.BoolVar(&charsFlag, "m", false, count.UsageMsg)

	flag.Parse()

	countFunc, err := count.DefineFunc(linesFlag, wordsFlag, charsFlag)
	if err != nil {
		log.Fatal(err.Error(), "\n", count.UsageMsg)
	}
	result := map[string]int{}
	lock := sync.Mutex{}
	wg := sync.WaitGroup{}

	for _, file := range flag.Args() {
		wg.Add(1)
		go func(file string) {
			defer wg.Done()
			counterResult, err := countFunc(file)
			if err != nil {
				log.Printf("Error processing file %s: %v", file, err)
				return
			}
			lock.Lock()
			result[file] = counterResult
			lock.Unlock()
		}(file)
	}
	wg.Wait()

	for _, file := range flag.Args() {
		if _, ok := result[file]; ok {
			fmt.Printf("%d\t%s\n", result[file], file)
		}
	}
}
