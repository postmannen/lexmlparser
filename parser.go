/*
	This package takes the tokens produced by the lexml package and creates a Go struct of the parsed values

	tokenStartTag      TokenType = "tokenStartTag"      // <tag> || <
	tokenEndTag        TokenType = "tokenEndTag"        // </tag> || />
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
	//Create a buffered reader of the channel. The .Next method will move to the next value from input channel.
	b := NewBuffer(30)
	b.Start(tCh)

	depth := 0 //used to indicate what level or sub level we are in the struct/tag

	var project string
	var class []string

	for v := range b.ChOut {
		switch v.TokenType {
		case "tokenStartTag":
			if depth == 0 {
				if v.TokenText == "project" {
					project = b.Slice[2].TokenText
					//fmt.Println("found project token, make a project struct", project)
				}
			}
			if depth == 1 {
				if v.TokenText == "class" {
					class = append(class, b.Slice[2].TokenText)
					//fmt.Printf("func (t %v) ", b.Slice[3].TokenText)
				}
			}

			if depth == 2 {
				if v.TokenText == "cmd" {
					fmt.Printf("func (t %v) %v(YYYY,YYYY) {\n", class[len(class)-1], b.Slice[2].TokenText)
				}
			}

			if depth == 3 {
				if v.TokenText == "arg" {
					type arg struct {
						name    string
						argType string
					}
					args := []arg{}
					for i := 1; b.Slice[i].TokenType != "tokenEndTag"; i++ {
						if b.Slice[i].TokenText == "name" {
							a := arg{name: b.Slice[i+1].TokenText, argType: b.Slice[i+3].TokenText}
							args = append(args, a)
							//fmt.Println("argument name = ", b.Slice[i+1].TokenText)
							//fmt.Println("argument type = ", b.Slice[i+3].TokenText)
						}

					}
					if len(args) != 0 {
						fmt.Println(args)
					}
				}
			}

			depth++
		case "tokenEndTag":
			depth--
			if depth == 1 {

			}

		case "tokenArgumentName":
			depth++
			depth--
		case "tokenArgumentValue":
			//if strings.Contains(strings.ToLower(v.TokenText), "state") {
			//	if b.Slice[2].TokenText == "id" {
			//		fmt.Println(v.TokenText)
			//	}
			//}
		case "tokenDescription":
		case "tokenEOF":
		case "tokenJustText":
		}

		b.ReadNext()
	}

	fmt.Println("------------------------------------")
	//create the main type
	fmt.Printf("type %v struct {\n", project)
	{
		for _, v := range class {
			fmt.Printf("\t%v\n", v)
		}
	}
	fmt.Println("}")

	//create the class types, like Piloting, PCMD etc....
	for _, v := range class {
		fmt.Printf("type %v struct {}\n", v)
	}
	fmt.Println()

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
