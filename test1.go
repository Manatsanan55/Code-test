package main

import "os"

import "fmt"

func walkFn(path string, info os.FileInfo, err error) error {
	fmt.Println(path)
	return nil
}
