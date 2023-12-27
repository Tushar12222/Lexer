package utils

import (
	"fmt"
	"lexer/models"
	"os"
	"strconv"
)

func Tokenize(currSource *models.Source) {
  for !currSource.IsAtEnd() {
    currSource.Start = currSource.Current
    scanToken(currSource)
  }
  var tokenType models.TokenType = models.EOF
  currSource.ListOftokens = append(currSource.ListOftokens, models.NewToken(tokenType.ToString(), "", "", currSource.Line))
}

func scanToken(currSource *models.Source) {
  var char byte = currSource.Advance()
  switch char {
      case '(': currSource.AddToken(models.LEFT_BRACE, nil)
      case ')': currSource.AddToken(models.RIGHT_BRACE, nil)
      case '{': currSource.AddToken(models.LEFT_PAREN, nil)
      case '}': currSource.AddToken(models.RIGHT_PAREN, nil)
      case ',': currSource.AddToken(models.COMMA, nil)
      case '.': currSource.AddToken(models.DOT, nil)
      case '-': currSource.AddToken(models.MINUS, nil)
      case '+': currSource.AddToken(models.PLUS, nil)
      case ';': currSource.AddToken(models.SEMICOLON, nil)
      case '*': currSource.AddToken(models.STAR, nil)
      case '!': currSource.AddToken(ternary(currSource.Match('='), models.BANG_EQUAL, models.BANG), nil)
      case '=': currSource.AddToken(ternary(currSource.Match('='), models.EQUAL_EQUAL, models.EQUAL), nil)
      case '<': currSource.AddToken(ternary(currSource.Match('='), models.LESS_EQUAL, models.LESS), nil)
      case '>': currSource.AddToken(ternary(currSource.Match('='), models.GREATER_EQUAL, models.GREATER), nil)
      case '/': 
        if currSource.Match('/') {
          for currSource.Peek() != '\n' && !currSource.IsAtEnd() {
            currSource.Advance()
          }
        } else {
            currSource.AddToken(models.SLASH, nil)
        }
      case  ' ', '\r', '\t':
      case '\n':  currSource.Line += 1
      case '"': currSource.String()
      default:
        if isDigit(char) {
          number(currSource)
        } else if (isAlpha(char)) {
            identifier(currSource)    
        } else {
            fmt.Println(fmt.Errorf("Unexpected character %d" , char))
            os.Exit(1)
        }
  }
}


func ternary(condition bool, a models.TokenType, b models.TokenType) models.TokenType {
  if condition {
    return a
  }
  return b
}


func isDigit(char byte) bool {
  return char >= '0' && char <= '9'
}

func number(currSource *models.Source) {
  for isDigit(currSource.Peek()) {
    currSource.Advance()
  }

  if currSource.Peek() == '.' && isDigit(peekNext(currSource)) {
    currSource.Advance()

    for isDigit(currSource.Peek()) {
      currSource.Advance()
    }
  }

  value , err := strconv.ParseFloat(currSource.Source[currSource.Start:currSource.Current], 64)
  if err != nil {
    fmt.Println("Error parsing string to float:" , err)
    os.Exit(1)
  }
  currSource.AddToken(models.NUMBER, value)
}


func peekNext(currSource *models.Source) byte {
  if currSource.Current + 1 >= len(currSource.Source) {
    return '\x00' 
  }

  return currSource.Source[currSource.Current + 1]
}

func isAlpha(char byte) bool {
    return (char >= 'a' && char <= 'z') ||
           (char >= 'A' && char <= 'Z') ||
            char == '_'
}

func identifier(currSource *models.Source) {
  for isAlphaNumeric(currSource.Peek()) {
    currSource.Advance()
  }

  mapping := map[string]models.TokenType{
    "and": models.AND,
    "class": models.CLASS,
    "else": models.ELSE,
    "false": models.FALSE,
    "for": models.FOR,
    "fun": models.FUN,
    "if": models.IF,
    "nil": models.NIL,
    "or": models.OR,
    "print": models.PRINT,
    "return": models.RETURN,
    "super": models.SUPER,
    "this": models.THIS,
    "true": models.TRUE,
    "var": models.VAR,
    "while": models.WHILE,
  }

  var text string = currSource.Source[currSource.Start : currSource.Current]
  tokenType , ok := mapping[text]

  if !ok {
    tokenType = models.IDENTIFIER
  }


  currSource.AddToken(tokenType, nil)
}

func isAlphaNumeric(char byte) bool {
  return isAlpha(char) || isDigit(char)
}
