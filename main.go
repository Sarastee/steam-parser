package main

import (
	"bufio"
	"fmt"
	"golang.design/x/clipboard"
	"log"
	"regexp"
	"strings"
)

type AccountNodes struct {
	AccountNodes []Account `json:"accountNodes"`
}

type Account struct {
	Login          string `json:"login"`
	Password       string `json:"password"`
	SharedSecret   string `json:"shared_secret"`
	IdentitySecret string `json:"identity_secret"`
	SteamID        string `json:"steamID"`
	Offer          string `json:"offer"`
}

func main() {
	input := readFromBuffer()

	lines, err := StringToLines(input)
	if err != nil {
		log.Fatal(err)
	}

	sliceSize := len(lines) / 7

	data := AccountNodes{}
	counter := 0
	i := 0
	data.AccountNodes = make([]Account, sliceSize)

	var buf []byte
	for _, line := range lines {

		line = Parse(line, `(?m)[\w]+[:][\s]+`, "")
		line = Parse(line, `[\s]`, "")

		switch counter {
		case 0:
			data.AccountNodes[i].Login = line
			counter++
		case 1:
			data.AccountNodes[i].Password = line
			counter++
		case 2:
			data.AccountNodes[i].SharedSecret = line
			counter++
		case 3:
			data.AccountNodes[i].IdentitySecret = line
			counter++
		case 4:
			data.AccountNodes[i].SteamID = line
			counter++
		case 5:
			data.AccountNodes[i].Offer = line
			counter++
		default:
			counter = 0
			temp := fmt.Sprintf("%s:%s:::%s:::%s:::::%s:::::%s\n",
				data.AccountNodes[i].Login,
				data.AccountNodes[i].Password,
				data.AccountNodes[i].SharedSecret,
				data.AccountNodes[i].IdentitySecret,
				data.AccountNodes[i].SteamID,
				data.AccountNodes[i].Offer)

			buf = append(buf, temp...)
			writeInBuffer(buf)
			i++
		}
	}
}

func readFromBuffer() string {
	err := clipboard.Init()
	if err != nil {
		panic(err)
	}
	data := clipboard.Read(clipboard.FmtText)
	return string(data)
}

func writeInBuffer(data []byte) {
	clipboard.Write(clipboard.FmtText, data)
}
func StringToLines(s string) (lines []string, err error) {
	scanner := bufio.NewScanner(strings.NewReader(s))
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	err = scanner.Err()
	return
}

func Parse(data string, pattern string, replacement string) string {
	deleteInfo := regexp.MustCompile(pattern)
	exitData := deleteInfo.ReplaceAllString(data, replacement)
	return exitData
}
