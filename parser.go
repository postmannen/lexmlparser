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

// Start will start the lexml parser. Takes a channel of tokens as it's input.
func Start(tCh chan lexml.Token) {
	// Create a buffered reader of the channel. The .Next method will move to
	// the next value from input channel. The buffered reader will let us
	// look at the values that are comming ahead of where we are right now.
	buf := NewBuffer(30)
	buf.Start(tCh)

	tagStack := newTagStack()

	// Depth is used to indicate what level or sub level we are in the struct/tag
	// to keep track of if we are working on a tag within another tag, and so on.
	// We add 1 to the depth for each tag we find, and subtract by 1 for each
	// closing tag.
	depth := 0

	fmt.Println("package main")
	fmt.Println()

	fmt.Println("import (")
	fmt.Println(`	"fmt"`)
	fmt.Println(")")
	fmt.Println()
	fmt.Println("type projectDef uint8 ")
	fmt.Println("type classDef uint8")
	fmt.Println("type cmdDef uint16")
	fmt.Println()
	fmt.Println("type command struct {")
	fmt.Println("	project projectDef")
	fmt.Println("	class   classDef")
	fmt.Println("	cmd     cmdDef")
	fmt.Println("}")
	fmt.Println()

	// Range over the ChOut of buf, where ChOut is an unbuffered channel,
	// and we can pick one value at a time.
	for v := range buf.ChOut {
		switch v.TokenType {
		// Check all the start tags.
		case "tokenStartTag":
			fmt.Println("startTag-------------------------------------------------------", v)
			fmt.Printf("depth = %v, startTag found : %v, adding to depth.\n", depth, v.TokenText)
			depth++
			tagStack.push(buf.Slice[2].TokenText)
			fmt.Println("Depth is now = ", depth)

		case "tokenEndTag":
			fmt.Println("endTag-------------------------------------------------------", v)
			fmt.Printf("depth = %v, endtag found : %v, subtracting to depth.\n", depth, v.TokenText)

			depth--
			tagStack.pop()
			fmt.Println("Depth is now = ", depth)
		case "tokenArgumentName":
		case "tokenArgumentValue":
		case "tokenDescription":
		case "tokenEOF":
		case "tokenJustText":
		}

		// Read the next token from the buffered channel.
		buf.ReadNext()
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

// tagStack will keep track of where we are working in the iteration,
type tagStack struct {
	data []string
}

// newTagStack is a push/pop storage for tags.
func newTagStack() *tagStack {
	return &tagStack{
		//data: make([]string, 0, 100),
	}
}

// push will add another item to the end of the stack with a normal append
func (s *tagStack) push(d string) {
	s.data = append(s.data, d)
	fmt.Println("DEBUG: Put on stack : ", s)
}

// pop will remove the last element of the stack
func (s *tagStack) pop() {
	fmt.Println("DEBUG: Before pop:", s)
	last := len(s.data)
	// ---
	if len(s.data) == 0 {
		return
	}
	s.data = append(s.data[0:0], s.data[:last-1]...)
	fmt.Println("DEBUG: After pop:", s)

}
