package main

import (
	"io/ioutil"
	"os"
	"sync"

	"golang.org/x/crypto/ssh"
)

var (
	hostPrivateKeySigner ssh.Signer
)

func init() {
	keyPath := "/etc/kessha/id_rsa"
	if len(os.Getenv("KESSHABOT_RSAID")) > 0 {
		keyPath = os.Getenv("KESSHABOT_RSAID")
	}

	hostPrivateKey, err := ioutil.ReadFile(keyPath)
	if err != nil {
		panic(err)
	}

	hostPrivateKeySigner, err = ssh.ParsePrivateKey(hostPrivateKey)
	if err != nil {
		panic(err)
	}
}

func main() {
	// Lets setup twitter and the shits
	setupTwitter()

	wg := sync.WaitGroup{}
	wg.Add(1)
	go startSSH()

	wg.Add(1)
	go startDNS()

	wg.Wait()
}
