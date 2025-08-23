package main

import (
	"bin/file"
	"fmt"
)

func main() {
	data, err := file.ReadFile("bins.json", "bins/")

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(file.IsJsonFile(data))

}
