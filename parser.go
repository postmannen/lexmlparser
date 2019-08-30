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
	"os"
	"strings"
	"unicode"
	"unicode/utf8"

	"github.com/go-acme/lego/log"

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

// tokenChannelbufferSize is set way too big. I believe 30-35 would be ok for this xml.
const tokenChannelbufferSize = 100

// parser will hard the state of the parsing variables.
type parser struct {
	// variablesForMap is a slice and are the collection of the tags used to form a map value
	// in the parsed output (and not a map value in this code here).
	// Essentially it will contains a command variable, and will be used at the end of the
	// code to create the key/values of the map structure in the output.
	variablesForMap []string
	// commandConstants/classConstants is a store for all the constants parsed, and are used in the code
	// to check if a constant with the same name has previosly beeing parsed to avoid
	// duplicated
	commandConstants map[string]bool
	classConstants   map[string]bool
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
	// duplicateClassCh is used to send a signal to printing functions for class variables
	// that there is a duplicate to make the variable unique.
	duplicateClassCh chan bool
	// output is where to redirect the output of the printing.
	output *os.File
}

type goType struct {
	name   string
	length string
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
func newParser(outFh *os.File) *parser {
	return &parser{
		variablesForMap:  []string{},
		commandConstants: map[string]bool{},
		classConstants:   map[string]bool{},
		tagStack:         newTagStack(),
		depth:            0,
		droneTypesToGoTypes: map[string]goType{
			"u8":     goType{name: "uint8", length: "1"},
			"i8":     goType{name: "int8", length: "1"},
			"u16":    goType{name: "uint16", length: "2"},
			"i16":    goType{name: "int16", length: "2"},
			"u32":    goType{name: "uint32", length: "4"},
			"i32":    goType{name: "int32", length: "4"},
			"u64":    goType{name: "uint64", length: "8"},
			"i64":    goType{name: "int64", length: "8"},
			"float":  goType{name: "float32", length: "4"},
			"double": goType{name: "float64", length: "8"},
			"string": goType{name: "string", length: "0"},
			"enum":   goType{name: "uint32", length: "4"},
		},
		duplicateClassCh: make(chan bool, 2),
		output:           outFh,
	}
}

// Start will start the lexml parser. Takes a channel of tokens as it's input.
func Start(tCh chan lexml.Token, outFh *os.File) {
	// Create a new parser
	p := newParser(outFh)

	// Create a buffered reader of the channel. The .Next method will move to
	// the next value from input channel. The buffered reader will let us
	// look at the values that are comming ahead of where we are right now.
	buf := NewBuffer(tokenChannelbufferSize)
	buf.Start(tCh)

	fmt.Fprintln(p.output, "package main")
	fmt.Fprintln(p.output)

	p.printTopDeclarations()

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

	p.printBuiltinFunctions()

	p.printFuncgetLengthOfStringData()

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
				// We need a buffer of all the arguments that belong to that specific cmd,
				// since the use of the arguments will be mixed with the cmd in the
				// generated output text.
				argBuf, err := p.newArgBufferForCmd(buf)
				if err != nil {
					log.Println("error: newArgBufferForCmd: ", err)
				}

				//fmt.Printf("------------ARGBUFFER----------- %+v\n", argBuf)

				p.doTagCommand(tmpBuf1, tmpBuf2, id, argBuf)
			}

		}
	}

}

// doTagProject will do all the parsing of a project tag.
func (p *parser) doTagProject(tmpBuf1 []lexml.Token, tmpBuf2 []lexml.Token, id string) {
	for _, v := range tmpBuf1 {
		if v.TokenType == tokenDescription {
			fmt.Fprintf(p.output, "// %v\n", v.TokenText)
			break
		}
	}

	name := tmpBuf1[2]
	fmt.Fprintf(p.output, "const project%v projectDef = %v\n", name.TokenText, id)
}

