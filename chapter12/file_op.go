package main

import (
	"os"
	"fmt"
	"io"
	"bufio"
	"io/ioutil"
)

func main() {
	filePath := "README.md";
	readFile_raw(filePath);
	readFile_string(filePath);
}

func readFile_string(filePath string) {
	content, error :=ioutil.ReadFile(filePath);
	if error != nil {
		fmt.Printf("Failed to open file '%s'\n", error);
		return;
	}
	fmt.Println("\n===========\nreadFile with ioutil\n");
	fmt.Printf("file Content: %s\n", string(content));
}

func readFile_raw(filePath string){
	fmt.Println("readFile with raw bufIo")
	inputFile, inputError := os.Open(filePath);
	if inputError != nil {
		fmt.Printf("Failed to open file '%s'\n", inputError);
		return;
	}

	defer inputFile.Close();

	inputReader := bufio.NewReader(inputFile);
	for i:=0; ; i++ {
		content, _,readError := inputReader.ReadLine();
		if readError == io.EOF {
			return
		}
		fmt.Printf("line %d: %s\n", i, content);
	}
}
