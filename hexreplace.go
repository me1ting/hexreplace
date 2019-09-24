package main

import (
	"bytes"
	hex2 "encoding/hex"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

var oldHex = flag.String("o", "", "old hex")
var newHex = flag.String("n", "", "new hex")

func main() {
	newUsage()
	flag.Parse()
	var file, save string
	if len(flag.Args()) == 0 {
		log.Fatalf("Try 'hexreplace --help' for more information.")
	} else if len(flag.Args()) == 0 {
		file = flag.Arg(0)
		save = file + "_new"
	} else {
		file = flag.Arg(0)
		save = flag.Arg(1)
	}

	if *oldHex == "" {
		log.Fatalf("Old hex can't be empty")
	}

	replaceHex(file, save, *oldHex, *newHex)
}

func newUsage() {
	flag.Usage = func() {
		fmt.Fprintf(flag.CommandLine.Output(), "Usage: %s\n", "hexreplace [OPTION]... SOURCE DEST")
		fmt.Fprintf(flag.CommandLine.Output(), "Replace hex in SOURCE, save as DEST.\n")
		fmt.Fprintf(flag.CommandLine.Output(), "\n")
		fmt.Fprintf(flag.CommandLine.Output(), "OPTION:\n")
		flag.PrintDefaults()
	}
}

func replaceHex(f string, save string, old string, new string) {
	oldBytes := hex2bytes(old)
	newBytes := hex2bytes(new)

	fileBytes, err := ioutil.ReadFile(f)
	check(err)
	if count := bytes.Count(fileBytes, oldBytes); count > 0 {
		log.Printf("find old hex at %d places\n", count)
		newFileBytes := bytes.ReplaceAll(fileBytes, oldBytes, newBytes)
		err = ioutil.WriteFile(save, newFileBytes, os.ModePerm)
		check(err)
	}
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func hex2bytes(hex string) []byte {
	src := []byte(hex)
	maxEnLen := hex2.DecodedLen(len(src))
	dst := make([]byte, maxEnLen)
	_, err := hex2.Decode(dst, src)
	check(err)
	return dst
}
