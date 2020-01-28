package main

import (
	"fmt"
	"os"

	"github.com/gaborszakacs/who-is-on-call/announcement"
)

func main() {
	announcer := announcement.Announcer{}
	err := announcer.DoItWithDrama()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
