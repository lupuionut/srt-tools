package main

import (
	"fmt"
	"os"
	"regexp"
)

type Block struct {
	Counter int
	Start   string
	End     string
	Text    string
}

// read a single block from file
// a block is delimited by a new line
// return the block and the value where new block starts
func readUntilNewLine(file *os.File, from int) ([]byte, int) {
	var last string
	var block_text []byte
	condition := true
	stat, _ := file.Stat()

	if from == int(stat.Size()) {
		return block_text, -1
	}

	for condition {
		text := make([]byte, 1)
		file.ReadAt(text, int64(from))
		c := text[0]
		block_text = append(block_text, c)
		if string(c) == "\n" {
			if last == "\n" {
				condition = false
			} else {
				last = "\n"
			}
		} else {
			last = ""
		}
		from++
	}

	return block_text, from
}

func formatBlock(text string) (Block, error) {
	regex, err :=
		regexp.Compile(`(\d\n)(\d{2}:\d{2}:\d{2},\d{1,3}.*\d{2}:\d{2},\d{1,3}\n)((.|\n)*)`)

	if err != nil {
		return Block{}, err
	}

	matches := regex.FindAllStringSubmatch(text, -1)
	fmt.Println(len(matches[0]))
	return Block{}, nil
}

func mutateBlock(block Block) (Block, error) {

	return Block{}, nil
}

func blockToString(block Block) (string, error) {

	return "", nil
}

func appendToFile(out *os.File, text string) {

}

func shiftTime(file *os.File, delay int, counter int) {
	position := 0
	var text []byte
	for position != -1 {
		text, position = readUntilNewLine(file, position)
		formatBlock(string(text))
	}
}
