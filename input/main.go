package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

func Parse(buf []int) string {
	return ""
}

func main() {
	inputReader := bufio.NewReader(os.Stdin)
	buf, err := readIntSlice(inputReader)
	if err != nil {
		return
	}

	fmt.Println(Parse(buf))
}

func readIntSlice(reader *bufio.Reader) ([]int, error) {
	lineBuf, err := reader.ReadString('\n')
	if err != nil && err != io.EOF {
		return nil, fmt.Errorf(err.Error())
	}

	lineBuf = strings.TrimRight(lineBuf, "\n")
	line := strings.Split(lineBuf, " ")
	var result []int
	for _, v := range line {
		i, err := strconv.ParseInt(v, 16, 32)
		if err != nil {
			return nil, fmt.Errorf(err.Error())
		}
		result = append(result, int(i))
	}
	return result, nil
}
