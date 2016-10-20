package log

import (
	"bufio"
	"fmt"
	"os"
)

var bw *bufio.Writer
var fp *os.File
var err error

func init() {
	fp, err = os.Create("./test.log")
	if err != nil {
		fmt.Println(err)
	}
	writer := bufio.NewWriter(fp)
	bw = bufio.NewWriter(writer)
}

func D(v ...interface{}) {
	bw.WriteString(fmt.Sprintln(v...))
}

func Close() {
	bw.Flush()
	fp.Close()
}
