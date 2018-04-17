package main

import (
	"encoding/json"
	"fmt"
	"os"
	"encoding/gob"
)

type Address struct {
	Type    string
	City    string
	Country string
}

type VCard struct {
	FirstName string
	LastName  string
	Addresses []*Address
	Remark    string
}

func main() {
	// JSON
	pa := &Address{"private", "Aartselaar", "Belgium"}
	wa := &Address{"work", "Boom", "Belgium"}
	vc := VCard{"Jan", "Kersschot", []*Address{pa, wa}, "none"}
	json, _ := json.Marshal(vc);
	fmt.Printf("jsonString:%s\n", json);

	// GOB
	gobFile, _ := os.OpenFile("gobFile.gob", os.O_CREATE|os.O_WRONLY, 0666);
	defer gobFile.Close();
	encoder := gob.NewEncoder(gobFile);
	error := encoder.Encode(vc);
	if error != nil {
		fmt.Printf("error during encode %s\n", error);
	}
}
