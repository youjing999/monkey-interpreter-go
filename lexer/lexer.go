package lexer

import "monkey-interpreter-go/token"

/**
词法分析器除了查看当前字符,还需要进一步“查看”字符串,即查看字符串中的下一个字符。
readPosition 始终指向所输入字符串中的“下一个”字符, position 则指向所输入字符串中与ch 字节对应的字符。
*/

type Lexer struct {
	input        string
	position     int  // 所输入字符串中的当前位置(指向当前字符)
	readPosition int  // 所输入字符串中的当前读取位置(指向当前字符之后的一个字符)
	ch           byte // 当前正在查看的字符
}

func New(input string) *Lexer {
	l := &Lexer{input: input}
	l.readChar()
	return l
}

// readChar 的目的是读取 input 中的下一个字符,并前移其在 input 中的位置。
func (l *Lexer) readChar() {
	if l.readPosition >= len(l.input) {
		l.ch = 0
	} else {
		l.ch = l.input[l.readPosition]
	}
	l.position = l.readPosition
	l.readPosition += 1
}

func (l *Lexer) NextToken() token.Token {
	var tok token.Token
	switch l.ch {
	case '=':
		tok = newToken(token.ASSIGN, l.ch)
	case ';':
		tok = newToken(token.SEMICOLON, l.ch)
	case '(':
		tok = newToken(token.LPAREN, l.ch)
	case ')':
		tok = newToken(token.RPAREN, l.ch)
	case ',':
		tok = newToken(token.COMMA, l.ch)
	case '+':
		tok = newToken(token.PLUS, l.ch)
	case '{':
		tok = newToken(token.LBRACE, l.ch)
	case '}':
		tok = newToken(token.RBRACE, l.ch)
	case 0:
		tok.Literal = ""
		tok.Type = token.EOF
	}
	l.readChar()
	return tok
}

func newToken(tokenType token.TokenType, ch byte) token.Token {
	return token.Token{Type: tokenType, Literal: string(ch)}
}
