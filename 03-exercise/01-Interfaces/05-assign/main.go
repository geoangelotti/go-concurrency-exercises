package main

import (
	"bytes"
	"fmt"
	"io"
	"os"
	"time"
)

func (t *time.Duration) Write(p []byte) (n int, err error) {
	return 0, nil
}

func main() {
	var w io.Writer
	w = os.Stdout
	w = new(bytes.Buffer)
	w = time.Second
	fmt.Println(w)
}
