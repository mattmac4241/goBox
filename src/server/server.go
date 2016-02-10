package main

import (
	"fmt"
	"helper"
	"io/ioutil"
	"log"
	"net"
	"strings"
)

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
	message := helper.DecodeMessage(c)
	evalInput(message, c)
}

func evalInput(message helper.Message, c net.Conn) {
	switch message.Command {
	case "List":
		list(c)
	case "Upload":
		upload(message, c)
	case "Download":
		download(message, c)
	default:
		fmt.Println("")
	}
}

func getFiles() string {
	var fileList []string
	files, _ := ioutil.ReadDir("./serverFiles")
	for _, f := range files {
		fil := fmt.Sprintf("%s\n", f.Name())
		fileList = append(fileList, fil)
	}

	return strings.Join(fileList, "")
}

func list(c net.Conn) {
	files := getFiles()
	mess := helper.CreateMessage([]byte(files), "List", "")
	helper.EncodeMessage(*mess, c)
}

func upload(message helper.Message, c net.Conn) {
	name := fmt.Sprintf("./%s/%s", "serverFiles", message.FileName)
	helper.WriteFile(name, message.Content)
}

func download(message helper.Message, c net.Conn) {
	file, err := helper.GetFile(message.FileName)
	if err != nil {
		mess := helper.CreateMessage(nil, "Download", "")
		helper.EncodeMessage(*mess, c)
	} else {
		mess := helper.CreateMessage(file, "Download", message.FileName)
		helper.EncodeMessage(*mess, c)
	}
}
