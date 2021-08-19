package ssh

import (
	"fmt"
	"github.com/sheikh1309/filesy/config"
	"golang.org/x/crypto/ssh"
)

type Connection struct {
	client *ssh.Client
	session *ssh.Session
}

func connectToHost(credentials config.Credentials) Connection {
	sshConfig := &ssh.ClientConfig{
		User: credentials.User,
		Auth: []ssh.AuthMethod{ssh.Password(credentials.Password)},
	}
	sshConfig.HostKeyCallback = ssh.InsecureIgnoreHostKey()

	client, err := ssh.Dial("tcp", fmt.Sprintf("%v:%v", credentials.Host, credentials.Port), sshConfig)
	if err != nil {
		panic(fmt.Sprintf("Fail to connect: %v", err))
	}

	session, err := client.NewSession()
	if err != nil {
		client.Close()
		panic(fmt.Sprintf("Fail to Create Session: %v", err))
	}

	return Connection{ client, session }
}