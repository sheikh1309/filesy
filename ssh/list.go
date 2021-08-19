package ssh

import (
	"fmt"
	"github.com/sheikh1309/filesy/config"
)

func List(credentials config.Credentials, dir string) []byte {
	return run(credentials, fmt.Sprintf("ls -lh %v", dir))
}