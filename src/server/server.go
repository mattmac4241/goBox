package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net"
	"strings"
)

type Message struct {
	Content, Command string
}

func main() {
	listener, err := net.Listen("tcp", "localhost:8000")
	if err != nil {
		log.Fatal(err)
	}

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Print(err)
			continue
		}
		go handleConn(conn)
	}
}

func handleConn(c net.Conn) {
	defer c.Close()
	input := bufio.NewScanner(c)
	for input.Scan() {
		in := input.Bytes()
		fmt.Fprintln(c, evalInput(in))
	}
}

func evalInput(input []byte) string {
	m := decodeInput(input)
	if m.Command == "List" {
		return getFiles()
	}
	return "command not found"
}

func decodeInput(input []byte) Message {
	jsonStream := string(input)
	dec := json.NewDecoder(strings.NewReader(jsonStream))
	var m Message
	if err := dec.Decode(&m); err == io.EOF {
		fmt.Print("")
	} else if err != nil {
		log.Fatal(err)
	}
	return m
}

func getFiles() string {
	var fileList []string
	files, _ := ioutil.ReadDir("./serverFiles")
	for _, f := range files {
		fil := fmt.Sprintf("%s\n", f.Name())
		fileList = append(fileList, fil)
	}

	return strings.Join(fileList, " ")
}
