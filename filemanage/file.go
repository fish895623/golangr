package filemanage

import "os"

func Readfile() string {
	dat, _ := os.ReadFile("../data/write")
	return string(dat)
}

func CheckFileExists() bool {
	if _, err := os.Stat("../data/write"); err == nil {
		return true
	} else {
		return false
	}
}

func CreateFile() {
	os.Create("../data/write")
}

func Writefile(a string) {
	if !CheckFileExists() {
		CreateFile()
	}

	data := []byte(a)
	os.WriteFile("../data/write", data, 0644)
}
