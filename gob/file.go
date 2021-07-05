package gob

import (
	gb "encoding/gob"
	"os"
	"strings"
)

// Dumps -
func Dumps(fileSuffix string, e interface{}) error {
	f, _ := os.Executable()
	fl, err := os.OpenFile(strings.Join([]string{f, fileSuffix, "gob"}, "."), os.O_CREATE|os.O_TRUNC|os.O_WRONLY, 0644)
	if err != nil {
		return err
	}
	defer fl.Close()
	return gb.NewEncoder(fl).Encode(e)
}

// Loads -
func Loads(fileSuffix string, e interface{}) error {
	f, _ := os.Executable()
	fl, err := os.OpenFile(strings.Join([]string{f, fileSuffix, "gob"}, "."), os.O_RDONLY, 0644)
	if err != nil {
		return err
	}
	defer fl.Close()
	return gb.NewDecoder(fl).Decode(e)
}
