/*
This main.go file will typically be the parser, but here we just print out the tokens that is received.
*/
package main

import (
	"flag"
	"log"
	"os"

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
	outFileName := flag.String("outFile", "", "file name to write to")

	flag.Parse()

	// Open the file to read from.
	inFh, err := os.Open(*inFileName)
	if err != nil {
		log.Fatal("Error: opening file: ", err)
	}

	var outFh *os.File

	if *writeMode == "file" {
		if *outFileName == "" {
			log.Println("error: You have to specify a filename with the -outFile parameter")
			os.Exit(1)
		}

		outFh, err = os.Create(*outFileName)
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
