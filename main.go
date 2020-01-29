package main

import (
	"fmt"
	s "strings"
)

func main() {
	fmt.Println("Hello, playground")
	var wd ="abca bbcd dddd"
	ww := s.Split(wd, " ")
	fmt.Println(ww)
	
	
	//for i:=0, wdd := range ww{
	for wdd := range ww{
		fmt.Println(ww[wdd])
		fix(ww[wdd])
	}
}

func fix( w string){
	fmt.Println("===",w)
	
	//temp = make( map[string] int )
	
	//s := string([]byte{65, 66, 67, 226, 130, 172})
	//s := string([]byte{65, 66, 67, 226, 130, 172})
	
	for ww, wd := range w {
		fmt.Println("================")
		fmt.Println( ww)
		fmt.Println(  w[ww]  )
		fmt.Println("+++", wd )
		fmt.Println( s.Count(wd, ww)   )
		//fmt.Printf( "%x ", w[ww])
		//fmt.Println("i = ", []byte( s.Index(w[ww])) , w)
		//fmt.Println("i = ", string( s.Index(w[ww])) , w)	
	}
	
}



//abcde
//abccde
//abcde

//abcd  

//abca bbcd dddd
//abc bcd d
