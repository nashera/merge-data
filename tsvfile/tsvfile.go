package tsvfile

import (
	"bufio"
	"fmt"
	"os"
)

// TsvLines Get Lines From Tsv file
func TsvLines(filename string) (chan string, chan error, error) {
	fmt.Println(filename)
	file, err := os.Open(filename)
	if err != nil {
		return nil, nil, err
	}

	ch := make(chan string)
	errs := make(chan error)
	go func(ch chan string, file *os.File, errs chan error) {
		defer func(ch chan string, file *os.File, errs chan error) {
			close(ch)
			close(errs)
			file.Close()
		}(ch, file, errs)
		contents := bufio.NewScanner(file)
		cbuffer := make([]byte, 0, bufio.MaxScanTokenSize)
		contents.Buffer(cbuffer, bufio.MaxScanTokenSize*50) // Otherwise long lines crash the scanner.
		var (
			err error
		)
		for contents.Scan() {
			ch <- contents.Text()
		}
		if err = contents.Err(); err != nil {
			errs <- err
			return
		}
	}(ch, file, errs)

	return ch, errs, nil
}
