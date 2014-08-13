package main

import (
	"crypto/rand"
	"flag"
	"fmt"

	"code.google.com/p/go.crypto/bcrypt"
)

func main() {
	algorithm := flag.String("a", "bcrypt", "Algorithm to use for encrypting the password")
	cost := flag.Int("c", 10, "The cost for the algorithm")
	pwlen := flag.Int("l", 8, "The length of the generated password")
	flag.Parse()

	pw := make([]byte, *pwlen)
	rand.Read(pw)

	var err error
	var hpw []byte

	switch *algorithm {
	case "bcrypt":
		hpw, err = bcrypt.GenerateFromPassword(pw, *cost)
		if err != nil {
			fmt.Println("error hashing password:", err)
		}
	}

	fmt.Println(string(pw), string(hpw))
}
