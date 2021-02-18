package main

import (
	"golang-verse/code" // Import the path of the folder containing the code that want to be run 
	"fmt"
)

// To run the code on folder "Code", write the parameter needed here
func main(){
	fmt.Println(code.Utama("1 10")) // "Utama" on "...code.Utama" refers to palindrome.go function
}