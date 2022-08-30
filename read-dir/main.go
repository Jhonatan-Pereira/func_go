package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	dir, err := os.Open("/tmp")
	if err != nil {
		panic(err)
	}

	for {
		files, err := dir.Readdir(1)
		if err != nil {
			if err == io.EOF {
				break
			}
			fmt.Printf("Error: %v", err)
			continue
		}

		fmt.Printf(files[0].Name())
	}
}
