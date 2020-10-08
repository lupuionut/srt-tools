package main

import (
	"os"
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
		return block_text, 0
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
