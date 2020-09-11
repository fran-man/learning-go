package main

import (
	"fmt"
	"net/rpc"
)

func clnt() {
	conn, err := rpc.Dial("tcp", "localhost:1234")
	if err != nil {
		fmt.Println("ERROR:", err)
		return
	}
	var result string
	callErr := conn.Call("Server.Alphabetize", "This is a totally scrambled string!!", &result)
	if callErr != nil {
		fmt.Println("ERROR:", callErr)
		return
	}

	fmt.Println("Received alphabetized string:", result)
}

func main() {
	clnt()
}
