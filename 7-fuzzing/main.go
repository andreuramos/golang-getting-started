package main

import (
	"errors"
	"fmt"
	"unicode/utf8"
)

func main() {
	input := "The quick brown fox jumped over the lazy dog"
	rev, revErr := Reverse(input)
	doubleRev, doubleRevErr := Reverse(rev)

	fmt.Printf("original: %q\n", input)
	fmt.Printf("reverse: %q, err: %q\n", rev, revErr)
	fmt.Printf("reeversed again: %q, err: %q\n", doubleRev, doubleRevErr)
}

func Reverse(s string) (string, error) {
	if !utf8.ValidString(s){
		return s, errors.New("input is not valid UTF-8")
	}
    r := []rune(s)
    for i, j := 0, len(r)-1; i < len(r)/2; i, j = i+1, j-1 {
        r[i], r[j] = r[j], r[i]
	}
    return string(r), nil
}