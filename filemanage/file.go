package filemanage

import "os"

func Readfile() string {
	dat, _ := os.ReadFile("./write")
	return string(dat)
}

func Writefile(a string) {
	data := []byte(a)
	os.WriteFile("./write", data, 0644)
}
