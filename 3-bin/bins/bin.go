package bins

import (
	"encoding/json"
	"errors"
	"strings"
	"time"
)

type Bin struct {
	Id        string    `json:"id"`
	Private   bool      `json:"private"`
	CreatedAt time.Time `json:"created_at"`
	Name      string    `json:"name"`
}

type BinList []Bin

func NewBin(name string) (*Bin, error) {
	if strings.TrimSpace(name) == "" {
		return nil, errors.New("NAME_IS_EMPTY")
	}

	return &Bin{
		Name:      name,
		CreatedAt: time.Now(),
	}, nil
}

func (bin *Bin) GetBytes() ([]byte, error) {
	bytes, err := json.Marshal(bin)
	if err != nil {
		return nil, errors.New("CANT_SERIALIZE_BIN: " + err.Error())
	}
	return bytes, nil
}

func (list *BinList) GetBytes() ([]byte, error) {
	bytes, err := json.Marshal(list)
	if err != nil {
		return nil, errors.New("CANT_SERIALIZE_BIN: " + err.Error())
	}
	return bytes, nil
}

func (list *BinList) Append(bin *Bin) {
	*list = append(*list, *bin)
}
