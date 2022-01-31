package main

import (
	"io"
	"os"
	"testing"
)

func BenchmarkRegex(b *testing.B) {
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

	for i := 0; i < b.N; i++ {
		findLetterRegex(body)
	}
}

func BenchmarkFind(b *testing.B) {
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

	for i := 0; i < b.N; i++ {
		findLetter(body)
	}
}
