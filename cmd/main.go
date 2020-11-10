/*
This main.go file will typically be the parser, but here we just print out the tokens that is received.
*/
package main

import (
	"flag"
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/postmannen/lexmlparser"

	"github.com/postmannen/lexml"
)

func main() {
	//defer profile.Start().Stop()

	// Check arguments given at start.
	a := os.Args
	if len(a) < 2 {
		log.Fatal("Specify an xml file\n")

	}

	inFileName := flag.String("inFile", "", "file name to read from")
	writeMode := flag.String("writeMode", "stdout", "stdout/file")

	flag.Parse()

	fn := filepath.Base(*inFileName)
	fileName := strings.Split(fn, ".")
	log.Printf("fileName = %v\n", fileName)

	// Open the file to read from.
	inFh, err := os.Open(*inFileName)
	if err != nil {
		log.Fatal("Error: opening file: ", err)
	}

	var outFh *os.File

	if *writeMode == "file" {
		outFh, err = os.Create(fileName[0] + ".go")
		if err != nil {
			log.Println("error: failed to open file for writing: ", err)
		}

		defer outFh.Close()

	} else {
		outFh = os.Stdout
	}

	// Start the lexer which will lex trough the xml file given
	// as an input argument,
	// and return tokens of what is being lexed back on a channel.
	tCh := lexml.LexStart(inFh)

	// Start the parser, and give it the token channel from
	// the lexer as it's input.
	lexmlparser.Start(tCh, outFh)
}
