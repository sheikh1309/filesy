package ssh

import (
	"github.com/sheikh1309/filesy/config"
	"fmt"
)

func run(credentials config.Credentials, cmd string) []byte {
	connection := connectToHost(credentials)
	out, err := connection.session.CombinedOutput(cmd)
	if err != nil {
		panic(fmt.Sprintf("Fail to run ssh command, %v", err))
	}
	err = connection.client.Close()
	if err != nil {
		panic(fmt.Sprintf("Fail to close ssh client, %v", err))
	}
	return out
}
