package main

import (
	"fmt"
	"lang.yottadb.com/go/yottadb"
)
type A struct {
	S string
}
func main() {

	err := yottadb.SetValE(yottadb.NOTTP, nil, "something", "TEST", []string{"0001"})
	if err != nil {
		panic(err)
	}

	// uncomment this line for make it work as expect
	// yottadb.Exit()

	defer func() {
		if r := recover(); r != nil {
			fmt.Println("panic recovery!")
			fmt.Println(r)
		}
	}()
	var a *A
	fmt.Println(a.S)
}
