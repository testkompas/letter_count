package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"regexp"
)

var (
	letters     = []byte("ABCDEFGHIJKLMNOPQRSTUVWXYZ")
	letterCount = make(map[byte]int)
)

func main() {
	fi, err := os.Open("./article.txt")
	if err != nil {
		panic(err)
	}

	defer func() {
		if err := fi.Close(); err != nil {
			panic(err)
		}
	}()

	body, err := io.ReadAll(fi)
	if err != nil {
		panic(err)
	}

	fo, err := os.Create("./output.txt")
	if err != nil {
		panic(err)
	}
	defer func() {
		if err := fo.Close(); err != nil {
			panic(err)
		}
	}()
	writer := bufio.NewWriter(fo)

	findLetterRegex(body)

	for _, v := range letters {
		data := fmt.Sprintf("%s = %d\n", string(v), letterCount[v])

		writer.WriteString(data)
		if e := writer.Flush(); e != nil {
			panic(e)
		}
	}
}

func findLetterRegex(body []byte) {
	for _, v := range letters {
		re := regexp.MustCompile(`[` + string(v) + `]`)
		find := re.FindAll(body, -1)
		letterCount[v] = len(find)
	}
}

func findLetter(body []byte) {
	for _, v := range letters {
		if _, ok := letterCount[v]; !ok {
			letterCount[v] = 0
		}
	}

	for _, v := range body {
		if _, ok := letterCount[v]; ok {
			letterCount[v]++
		}
	}
}
