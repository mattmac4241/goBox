package helper

import (
	"errors"
	"io/ioutil"
)

//retrieve a file, check if it is there and then get body of file
func GetFile(file string) ([]byte, error) {
	body, err := ioutil.ReadFile(file)
	if err != nil {
		return nil, errors.New("Failed to open file")
	}
	return body, nil
}
