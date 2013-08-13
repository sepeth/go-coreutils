package main

import (
	"flag"
	"fmt"
	"os"
	"os/user"
)

var (
	help_text = `Usage: whoami [OPTION]...
Print the user name associated with the current effective user ID.
Same as id -un.

      --help     display this help and exit
      --version  output version information and exit`

	version_text = `go-whoami (go-coreutils) 0.1`

	help    = flag.Bool("help", false, help_text)
	version = flag.Bool("version", false, version_text)
)

func main() {
	flag.Parse()
	if *help {
		fmt.Println(help_text)
		os.Exit(0)
	}
	if *version {
		fmt.Println(version_text)
		os.Exit(0)
	}
	current_user, err := user.Current()
	if err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}
	fmt.Println(current_user.Username)
	os.Exit(0)
}
