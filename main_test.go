package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"testing"
)

func testRename(t *testing.T, src string) {
	_ = tmpFile()
}

func tmpFile() *os.File {
	tmpFile, err := ioutil.TempFile(os.TempDir(), "prefix-")
	if err != nil {
		log.Fatal("Cannot create temporary file", err)
		return nil
	}

	defer os.Remove(tmpFile.Name())
	fmt.Println("Created File: " + tmpFile.Name())
	return tmpFile
}
