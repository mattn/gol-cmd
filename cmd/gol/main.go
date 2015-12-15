package main

import (
	"fmt"
	"io"
	"os"
	"time"
)

var out io.Writer = os.Stdout

func main() {
	fmt.Fprintln(out, "\x1b[2J")
	for i := 75; i > 0; i-- {
		fmt.Fprintf(out, "\x1b[10;%dH\x1bK", i)
		fmt.Print(`ʕ◔ϖ◔ʔ `)
		time.Sleep(50 * time.Millisecond)
	}
}
