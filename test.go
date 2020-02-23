package main

import (
	"log"
	"os"
)

func main() {
	_, err := os.Stat("C:/New Folder")
	if os.IsNotExist(err) {
		errDir := os.MkdirAll("C:/New folder", 0755)
		log.Fatal(err)
	}

}
