package main

import (
	"flag"
	"fmt"
	"os"
	"time"
)

// text are copied from original coreutil sleep
var (
	help = `Usage: sleep NUMBER[SUFFIX]...
  or:  sleep OPTION
Pause for NUMBER seconds.  SUFFIX may be 's' for seconds (the default),
'm' for minutes, 'h' for hours or 'd' for days.  Unlike most implementations
that require NUMBER be an integer, here NUMBER may be an arbitrary floating
point number.  Given two or more arguments, pause for the amount of time
specified by the sum of their values.

      --help     display this help and exit
      --version  output version information and exit

Report sleep bugs to bug-coreutils@gnu.org
GNU coreutils home page: <http://www.gnu.org/software/coreutils/>
General help using GNU software: <http://www.gnu.org/gethelp/>
For complete documentation, run: info coreutils 'sleep invocation'
`
	version = `sleep (GNU coreutils) 8.20
Copyright (C) 2012 Free Software Foundation, Inc.
License GPLv3+: GNU GPL version 3 or later <http://gnu.org/licenses/gpl.html>.
This is free software: you are free to change and redistribute it.
There is NO WARRANTY, to the extent permitted by law.

Written by Jim Meyering and Paul Eggert.
`

	helpFlag    = flag.Bool("help", false, help)
	versionFlag = flag.Bool("version", false, version)
)

func usage() {
	fmt.Printf("sleep: missing operand\nTry 'sleep --help' for more information.\n")
}

func main() {
	flag.Parse()
	if flag.NArg() == 0 && flag.NFlag() == 0 {
		usage()
		os.Exit(1)
	}

	if *helpFlag {
		fmt.Println(help)
		os.Exit(0)
	}

	if *versionFlag {
		fmt.Println(version)
		os.Exit(0)
	}

	var total time.Duration

	// coreutil's sleep says: "Given two or more arguments, pause for the amount
	// of time specified by the sum of their value"
	for i := 0; i < flag.NArg(); i++ {
		d, err := time.ParseDuration(flag.Arg(i))
		if err != nil {
			fmt.Printf("sleep: invalid time interval ‘%s’\n", flag.Arg(i))
			os.Exit(1)
		}

		total = total + d
	}

	// sleep for a total time of passed times
	time.Sleep(total)
}
