package models

import (
	"fmt"
	"os"
)

type Source struct {
	ListOftokens []Token
	Source       string
	Start        int
	Current      int
	Line         int
}

func NewSource(source string, start int, current int, line int) Source {
	return Source{
		[]Token{},
		source,
		start,
		current,
		line,
	}
}

func (source *Source) AddToken(tokenType TokenType, literal interface{}) {
	var text string = source.Source[source.Start:source.Current]
	var token Token = NewToken(tokenType.ToString(), literal, text, source.Line) 
	source.ListOftokens = append(source.ListOftokens, token)
}

func (currSource *Source) IsAtEnd() bool {
	return currSource.Current >= len([]rune(currSource.Source))
}

func (currSource *Source) Advance() byte {
	var char byte = currSource.Source[currSource.Current]
	currSource.Current += 1
	return char
}

func (currSource *Source) Match(expected byte) bool {
	if currSource.IsAtEnd() {
		return false
	}

	if currSource.Source[currSource.Current] != expected {
		return false
	}

	currSource.Current += 1

	return true
}

func (currSource *Source) Peek() byte {
	if currSource.IsAtEnd() {
		return '\x00'
	}
	return currSource.Source[currSource.Current]
}

func (currSource *Source) String() {
	for currSource.Peek() != '"' && !currSource.IsAtEnd() {
		if currSource.Peek() == '\n' {
			currSource.Line += 1
		}
		currSource.Advance()
	}

	if currSource.IsAtEnd() {
    fmt.Println("Unterminated String at line: ", currSource.Line)
    os.Exit(1)
	}

  currSource.Advance()

  var text string = currSource.Source[currSource.Start + 1 : currSource.Current - 1]

  currSource.AddToken(STRING , text)
}
