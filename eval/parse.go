package eval

import (
   "fmt"
   "strconv"
   "strings"
   "text/scanner"
)

type lexer struct {
   scan scanner.Scanner
   token rune // current lookahead token
}

func (lex *lexer) next() { lex.token = lex.scan.Scan() }
func (lex *lexer) text() string { return lex.scan.TokenText() }

type lexPanic string

// describe returns a string describing the current token, for use in errors.
func (lex *lexer) describe() string {
   switch lex.token {
   case scanner.EOF:
      return "end of file"
   case scanner.Ident:
      return fmt.Sprintf("identifier %s", lex.text())
   case scanner.Int, scanner.Float:
      return fmt.Sprintf("number %s", lex.text())
   }
   return fmt.Sprintf("%q", rune(lex.token))
}

func precedence(op rune) int {
	switch op {
	case '*', '/':
		return 2
	case '+', '-':
		return 1
	}
	return 0
}

func Parse(input string) (_ Expr, err error) {
	defer func() {
		switch x := recover().(type) {
		case nil:
			// no panic
		case lexPanic:
			err = fmt.Errorf("%s", x)
		default:
			// unexpected panic: resume state of panic.
			panic(x)
		}
	}()
	lex := new(lexer)
	lex.scan.Init(strings.NewReader(input))
	lex.scan.Mode = scanner.ScanIdents | scanner.ScanInts | scanner.ScanFloats
	lex.next() // initial lookahead
	e := parseExpr(lex)
	if lex.token != scanner.EOF {
		return nil, fmt.Errorf("unexpected %s", lex.describe())
	}
	return e, nil
}

 
