package log

import (
	"bufio"
	"fmt"
	"os"
	"time"
)

var bw *bufio.Writer
var fp *os.File
var err error

func init() {
	filename := time.Now().Format("log/2006_0102_1504_05.log")
	fp, err = os.Create(filename)
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
