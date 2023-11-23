package main

import (
	"bufio"
	"errors"
	"fmt"
	"log"
	"net"
	"net/http"
)

type Ws struct{
	conn net.Conn
	bufrw *bufio.ReadWriter
	header http.Header
	status int
}

func main() {
	fmt.Println("Hello new Server")
	setup()
	http.ListenAndServe(":8000",nil)
}

func homepage(w http.ResponseWriter, r *http.Request){
	w.Write([]byte("Welcome to the homepage"))
}

func websocketServer(w http.ResponseWriter, r *http.Request){
	w.Write([]byte("Welcome to the Websocket page"))
	//tcp connection is established, hijack the connection
	ws, err :=NewWSHandler(w,r)
	if err != nil{
		log.Fatal(err)
		return
	}

	// get the handshake header ands end back the handshake header
	err= ws.Handshake()
	if err!=nil{
		log.Fatal(err)
		return 
	}

	defer ws.conn.Close()
	
}

func NewWSHandler(w http.ResponseWriter, r * http.Request) (*Ws, error){
	hj, ok := w.(http.Hijacker)
	if !ok{
		return nil,errors.New("Unable to Establish Connection")
	}
	conn, bufrw, err := hj.Hijack()
	if err!=nil{
		return nil, err
	}

	return &Ws{conn, bufrw, r.Header, 200}, nil
}

func (ws *Ws) Handshake() error{
	

	return nil
}

func setup(){
	http.HandleFunc("/" , homepage)
	http.HandleFunc("/ws", websocketServer)
}