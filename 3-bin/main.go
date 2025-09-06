package main

import (
	"bin/file"
	"bin/storage"
	"fmt"
)

func main() {

	storageAccess := storage.NewStorage(file.NewFileStorageWithDefault())
	list, err := storageAccess.ReadBinList()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(list)

}
