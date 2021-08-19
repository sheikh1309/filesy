package ssh

import (
	"fmt"
	"github.com/sheikh1309/filesy/config"
)

func List(credentials config.Credentials, dir string)  {
	res := run(credentials, fmt.Sprintf("ls %v", dir))
	fmt.Println(string(res))
}