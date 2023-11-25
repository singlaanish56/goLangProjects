package main

import (
	"bufio"
	"crypto/sha1"
	"encoding/base64"
	"errors"
	"fmt"
	"log"
	"net"
	"net/http"
	"strings"
)

type Ws struct{
	conn net.Conn
	bufrw *bufio.ReadWriter
	header http.Header
	status int
}

type Frame struct{
	IsFragment bool
	Opcode byte
	Reserved byte
	IsMasked bool
	Length int
	Payload []byte
}

func main() {
	fmt.Println("Hello new Server")
	setup()
	http.ListenAndServe(":8000",nil)
}

//route setup
func setup(){
	http.HandleFunc("/" , homepage)
	http.HandleFunc("/ws", websocketServer)
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

//server-side handshake
func NewWSHandler(w http.ResponseWriter, r * http.Request) (*Ws, error){
	hj, ok := w.(http.Hijacker)
	if !ok{
		return nil,errors.New("unable to establish connection")
	}
	conn, bufrw, err := hj.Hijack()
	if err!=nil{
		return nil, err
	}

	return &Ws{conn, bufrw, r.Header, 200}, nil
}

func (ws *Ws) Handshake() error{
	hashWebSocketKey :=  getHash(ws.header.Get("Sec-Websocket-Key"))

	handhsakeheader := []string{
		"HTTP/1.1 101 Switching Protocols",
		"Upgarde: websocket",
		"Connection: Upgrade",
		"Sec-WebSocket-Accept: " + hashWebSocketKey,
		"",
		"",
	}

	_, err := ws.bufrw.Write([]byte(strings.Join(handhsakeheader, "\r\n")))

	return err
}

func getHash(key string) string{
	h := sha1.New()
	h.Write([]byte(key))
	h.Write([]byte("258EAFA5-E914-47DA-95CA-C5AB0DC85B11"))
	return base64.StdEncoding.EncodeToString(h.Sum(nil))
}

//recieve the data
func (ws *Ws) Recv() (Frame, error){
	frame := Frame{}

	hear, err := ws.bufrw.Read()
}