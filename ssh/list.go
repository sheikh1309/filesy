package ssh

import (
	"fmt"
	"github.com/sheikh1309/filesy/config"
)

func List(credentials config.Credentials, dir string) []byte {
	return run(credentials, fmt.Sprintf("ls -lhF %v | sed '/total/ d'", dir))
}

func Tree(credentials config.Credentials, dir string) []byte {
	return run(credentials, fmt.Sprintf("ls -RlFh %v | sed '/total/ d'", dir))
}