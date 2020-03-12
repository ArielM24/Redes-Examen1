package main

import(
	"fmt"
	"net"
	"os"
	"./soup"
)

func main() {
	server, errs := net.Listen("tcp",":2000")

	if errs != nil {
		fmt.Println("Error listening: ",errs)
		os.Exit(1)
	}
	fmt.Println("Server started! Waiting for connections...")

	for {
		connection, errc := server.Accept()
		if errc != nil {
			fmt.Println("Error: ",errc)
			os.Exit(1)
		}
		fmt.Println("Client connected")
		readOptions(connection)
		defer connection.Close()
	}
	fmt.Println("Connection finished!")
}

func readOptions(conn net.Conn) {
	op := make([]byte,1)
	typ := make([]byte,1)
	diff := make([]byte,1)
	name := make([]byte,15)

	conn.Read(op)
	conn.Read(typ)
	conn.Read(diff)
	conn.Read(name)

	fmt.Println("read:",op,typ,diff,string(name))
	if op[0] == 0 {
		conn.Close()
		fmt.Println("Connection finished!")
	}else if op[0] == 1 {
		makeGame(conn,typ[0],diff[0],string(name))
	}
}

func makeGame(conn net.Conn, typ, diff byte, name string) {
	m := soup.MakeRandomMap(typ)
	board := soup.FillBoard(m)
	sendWords(conn, m)
	fmt.Println(len(board))
}

func sendWords(conn net.Conn, words map[string][]int){
	for w,_ := range words {
		aux := fmt.Sprintf("%15s",w)
		conn.Write([]byte(aux))
	}
}