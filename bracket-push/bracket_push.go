package brackets

import "errors"

const testVersion = 4

type stack struct {
	nodes []rune
	count int
}

func (s *stack) push(node rune) {
	s.nodes = append(s.nodes, node)
	s.count++
}

func (s *stack) pop() (rune, error) {
	if s.count == 0 {
		return '0', errors.New("Error: cannot pop from stack with no elements")
	}
	node := s.nodes[s.count-1]
	s.nodes = s.nodes[:s.count-1]
	s.count--
	return node, nil
}

var brackets = map[rune]rune{
	'{': '}',
	'[': ']',
	'(': ')',
}

// Bracket returns true if all brackets and braces match
func Bracket(str string) (bool, error) {
	var bracketsStack stack
	for _, r := range str {
		if closingBracket := brackets[r]; closingBracket > 0 {
			bracketsStack.push(r)
		} else {
			bracket, _ := bracketsStack.pop()
			if r != brackets[bracket] {
				return false, nil
			}
		}
	}
	if bracketsStack.count > 0 {
		return false, nil
	}
	return true, nil
}
