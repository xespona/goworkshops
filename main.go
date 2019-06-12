package main

import (
	"fmt"
	"github.com/xespona/goworkshops/crypto"
)

func main() {
	a := "FOR A MOMENT, NOTHING HAPPENED. THEN, AFTER A SECOND OR SO, NOTHING CONTINUED TO HAPPEN"
	s := crypto.Encrypt(a)

	fmt.Println(s)
	fmt.Println(crypto.Decrypt(s))
}
