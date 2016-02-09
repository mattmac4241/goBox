package helper

import (
	"encoding/gob"
	"errors"
	"io/ioutil"
	"log"
	"net"
	"os"
)

type Message struct {
	Content  []byte
	Command  string
	FileName string
}

//retrieve a file, check if it is there and then get body of file
func GetFile(file string) ([]byte, error) {
	body, err := ioutil.ReadFile(file)
	if err != nil {
		return nil, errors.New("Failed to open file")
	}
	return body, nil
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
