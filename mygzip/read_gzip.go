import (
	"bufio"
	"compress/gzip"
	"flag"
	"fmt"
	"log"
	"os"
	"strings"
)


// GZLines iterates over lines of a file that's gzip-compressed.
// Iterating lines of an io.Reader is one of those things that Go
// makes needlessly complex.
func GZLines(filename string) (chan []byte, chan error, error) {
	rawf, err := os.Open(filename)
	if err != nil {
		return nil, nil, err
	}
	rawContents, err := gzip.NewReader(rawf)
	if err != nil {
		return nil, nil, err
	}
	contents := bufio.NewScanner(rawContents)
	cbuffer := make([]byte, 0, bufio.MaxScanTokenSize)
	contents.Buffer(cbuffer, bufio.MaxScanTokenSize*50) // Otherwise long lines crash the scanner.
	ch := make(chan []byte)
	errs := make(chan error)
	go func(ch chan []byte, errs chan error, contents *bufio.Scanner) {
		defer func(ch chan []byte, errs chan error) {
			close(ch)
			close(errs)
		}(ch, errs)
		var (
			err error
		)
		for contents.Scan() {
			ch <- contents.Bytes()
		}
		if err = contents.Err(); err != nil {
			errs <- err
			return
		}
	}(ch, errs, contents)
	return ch, errs, nil
}
