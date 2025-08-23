package storage

import (
	"bin/bins"
	"bin/file"
	"encoding/json"
	"errors"
)

const (
	binStorageDir = "bins/"
	binFileName   = "bins.json"
)

func SaveBinList(list bins.BinList) error {
	listBytes, err := json.Marshal(list)
	if err != nil {
		return errors.New("CANT_SAVE_BINLIST_FROM_JSON: " + err.Error())
	}

	err = file.SaveFile(listBytes, binFileName, binStorageDir)
	if err != nil {
		return errors.New("CANT_SAVE_BINLIST: " + err.Error())
	}

	return nil
}

func ReadBinList() (*bins.BinList, error) {
	data, err := file.ReadFile(binFileName, binStorageDir)
	if err != nil {
		return nil, errors.New("CANT_READ_BINLIST_FROM_FILE: " + err.Error())
	}

	var result bins.BinList
	err = json.Unmarshal(data, &result)
	if err != nil {
		return nil, errors.New("CANT_READ_BIN_FROM_JSON: " + err.Error())
	}
	return &result, nil
}
