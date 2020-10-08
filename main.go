package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	args := os.Args[1:]
	if len(args) == 0 {
		log.Fatal("Not enough arguments")
	}
	counter, delay := parseArgs(args)
	fmt.Println(counter, delay)
}

func parseArgs(args []string) (int, int) {
	length := len(args)
	counter := 0
	delay := 0

	for key, arg := range args {
		switch arg {
		case "-d":
			if key+1 < length {
				delay, _ = strconv.Atoi(args[key+1])
			}
			break

		case "-c":
			if key+1 < length {
				counter, _ = strconv.Atoi(args[key+1])
			}
			break
		}
	}
	return counter, delay
}
