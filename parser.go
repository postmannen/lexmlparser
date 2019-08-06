/*
Package lexmlparser ,
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

var depthMap = map[int]string{
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

// Start will start the lexml parser. Takes a channel of tokens as it's input.
func Start(tCh chan lexml.Token) {
	// Create a buffered reader of the channel. The .Next method will move to
	// the next value from input channel. The buffered reader will let us
	// look at the values that are comming ahead of where we are right now.
	buf := NewBuffer(30)
	buf.Start(tCh)

	depth := 0 //used to indicate what level or sub level we are in the struct/tag

	var project string
	var class []string

	for v := range buf.ChOut {
		switch v.TokenType {
		case "tokenStartTag":
			//project tag
			if depth == 0 {
				if v.TokenText == "project" {
					project = buf.Slice[2].TokenText
					//fmt.Println("found project token, make a project struct", project)
				}
			}

			//class tag
			if depth == 1 {
				if v.TokenText == "class" {
					class = append(class, buf.Slice[2].TokenText)
					//fmt.Printf("func (t %v) ", b.Slice[3].TokenText)
				}
			}

			//cmd tag
			if depth == 2 {
				if v.TokenText == "cmd" {
					fmt.Printf("\nfunc (t %v) %v", class[len(class)-1], buf.Slice[2].TokenText)
				}
			}

			//arg tag
			if depth == 3 {
				if v.TokenText == "arg" {
					type arg struct {
						name    string
						argType string
					}
					args := []arg{}
					for i := 1; buf.Slice[i].TokenType != "tokenEndTag"; i++ {
						if buf.Slice[i].TokenText == "name" {
							a := arg{name: buf.Slice[i+1].TokenText, argType: buf.Slice[i+3].TokenText}
							args = append(args, a)
						}

					}
					//print all the args, TODO: move this one to tokenEndTag
					if len(args) != 0 {
						fmt.Print("(")
						for _, a := range args {
							fmt.Printf("%v %v\n", a.name, a.argType)
						}
					}
				}
			}

			depth++
		case "tokenEndTag":
			depth--
		case "tokenArgumentName":
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

		buf.ReadNext()
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
