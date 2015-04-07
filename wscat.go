package main

import "os"
import "fmt"
import "golang.org/x/net/websocket"
import "log"

func main() {
	if os.Args[0] == "" || len(os.Args) < 2 || len(os.Args) >= 3 {
		log.Fatal("invalid argument")
		os.Exit(-1)
	}

	origin := "https://kkdev.org/?wscatinit"
	url := os.Args[1]

	if len(os.Args) == 2 {
		origin = ""
	}

}
