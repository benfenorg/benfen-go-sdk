package main

import (
	"fmt"
	"io/ioutil"
)

func ReadFile() string {

	// file path.
	filePath := "bfc.keystore"

	// read file
	content, err := ioutil.ReadFile(filePath)
	if err != nil {
		fmt.Println("Can not read file : ", err)
		return ""
	}

	// convert to string and output
	fmt.Println(string(content))
	return string(content)
}

func Write_file(content []byte, fileName string) {

	filepath := fileName
	err := ioutil.WriteFile(filepath, []byte(content), 0644)
	if err != nil {
		fmt.Println("Can not write file : ", err)
	}

	fmt.Println("file write success!")
}
