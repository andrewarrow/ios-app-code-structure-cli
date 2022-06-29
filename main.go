package main

import (
	"iacs/ios"
	"math/rand"
	"os"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())

	if len(os.Args) == 1 {
		PrintHelp()
		return
	}
	command := os.Args[1]

	if command == "scan" {
		dir := os.Args[2]
		ios.Scan(dir)
	} else if command == "ls" {
	}
}
