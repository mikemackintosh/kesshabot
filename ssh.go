package main

import (
	"crypto/md5"
	"fmt"
	"log"
	"net"
	"os"

	"golang.org/x/crypto/ssh"
)

func startSSH() {
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
			log.Printf("[Error] Accepting Socket: %s", err)
			conn.Close()
			continue
		}

		// From a standard TCP connection to an encrypted SSH connection
		sshConn, _, _, err := ssh.NewServerConn(conn, &config)
		defer sshConn.Close()
		if err != nil {
			log.Printf("[Error] Creating New SSH Connection: %s", err)
			conn.Close()
			continue
		}

		log.Println("Connection from", sshConn.RemoteAddr())
		sshConn.Close()
	}
}

func keyAuth(conn ssh.ConnMetadata, key ssh.PublicKey) (*ssh.Permissions, error) {
	md5sum := md5.Sum(key.Marshal())

	msg := fmt.Sprintf("PubKey Auth Attempted:\n💻 %s\n👤 %s\n🔐 %s\n", conn.RemoteAddr(), conn.User(), rfc4716hex(md5sum[:]))
	log.Println(msg)
	_, _, err := twClient.Client.Statuses.Update(msg, nil)
	if err != nil {
		log.Printf("[Error] pubkeyAuth: %s\n", err)
	}
	return nil, fmt.Errorf("Never auth")
}

// Pass auth will record password authentication attempts to the server
func passAuth(conn ssh.ConnMetadata, password []byte) (*ssh.Permissions, error) {
	msg := fmt.Sprintf("Password Auth Attempted:\n💻 %s\n👤 %s\n🔐 %s\n", conn.RemoteAddr(), conn.User(), string(password))
	log.Println(msg)
	_, _, err := twClient.Client.Statuses.Update(msg, nil)
	if err != nil {
		log.Printf("[Error] passAuth: %s\n", err)
	}
	return nil, fmt.Errorf("Never auth")
}

func rfc4716hex(data []byte) string {
	var fingerprint string
	for i := 0; i < len(data); i++ {
		fingerprint = fmt.Sprintf("%s%0.2x", fingerprint, data[i])
		if i != len(data)-1 {
			fingerprint = fingerprint + ":"
		}
	}
	return fingerprint
}
