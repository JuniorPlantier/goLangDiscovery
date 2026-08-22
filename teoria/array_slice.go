package main

import "fmt"

func main() {

	// var arr [2]string

	s := make([]string, 1)
	s[0] = "Danielzinho"
	s = append(s, "Arthur")
	s = append(s, "Rafael")

	fmt.Println(s, len(s))
}
