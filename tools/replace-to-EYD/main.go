package main

import (
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

type ReplacerWord struct {
	Start    string
	Replacer string
}

var listReplacerWord = []ReplacerWord{{"diantara", "di antara"}, {"diakhir", "di akhir"}, {"diatas", "di atas"}, {"diawal", "di awal"}, {"dibawah", "di bawah"}, {"dibelakang", "di belakang"}, {"didalam", "di dalam"}, {"didekat", "di dekat"}, {"didepan", "di depan"}, {"dikanan", "di kanan"}, {"dikiri", "di kiri"}, {"diluar", "di luar"}, {"dimana", "di mana"}, {"disamping", "di samping"}, {"disaat", "di saat"}, {"disana", "di sana"}, {"disebelah", "di sebelah"}, {"disini", "di sini"}, {"disisi", "di sisi"}, {"disitu", "di situ"}, {"ditengah", "di tengah"}, {"ditengah-tengah", "di tengah-tengah"}, {"ditiap", "di tiap"}, {"ditiap-tiap", "di tiap-tiap"},
	{"Diantara", "Di antara"}, {"Diakhir", "Di akhir"}, {"Diatas", "Di atas"}, {"Diawal", "Di awal"}, {"Dibawah", "Di bawah"}, {"Dibelakang", "Di belakang"}, {"Didalam", "Di dalam"}, {"Didekat", "Di dekat"}, {"Didepan", "Di depan"}, {"Dikanan", "Di kanan"}, {"Dikiri", "Di kiri"}, {"Diluar", "Di luar"}, {"Dimana", "Di mana"}, {"Disamping", "Di samping"}, {"Disaat", "Di saat"}, {"Disana", "Di sana"}, {"Disebelah", "Di sebelah"}, {"Disini", "Di sini"}, {"Disisi", "Di sisi"}, {"Disitu", "Di situ"}, {"Ditengah", "Di tengah"}, {"Ditengah-tengah", "Di tengah-tengah"}, {"Ditiap", "Di tiap"}, {"Ditiap-tiap", "Di tiap-tiap"},
}

func main() {
	out, err := exec.Command("git", "diff", "--name-only", "HEAD^").Output()
	if err != nil {
		log.Println("Error exec command", err.Error())
		panic(err)
	}

	listFileName := string(out)

	for _, fileName := range strings.Split(listFileName, "\n") {
		if fileName != "" && filepath.Ext(fileName) == ".md" {
			MDText, err := OpenFile(fileName)
			if err != nil {
				log.Println("Error open file", err.Error())
				panic(err)
			}
			newMDText := ReplaceToEYD(MDText)
			err = OverwriteText(newMDText, fileName)
			if err != nil {
				log.Println("Error Overwrite text file", err.Error())
				panic(err)
			}
		}
	}
}

func ReplaceToEYD(textFile string) string {
	for i := 0; i < len(listReplacerWord); i++ {
		textFile = strings.Replace(textFile, listReplacerWord[i].Start, listReplacerWord[i].Replacer, -1)
	}

	return textFile
}

func OpenFile(path string) (string, error) {
	f, err := os.Open(path)
	if err != nil {
		return "", err
	}

	defer f.Close()

	b, err := ioutil.ReadAll(f)
	if err != nil {
		return "", err
	}

	textFile := string(b)

	return textFile, nil
}

func OverwriteText(textFile string, path string) error {
	f, err := os.OpenFile(path, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0755)
	if err != nil {
		return err
	}

	defer f.Close()

	_, err = f.WriteString(textFile)
	if err != nil {
		return err
	}

	return nil
}
