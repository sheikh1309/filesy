package ssh

import (
	"fmt"
	"github.com/sheikh1309/filesy/config"
)

func Remove(credentials config.Credentials, name string, force bool, recursive bool)  {
	options := ""
	if force {
		options += "-f "
	}
	if recursive {
		options += "-r "
	}
	run(credentials, fmt.Sprintf("rm %v %v", options, name))
}