package main

import (
	"encoding/gob"
	"flag"
	"fmt"
	"helper"
	"log"
	"net"
)

func main() {
	uploadFile := flag.String("u", "", "Used to upload file")
	downloadFile := flag.String("d", "", "Name of file to download")
	listFlag := flag.Bool("l", false, "List files server contains")
	flag.Parse()
	upload(*uploadFile)
	list(*listFlag)
	download(*downloadFile)

}

func upload(fileName string) {
	if containsFile(fileName) {
		l, err := helper.GetFile(fileName)
		if err != nil {
			fmt.Println(err)
		} else {
			fmt.Println(len(l))
			message := helper.CreateMessage(l, "Upload", fileName)
			connection := getConnection("localhost", "8000")
			encoder := gob.NewEncoder(connection)
			encoder.Encode(message)
			connection.Close()
		}
	}
}

func download(file string) {
	if containsFile(file) {
		message := helper.CreateMessage(nil, "Download", file)
		connection := getConnection("localhost", "8000")
		helper.EncodeMessage(*message, connection)
		mess := helper.DecodeMessage(connection)
		if mess.FileName == "" {
			fmt.Println("NO file found")
		} else {
			file := fmt.Sprintf("./%s/%s", "clientFiles", mess.FileName)
			helper.WriteFile(file, mess.Content)
		}

		connection.Close()
	}
}

//Return a list of all files contained in the server
func list(listFlag bool) {
	if listFlag == true {
		message := helper.CreateMessage(nil, "List", "")
		connection := getConnection("localhost", "8000")
		helper.EncodeMessage(*message, connection)
		mess := helper.DecodeMessage(connection)
		fmt.Println(string(mess.Content))
		connection.Close()
	}
}

func containsFile(file string) bool {
	if file == "" {
		return false
	}
	return true
}

func getConnection(hostname, port string) net.Conn {
	connection := fmt.Sprintf("%s:%s", hostname, port)
	listener, err := net.Dial("tcp", connection)
	if err != nil {
		log.Fatal(err)
	}
	return listener
}
