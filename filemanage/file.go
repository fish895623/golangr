package filemanage

import "os"

func Readfile() string {
	dat, _ := os.ReadFile("../data/write")
	return string(dat)
}

func Writefile(a string) {
	data := []byte(a)
	os.WriteFile("../data/write", data, 0644)
}
