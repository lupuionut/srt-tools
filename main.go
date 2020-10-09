package main

import (
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
	defer file.Close()
	if err != nil {
		log.Fatal(err)
	}
	shiftTime(file, delay, counter)
}

func parseArgs(args []string) (int, float64, string) {
	length := len(args)
	counter := 0
    delay := float64(0)
	in := ""

	for key, arg := range args {
		switch arg {
		case "delay":
			if key+1 < length {
                delay, _ = strconv.ParseFloat(args[key+1], 32)
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
