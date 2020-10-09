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
    shift(counter, delay, path)
}

func shift(counter int, delay float64, path string) {
    outPath := path + "-resync"
	in, err := os.OpenFile(path, os.O_RDONLY, 0755)
	defer in.Close()
	if err != nil {
		log.Fatal(err)
	}
    out, err := os.Create(outPath)
    defer out.Close()
    if err != nil {
        log.Fatal(err)
    }
    ShiftTime(in, out, delay, counter)
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
