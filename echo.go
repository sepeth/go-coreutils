package main

import (
	"flag"
	"fmt"
	"os"
	"strings"
)

func usage() {
	fmt.Println("Usage: echo [options] [string ...]")
}

var enableEscapeChars = flag.Bool("e", false, "Enable escape characters")
var omitNewline = flag.Bool("n", false, "don't print trailing newline")
var disableEscapeChars = flag.Bool("E", true, "Disable escape characters")

func main() {
	flag.Parse()
	concatenated := strings.Join(flag.Args(), " ")
	a := []rune(concatenated)
	length := len(a)
	ai := 0
	if (*enableEscapeChars == true || *disableEscapeChars == false) && length != 0 {
		for i := 0; i < length; {
			c := a[i]
			i++
			if c == '\\' && i < length {
				c = a[i]
				i++
				switch c {
				case 'a':
					c = '\a'
					break
				case 'b':
					c = '\b'
					break
				case 'c':
					os.Exit(0)
				case 'e':
					c = '\x1B'
					break
				case 'f':
					c = '\f'
					break
				case 'n':
					c = '\n'
					break
				case 'r':
					c = '\r'
					break
				case 't':
					c = '\t'
					break
				case 'v':
					c = '\v'
					break
				case '\\':
					c = '\\'
					break
				case 'x':
					c = a[i]
					i++
					if '9' >= c && c >= '0' && i < length {
						hex := (c - '0')
						c = a[i]
						i++
						if '9' >= c && c >= '0' && i < length {
							c = 16*(c-'0') + hex
						}
					}
					break
				}
			}
			a[ai] = c
			ai++
		}
	}
	fmt.Print(string(a[:ai]))
	if *omitNewline == false {
		fmt.Print("\n")
	}
}
