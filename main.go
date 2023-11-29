package main

import (
	"fmt"
	"os"
)

type Friend struct {
	id         uint
	name       string
	address    string
	occupation string
	reason     string
}

var friends []Friend = []Friend{
	{
		id:         1,
		name:       "Anto",
		address:    "Kuningan",
		occupation: "Software Engineer",
		reason:     "Karena golang laku dipasaran",
	},
	{
		id:         2,
		name:       "Yuni",
		address:    "Cipete",
		occupation: "QA",
		reason:     "Karena golang populer",
	},
	{
		id:         3,
		name:       "Bruno",
		address:    "Indrapasta",
		occupation: "Software Engineer",
		reason:     "Karena golang punya concurrency",
	},
}

func main() {
	var arg = os.Args

	if len(arg) < 2 {
		fmt.Println("Masukkan Nama ")

	} else {
		findMyFriend(arg[1])
	}
}

func findMyFriend(name string) {
	found := false
	for _, eachFriend := range friends {
		if name == eachFriend.name {
			fmt.Println("ID:", eachFriend.id)
			fmt.Println("Nama:", eachFriend.name)
			fmt.Println("Alamat:", eachFriend.address)
			fmt.Println("Pekerjaan:", eachFriend.occupation)
			fmt.Println("Alasan:", eachFriend.reason)
			found = true
			break
		}
	}

	if !found {
		fmt.Println("Nama tidak ditemukan")
	}
}
