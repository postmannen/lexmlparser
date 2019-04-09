/*
	This package takes the tokens produced by the lexml package and creates a Go struct of the parsed values

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
	"strings"

	"github.com/postmannen/lexml"
)

var depthMap map[int]string = map[int]string{
	0: "",
	1: "	",
	2: "		",
	3: "			",
	4: "				",
	5: "					",
	6: "						",
	7: "							",
	8: "								",
	9: "									",
}

//Start will start the lexml parser. Takes a channel of tokens as it's input.
func Start(tCh chan lexml.Token) {
	b := NewBuffer(10)
	b.Start(tCh)

	depth := 0 //used to indicate what level or sub level we are in the struct/tag

	for v := range b.ChOut {
		switch v.TokenType {
		case "tokenStartTag":

			if depth == 0 {
			}
			if depth > 0 {
			}
			depth++
		case "tokenEndTag":
			depth--
		case "tokenArgumentFound":
		case "tokenArgumentName":
			depth++
			depth--
		case "tokenArgumentValue":
			if strings.Contains(strings.ToLower(v.TokenText), "state") {
				if b.Slice[2].TokenText == "id" {
					fmt.Println(v.TokenText)
				}
			}
		case "tokenDescription":
		case "tokenEOF":
		case "tokenJustText":
		}

		b.ReadNext()
	}
}

//createLiterals will create a literal string. Takes []string and return a string to the caller.
func createLiteral(s []string) string {
	var tmpString string

	for i := 0; i < len(s); i++ {
		if i < len(s)-2 {
			v := strings.TrimSpace(s[i])
			tmpString = fmt.Sprintf("%v%v.", tmpString, v)
		}
		if i == len(s)-2 {
			v := strings.TrimSpace(s[i])
			tmpString = fmt.Sprintf("%v%v", tmpString, v)
		}
		if i == len(s)-1 {
			v := strings.TrimSpace(s[i])
			tmpString = fmt.Sprintf(`%v="%v"`, tmpString, v)
		}
	}
	return tmpString

}
