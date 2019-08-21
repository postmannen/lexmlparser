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
	"unicode"
	"unicode/utf8"

	"github.com/postmannen/lexml"
)

// Define all the toke types.
const tokenStartTag lexml.TokenType = "tokenStartTag"           // <tag> || <
const tokenEndTag lexml.TokenType = "tokenEndTag"               // </tag> || />
const tokenArgumentName lexml.TokenType = "tokenArgumentName"   // name is infront of a = sign
const tokenArgumentValue lexml.TokenType = "tokenArgumentValue" // value is after a = sign
const tokenDescription lexml.TokenType = "tokenDescription"     // Description, just text between tags
const tokenEOF lexml.TokenType = "tokenEOF"                     //End Of File
const tokenJustText lexml.TokenType = "tokenJustText"           //just text, no start or end tag

const tokenChannelbufferSize = 60

// parser will hard the state of the parsing variables.
type parser struct {
	// variablesForMap is a slice and are the collection of the tags used to form a map value
	// in the parsed output (and not a map value in this code here).
	// Essentially it will contains a command variable, and will be used at the end of the
	// code to create the key/values of the map structure in the output.
	variablesForMap []string
	// commandConstants is a store for all the constants parsed, and are used in the code
	// to check if a constant with the same name has previosly beeing parsed to avoid
	// duplicated
	commandConstants map[string]bool
	// tagStack , are a push/pop storage for stack values.
	// The contents of the tag stack is used to create names
	// that consists of several tag names.
	tagStack *tagStack
	// depth is used to indicate what level or sub level we are in the struct/tag
	// to keep track of if we are working on a tag within another tag, and so on.
	// We add 1 to the depth for each tag we find, and subtract by 1 for each
	// closing tag.
	depth int
	// droneTypesToGoTypes is a map used to know how to map the types found in the xml like
	// u8/i8/float etc to they're go equivalent.
	droneTypesToGoTypes map[string]goType
}

type goType struct {
	name   string
	length int
}

/*
u8 1 unsigned 8bit value
i8 1 signed 8bit value
u16 2 unsigned 16bit value
i16 2 signed 16bit value
u32 4 unsigned 32bit value
i32 4 signed 32bit value
u64 8 unsigned 64bit value
i64 8 signed 64bit value
float 4 IEEE-754 single precision
double 8 IEEE-754 double precision
string * Null terminated string (C-String)
(Variable size)
enum 4 Per command defined enum
*/

// newParser will return a new *parser struct that will hold the state of the
// parsing while parsing.
func newParser() *parser {
	return &parser{
		variablesForMap:  []string{},
		commandConstants: map[string]bool{},
		tagStack:         newTagStack(),
		depth:            0,
		droneTypesToGoTypes: map[string]goType{
			"u8":     goType{name: "uint8", length: 1},
			"i8":     goType{name: "int8", length: 1},
			"u16":    goType{name: "uint16", length: 2},
			"i16":    goType{name: "int16", length: 2},
			"u32":    goType{name: "uint32", length: 4},
			"i32":    goType{name: "int32", length: 4},
			"u64":    goType{name: "uint64", length: 8},
			"i64":    goType{name: "int64", length: 8},
			"float":  goType{name: "float32", length: 4},
			"double": goType{name: "float64", length: 8},
			"string": goType{name: "string", length: 0},
			"enum":   goType{name: "uint32", length: 4},
		},
	}
}

// Start will start the lexml parser. Takes a channel of tokens as it's input.
func Start(tCh chan lexml.Token) {
	// Create a new parser
	p := newParser()

	// Create a buffered reader of the channel. The .Next method will move to
	// the next value from input channel. The buffered reader will let us
	// look at the values that are comming ahead of where we are right now.
	buf := NewBuffer(tokenChannelbufferSize)
	buf.Start(tCh)

	fmt.Println("package main")
	fmt.Println()

	printTopDeclarations()

	// Range over the ChOut of buf, where ChOut is an unbuffered channel,
	// and we can pick one value at a time.
	for v := range buf.ChOut {
		switch v.TokenType {

		// Everything we want parse into something else starts with
		// a start tag. If a start tag is found we should check
		// for tags between startTag -> endTag for the type we are
		// looking for.
		//
		// Check all the start tags.
		case tokenStartTag:
			p.doTokenTagStart(buf)
		// Check all the end tags
		case tokenEndTag:
			//*fmt.Println("endTag-------------------------------------------------------", v)
			//*fmt.Printf("depth = %v, endtag found : %v, subtracting to depth.\n", depth, v.TokenText)

			p.depth--
			p.tagStack.pop()
			//*fmt.Println("Depth is now = ", depth)
		}

		// Read the next token from the buffered channel.
		buf.ReadNext()
	}

	p.printMapDeclaration()

}

