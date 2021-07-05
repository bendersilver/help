package gob

import (
	gb "encoding/gob"
	"os"
)

// Dumps -
func Dumps(fileName string, e interface{}) error {
	fl, err := os.OpenFile(fileName, os.O_CREATE|os.O_TRUNC|os.O_WRONLY, 0644)
	if err != nil {
		return err
	}
	defer fl.Close()
	return gb.NewEncoder(fl).Encode(e)
}

// Loads -
func Loads(fileName string, e interface{}) error {
	fl, err := os.OpenFile(fileName, os.O_RDONLY, 0644)
	if err != nil {
		return err
	}
	defer fl.Close()
	return gb.NewDecoder(fl).Decode(e)
}
