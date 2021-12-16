package utility

import (
	"testing"
	"strings"
)


// test JoinLines Function
// this function take as a input a slice of string and returns
// a string in which all of the elements of the slice are joined
// with the newline 
func TestJoinLines(t *testing.T)  {
	table := []struct {
		name     string
		input    []string
		expected string
	}{
		{"a|b|c", []string{"a","b","c"}, "a\nb\nc\n"},
		{"hello|how|are|you", []string{"hello","how","are","you"}, "hello\nhow\nare\nyou\n"},
	}
	for _, v := range table {
		
		actual := JoinLines(v.input)
		if actual != v.expected {
			t.Errorf("expected %v, got %v", v.expected, actual)
		}
		
	}
}
// test Readlines Function
// this function reads lines of text from a Reader source 
// and add them to a slice. when it encounters EOF it returns
// the slice of strings which the function read.
func TestReadLines(t *testing.T) {
	table := []struct {
		name     string
		input    string
		expected []string
	}{
		{"ciao|come|va", "ciao\ncome\nva", []string{"ciao", "come", "va"}},
		{"string|to|read|now!", "string\nto\nread\nnow!", []string{"string","to","read","now!"}},
	}
	for _, v := range table {
		
		actual := ReadLines(strings.NewReader(v.input))
		for i:=0; i < len(actual); i++{
			if actual[i] != v.expected[i] || len(actual) != len(v.expected){
				t.Errorf("expected %v, got %v", v.expected, actual)
			}
		}
		// if actual != v.expected {
		// 	
		// }
		
	}
}