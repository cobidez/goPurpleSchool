package file

import (
	"bin/bins"
	"encoding/json"
	"errors"
	"os"
	"path/filepath"
)

const (
	defaultBinFileName    = "bins.json"
	defaultBinStoragePath = "bins/"
)

type fileStorage struct {
	fileName string
	filePath string
}

func NewFileStorage(fileName, filePath string) *fileStorage {
	return &fileStorage{
		fileName: fileName,
		filePath: fileName,
	}
}

func NewFileStorageWithDefault() *fileStorage {
	return &fileStorage{
		fileName: defaultBinFileName,
		filePath: defaultBinStoragePath,
	}
}

func (fs *fileStorage) Write(list bins.BinList) error {
	var content []byte
	content, err := json.Marshal(list)
	if err != nil {
		return errors.New("CANT_SAVE_BINLIST_FROM_JSON: " + err.Error())
	}

	dir := filepath.Dir(fs.filePath)

	err = os.MkdirAll(dir, os.ModePerm)
	if err != nil {
		return errors.New("CANT_CREATE_DIR: " + err.Error())
	}

	file, err := os.OpenFile(dir+"/"+fs.fileName,
		os.O_CREATE|os.O_WRONLY|os.O_TRUNC,
		os.ModePerm)
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

func (fs *fileStorage) Read() (bins.BinList, error) {
	dir := filepath.Dir(fs.filePath)
	data, err := os.ReadFile(dir + "/" + fs.fileName)
	if err != nil {
		return nil, errors.New("CANT_READ_FILE: " + err.Error())
	}

	list := bins.BinList{}
	err = json.Unmarshal(data, &list)
	if err != nil {
		return nil, errors.New("CANT_READ_FILE: " + err.Error())
	}

	return list, nil
}

func IsJsonFile(file []byte) bool {
	var result any
	err := json.Unmarshal(file, &result)
	return err == nil
}
