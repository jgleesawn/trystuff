package main

import (
	"math/rand"
	"fmt"
	"crypto/sha512"
	"strconv"
)

func main() {
	msg := make([]byte,2048)
	for i := 0; i<2048; i++ {
		msg[i] = strconv.Itoa(rand.Int()%10)[0]
	}

	out := sha512.Sum512(msg)
	fmt.Println(len(out))
/*
	msghash.Write(msg[0:1024])
	out1 := msghash.checkSum()
	fmt.Println(len(out1))
	msg2 := "Test1"

	msghash.Write([]byte(msg2))
	out2 := msghash.checkSum()
	fmt.Println(out2,len(out2))

	msghash.Write([]byte(msg2))
	out2 = msghash.checkSum()
	fmt.Println(out2,len(out2))
*/

}

