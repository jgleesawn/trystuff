package main

import (
	"fmt"
	"encoding/json"
	"bytes"
)

type Struct struct{
	Field0	[]int
	Field1	int
	Field2	[]int
}

func main() {
	s := Struct{make([]int,3),3,make([]int,3)}
	a := []string{"test","again","3rd"}
	s.Field0[0] = 3

	var buffer []byte
	buf := bytes.NewBuffer(buffer)

	Enc := json.NewEncoder(buf)
	Enc.Encode(s)
	Enc.Encode(a)
	a[0] = "change"
	
	s = Struct{make([]int,3),3,make([]int,3)}
	Dec := json.NewDecoder(buf)
	err := Dec.Decode(&s)
	if err != nil {
		panic(err)
	}
	err = Dec.Decode(&a)
	if err != nil {
		panic(err)
	}
	
	fmt.Println(a)
	fmt.Println(buf.Bytes())
	fmt.Println(s)
	fmt.Println(buf.Bytes())
}
