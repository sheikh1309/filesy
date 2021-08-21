package ssh

import (
	"fmt"
	"github.com/sheikh1309/filesy/config"
)

func Copy(credentials config.Credentials, source string, dest string, recursive bool)  {
	options := ""
	if recursive {
		options += "-R "
	}
	run(credentials, fmt.Sprintf("cp %v %v %v", options, source, dest))
}