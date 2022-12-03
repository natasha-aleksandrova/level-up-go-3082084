package main

import (
	"flag"
	"log"
)

type operatorType int

const (
	openBracket operatorType = iota
	closedBracket
	otherOperator
)

type stack struct {
	items []rune
}

func (s *stack) push(item rune) {
	s.items = append(s.items, item)
}

func (s *stack) pop() *rune {
	if len(s.items) == 0 {
		return nil
	}
	lastIdx := len(s.items) - 1
	lastItem := s.items[lastIdx]
	s.items = s.items[:lastIdx]
	return &lastItem
}

// bracketPairs is the map legal bracket pairs
var bracketPairs = map[rune]rune{
	'(': ')',
	'[': ']',
	'{': '}',
}

func getOperatorType(op rune) operatorType {
	for opening, closing := range bracketPairs {
		switch op {
		case opening:
			return openBracket
		case closing:
			return closedBracket
		}
	}
	return otherOperator
}
func isOpeningBracket(ch rune) bool {
	for _, v := range bracketPairs {
		if ch == v {
			return true
		}
	}
	return false
}

func isClosingBracket(ch rune) bool {
	for k := range bracketPairs {
		if ch == k {
			return true
		}
	}
	return false
}

// isBalanced returns whether the given expression
// has balanced brackets.
func isBalanced(expr string) bool {
	s := stack{}
	for _, ch := range expr {
		switch getOperatorType(ch) {
		case openBracket:
			s.push(ch)
		case closedBracket:
			last := s.pop()
			if last == nil || bracketPairs[*last] != ch {
				return false
			}
		}
	}
	return len(s.items) == 0
}

// printResult prints whether the expression is balanced.
func printResult(expr string, balanced bool) {
	if balanced {
		log.Printf("%s is balanced.\n", expr)
		return
	}
	log.Printf("%s is not balanced.\n", expr)
}

func main() {
	expr := flag.String("expr", "", "The expression to validate brackets on.")
	flag.Parse()
	printResult(*expr, isBalanced(*expr))
}
