package main

import "fmt"

func main() {
	a := "FOR A MOMENT, NOTHING HAPPENED. THEN, AFTER A SECOND OR SO, NOTHING CONTINUED TO HAPPEN"
	s := Encrypt(a)

	fmt.Println(s)
	fmt.Println(Decrypt(s))
}
