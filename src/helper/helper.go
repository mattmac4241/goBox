package helper

import (
	"bufio"
	"encoding/gob"
	"errors"
	"io"
	"log"
	"net"
	"os"
)

type Message struct {
	Content  []byte
	Command  string
	FileName string
}

func GetFile(file string) ([]byte, error) {
	var content []byte
	f, err := os.Open(file)
	if err != nil {
		return nil, errors.New("Failed to open file")
	}
	r := bufio.NewReader(f)
	buf := make([]byte, 2048)
	for {
		// read a chunk
		n, err := r.Read(buf)
		if err != nil && err != io.EOF {
			panic(err)
		}
		content = append(content, buf[:n]...)
		if n == 0 {
			break
		}
	}
	return content, nil

}

func WriteFile(fileName string, content []byte) {
	file, err := os.OpenFile(
		fileName,
		os.O_WRONLY|os.O_TRUNC|os.O_CREATE,
		0666,
	)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	// Write bytes to file
	bytesWritten, err := file.Write(content)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Wrote %d bytes.\n", bytesWritten)

}

func EncodeMessage(message Message, c net.Conn) {
	encoder := gob.NewEncoder(c)
	encoder.Encode(message)
}

func DecodeMessage(c net.Conn) Message {
	dec := gob.NewDecoder(c)
	message := &Message{}
	dec.Decode(message)
	return *message
}

func CreateMessage(content []byte, command, fileName string) *Message {
	message := new(Message)
	message.Command = command
	message.Content = content
	message.FileName = fileName
	return message
}
