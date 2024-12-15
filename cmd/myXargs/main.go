package main

import (
	"day02/internal/xargs"
	"log"
	"os"
)

func main() {
	Xargs, err := xargs.Init()
	if err != nil {
		log.Fatal(err)
	}
	Xargs.Exec(os.Stdout)
}
