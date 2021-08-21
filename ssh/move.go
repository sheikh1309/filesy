package ssh

import (
	"fmt"
	"github.com/sheikh1309/filesy/config"
)

func Move(credentials config.Credentials, source string, dest string)  {
	run(credentials, fmt.Sprintf("mv %v %v", source, dest))
}