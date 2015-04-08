package main

import "os"
import "fmt"
import "golang.org/x/net/websocket"
import "log"
import "net/url"
import "strings"

func ws_in(ws *websocket.Conn) {

	var msg string

	for true {

		err := websocket.Message.Receive(ws,&msg)

		if err != nil {
			log.Fatal(err)
		}

		//s := msg

		//log.Println("Recv:" + s)

		if len(msg) != 0 {

			fmt.Println(msg)

		} else {

			fmt.Println("(Empty)")

		}
	}

}

func ws_out(ws *websocket.Conn) {
	for true {
		var so string
		_, err := fmt.Scanf("%s\n", &so)
		if err != nil {
			log.Fatal(err)
		}
		//log.Println("Sending:" + so)
		websocket.Message.Send(ws,so)

		if err != nil {
			log.Fatal(err)
		}

		//log.Println("Sent:" + so)
	}
}

func ws_connd(ws *websocket.Conn) {
	go ws_in(ws)

	log.Println("Ready")

	ws_out(ws)
}

func main() {

	if os.Args[0] == "" || len(os.Args) < 2 || len(os.Args) > 3 {
		log.Fatal("invalid argument")
		os.Exit(-1)
	}

	origin := ""
	urli := os.Args[1]

	if len(os.Args) == 2 {
		//origin = ""
	} else {
		origin = os.Args[2]
	}

	log.Println("Connecting to " + urli)

	if origin == "" {
		oriu,_:=url.Parse(urli)
		 switch strings.ToUpper(oriu.Scheme){
			case "WSS":
			oriu.Scheme="https"
			case "WS":
			oriu.Scheme="http"
		}
		origin=oriu.String()
		ws, err := websocket.Dial(urli, "", origin)
		if err != nil {
			log.Fatal(err)
		}
		ws_connd(ws)

	} else {
		ws, err := websocket.Dial(urli, "", origin)
		if err != nil {
			log.Fatal(err)
		}
		ws_connd(ws)

	}

}
