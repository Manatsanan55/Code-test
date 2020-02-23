package main

import (
	"fmt"
	"os"
	"path/filepath"
)

func walkFn(path string, info os.FileInfo, err error) error {
	fmt.Println(path)
	return nil
}

func main() {
	filepath.Walk("C:/New folder", walkFn)
	slash := string(filepath.Separator)
	dirname := "." +slash
	fmt.Printf("dirname is: %s\n", dirname)
	d, err := os.Open(dirname)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	filesinfo, err := d.Readdir(-1)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	for _, fi := range filesinfo {
		if fi.Mode().IsRegular() {
			fmt.Printf("-----\nFileName: %s \nSize: \t%d bytes\n",
			    fi.Name(), fi.Size())
		}
	}
}