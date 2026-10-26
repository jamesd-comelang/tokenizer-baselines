package main

import (
	"fmt"
	"os"
	"strings"
)

func tkByte(b []byte, i int) int {
	if i < 0 || i >= len(b) {
		return 0
	}
	return int(b[i])
}

func tkIsSpace(c int) bool { return c == 32 || (c >= 9 && c <= 13) }
func tkIsDigit(c int) bool { return c >= '0' && c <= '9' }
func tkIsXdigit(c int) bool {
	if tkIsDigit(c) {
		return true
	}
	if c >= 'a' && c <= 'f' {
		return true
	}
	return c >= 'A' && c <= 'F'
}
func tkIsIdchar(c int) bool {
	if c >= 0x80 {
		return true
	}
	if tkIsDigit(c) {
		return true
	}
	if c >= 'A' && c <= 'Z' {
		return true
	}
	if c >= 'a' && c <= 'z' {
		return true
	}
	return c == '_' || c == '$'
}

func tkClass(c int) int {
	if c == 0 {
		return CC_NUL
	}
	if c == 0xEF {
		return CC_BOM
	}
	if c >= 0x80 {
		return CC_ID
	}
	if c == 'x' || c == 'X' {
		return CC_X
	}
	if (c >= 'A' && c <= 'W') || (c >= 'a' && c <= 'w') {
		return CC_KYWD0
	}
	if (c >= 'Y' && c <= 'Z') || (c >= 'y' && c <= 'z') || c == '_' {
		return CC_KYWD
	}
	if tkIsDigit(c) {
		return CC_DIGIT
	}
	if c == '$' {
		return CC_DOLLAR
	}
	if c == '@' || c == '#' || c == ':' {
		return CC_VARALPHA
	}
	if c == '?' {
		return CC_VARNUM
	}
	if tkIsSpace(c) {
		if c == 11 {
			return CC_ILLEGAL
		}
		return CC_SPACE
	}
	if c == '"' || c == 39 || c == '`' {
		return CC_QUOTE
	}
	if c == '[' {
		return CC_QUOTE2
	}
	if c == '|' {
		return CC_PIPE
	}
	if c == '-' {
		return CC_MINUS
	}
	if c == '<' {
		return CC_LT
	}
	if c == '>' {
		return CC_GT
	}
	if c == '=' {
		return CC_EQ
	}
	if c == '!' {
		return CC_BANG
	}
	if c == '/' {
		return CC_SLASH
	}
	if c == '(' {
		return CC_LP
	}
	if c == ')' {
		return CC_RP
	}
	if c == ';' {
		return CC_SEMI
	}
	if c == '+' {
		return CC_PLUS
	}
	if c == '*' {
		return CC_STAR
	}
	if c == '%' {
		return CC_PERCENT
	}
	if c == ',' {
		return CC_COMMA
	}
	if c == '&' {
		return CC_AND
	}
	if c == '~' {
		return CC_TILDA
	}
	if c == '.' {
		return CC_DOT
	}
	return CC_ILLEGAL
}

func tkIdLen(b []byte, pos int) int {
	j := pos + 1
	for tkIsIdchar(tkByte(b, j)) {
		j++
	}
	return j - pos
}

func tkScanDigits(b []byte, j, t int) (int, int) {
	for {
		c := tkByte(b, j)
		if tkIsDigit(c) {
			j++
		} else if c == '_' {
			t = TK_QNUMBER
			j++
		} else {
			break
		}
	}
	return j, t
}

func tkNumber(b []byte, pos int) (int, int) {
	t := TK_INTEGER
	var j int
	if tkByte(b, pos) == '0' && (tkByte(b, pos+1) == 'x' || tkByte(b, pos+1) == 'X') && tkIsXdigit(tkByte(b, pos+2)) {
		j = pos + 3
		for {
			c := tkByte(b, j)
			if tkIsXdigit(c) {
				j++
			} else if c == '_' {
				t = TK_QNUMBER
				j++
			} else {
				break
			}
		}
	} else {
		j, t = tkScanDigits(b, pos, t)
		if tkByte(b, j) == '.' {
			if t == TK_INTEGER {
				t = TK_FLOAT
			}
			j++
			j, t = tkScanDigits(b, j, t)
		}
		e := tkByte(b, j)
		if (e == 'e' || e == 'E') && (tkIsDigit(tkByte(b, j+1)) || ((tkByte(b, j+1) == '+' || tkByte(b, j+1) == '-') && tkIsDigit(tkByte(b, j+2)))) {
			if t == TK_INTEGER {
				t = TK_FLOAT
			}
			j += 2
			j, t = tkScanDigits(b, j, t)
		}
	}
	for tkIsIdchar(tkByte(b, j)) {
		t = TK_ILLEGAL
		j++
	}
	return t, j - pos
}

