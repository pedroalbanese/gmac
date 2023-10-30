# GMAC
[![ISC License](http://img.shields.io/badge/license-ISC-blue.svg)](https://github.com/pedroalbanese/gmac/blob/master/LICENSE.md) 
[![GoDoc](https://godoc.org/github.com/pedroalbanese/gmac?status.png)](http://godoc.org/github.com/pedroalbanese/gmac)
[![GitHub downloads](https://img.shields.io/github/downloads/pedroalbanese/gmac/total.svg?logo=github&logoColor=white)](https://github.com/pedroalbanese/gmac/releases)
[![Go Report Card](https://goreportcard.com/badge/github.com/pedroalbanese/gmac)](https://goreportcard.com/report/github.com/pedroalbanese/gmac)
[![GitHub go.mod Go version](https://img.shields.io/github/go-mod/go-version/pedroalbanese/gmac)](https://golang.org)
[![GitHub release (latest by date)](https://img.shields.io/github/v/release/pedroalbanese/gmac)](https://github.com/pedroalbanese/gmac/releases)

### RFC 4543 Galois Message Authentication Code (GMAC) (NIST SP800-38D)

## Usage
```go
package main

import (
	"crypto/aes"
	"encoding/hex"
	"fmt"
	"log"

	"github.com/pedroalbanese/gmac"
)

func main() {
	// Example using AES block cipher
	key, _ := hex.DecodeString("00000000000000000000000000000000")
	nonce, _ := hex.DecodeString("00000000000000000000000000000000")
	message := []byte("Yoda said, do or do not. There is no try.")

	block, _ := aes.NewCipher(key)
	gmac, err := gmac.CalculateGMAC(block, nonce, message)
	if err != nil {
		log.Fatal(err)
	}

	// Print out GMAC
	fmt.Printf("GMAC: %x\n", gmac)
}
```

Result: `GMAC: e7ee2c63b4dc328eed4a86b3fb3490af`

## License

This project is licensed under the ISC License.

#### Copyright (c) 2020-2023 Pedro F. Albanese - ALBANESE Research Lab.

