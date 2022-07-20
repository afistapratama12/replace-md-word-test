package main

import (
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"strings"
)

type ReplacerWord struct {
	Start    string
	Replacer string
}

var listReplacerWord = []ReplacerWord{{"disini", "di sini"}, {"disana", "di sana"}, {"disitu", "di situ"}, {"didalam", "di dalam"}, {"diatas", "di atas"}, {"dibawah", "di bawah"},
	{"Disini", "Di sini"}, {"Disana", "Di sana"}, {"Disitu", "Di situ"}, {"Didalam", "Di dalam"}, {"Diatas", "Di atas"}, {"Dibawah", "Di bawah"},
}

func main() {
	out, err := exec.Command("git", "diff", "--name-only", "HEAD^").Output()
	if err != nil {
		log.Println("Error exec command", err.Error())
		panic(err)
	}

	listFileName := string(out)

	for _, fileName := range strings.Split(listFileName, "\n") {
		if fileName != "" && fileName[len(fileName)-3:] == ".md" {
			MDText := OpenFile(fileName)
			newMDText := ReplaceToEYD(MDText)
			OverwriteText(newMDText, fileName)
		}
	}
}

func ReplaceToEYD(textFile string) string {
	for i := 0; i < len(listReplacerWord); i++ {
		textFile = strings.Replace(textFile, listReplacerWord[i].Start, listReplacerWord[i].Replacer, -1)
	}

	return textFile
}

func OpenFile(path string) string {
	f, err := os.Open(path)
	if err != nil {
		log.Println("Error open file", err.Error())
		panic(err)
	}

	defer f.Close()

	b, err := ioutil.ReadAll(f)
	if err != nil {
		log.Println("Error ioutil Readall", err.Error())
		panic(err)
	}

	textFile := string(b)

	return textFile
}

func OverwriteText(textFile string, path string) {
	f, err := os.OpenFile(path, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0755)
	if err != nil {
		log.Println("Error open file", err.Error())
		panic(err)
	}

	defer f.Close()

	_, err = f.WriteString(textFile)
	if err != nil {
		log.Println("Error write new text to file", err.Error())
		panic(err)
	}
	return
}
