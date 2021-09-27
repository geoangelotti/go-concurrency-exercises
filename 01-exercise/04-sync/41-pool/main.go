package main

import (
	"bytes"
	"fmt"
	"io"
	"os"
	"sync"
	"time"
)

var bufferPool = sync.Pool{
	New: func() interface{} {
		fmt.Println("Allocating new buffer")
		return new(bytes.Buffer)
	},
}

func log(w io.Writer, val string) {
	var b = bufferPool.Get().(*bytes.Buffer)

	b.WriteString(time.Now().Format("15:04:05"))
	b.WriteString(" : ")
	b.WriteString(val)
	b.WriteString("\n")

	w.Write(b.Bytes())

	bufferPool.Put(b)
}

func main() {
	log(os.Stdout, "debug-string1")
	log(os.Stdout, "debug-string2")
}
