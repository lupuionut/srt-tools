package main

import (
	"errors"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
	"time"
)

type Block struct {
	Counter int
	Start   time.Time
	End     time.Time
	Text    string
}

// read a single block from file
// a block is delimited by a new line
// return the block and the value where new block starts
func readUntilNewLine(file *os.File, from int) ([]byte, int) {
	delimiter := []string{" ", " ", " "}
	var block_text []byte
	condition := true
	stat, _ := file.Stat()

	if from == int(stat.Size()) {
		return block_text, -1
	}

	for condition && from < int(stat.Size()) {
		text := make([]byte, 1)
		file.ReadAt(text, int64(from))
		c := text[0]

		block_text = append(block_text, c)
		delimiter = pushToDelimiter(delimiter, string(c))
		if strings.Join(delimiter, "") == "\n\r\n" ||
			strings.Join(delimiter[0:2], "") == "\n\n" {
			condition = false
		}
		from++
	}
	return block_text, from
}

func pushToDelimiter(delimiter []string, s string) []string {
	delimiter[0] = delimiter[1]
	delimiter[1] = delimiter[2]
	delimiter[2] = s
	return delimiter
}

func formatBlock(text string) (Block, error) {
	regex, err :=
		regexp.Compile(`(\d{1,}[\r|\n]*)(\d{2}:\d{2}:\d{2},\d{1,3}.*\d{2}:\d{2},\d{1,3}[\n|\r]*)((.|\n|\r)*)`)

	if err != nil {
		return Block{}, err
	}

	matches := regex.FindAllStringSubmatch(text, -1)
	if len(matches) == 0 {
		return Block{}, errors.New("invalid block")
	} else {
		counter, err := strconv.Atoi(strings.TrimSpace(matches[0][1]))
		start, end, err := extractTimes(matches[0][2])
		if err != nil {
			return Block{}, err
		}
		return Block{
			Counter: counter,
			Start:   start,
			End:     end,
			Text:    matches[0][3],
		}, nil
	}
}

func extractTimes(text string) (time.Time, time.Time, error) {
	regex, err :=
		regexp.Compile(`(\d{2,}:\d{2,}:\d{2,},\d*) --> (\d{2,}:\d{2,}:\d{2,},\d*)`)

	if err != nil {
		return time.Now(), time.Now(), errors.New("invalid time format")
	}

	matches := regex.FindAllStringSubmatch(text, -1)
	if len(matches) == 0 {
		return time.Now(), time.Now(), errors.New("no time strings found")
	} else {
		startText := strings.Replace(matches[0][1], ",", ".", 1)
		endText := strings.Replace(matches[0][2], ",", ".", 1)
		start, _ := time.Parse(time.StampMilli, "Jan 1 "+startText)
		end, _ := time.Parse(time.StampMilli, "Jan 1 "+endText)
		return start, end, nil
	}
}

func mutateBlock(block Block, delay float64, counter int) (Block, error) {
	delayDuration := time.Duration(delay * 1000)
	currentMilliseconds :=
		block.Start.Nanosecond()/1000000 +
			block.Start.Second()*1000 +
			block.Start.Minute()*60*1000 +
			block.Start.Hour()*3600*1000

	newstart := block.Start
	newend := block.End
	if float64(currentMilliseconds)+(delay*1000) > 0 {
		newstart = block.Start.Add(time.Millisecond * delayDuration)
		newend = block.End.Add(time.Millisecond * delayDuration)
	}

	newcounter := block.Counter + counter
	if newcounter < 0 {
		newcounter = 0
	}
	newblock := Block{
		Counter: newcounter,
		Text:    block.Text,
		Start:   newstart,
		End:     newend,
	}
	return newblock, nil
}

func blockToString(block Block) (string, error) {
	var text string
	text += strconv.Itoa(block.Counter) + "\n"
	text += timeToString(block.Start) + " --> " + timeToString(block.End) + "\n"
	text += block.Text
	return text, nil
}

func timeToString(t time.Time) string {
	hour := strconv.Itoa(t.Hour())
	if len(hour) < 2 {
		hour = "0" + hour
	}
	minutes := strconv.Itoa(t.Minute())
	if len(minutes) < 2 {
		minutes = "0" + minutes
	}
	seconds := strconv.Itoa(t.Second())
	if len(seconds) < 2 {
		seconds = "0" + seconds
	}
	milliseconds := strconv.Itoa(t.Nanosecond() / 1000000)
	if len(milliseconds) < 3 {
		milliseconds = strings.Repeat("0", 3-len(milliseconds)) + milliseconds
	}
	return hour + ":" + minutes + ":" + seconds + "," + milliseconds
}

func ShiftTime(file *os.File, out *os.File, delay float64, counter int) {
	position := 0
	var text []byte
	for position != -1 {
		text, position = readUntilNewLine(file, position)
		block, err := formatBlock(string(text))
		if err == nil {
			block, err = mutateBlock(block, delay, counter)
			newline, _ := blockToString(block)
			out.Write([]byte(newline))
		}
	}
	fmt.Println("resync complete")
}
