package file

import (
	"encoding/json"
	"errors"
	"os"
	"path/filepath"
)

func SaveFile(content []byte, name string, path string) error {
	dir := filepath.Dir(path)
	err := os.MkdirAll(dir, os.ModePerm)
	if err != nil {
		return errors.New("CANT_CREATE_DIR: " + err.Error())
	}

	file, err := os.OpenFile(dir+"/"+name, os.O_CREATE|os.O_WRONLY|os.O_TRUNC, os.ModePerm)
	if err != nil {
		return errors.New("CANT_OPEN_CREATE_FILE: " + err.Error())
	}
	defer file.Close()

	_, err = file.Write(content)
	if err != nil {
		return errors.New("CANT_WRITE_FILE: " + err.Error())
	}

	return nil
}

func ReadFile(name string, path string) ([]byte, error) {
	dir := filepath.Dir(path)
	data, err := os.ReadFile(dir + "/" + name)
	if err != nil {
		return nil, errors.New("CANT_READ_FILE: " + err.Error())
	}

	return data, nil
}

func IsJsonFile(file []byte) bool {
	var result any
	return json.Unmarshal(file, result) == nil
}
