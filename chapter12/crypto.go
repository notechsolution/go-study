package main

import (
	"io"
	"fmt"
	"crypto/sha1"
	"crypto/sha256"
)

func main() {
	hasher := sha1.New();
	io.WriteString(hasher, "I am the miner");
	fmt.Printf("hash of 'I am the miner is' %x\n", hasher.Sum([]byte{}));
	fmt.Printf("hash of 'I am the miner is' %x\n", hasher.Sum([]byte{}));

	hasher256 := sha256.New();
	io.WriteString(hasher, "I am the miner");
	fmt.Printf("sha256 of 'I am the miner is' %x\n", hasher256.Sum([]byte{}));
	fmt.Printf("sha256 of 'I am the miner is' %x\n", hasher256.Sum([]byte{}));
}
