package reader

import (
	"os"
	"path/filepath"
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
