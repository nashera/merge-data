package civic

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

// VarCivic Variant in civic
type VarCivic struct {
	VariantID  string
	Chromosome string
	ChromStart string
	ChromEnd   string
	Name       string
	EntrezName string
	Score      string
	Hgvs       string
}

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

// ReadCivic 读取civic数据库数据
func ReadCivic(civicPath string) []*VarCivic {
	fmt.Println(civicPath)
	lines, errors, err := TsvLines(civicPath)
	if err != nil {
		log.Fatal(err)
	}
	go func(errs chan error) {
		err := <-errs
		if err != nil {
			log.Fatal(err)
		}
	}(errors)
	var varList []*VarCivic
	for line := range lines {
		w := strings.Split(line, "\t")
		varList = append(varList,
			&VarCivic{
				VariantID:  strings.Trim(w[0], "\""),
				Chromosome: strings.Trim(w[1], `"`),
				ChromStart: strings.Trim(w[2], `"`),
				ChromEnd:   strings.Trim(w[3], `"`),
				Name:       strings.Trim(w[4], `"`),
				EntrezName: strings.Trim(w[5], `"`),
				Score:      strings.Trim(w[6], `"`),
				Hgvs:       strings.Trim(w[7], `"`),
			})

		// fmt.Printf("%+v\n", string(line))
	}
	return varList

}
