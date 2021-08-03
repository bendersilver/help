package gob

import (
	gb "encoding/gob"
	"fmt"
	"os"
	"path"
)

func getName(suffix string) string {
	f, _ := os.Executable()
	var name string = fmt.Sprintf("%s.%s.gob", path.Base(f), suffix)
	if v, ok := os.LookupEnv("GOB_DIR"); ok {
		if _, err := os.Stat(v); os.IsNotExist(err) {
			os.Mkdir(v, os.ModePerm)
		}
		return path.Join(v, name)
	}
	return path.Join(path.Dir(f), name)
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
