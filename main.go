package main

import (
	"bufio"
	"compress/gzip"
	"flag"
	"fmt"
	"log"
	"os"
	"strings"
)

const (
	// Version 版本号
	Version = "v1.0"
)

var (
	cosmicTsv string
)

// CosmicVar variant in cosmic database


func init() {
	flag.StringVar(&cosmicTsv, "cosmic", "/cgdata/zhangxi/project/cosmic/test.tsv.gz", "the path ot cosmicTSV")
	flag.Parse()
}



func main() {
	varList := ReadCosmic(cosmicTsv)
	for _, variant := range varList {
		fmt.Println(variant)
	}

}
