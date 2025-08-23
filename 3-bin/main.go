package main

import (
	"bin/storage"
	"fmt"
)

func main() {
	binList, _ := storage.ReadBinList()
	fmt.Println(binList)

}
