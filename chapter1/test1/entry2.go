package test1

import "fmt"

func CallPrint(s string){
	fmt.Println("call print string from entry2's public function")
	printString(s)
}