func tkVariable(b []byte, pos int) (int, int) {
	t := TK_VARIABLE
	cnt := 0
	j := pos + 1
	done := false
	for !done {
		c := tkByte(b, j)
		if c == 0 {
			done = true
		} else if tkIsIdchar(c) {
			cnt++
			j++
		} else if c == '(' && cnt > 0 {
			j++
			for {
				d := tkByte(b, j)
				if d == 0 || tkIsSpace(d) || d == ')' {
					break
				}
				j++
			}
			if tkByte(b, j) == ')' {
				j++
			} else {
				t = TK_ILLEGAL
			}
			done = true
		} else if c == ':' && tkByte(b, j+1) == ':' {
			j += 2
		} else {
			done = true
		}
	}
	if cnt == 0 {
		t = TK_ILLEGAL
	}
	return t, j - pos
}

func tkScan(b []byte, pos int) (int, int) {
	c0 := tkByte(b, pos)
	cls := tkClass(c0)
	if cls == CC_SPACE {
		j := pos + 1
		for tkIsSpace(tkByte(b, j)) {
			j++
		}
		return TK_SPACE, j - pos
	}
	if cls == CC_BOM {
		if tkByte(b, pos+1) == 0xBB && tkByte(b, pos+2) == 0xBF {
			return TK_SPACE, 3
		}
		return TK_ID, tkIdLen(b, pos)
	}
	if cls == CC_NUL {
		return TK_ILLEGAL, 0
	}
	if cls == CC_MINUS {
		if tkByte(b, pos+1) == '-' {
			j := pos + 2
			for tkByte(b, j) != '\n' && tkByte(b, j) != 0 {
				j++
			}
			return TK_COMMENT, j - pos
		}
		if tkByte(b, pos+1) == '>' {
			n := 2
			if tkByte(b, pos+2) == '>' {
				n = 3
			}
			return TK_PTR, n
		}
		return TK_MINUS, 1
	}
	if cls == CC_LP {
		return TK_LP, 1
	}
	if cls == CC_RP {
		return TK_RP, 1
	}
	if cls == CC_SEMI {
		return TK_SEMI, 1
	}
	if cls == CC_PLUS {
		return TK_PLUS, 1
	}
	if cls == CC_STAR {
		return TK_STAR, 1
	}
	if cls == CC_PERCENT {
		return TK_REM, 1
	}
	if cls == CC_COMMA {
		return TK_COMMA, 1
	}
	if cls == CC_AND {
		return TK_BITAND, 1
	}
	if cls == CC_TILDA {
		return TK_BITNOT, 1
	}
	if cls == CC_SLASH {
		if tkByte(b, pos+1) != '*' || tkByte(b, pos+2) == 0 {
			return TK_SLASH, 1
		}
		j := pos + 2
		for tkByte(b, j) != 0 {
			if tkByte(b, j) == '*' && tkByte(b, j+1) == '/' {
				return TK_COMMENT, j + 2 - pos
			}
			j++
		}
		return TK_COMMENT, j - pos
	}
	if cls == CC_EQ {
		if tkByte(b, pos+1) == '=' {
			return TK_EQ, 2
		}
		return TK_EQ, 1
	}
	if cls == CC_LT {
		c1 := tkByte(b, pos + 1)
		if c1 == '=' {
			return TK_LE, 2
		}
		if c1 == '>' {
			return TK_NE, 2
		}
		if c1 == '<' {
			return TK_LSHIFT, 2
		}
		return TK_LT, 1
	}
	if cls == CC_GT {
		c1 := tkByte(b, pos + 1)
		if c1 == '=' {
			return TK_GE, 2
		}
		if c1 == '>' {
			return TK_RSHIFT, 2
		}
		return TK_GT, 1
	}
	if cls == CC_BANG {
		if tkByte(b, pos+1) != '=' {
			return TK_ILLEGAL, 1
		}
		return TK_NE, 2
	}
	if cls == CC_PIPE {
		if tkByte(b, pos+1) != '|' {
			return TK_BITOR, 1
		}
		return TK_CONCAT, 2
	}
	if cls == CC_QUOTE {
		delim := c0
		j := pos + 1
		closed := false
		for tkByte(b, j) != 0 {
			if tkByte(b, j) == delim {
				if tkByte(b, j+1) == delim {
					j += 2
				} else {
					closed = true
					break
				}
			} else {
				j++
			}
		}
		if closed {
			if delim == 39 {
				return TK_STRING, j + 1 - pos
			}
			return TK_ID, j + 1 - pos
		}
		return TK_ILLEGAL, j - pos
	}
	if cls == CC_QUOTE2 {
		j := pos + 1
		for tkByte(b, j) != 0 && tkByte(b, j) != ']' {
			j++
		}
		if tkByte(b, j) == ']' {
			return TK_ID, j + 1 - pos
		}
		return TK_ILLEGAL, j - pos
	}
	if cls == CC_DOT {
		if !tkIsDigit(tkByte(b, pos + 1)) {
			return TK_DOT, 1
		}
		return tkNumber(b, pos)
	}
	if cls == CC_DIGIT {
		return tkNumber(b, pos)
	}
	if cls == CC_VARNUM {
		j := pos + 1
		for tkIsDigit(tkByte(b, j)) {
			j++
		}
		return TK_VARIABLE, j - pos
	}
	if cls == CC_DOLLAR || cls == CC_VARALPHA {
		return tkVariable(b, pos)
	}
	if cls == CC_KYWD0 {
		if tkClass(tkByte(b, pos+1)) > CC_KYWD {
			return TK_ID, tkIdLen(b, pos)
		}
		j := pos + 2
		for tkClass(tkByte(b, j)) <= CC_KYWD && tkByte(b, j) != 0 {
			j++
		}
		if tkIsIdchar(tkByte(b, j)) {
			return TK_ID, tkIdLen(b, pos)
		}
		lw := strings.ToLower(string(b[pos:j]))
		if v, ok := kwMap[lw]; ok {
			return v, j - pos
		}
		return TK_ID, j - pos
	}
	if cls == CC_X {
		if tkByte(b, pos+1) == 39 {
			j := pos + 2
			for tkIsXdigit(tkByte(b, j)) {
				j++
			}
			bad := false
			if tkByte(b, j) != 39 || (j-pos)%2 == 1 {
				bad = true
			}
			if bad {
				for tkByte(b, j) != 0 && tkByte(b, j) != 39 {
					j++
				}
			}
			if tkByte(b, j) != 0 {
				j++
			}
			if bad {
				return TK_ILLEGAL, j - pos
			}
			return TK_BLOB, j - pos
		}
		return TK_ID, tkIdLen(b, pos)
	}
	if cls == CC_KYWD || cls == CC_ID {
		return TK_ID, tkIdLen(b, pos)
	}
	return TK_ILLEGAL, 1
}

func main() {
	data, err := os.ReadFile("workload.sql")
	if err != nil {
		fmt.Println("cannot open workload.sql")
		os.Exit(1)
	}
	lines := strings.Split(string(data), "\n")
	if len(lines) > 0 && lines[len(lines)-1] == "" {
		lines = lines[:len(lines)-1]
	}
	nlines := len(lines)
	var count, tsum, lsum int64
	for pass := 0; pass < 20; pass++ {
		for _, line := range lines {
			b := []byte(line)
			pos := 0
			for pos < len(b) {
				t, n := tkScan(b, pos)
				if n == 0 {
					break
				}
				count++
				tsum += int64(t)
				lsum += int64(n)
				pos += n
			}
		}
	}
	fmt.Printf("Go: %d tokens, typesum=%d lensum=%d (20 passes, %d lines)\n", count, tsum, lsum, nlines)
}
