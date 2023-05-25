package main

import (
	"fmt"
	"golang.design/x/clipboard"
	"regexp"
)

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

	blank := regexp.MustCompile("Login:\\s*(\\w+)\\s*\\nPassword:\\s*(\\w+)\\s*shared_secret:\\s*(.+=)\\s*identity_secret:\\s*(.+=)\\s*SteamID:\\s*(\\w+)\\s*\\nOffer:\\s*(.+\\S)")
	res := blank.FindAllStringSubmatch(input, -1)

	var buf []byte
	accounts := make([]Account, 0, 1)

	for i := 0; i < len(res); i++ {
		accounts = append(accounts, Account{
			Login:          res[i][1],
			Password:       res[i][2],
			SharedSecret:   res[i][3],
			IdentitySecret: res[i][4],
			SteamID:        res[i][5],
			Offer:          res[i][6],
		})

		temp := fmt.Sprintf("%s:%s:::%s:::%s:::::%s:::::%s\n",
			accounts[i].Login,
			accounts[i].Password,
			accounts[i].SharedSecret,
			accounts[i].IdentitySecret,
			accounts[i].SteamID,
			accounts[i].Offer)
		buf = append(buf, temp...)

		writeInBuffer(buf)
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
