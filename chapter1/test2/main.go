package main

import "fmt"

var array1 = [3]*int{new(int), new(int), new(int)}

func main() {
	*array1[0] = 1
	*array1[1] = 2
	*array1[2] = 3
	array2 := array1
	*array1[0] = 5
	fmt.Println(*array2[0])
}
