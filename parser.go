/*
	tokenStartTag      TokenType = "tokenStartTag"      // <tag> || <
	tokenEndTag        TokenType = "tokenEndTag"        // </tag> || />
	tokenArgumentFound TokenType = "tokenArgumentFound" // =
	tokenArgumentName  TokenType = "tokenArgumentName"  // name is infront of a = sign
	tokenArgumentValue TokenType = "tokenArgumentValue" // value is after a = sign
	tokenDescription   TokenType = "tokenDescription"   // Description, just text between tags
	tokenEOF           TokenType = "tokenEOF"           //End Of File
	tokenJustText      TokenType = "tokenJustText"      //just text, no start or end tag
*/
package lexmlparser

import (
	"fmt"

	"github.com/postmannen/lexml"
)

//Start will start the lexml parser
func Start(tCh chan lexml.Token) {
	for v := range tCh {
		switch v.TokenType {
		case "tokenStartTag     ":
			fmt.Println("*readToken from channel * ", v.TokenType, ", tokenText = ", v.TokenText)
		case "tokenEndTag       ":
			fmt.Println("*readToken from channel * ", v.TokenType, ", tokenText = ", v.TokenText)
		case "tokenArgumentFound":
			fmt.Println("*readToken from channel * ", v.TokenType, ", tokenText = ", v.TokenText)
		case "tokenArgumentName ":
			fmt.Println("*readToken from channel * ", v.TokenType, ", tokenText = ", v.TokenText)
		case "tokenArgumentValue":
			fmt.Println("*readToken from channel * ", v.TokenType, ", tokenText = ", v.TokenText)
		case "tokenDescription  ":
			fmt.Println("*readToken from channel * ", v.TokenType, ", tokenText = ", v.TokenText)
		case "tokenEOF          ":
			fmt.Println("*readToken from channel * ", v.TokenType, ", tokenText = ", v.TokenText)
		case "tokenJustText     ":
			fmt.Println("*readToken from channel * ", v.TokenType, ", tokenText = ", v.TokenText)
		}
	}
}
