package main

import (
	"day02/internal/rotate"
	"flag"
	"log"
	"sync"
)

func main() {

	var archiveFolderFlag string
	flag.StringVar(&archiveFolderFlag, "a", "", "usage: ./myRotate [[-a] folder/path] file.log file2.log...")
	flag.Parse()

	files := flag.Args()

	wg := sync.WaitGroup{}
	for _, currentFile := range files {
		wg.Add(1)
		go func(filename string) {
			defer wg.Done()
			err := rotate.PackFile(currentFile, archiveFolderFlag)
			if err != nil {
				log.Fatal(err)
			}
		}(currentFile)
	}
	wg.Wait()
}
