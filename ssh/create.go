package ssh

import (
	"fmt"
	"github.com/sheikh1309/filesy/config"
)

func CreateFile(credentials config.Credentials, filename string)  {
	run(credentials, fmt.Sprintf("touch %v", filename))
}

func CreateDir(credentials config.Credentials, dir string)  {
	run(credentials, fmt.Sprintf("mkdir %v", dir))
}