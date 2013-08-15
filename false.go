package main

import (
	"flag"
	"fmt"
	"os"
)

var (
	help_text = `Usage: /bin/false [ignored command line arguments]
  or:  /bin/false OPTION
Exit with a status code indicating failure.

      --help     display this help and exit
      --version  output version information and exit

NOTE: your shell may have its own version of false, which usually supersedes
the version described here.  Please refer to your shell's documentation
for details about the options it supports.`

	help = flag.Bool("help", false, help_text)

	version_text = "false (go-coreutils) 0.1"
	version      = flag.Bool("version", false, version_text)
)

func main() {
	flag.Parse()

	if flag.NFlag() > 1 {
		os.Exit(-1)
	}

	if *help {
		fmt.Println(help_text)
		os.Exit(-1)
	}

	if *version {
		fmt.Println(version_text)
		os.Exit(-1)
	}
	os.Exit(-1)
}
