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

//Start will start the lexml parser

func Start(tCh chan lexml.Token) {
	depth := 0             //used to indicate what level or sub level we are in the struct/tag
	literal := []string{}  //used for a single literal
	literals := []string{} //used for storing all the literals

	for v := range tCh {
		//fmt.Println("* depth = ", depth)
		switch v.TokenType {
		case "tokenStartTag":
			if depth == 0 {
				//fmt.Printf("*DEPTH=%v* type %v struct {\n", depth, v.TokenText)
				fmt.Printf("type %v struct {\n", v.TokenText)
				literals = append(literals, fmt.Sprintf("var %v %v", v.TokenText, v.TokenText))
				literal = append(literal, fmt.Sprintf("%v", v.TokenText))
			}
			if depth > 0 {
				//fmt.Printf("*DEPTH=%v* %v %v struct {\n", depth, depthMap[depth], v.TokenText)
				fmt.Printf("%v %v struct {\n", depthMap[depth], v.TokenText)
				literal = append(literal, v.TokenText)
			}
			depth++
		case "tokenEndTag":
			depth--
			//fmt.Printf("*DEPTH=%v* %v }\n", depth, depthMap[depth])
			fmt.Printf("%v }\n", depthMap[depth])
			literal = literal[:depth]
		case "tokenArgumentFound":
		case "tokenArgumentName":
			depth++
			//fmt.Printf("*DEPTH=%v* %v %v", depth, depthMap[depth], v.TokenText)
			fmt.Printf("%v %v", depthMap[depth], v.TokenText)
			literal = append(literal, v.TokenText)
			depth--
		case "tokenArgumentValue":
			fmt.Printf(" %v\n", "string")
			literal = append(literal, v.TokenText)
			//fmt.Println(literal)
			//fmt.Println("literal = ", literal)
			literals = append(literals, createLiteral(literal))
			literal = literal[:depth]

		case "tokenDescription":
			//fmt.Println(v.TokenText)
		case "tokenEOF":
		case "tokenJustText":
		}
	}
	for _, lit := range literals {
		fmt.Println(lit)
	}
}

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