// doTokenTagStart will do all the parsing of a tagStart.
func (p *parser) doTokenTagStart(buf *Buffer) {
	//*fmt.Println("startTag-------------------------------------------------------", v)
	//*fmt.Printf("depth = %v, startTag found : %v, adding to depth.\n", depth, v.TokenText)
	//
	// Push the name of the tag found on the tag Stack.
	p.depth++
	p.tagStack.push(buf.Slice[2].TokenText)
	//*fmt.Println("Depth is now = ", depth)

	// Get the first 2 sequences of tokens that have a start and stop tag in the buffer.
	tmpBuf1, tmpBuf2 := newPartialBuffer(buf)

	//Remove later, just for showing values
	//fmt.Println()
	//for _, v := range tmpBuf1 {
	//	fmt.Printf("*** tmpBuf1 : %v\n", v)
	//}
	//
	////Remove later, just for showing values
	//fmt.Println()
	//for _, v := range tmpBuf2 {
	//	fmt.Printf("*** tmpBuf2 : %v\n", v)
	//}

	// Range the buffer for this specific token
	for i, v := range tmpBuf1 {
		//*fmt.Printf("tmpBuf : %+v\n", v)
		// If there is an id value we will know that it is a project/class/cmd tag.
		if v.TokenText == "id" {
			id := tmpBuf1[i+1].TokenText

			//Check if it is either project, class or cmd tag.
			switch tmpBuf1[0].TokenText {

			case "project":
				p.doTagProject(tmpBuf1, tmpBuf2, id)
			case "class":
				p.doTagClass(tmpBuf1, tmpBuf2, id)
			case "cmd":
				p.doTagCommand(tmpBuf1, tmpBuf2, id)
			}

		}
	}

}

// doTagProject will do all the parsing of a project tag.
func (p *parser) doTagProject(tmpBuf1 []lexml.Token, tmpBuf2 []lexml.Token, id string) {
	for _, v := range tmpBuf1 {
		if v.TokenType == tokenDescription {
			fmt.Printf("// %v\n", v.TokenText)
			break
		}
	}

	name := tmpBuf1[2]
	fmt.Printf("	const %v projectDef = %v\n", lowerFirstCharacter(name.TokenText), id)
}

// doTagClass will do all the parsing of a class tag.
func (p *parser) doTagClass(tmpBuf1 []lexml.Token, tmpBuf2 []lexml.Token, id string) {
	// Check if there is a tokenDescription tag
	for _, v := range tmpBuf1 {
		if v.TokenType == tokenDescription {
			fmt.Printf("// %v\n", v.TokenText)
			break
		}
	}

	name := tmpBuf1[2]
	fmt.Printf("const %v classDef = %v\n", lowerFirstCharacter(name.TokenText), id)

}

