package gob

import (
	gb "encoding/gob"
	"os"
	"path"
	"strings"
)

func getName(suffix string) string {
	var name string
	f, _ := os.Executable()
	if v, ok := os.LookupEnv("GOB_DIR"); ok {
		if _, err := os.Stat(v); os.IsNotExist(err) {
			os.Mkdir(v, os.ModePerm)
		}
		name = path.Join(v, strings.Join([]string{path.Base(f), suffix, "gob"}, "."))
	} else {
		name = strings.Join([]string{f, suffix, "gob"}, ".")
	}
	return name
}

// Dumps -
func Dumps(fileSuffix string, e interface{}) error {

	fl, err := os.OpenFile(getName(fileSuffix), os.O_CREATE|os.O_TRUNC|os.O_WRONLY, 0644)
	if err != nil {
		return err
	}
	defer fl.Close()
	return gb.NewEncoder(fl).Encode(e)
}

// Loads -
func Loads(fileSuffix string, e interface{}) error {
	fl, err := os.OpenFile(getName(fileSuffix), os.O_RDONLY, 0644)
	if err != nil {
		return err
	}
	defer fl.Close()
	return gb.NewDecoder(fl).Decode(e)
}
