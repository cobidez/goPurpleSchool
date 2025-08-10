package bins

import (
	"errors"
	"strings"
	"time"
)

// Types
type Bin struct {
	id        string
	private   bool
	createdAt time.Time
	name      string
}

type BinList []Bin

// Funcs
func newBin(name string) (*Bin, error) {
	if strings.TrimSpace(name) == "" {
		return nil, errors.New("NAME_IS_EMPTY")
	}

	return &Bin{
		name:      name,
		createdAt: time.Now(),
	}, nil
}
