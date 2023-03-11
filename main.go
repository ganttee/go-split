package main

import (
	"flag"
	"fmt"
	"io/ioutil"
)

var (
	inFile  string
	outFile string
	before  int
	after   int
)

func init() {
	flag.StringVar(&inFile, "infile", "", "input file, eg: -infile XXX")
	flag.StringVar(&outFile, "outfile", "", "output file, eg: -outfile XXX")
	flag.IntVar(&before, "before", 0, "before offset bytes, eg: -before 10")
	flag.IntVar(&after, "after", 0, "after offset bytes, eg: -after 10")
}

func main() {
	flag.Parse()

	inf, err := ioutil.ReadFile(inFile)
	if err != nil {
		fmt.Println(err)
		return
	}

	bufLen := len(inf)
	exBuf := inf[before:(bufLen - after)]

	err = ioutil.WriteFile(outFile, exBuf, 0644)
	if err != nil {
		fmt.Println("write error ", err)
		return
	}

	fmt.Println("Original File Size:", bufLen)
	fmt.Println("Anterior Offset Bytes:", before)
	fmt.Println("Back Offset Bytes:", after)
	fmt.Println("Write Byte Size:", len(exBuf))
}
