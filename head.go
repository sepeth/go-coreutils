package main

import (
	"bufio"
	"bytes"
	"flag"
	"fmt"
	"io"
	"math"
	"net"
	"os"
	"regexp"
	"strconv"
)

// TODO: these needs long opts too...
// TODO: coreutils has this
// K may have a multiplier suffix:\n\
// b 512, kB 1000, K 1024, MB 1000*1000, M 1024*1024,\n\
// GB 1000*1000*1000, G 1024*1024*1024, and so on for T, P, E, Z, Y.\n\

var flagBytes = flag.String("c", "0", "number of bytes")
var flagLines = flag.String("n", "10", "number of lines")
var flagQuiet = flag.Bool("q", false, "quiet")
var flagVerbose = flag.Bool("v", false, "verbose")

func openFile(s string) (io.ReadWriteCloser, error) {
	fi, err := os.Stat(s)
	if err != nil {
		return nil, err
	}
	if fi.Mode()&os.ModeSocket != 0 {
		return net.Dial("unix", s)
	}
	return os.Open(s)
}

/**
print n lines from head to tail
*/
func readLines(lineCount int, r io.Reader, w io.Writer) (err error) {
	var line string
	br := bufio.NewReader(r)
	nr := 0
	for {
		line, err = br.ReadString('\n')
		fmt.Fprint(w, line)
		nr++
		if nr >= lineCount {
			return
		}

		if err != nil {
			return
		}
	}
}

/**
read n bytes from head
**/
func readBytes(byteCount int, r io.Reader, w io.Writer) (err error) {
	br := bufio.NewReader(r)
	out := bufio.NewWriter(w)

	var c byte
	nr := 0
	for {
		// we read bytes, not chars, because head does it like that
		// for multibyte strings
		// so
		// echo 'ıy' | head -n 2
		// is ı not ıy
		// in both.
		c, err = br.ReadByte()
		out.WriteByte(c)
		nr++
		if nr >= byteCount {
			out.Flush()
			return
		}
	}
}

func elideTailBytes(elideCount int, r io.Reader, w io.Writer) (err error) {
	/**
	 * see elide tail lines
	 **/
	var buffer bytes.Buffer
	var c byte

	out := bufio.NewWriter(w)
	br := bufio.NewReader(r)
	nr := 0
	for {
		c, err = br.ReadByte()
		buffer.WriteByte(c)
		if err == io.EOF {
			// end of file reached

			// the count of lines in the file is 10 and we want -20 lines
			// we exit
			if nr < -1*elideCount {
				return
			}

			// how many lines do we need to print
			pCnt := nr + elideCount
			t := 1
			for {
				c, err = buffer.ReadByte()
				out.WriteByte(c)
				if t >= pCnt {
					out.Flush()
					return
				}
				t++
			}
			fmt.Println("eof")
			return
		} else if err != nil {
			return
		}
		nr++
	}
}

/**
print all but the last K lines of each file

eg:
if the file has these lines,
1
2
3

head -n -1 should print
1
2

**/
func elideTailLines(elideCount int, r io.Reader, w io.Writer) (err error) {
	/**
	 * please see elide_tail_lines_seekable at
	 * https://github.com/goj/coreutils/blob/rm-d/src/head.c for
	 * a much better implementation, this is just a naive implementation,
	 * we buffer all input though fd might be seekable, in that case,
	 * we have to use another function and move the fd back.
	 * though in many cases - the file is small or -n is not a ridicilious amount - ,
	 * this is faster than moving the fd forth and back
	 */
	var line string
	var bufline string
	var buffer bytes.Buffer

	br := bufio.NewReader(r)
	nr := 0
	for {
		line, err = br.ReadString('\n')
		buffer.WriteString(line)
		if err == io.EOF {
			// end of file reached

			// the count of lines in the file is 10 and we want -20 lines
			// we exit
			if nr < -1*elideCount {
				return
			}

			// how many lines do we need to print
			pCnt := nr + elideCount
			t := 1
			for {
				bufline, err = buffer.ReadString('\n')
				fmt.Print(bufline)
				if t >= pCnt {
					return
				}
				t++
			}
			fmt.Println("eof")
			return
		} else if err != nil {
			return
		}
		nr++
	}
}

func SuffixedArgToInt(arg string) int {
	i, err := strconv.Atoi(arg)
	if err == nil {
		return i
	}
	// K may have a multiplier suffix:\n\
	// b 512, kB 1000, K 1024, MB 1000*1000, M 1024*1024,\n\
	// GB 1000*1000*1000, G 1024*1024*1024, and so on for T, P, E, Z, Y.\n\

	// gnu-coreutils uses xstrtoumax from https://github.com/gagern/gnulib/blob/master/lib/xstrtol.c
	// this is just a similar and not so good implementation of it - cant
	// find anything similar in go lib, but this covers most of it
	// glad that we dont have MiB in here
	re := regexp.MustCompile("([0-9]+?)([bkmBKMGTPEZY]+)")
	matches := re.FindAllStringSubmatch(arg, -1)

	if len(matches) == 0 {
		return 0
	}

	number, err := strconv.Atoi(matches[0][1])
	if err != nil {
		return 0
	}

	multiplier := matches[0][2]
	var power float64

	switch multiplier {
	case "b":
		power = 512
	case "k":
		power = 1024
	case "kB":
		power = 1000
	case "MB":
		power = 1000 * 1000
	case "GB":
		power = math.Pow(1000, 3)
	case "TB": // terra
		power = math.Pow(1024, 4)
	case "PB": // peta
		power = math.Pow(1024, 5)
	case "EB": // exa
		power = math.Pow(1024, 6)
	case "ZB": // zetta
		power = math.Pow(1024, 7)
	case "YB": // yotta
		power = math.Pow(1024, 8)
	case "M":
		power = 1024 * 1024
	case "G":
		power = 1024 * 1024 * 1024
	case "T": // terra
		power = math.Pow(1024, 4)
	case "P": // peta
		power = math.Pow(1024, 5)
	case "E": // exa
		power = math.Pow(1024, 6)
	case "Z": // zetta
		power = math.Pow(1024, 7)
	case "Y": // yotta
		power = math.Pow(1024, 8)
	default:
		power = 0

	}
	return number * int(power)
}

func main() {
	flag.Parse()

	bytes := SuffixedArgToInt(*flagBytes)
	lines := SuffixedArgToInt(*flagLines)

	// if there are no args, head should wait until we get
	// somthing from the stdin.
	args := flag.Args()
	if len(args) == 0 {

		// why we cant have a nil in flags' default ??
		if bytes > 0 {
			readBytes(bytes, os.Stdin, os.Stdout)
		} else if bytes < 0 {
			elideTailBytes(bytes, os.Stdin, os.Stdout)
		} else if lines < 0 {
			elideTailLines(lines, os.Stdin, os.Stdout)
		} else {
			readLines(lines, os.Stdin, os.Stdout)
		}

	} else {
		for _, fname := range flag.Args() {
			f, err := openFile(fname)
			defer f.Close()
			if err != nil {
				fmt.Fprintln(os.Stderr, err)
				continue
			}

			// why we cant have a nil in flags' default ??
			if bytes > 0 {
				readBytes(bytes, f, os.Stdout)
			} else if bytes < 0 {
				elideTailBytes(bytes, f, os.Stdout)
			} else if lines < 0 {
				elideTailLines(lines, f, os.Stdout)
			} else {
				readLines(lines, f, os.Stdout)
			}
		}
	}
}
