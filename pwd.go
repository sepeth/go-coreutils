// pwd program for go-coreutils - fka

package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	pwd, err := os.Getwd()
	if err != nil {
		log.Fatalln(err)
	} else {
		fmt.Println(pwd)
	}
}
