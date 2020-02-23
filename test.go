package main

import (
	"os"
)

func main() {
	_, err := os.Stat("C:/New Folder")
	if os.IsNotExist(err) {
	}

}
