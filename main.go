package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net"
	"os"

	"github.com/dghubble/go-twitter/twitter"
	"github.com/dghubble/oauth1"
	"golang.org/x/crypto/ssh"
)

var (
	hostPrivateKeySigner ssh.Signer
	twClient             = Twitter{}
)

type Twitter struct {
	Client *twitter.Client
}

func init() {
	keyPath := "./id_rsa"

	hostPrivateKey, err := ioutil.ReadFile(keyPath)
	if err != nil {
		panic(err)
	}

	hostPrivateKeySigner, err = ssh.ParsePrivateKey(hostPrivateKey)
	if err != nil {
		panic(err)
	}
}

func keyAuth(conn ssh.ConnMetadata, key ssh.PublicKey) (*ssh.Permissions, error) {
	log.Println(conn.RemoteAddr(), "authenticate with", key.Type())
	log.Printf("sshcapture: ip=%s user=%s type=%s\n", conn.RemoteAddr(), conn.User(), key.Type())

	return nil, fmt.Errorf("user %s (key-type %s) is bullshit and you're an asshole.", conn.User(), key.Type())
}

func passAuth(conn ssh.ConnMetadata, password []byte) (*ssh.Permissions, error) {
	msg := fmt.Sprintf("IP:%s tried to log in with username %s and password %s\n", conn.RemoteAddr(), conn.User(), string(password))
	log.Println(msg)
	_, _, _ = twClient.Client.Statuses.Update(msg, nil)
	return nil, fmt.Errorf("user %s (password %s) is bullshit and you're an asshole.", conn.User(), string(password))
}

func main() {
	twconfig := oauth1.NewConfig("consumerKey", os.Getenv("TWITTER_CONSUMER"))
	token := oauth1.NewToken("accessToken", os.Getenv("TWITTER_TOKEN"))
	httpClient := twconfig.Client(oauth1.NoContext, token)
	twClient.Client = twitter.NewClient(httpClient)

	config := ssh.ServerConfig{
		PublicKeyCallback: keyAuth,
		PasswordCallback:  passAuth,
		ServerVersion:     "SSH-2.0-YOURMOM",
	}
	config.AddHostKey(hostPrivateKeySigner)

	port := "2022"
	if os.Getenv("PORT") != "" {
		port = os.Getenv("PORT")
	}
	socket, err := net.Listen("tcp", ":"+port)
	if err != nil {
		log.Printf("[Error] %s", err)
	}

	for {
		conn, err := socket.Accept()
		if err != nil {
			log.Printf("[Error] %s", err)
			continue
		}

		// From a standard TCP connection to an encrypted SSH connection
		sshConn, _, _, err := ssh.NewServerConn(conn, &config)
		if err != nil {
			log.Printf("[Error] %s", err)
			continue
		}

		log.Println("Connection from", sshConn.RemoteAddr())
		sshConn.Close()
	}
}
