package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"strings"
)

var (
	start   = []string{"disini", "disana", "disitu", "didalam", "diatas", "dibawah"}
	replace = []string{"di sini", "di sana", "di situ", "di dalam", "di atas", "di bawah"}
)

func main() {
	out, err := exec.Command("git", "diff", "--name-only", "HEAD^").Output()
	if err != nil {
		panic(err)
	}

	data := string(out)

	for _, line := range strings.Split(data, "\n") {
		if line != "" && line[len(line)-3:] == ".md" {
			// res := OpenFile(line)
			// newData := ReplaceToEYD(res)
			// WriteNewFile(newData, line)

			fmt.Println(line)
		}
	}
}

func ReplaceToEYD(data string) string {
	for i := 0; i < len(start); i++ {
		data = strings.Replace(data, start[i], replace[i], -1)
	}

	return data
}

func OpenFile(path string) string {
	f, err := os.Open(path)
	if err != nil {
		panic(err)
	}

	defer f.Close()

	b, err := ioutil.ReadAll(f)
	if err != nil {
		panic(err)
	}

	data := string(b)

	return data
}

func WriteNewFile(data string, path string) {
	f, err := os.Create(path)
	if err != nil {
		panic(err)
	}

	defer f.Close()

	f.WriteString(data)
}