// doTagCommand will do all the parsing of a command tag
func (p *parser) doTagCommand(tmpBuf1 []lexml.Token, tmpBuf2 []lexml.Token, id string) {
	// TODO : Implement detection of duplicate commands !!!

	// Check if there are comments to be printed for the command.
	//
	// The startToken..if found, is located in the 0'th position of the buffer.
	if tmpBuf2[0].TokenType == tokenStartTag && tmpBuf2[0].TokenText == "comment" {
		// Create the comments above the const declaration.
		for i, v := range tmpBuf2 {
			// We do not want the first value since it is a start tag.
			if i == 0 {
				continue
			}
			if v.TokenType == tokenArgumentName {
				fmt.Printf("// %v : ", v.TokenText)
			}
			if v.TokenType == tokenArgumentValue {
				fmt.Printf("%v, \n", v.TokenText)
			}
		}
	}

	// Create the variable name of the current project->class->command
	// content in the tagStack.
	var variableName string
	for i, v := range p.tagStack.data {
		// We do not want the first value naming the project
		// in the variableName value, only class+command.
		if i != 0 {
			variableName += v
		}
	}

	// The name of the command const is found at slice pos [2].
	cmdConstname := tmpBuf1[2]

	// Check if there have been any previous use of the same const.
	// If seen before, add DUPLICATE at the end of const name.
	_, ok := p.commandConstants[cmdConstname.TokenText]
	if ok {
		cmdConstname.TokenText += "DUPLICATE"
	}

	// Store the const to check for duplicates on later iterations.
	p.commandConstants[cmdConstname.TokenText] = true

	constName := lowerFirstCharacter(cmdConstname.TokenText)
	fmt.Printf("const %v cmdDef = %v\n", constName, id)
	fmt.Println()

	// Create the struct type command which will hold the decode methods
	// for the command
	fmt.Printf("type %v command\n", concatenateSlice(p.tagStack.data))
	fmt.Println()

	// Create the decode function for the command type
	fmt.Printf("func (a %v) decode() {\n", concatenateSlice(p.tagStack.data))
	fmt.Printf("//TODO: .............\n")
	txt := `fmt.Printf(".....we are now decoding the payload %v, which is of type %T\n", a, a)`
	fmt.Println(txt)
	txt = `fmt.Printf("%+v\n", a)`
	fmt.Println(txt)
	fmt.Printf("}\n")

	project := lowerFirstCharacter(p.tagStack.data[0])
	class := lowerFirstCharacter(p.tagStack.data[1])
	command := lowerFirstCharacter(p.tagStack.data[2])

	fmt.Println()
	fmt.Printf("var %v = %v {\n", lowerFirstCharacter(variableName), concatenateSlice(p.tagStack.data))
	fmt.Printf("project: %v,\n", project)
	fmt.Printf("class: %v,\n", class)
	fmt.Printf("cmd: %v,\n", command)
	fmt.Printf("}\n")
	fmt.Println()

	// store the variable name in a slice so we can use it
	// to create the map[command]decoder map later.
	p.variablesForMap = append(p.variablesForMap, variableName)
}

// lowerFirstCharacer, turns the first character of a string
// to lowercase.
func lowerFirstCharacter(s string) string {
	if s == "" {
		return ""
	}
	r, n := utf8.DecodeRuneInString(s)
	return string(unicode.ToLower(r)) + s[n:]
}

// printTopDeclarations will print things like package ...., func main,
// imports, etc....
func printTopDeclarations() {
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
}

// printMapDeclaration will print the whole map structure which
// maps all the command variables to it's type.
func (p *parser) printMapDeclaration() {
	// Map for storing the different commands for lookup.
	fmt.Println("type decoder interface {")
	fmt.Println("decode()")
	fmt.Println("}")
	fmt.Println()
	fmt.Println("var commandMap = map[command]decoder {")

	// Will go through the slice and pick out one variable
	// at a time and create the map value
	for _, v := range p.variablesForMap {
		fmt.Printf("command(%v) : %v,\n", lowerFirstCharacter(v), lowerFirstCharacter(v))
	}
	fmt.Println("}")
	fmt.Println()
}

// concatenateSlice will take all the string elements of
// a slice, and return them as a single string.
func concatenateSlice(s []string) string {
	var output string
	for _, v := range s {
		output += v
	}

	return output
}

// newPartialBuffer takes a *lexml.Buffer as input, and returns the first two
// portions of that buffer forming a start -> stop token sequence.
func newPartialBuffer(buf *Buffer) (firstBuffer []lexml.Token, secondBuffer []lexml.Token) {
	endTagPosition1 := 0
	endTagPosition2 := 0
	buf1 := buf.Slice

	for i, v := range buf1 {
		// If it is the first position in slice, just continue with the next iteration.
		if i == 0 {
			continue
		}

		// Check for a start or end tag. We also need to check if there are a
		// start tag after the first start tag at position 0 in the buffer,
		if v.TokenType == tokenEndTag || v.TokenType == tokenStartTag {
			endTagPosition1 = i
			break
		}
	}

	buf2 := buf1[endTagPosition1:]

	// Get next series of tokens betweem a start and stop tag
	for i, v := range buf2 {
		// If it is the first position in slice, just continue with the next iteration.
		if i == 0 {
			continue
		}

		// Check for a start or end tag. We also need to check if there are a
		// start tag after the first start tag at position 0 in the buffer,
		if v.TokenType == tokenEndTag || v.TokenType == tokenStartTag {
			endTagPosition2 = i
			break
		}
	}

	return buf1[:endTagPosition1], buf2[:endTagPosition2]
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
	//*fmt.Println("DEBUG: Put on stack : ", s)
}

// pop will remove the last element of the stack
func (s *tagStack) pop() {
	//*fmt.Println("DEBUG: Before pop:", s)
	last := len(s.data)
	// ---
	if len(s.data) == 0 {
		return
	}
	s.data = append(s.data[0:0], s.data[:last-1]...)
	//*fmt.Println("DEBUG: After pop:", s)

}
