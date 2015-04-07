package main

import "os"
import "fmt"
import "golang.org/x/net/websocket"
import "log"

func ws_in(ws *websocket.Conn) {
	var msg = make([]byte, 65536)
	for true {
		n, err := ws.Read(msg)
		if err != nil {
			log.Fatal(err)
		}
		s := string(msg)
		if n != 0 {
			fmt.Println(s)
		} else {
			fmt.Println("(Empty)")
		}
	}

}

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

	ws, err := websocket.Dial(url, "", origin)

	if err != nil {
		log.Fatal(err)
	}

	go ws_in(ws)

	for true {
		var so string
		_, err = fmt.Sscanln(so)

		if err != nil {
			log.Fatal(err)
		}

		ws.Write([]byte(so))

		if err != nil {
			log.Fatal(err)
		}
	}

}
