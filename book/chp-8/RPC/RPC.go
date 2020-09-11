package main

import (
	"fmt"
	"net"
	"net/rpc"
	"sort"
)

type Server struct{}

func serv() {
	rpc.Register(new(Server))
	lstn, listenEr := net.Listen("tcp", ":1234")
	if listenEr != nil {
		fmt.Println("ERROR:", listenEr)
		return
	}
	for {
		conn, connErr := lstn.Accept()
		if connErr != nil {
			continue
		}
		go rpc.ServeConn(conn)
	}
}

func main() {
	serv()
}

func (srv *Server) Alphabetize(input string, output *string) error {
	*output = string(
		sortRuneSlice(
			stringToChars(
				input)))
	return nil
}

func sortRuneSlice(runes []rune) []rune {
	sort.Slice(runes, func(i, j int) bool {
		return runes[i] < runes[j]
	})
	return runes
}

func stringToChars(stringToSplit string) (charArray []rune) {
	charArray = make([]rune, len(stringToSplit))
	for i, char := range stringToSplit {
		charArray[i] = char
	}
	return
}
