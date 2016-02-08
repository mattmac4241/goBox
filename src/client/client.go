package main

import (
	"flag"
	"fmt"
	"helper"
	"log"
	"net"
)

func main() {
	uploadFile := flag.String("u", "", "Used to upload file")
	downloadFile := flag.String("d", "", "Name of file to download")
	listFlag := flag.Bool("l", true, "List files server contains")
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
			connection := getConnection("localhost", "8000")
			connection.Write([]byte("test"))
			connection.Close()
			fmt.Println(string(l))
		}
	}
}

func download(file string) {
	if containsFile(file) {
		fmt.Println(file)

	}
}

func list(listFlag bool) {
	if listFlag == true {
		fmt.Println("List")
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
