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
	counter, delay, path := parseArgs(args)
	file, err := os.OpenFile(path, os.O_RDONLY, 0755)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(delay, counter)
    block, end := readUntilNewLine(file, 46)
    fmt.Println(string(block), end)
}

func parseArgs(args []string) (int, int, string) {
	length := len(args)
	counter := 0
	delay := 0
	in := ""

	for key, arg := range args {
		switch arg {
		case "delay":
			if key+1 < length {
				delay, _ = strconv.Atoi(args[key+1])
			}
			break

		case "counter":
			if key+1 < length {
				counter, _ = strconv.Atoi(args[key+1])
			}
			break

		case "in":
			if key+1 < length {
				in = args[key+1]
			}
			break
		}

	}
	return counter, delay, in
}
