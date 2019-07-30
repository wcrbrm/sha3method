package main

import (
	"bufio"
	"encoding/hex"
	"fmt"
	"github.com/ethereumproject/go-ethereum/crypto/sha3"
	"os"
)

func main() {
	hash := sha3.NewKeccak256()
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
	   hash.Write([]byte(scanner.Text()))
           buf := hash.Sum(nil)
	   fmt.Println(hex.EncodeToString(buf)[:8])
	}
	if scanner.Err() != nil {
	}
}


