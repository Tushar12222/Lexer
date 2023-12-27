package models

import "strconv"

type Token struct {
  Type string
  Literal interface{}
  Lexeme string
  Line int
}

func NewToken(tokenType string,literal interface{}, lexeme string, line int) Token {
  return Token{
    tokenType,
    literal,
    lexeme,
    line,
  }
}

func (token *Token) ToString() string {
  switch value := token.Literal.(type) {
    case string:
      return "Type: " + token.Type + ", Literal: " + value + ", Lexeme: " + token.Lexeme + ", Line: " + strconv.Itoa(token.Line)
    
    case float64:
      return "Type: " + token.Type + ", Literal: " + strconv.FormatFloat(value, 'f', -1, 64) + ", Lexeme: " + token.Lexeme + ", Line: " + strconv.Itoa(token.Line)
    
    default: 
      return "Type: " + token.Type + ", Literal: " + "" + ", Lexeme: " + token.Lexeme + ", Line: " + strconv.Itoa(token.Line)
  }
}
