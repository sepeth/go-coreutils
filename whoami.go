package main

import (
	"flag"
	"fmt"
	"os/user"
)

var (
	help_text = `Usage: whoami [OPTION]...
Print the user name associated with the current effective user ID.
Same as id -un.

      --help     display this help and exit
      --version  output version information and exit

Report whoami bugs to bug-coreutils@gnu.org
GNU coreutils home page: <http://www.gnu.org/software/coreutils/>
General help using GNU software: <http://www.gnu.org/gethelp/>
For complete documentation, run: info coreutils 'whoami invocation'`

	version_text = `go-whoami (go-coreutils) 0.1`

	help    = flag.Bool("help", false, help_text)
	version = flag.Bool("version", false, version_text)
)

func main() {
	flag.Parse()
	if *help {
		fmt.Println(help_text)
		return
	}
	if *version {
		fmt.Println(version_text)
		return
	}
	current_user, _ := user.Current()
	fmt.Println(current_user.Username)
}
