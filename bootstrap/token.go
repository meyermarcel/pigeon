package bootstrap

import (
	"fmt"

	"github.com/mna/pigeon/ast"
)

type tid int

const (
	invalid tid = iota - 1
	eof         // end-of-file token, id 0

	ident   tid = iota + 127 // identifiers follow the same rules as Go
	ruledef                  // rule definition token

	// literals
	char      // character literal, as in Go ('a'i?)
	str       // double-quoted string literal, as in Go ("string"i?)
	rstr      // back-tick quoted raw string literal, as in Go (`string`i?)
	class     // square-brackets character classes ([a\n\t]i?)
	lcomment  // line comment as in Go (// comment or /* comment */ with no newline)
	mlcomment // multi-line comment as in Go (/* comment */)
	code      // code blocks between '{' and '}'

	// operators and delimiters have the value of their char
	// smallest value in that category is 10, for '\n'
	eol         tid = '\n' // end-of-line token, required in the parser
	colon       tid = ':'  // separate variable name from expression ':'
	semicolon   tid = ';'  // optional ';' to terminate rules
	lparen      tid = '('  // parenthesis to group expressions '('
	rparen      tid = ')'  // ')'
	dot         tid = '.'  // any matcher '.'
	ampersand   tid = '&'  // and-predicate '&'
	exclamation tid = '!'  // not-predicate '!'
	question    tid = '?'  // zero-or-one '?'
	plus        tid = '+'  // one-or-more '+'
	star        tid = '*'  // zero-or-more '*'
	slash       tid = '/'  // ordered choice '/'
)

var lookup = map[tid]string{
	invalid:     "invalid",
	eof:         "eof",
	ident:       "ident",
	ruledef:     "ruledef",
	char:        "char",
	str:         "str",
	rstr:        "rstr",
	class:       "class",
	lcomment:    "lcomment",
	mlcomment:   "mlcomment",
	code:        "code",
	eol:         "eol",
	colon:       "colon",
	semicolon:   "semicolon",
	lparen:      "lparen",
	rparen:      "rparen",
	dot:         "dot",
	ampersand:   "ampersand",
	exclamation: "exclamation",
	question:    "question",
	plus:        "plus",
	star:        "star",
	slash:       "slash",
}

func (t tid) String() string {
	if s, ok := lookup[t]; ok {
		return s
	}
	return fmt.Sprintf("tid(%d)", t)
}

var blacklistedIdents = map[string]struct{}{
	// Go keywords http://golang.org/ref/spec#Keywords
	"break":       {},
	"case":        {},
	"chan":        {},
	"const":       {},
	"continue":    {},
	"default":     {},
	"defer":       {},
	"else":        {},
	"fallthrough": {},
	"for":         {},
	"func":        {},
	"go":          {},
	"goto":        {},
	"if":          {},
	"import":      {},
	"interface":   {},
	"map":         {},
	"package":     {},
	"range":       {},
	"return":      {},
	"select":      {},
	"struct":      {},
	"switch":      {},
	"type":        {},
	"var":         {},

	// predeclared identifiers http://golang.org/ref/spec#Predeclared_identifiers
	"any":        {},
	"bool":       {},
	"byte":       {},
	"comparable": {},
	"complex64":  {},
	"complex128": {},
	"error":      {},
	"float32":    {},
	"float64":    {},
	"int":        {},
	"int8":       {},
	"int16":      {},
	"int32":      {},
	"int64":      {},
	"rune":       {},
	"string":     {},
	"uint":       {},
	"uint8":      {},
	"uint16":     {},
	"uint32":     {},
	"uint64":     {},
	"uintptr":    {},
	"true":       {},
	"false":      {},
	"iota":       {},
	"nil":        {},
	"append":     {},
	"cap":        {},
	"clear":      {},
	"close":      {},
	"complex":    {},
	"copy":       {},
	"delete":     {},
	"imag":       {},
	"len":        {},
	"make":       {},
	"max":        {},
	"min":        {},
	"new":        {},
	"panic":      {},
	"print":      {},
	"println":    {},
	"real":       {},
	"recover":    {},
}

// Token is a syntactic token generated by the scanner.
type Token struct {
	id  tid
	lit string
	pos ast.Pos
}

var tokenStringLen = 50

func (t Token) String() string {
	v := t.lit
	if len(v) > tokenStringLen {
		v = v[:tokenStringLen/2] + "[...]" + v[len(v)-(tokenStringLen/2):]
	}
	return fmt.Sprintf("%s: %s %q", t.pos, t.id, v)
}