// doTagClass will do all the parsing of a class tag.
func (p *parser) doTagClass(tmpBuf1 []lexml.Token, tmpBuf2 []lexml.Token, id string) {
	// Check if there is a tokenDescription tag
	for _, v := range tmpBuf1 {
		if v.TokenType == tokenDescription {
			fmt.Fprintf(p.output, "// %v\n", v.TokenText)
			break
		}
	}

	//--

	// The name of the command const is found at slice pos [2].
	classConstName := tmpBuf1[2]

	// Check if there have been any previous use of the same const.
	// If seen before, add DUPLICATE at the end of const name.
	_, ok := p.classConstants[classConstName.TokenText]
	if ok {
		classConstName.TokenText = classConstName.TokenText + "DUPLICATE"
		// put a true value on the channel so we can signal to the function
		// writing out the variable that the class field name should be
		// postfixed with DUPLICATE
		p.duplicateClassCh <- true
	}

	// Store the const to check for duplicates on later iterations.
	p.classConstants[classConstName.TokenText] = true
	//--
	fmt.Fprintf(p.output, "const class%v classDef = %v\n", classConstName.TokenText, id)

}

// doTagCommand will do all the parsing of a command tag
func (p *parser) doTagCommand(tmpBuf1 []lexml.Token, tmpBuf2 []lexml.Token, id string, argBuf []argument) {
	// TODO: Add parsing of buffer="NON_ACK", and add a field in the command struct
	// that we can check to know when to send an ack or not.

	// -------------------------CREATE COMMENTS------------------------------------------

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
				fmt.Fprintf(p.output, "// %v : ", v.TokenText)
			}
			if v.TokenType == tokenArgumentValue {
				fmt.Fprintf(p.output, "%v, \n", v.TokenText)
			}
		}
	}

	// -------------------------CREATE COMMENTS, END------------------------------------------

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
		cmdConstname.TokenText = cmdConstname.TokenText + "DUPLICATE"
	}

	// Store the const to check for duplicates on later iterations.
	p.commandConstants[cmdConstname.TokenText] = true

	constName := cmdConstname.TokenText
	fmt.Fprintf(p.output, "const cmd%v cmdDef = %v\n", constName, id)
	fmt.Fprintln(p.output)

	// Create the struct type command which will hold the decode methods
	// for the command
	fmt.Fprintf(p.output, "type %v command\n", concatenateSlice(p.tagStack.data))
	fmt.Fprintln(p.output)

	// TODO: -----------Put in the argument checking and parsing here--------------------
	// TODO: Store all the arguments in an argument struct with fields needed, since we need to
	// iterate it here, and we also need to iterate it in the creation of the decode method below.
	//
	// Create a specific struct for a specific command, by adding Arguments to the end of the
	// command name.
	fmt.Fprintf(p.output, "type %v struct {\n", concatenateSlice(p.tagStack.data)+"Arguments")
	for _, v := range argBuf {
		fmt.Fprintf(p.output, "%v %v\n", v.name, v.goType)
	}
	fmt.Fprintln(p.output, "}")
	fmt.Fprintln(p.output)
	// TODO: Write out the name of the arguments and the Go equivalent of the type for the fields.

	// ----------------------------DECODE METHOD--------------------------------------------------
	// Create the decode function for the command type
	fmt.Fprintf(p.output, "func (a %v) decode(b []byte) interface{} {\n", concatenateSlice(p.tagStack.data))
	fmt.Fprintf(p.output, "//TODO: .............\n")
	//txt := `fmt.Printf(".....we are now decoding the payload %v, which is of type %T\n", a, a)`
	//fmt.Println(txt)
	//txt = `fmt.Printf("%+v\n", a)`
	//fmt.Println(txt)

	txt := "arg := " + concatenateSlice(p.tagStack.data) + "Arguments" + "{}"

	//if there is a string argument, add variables needed
	foundStringArg := false
	for _, v := range argBuf {
		if v.goType == "string" {
			foundStringArg = true
		}
	}

	if foundStringArg {
		fmt.Fprintln(p.output, "var stringEnd int")
		fmt.Fprintln(p.output, "var err error")
	}

	fmt.Fprintln(p.output, txt)

	if len(argBuf) != 0 {
		fmt.Fprintln(p.output, "var offset = 0")
		// Print the parsing of the types for the decode method
		for _, v := range argBuf {

			// The parsing of everything except a string is the same. Check if string...
			if v.goType != "string" {
				txt := "binary.Read(bytes.NewReader(b[offset:offset+" + v.length + "]), binary.LittleEndian, &arg." + v.name + ")"
				fmt.Fprintln(p.output, txt)

				// the linter complains for ´arg += 1´, so we add a check and replace it
				// with arg++ if the length == 1.
				if v.length != "1" {
					fmt.Fprintln(p.output, "offset += "+string(v.length))
				} else {
					fmt.Fprintln(p.output, "offset++ ")
				}
			} else if v.goType == "string" {
				fmt.Fprintln(p.output, `
				stringEnd, err = getLengthOfStringData(b[offset:])
				if err != nil {
					log.Println("error: ", err)
				}`)
				fmt.Fprintf(p.output, "arg.%v = string(b[offset:offset+stringEnd])\n", v.name)
				fmt.Fprintln(p.output, "offset += stringEnd")
			}

		}
	} else {
		fmt.Fprintln(p.output, "// No arguments to decode here !!")
	}

	fmt.Fprintln(p.output)
	fmt.Fprintln(p.output, "return arg")
	fmt.Fprintf(p.output, "}\n")

	// ----------------------------DECODE METHOD, END--------------------------------------------------

	// ----------------------------CREATE VAR--------------------------------------------------
	project := p.tagStack.data[0]
	class := p.tagStack.data[1]
	command := p.tagStack.data[2]

	fmt.Fprintln(p.output)
	fmt.Fprintf(p.output, "var %v = %v {\n", lowerFirstCharacter(variableName), concatenateSlice(p.tagStack.data))
	fmt.Fprintf(p.output, "project: project%v,\n", project)
	select {
	case <-p.duplicateClassCh:
		fmt.Fprintf(p.output, "class: class%vDUPLICATE,\n", class)
	default:
		fmt.Fprintf(p.output, "class: class%v,\n", class)
	}
	fmt.Fprintf(p.output, "cmd: cmd%v,\n", command)
	fmt.Fprintf(p.output, "}\n")
	fmt.Fprintln(p.output)

	// store the variable name in a slice so we can use it
	// to create the map[command]decoder map later.
	p.variablesForMap = append(p.variablesForMap, variableName)

	// ----------------------------CREATE VAR, END--------------------------------------------------
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

