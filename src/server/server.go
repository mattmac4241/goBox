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

func evalInput(message helper.Message, c net.Conn) string {
	if message.Command == "List" {
		files := getFiles()
		mess := helper.CreateMessage([]byte(files), "List", "")
		fmt.Println(mess)
		helper.EncodeMessage(*mess, c)
	} else if message.Command == "Upload" {
		name := fmt.Sprintf("./%s/%s", "serverFiles", message.FileName)
		helper.WriteFile(name, message.Content)
	} else if message.Command == "Download" {
		file, err := helper.GetFile(message.FileName)
		if err != nil {
			mess := helper.CreateMessage(nil, "Download", "")
			helper.EncodeMessage(*mess, c)
		} else {
			mess := helper.CreateMessage(file, "Download", message.FileName)
			helper.EncodeMessage(*mess, c)
		}

	}
	return "command not found"
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
