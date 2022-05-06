package main

import (
	"fmt"
	"strings"
	"time"

	"github.com/algorand/go-algorand-sdk/crypto"
	"github.com/algorand/go-algorand-sdk/mnemonic"
)

const prefix = "ALGO"
const workers = 20

func main() {
	fmt.Printf("Attempting to generate account with prefix '%s'...\n", prefix)
	startTime := time.Now()

	syncChan := make(chan string, 1)

	for i := 0; i < workers; i++ {
		go doWork(syncChan)
	}

	fmt.Println(<-syncChan)

	close(syncChan)

	fmt.Printf("Time taken: %s\n", time.Since(startTime))
}

func doWork(syncChan chan string) {
	for {
		account := crypto.GenerateAccount()
		passphrase, _ := mnemonic.FromPrivateKey(account.PrivateKey)
		myAddress := account.Address.String()
		if strings.HasPrefix(myAddress, prefix) {
			syncChan <- fmt.Sprintf("Address: '%s'\nPassphrase: '%s'", myAddress, passphrase)
		}
	}
}