func (p *parser) printBuiltinFunctions() {
	text := `
	// lenStringData takes a []byte which is the data for the arguments, and returns
	// the position of the 0 terminator for the string.
	// The []byte given as input will start looking from the beginning of the slice,
	// so the input slice should be sliced to start from the offset of the string.
	func lenStringData(b []byte) (int, error) {
		// Figure out the length of the string
		for i := 0; i < cap(b); i++ {
			//fmt.Printf("%+v, of type %T\n", b[i], b[i])

			//fmt.Println("i = ", i)
			if b[i] == 0 {
				//fmt.Println("lengthString = ", i)

				// add 1 to jump to the 0
				return i + 1, nil
			}

		}

		err := fmt.Errorf("no string bytes found, returning 0")
		return 0, err
	}
	`
	fmt.Fprintln(p.output, text)
}

// printTopDeclarations will print things like package ...., func main,
// imports, etc....
func (p *parser) printTopDeclarations() {
	fmt.Fprintln(p.output, "import (")
	fmt.Fprintln(p.output, `	"fmt"`)
	fmt.Fprintln(p.output, `	"bytes"`)
	fmt.Fprintln(p.output, `	"log"`)
	fmt.Fprintln(p.output, `	"encoding/binary"`)
	fmt.Fprintln(p.output, ")")
	fmt.Fprintln(p.output)
	fmt.Fprintln(p.output, "type projectDef uint8 ")
	fmt.Fprintln(p.output, "type classDef uint8")
	fmt.Fprintln(p.output, "type cmdDef uint16")
	fmt.Fprintln(p.output)
	fmt.Fprintln(p.output, "type command struct {")
	fmt.Fprintln(p.output, "	project projectDef")
	fmt.Fprintln(p.output, "	class   classDef")
	fmt.Fprintln(p.output, "	cmd     cmdDef")
	fmt.Fprintln(p.output, "}")
	fmt.Fprintln(p.output)
}

// TODO/NB: In the XML there is a tag value for NON_ACK and HIGH_PRIO.
// With the specification as follows:
/*
	4.4.1 buffer
	The value of this attribute can be either NON_ACK, ACK or HIGH_PRIO, defaulting to ACK if not given. It gives a hint about the destination buffer for the command.
	For the Bebop Drone, the NON_ACK buffers are 10 (c2d) and 127 (d2c),
	the ACK buffers are 11 (c2d) and 126 (d2c), and the HIGH_PRIO buffer is the
	12 (c2d).
	This is only a hint, and the product will decode any ARCommand on any
	ARNetwork buffer, as long as the buffer is not used for ARStream.
*/
// Based on the text above I find no reason to parse the ACK flag, since that need to
// be handled on the protocol level in the drone driver code based on what buffer the
// command was received on, and the info in the XML is for description.

