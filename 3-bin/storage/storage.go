package storage

import (
	"bin/bins"
	"errors"
)

type Storage struct {
	bins.ReadWriter
}

func NewStorage(rw bins.ReadWriter) *Storage {
	return &Storage{rw}
}

func (storage *Storage) SaveBinList(list bins.BinList) error {
	err := storage.Write(list)
	if err != nil {
		return errors.New("CANT_SAVE_BINLIST: " + err.Error())
	}

	return nil
}

func (storage *Storage) ReadBinList() (bins.BinList, error) {
	list, err := storage.Read()
	if err != nil {
		return nil, errors.New("CANT_READ_BINLIST_FROM_FILE: " + err.Error())
	}

	return list, nil
}
