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

	fileName := flag.String("fileName", "", "specify the filename to check")

	flag.Parse()

	// Open the file to read from.
	fh, err := os.Open(*fileName)
	if err != nil {
		log.Fatal("Error: opening file: ", err)
	}

	// Start the lexer which will lex trough the xml file given
	// as an input argument,
	// and return tokens of what is being lexed back on a channel.
	tCh := lexml.LexStart(fh)

	// Start the parser, and give it the token channel from
	// the lexer as it's input.
	lexmlparser.Start(tCh)

}
