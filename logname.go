package main

import (
	"flag"
	"fmt"
	"os"
	"os/user"
)

var help string = "Usage: logname \nPrint the name of the current user"

var helpFlag = flag.Bool("help", false, help)

func GetCurrentUser(username *string) error {

	current_user, err := user.Current()

	*username = current_user.Username

	if err != nil {
		return err
	}

	return nil
}

func main() {

	flag.Parse()

	if *helpFlag {
		fmt.Println(help)
		os.Exit(0)
	}

	if flag.NArg() > 0 {
		fmt.Println(help)
		os.Exit(0)
	}

	if flag.NArg() == 0 && flag.NFlag() == 0 {

		var username string

		err := GetCurrentUser(&username)

		if err != nil {
			fmt.Fprintln(os.Stderr, err)
		}

		fmt.Println(username)

	}

}
