package utils

import (
	"os"
	"path/filepath"
	"strings"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func Read(path string) string {
	absPath, _ := filepath.Abs("../advent-of-code-2023" + path)
	dat, err := os.ReadFile(absPath)
	check(err)
	return string(dat)
}

func ReadLines(path string) []string {
	return strings.Split(Read(path), "\n")
}
