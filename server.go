package main

import(
	"fmt"
	"net"
	"os"
	"./soup"
	"strings"
	"time"
	"io/ioutil"
	"strconv"
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

	if op[0] == 0 {
		conn.Close()
		fmt.Println("Connection finished!")
	}else if op[0] == 1 {
		conn.Read(typ)
		conn.Read(diff)
		conn.Read(name)
		fmt.Println("read:",op,typ,diff,string(name))
		makeGame(conn,typ[0],diff[0],string(name))
	}else if op[0] == 2 {
		watchScores(conn)
	}
}

func makeGame(conn net.Conn, typ, diff byte, name string) {
	m := soup.MakeRandomMap(typ)
	board := soup.FillBoard(m)
	soup.PrintBoard(board)
	sendWords(conn, m)
	sendBoard(conn,board)
	fmt.Println(len(board))
	play(conn,typ,diff,m,name)
}

func sendWords(conn net.Conn, words map[string][]int){
	for w,_ := range words {
		aux := fmt.Sprintf("%15s",w)
		conn.Write([]byte(aux))
	}
}

func sendBoard(conn net.Conn, board []int) {
	data := make([]byte,225)
	for i,v := range board {
		data[i] = byte(v)
	}
	conn.Write(data)
}

func play(conn net.Conn, typ, diff byte, m map[string][]int, name string) {
	found := 0
	var x1,x2 string
	buffx1 := make([]byte,3)
	buffx2 := make([]byte,3)

	start := time.Now()
	elapsed := time.Since(start)

	for found < 15 {
		conn.Read(buffx1)
		conn.Read(buffx2)

		x1 = string(buffx1)
		x2 = string(buffx2)

		x1 = strings.Replace(x1," ","",-1)
		x2 = strings.Replace(x2," ","",-1)

		fmt.Println(x1,x2)

		f,w := soup.IsThere(x1,x2,m)
		if f {
			conn.Write([]byte{1})
			found++
			m[w] = []int{0,0,0}
			w = fmt.Sprintf("%15s",w)
			conn.Write([]byte(w))
		}else{
			conn.Write([]byte{0})
		}
	}

	addScore(fmt.Sprintf("%s",start),fmt.Sprintf("%s",elapsed),name)
}

func addScore(s,e, name string) {
	f,_ := os.OpenFile("Scores.txt",os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	defer f.Close()
	f.WriteString(strings.Replace(name," ","",-1)+" "+s+" took "+e+"\n")
}

func watchScores(conn net.Conn) {
	data,_ := ioutil.ReadFile("Scores.txt")
	l1 := strconv.Itoa(len(data))
	l2 := byte(len(l1))
	fmt.Println(l1,l2,string(data))
	conn.Write([]byte{l2})
	conn.Write([]byte(l1))
	conn.Write(data)
}