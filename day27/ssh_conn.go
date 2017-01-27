package main

import (
	"io/ioutil"
	"log"

	"golang.org/x/crypto/ssh"
	"golang.org/x/oauth2"
)

// variable representing your token at digital ocean
var pat = "past here your token"

type TokenSource struct {
	AccessToken string
}

func (t *TokenSource) Token() (*oauth2.Token, error) {
	token := &oauth2.Token{
		AccessToken: t.AccessToken,
	}
	return token, nil
}
func PublicKeyFile(file string) ssh.AuthMethod {
	buffer, err := ioutil.ReadFile(file)
	if err != nil {
		log.Println(err)
		return nil
	}

	key, err := ssh.ParsePrivateKey(buffer)
	if err != nil {
		log.Println(err)
		return nil
	}
	return ssh.PublicKeys(key)
}

func main() {
	// Create an droplet
	// copy ip and paste below
	// don't forget to change pat variable and set your ssh key in the machine
	sshConfig := &ssh.ClientConfig{
		User: "root",
		Auth: []ssh.AuthMethod{
			PublicKeyFile("id_rsa"),
		},
	}
	connection, err := ssh.Dial("tcp", "paste here the ip:22", sshConfig)
	if err != nil {
		log.Fatalf("Failed to dial: %s", err)
	}
	session, err := connection.NewSession()
	if err != nil {
		log.Fatalf("Failed to create session: %s", err)
	}
	// to change the command that will be runned, only modify CombinedOutput
	output, err := session.CombinedOutput("ls -la")
	if err != nil {
		log.Fatal(err)
	}
	println(string(output))
}