// printMapDeclaration will print the whole map structure which
// maps all the command variables to it's type.
func (p *parser) printMapDeclaration() {
	// Map for storing the different commands for lookup.
	fmt.Fprintln(p.output, "type decoder interface {")
	fmt.Fprintln(p.output, "decode([]byte) interface{}")
	fmt.Fprintln(p.output, "}")
	fmt.Fprintln(p.output)
	fmt.Fprintln(p.output, "var commandMap = map[command]decoder {")

	// Will go through the slice and pick out one variable
	// at a time and create the map value
	for _, v := range p.variablesForMap {
		fmt.Fprintf(p.output, "command(%v) : %v,\n", lowerFirstCharacter(v), lowerFirstCharacter(v))
	}
	fmt.Fprintln(p.output, "}")
	fmt.Fprintln(p.output)
}

func (p *parser) printFuncgetLengthOfStringData() {
	txt := `
	func getLengthOfStringData(b []byte) (int, error) {
		// Figure out the length of the string
		for i := 0; i < cap(b); i++ {
			//fmt.Printf("%+v, of type %T\n", b[i], b[i])
	
			//fmt.Println("i = ", i)
			if b[i] == 0 {
				//fmt.Println("lengthString = ", i)
	
				// add 1 to jump to the 0
				return i + 1, nil
			}
	
		}
	
		err := fmt.Errorf("no string bytes found, returning 0")
		return 0, err
	}
	`
	fmt.Fprintln(p.output, txt)
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

// ---------------------------------------HERE----------------------------------------
type argument struct {
	name    string
	xmlType string
	goType  string
	length  string
}

// newArgBufferForCmd Will create a buffer starting at a cmd startTag, and ending
// at a cmd stopTag, so it will be simpler to parse out the arguments for a specific
// cmd.
func (p *parser) newArgBufferForCmd(buf *Buffer) (argBuffer []argument, err error) {
	//fmt.Println("---buf---", buf)

	foundCMDStartTag := false
	//find the position of start of cmd
	for i, v := range buf.Slice {
		if v.TokenType == tokenStartTag && v.TokenText == "cmd" {
			foundCMDStartTag = true
		}

		if foundCMDStartTag && v.TokenType == tokenEndTag && v.TokenText == "cmd" {
			// When reached the end of the cmd, we are done with all arguments,
			// and can return the argument slice []argument
			return argBuffer, nil
		}

		if buf.Slice[i].TokenText == "arg" && buf.Slice[i+1].TokenText == "name" {
			a := argument{}
			a.name = buf.Slice[i+2].TokenText
			// check if the name is == type, and add an X to not conflict with go's
			// type system.
			if a.name == "type" {
				a.name += "X"
			}

			// check if the name contains underscores, and if it does, remove them.
			underScore := strings.Contains(a.name, "_")
			if underScore {
				s := strings.Split(a.name, "_")
				a.name = concatenateSlice(s)
			}

			typ := buf.Slice[i+4].TokenText

			// lookup, and pick the needed values from the type specification map.
			v, ok := p.droneTypesToGoTypes[typ]
			if ok {
				a.xmlType = typ
				a.goType = v.name
				a.length = v.length
			}

			// TODO: Add the different <enum specifications> as comments, so the user
			// will know what values to enter, or what values where received.

			argBuffer = append(argBuffer, a)

			//fmt.Println("--------------------a.name---------------------------", a)
		}
	}

	if !foundCMDStartTag {
		return nil, fmt.Errorf("no start tags to parse arguments inside found")
	}

	return argBuffer, nil
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

	// Get next series of tokens between a start and stop tag
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

func getLengthOfStringData(b []byte) (int, error) {
	// Figure out the length of the string
	for i := 0; i < cap(b); i++ {
		//fmt.Printf("%+v, of type %T\n", b[i], b[i])

		//fmt.Println("i = ", i)
		if b[i] == 0 {
			//fmt.Println("lengthString = ", i)

			// add 1 to jump to the 0
			return i + 1, nil
		}

	}

	err := fmt.Errorf("no string bytes found, returning 0")
	return 0, err
}
