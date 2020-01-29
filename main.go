package main

import (
	"fmt"
	s "strings"
)

func main() {
	var wd = "abca bbcd dddd"
	ww := s.Split(wd, " ")
	fmt.Println(ww)

	for wdd := range ww {
		fmt.Println("===============")
		fmt.Println(ww[wdd], "=== ", Fix(ww[wdd]))
		fmt.Println("===============")
	}
}

func Fix(w string) string {
	var ts string

	for _, v := range w {
		if !s.Contains(ts, string(v)) {
			ts += string(v)
		}

	}
	return ts
}

//abcde
//abccde
//abcde

//abcd

//abca bbcd dddd
//abc bcd d
