package main

import(
	"fmt"
	"net"
	"./soup"
	"strings"
)

func main() {
	var diff,typ,op byte
	var name string
	conn, errc := net.Dial("tcp",":2000")
	if errc != nil {
		panic(errc)
	}

	defer conn.Close()

	fmt.Println("Selec an option")
	fmt.Println("0 -> exit")
	fmt.Println("1 -> play")
	fmt.Println("2 -> records")
	fmt.Scanf("%d", &op)

	if op > 0 {
		fmt.Println("Selec a modality")
		fmt.Println("0 -> topic")
		fmt.Println("1 -> anagram")
		fmt.Scanf("%d", &typ)

		fmt.Println("Selec a difficulty")
		fmt.Println("0 -> easy")
		fmt.Println("1 -> normal")
		fmt.Println("2 -> hard")
		fmt.Scanf("%d", &diff)

		fmt.Println("Enter your name")
		fmt.Scanf("%s",&name)
	}else{
		fmt.Println("Goodbye!")
	}
	writeOptions(conn, op, typ, diff, name)
	
}

func writeOptions(conn net.Conn, op, typ, diff byte, name string) {
	conn.Write([]byte{op})
	conn.Write([]byte{typ})
	conn.Write([]byte{diff})
	usr_name := fmt.Sprintf("%15s",name)
	conn.Write([]byte(usr_name))
	if op > 0 {
		words := readWords(conn,diff,typ)
		board := readBoard(conn)
		printWords(words,diff,typ,0)
		soup.PrintBoard(board)
		play(conn,words,board,typ,diff)
	}
}

func readWords(conn net.Conn, diff,typ byte) []string{
	words := make([]string,15)
	for i := 0; i < 15; i++ {
		aux := make([]byte,15)
		conn.Read(aux)
		w := strings.Replace(string(aux)," ","",-1)
		words[i] = w
	}
	return words
}

func printWords(words []string,typ, diff,err byte) {
	fmt.Println("Words:")
	for i := 0; i < 15; i++ {
		if diff == 0 {
			if typ == 0 {
				fmt.Println(i,":",words[i])
			}else if typ == 1{
				fmt.Println(i,":",soup.Anagram(words[i]))
			}
		}
		if diff == 2 {
			fmt.Println(i,":",len(words[i]))
		}
	}
}

func readBoard(conn net.Conn) []int{
	data := make([]byte,225)
	conn.Read(data)
	board := make([]int,225)
	for i,v := range data {
		board[i] = int(v)
	}
	return board
}

func play(conn net.Conn,words []string,board []int, diff, typ byte){
	found := 0
	var err byte = 0
	var x1, x2, w string
	fbuff := make([]byte,1)
	wbuff := make([]byte,15)
	for found < 15 {
		fmt.Print("x1:\t")
		fmt.Scanf("%s",&x1)

		fmt.Print("x2:\t")
		fmt.Scanf("%s",&x2)
		x1 = fmt.Sprintf("%3s",x1)
		x2 = fmt.Sprintf("%3s",x2)

		conn.Write([]byte(x1))
		conn.Write([]byte(x2))

		conn.Read(fbuff)
		if fbuff[0] == 1 {
			fmt.Print("Encontrada!  ")
			conn.Read(wbuff)
			w = strings.Replace(string(wbuff)," ","",-1)
			fmt.Println(w)
			found++
		}else{
			fmt.Println("No")
			err++
		}
		printWords(words,diff,typ,err)
		soup.PrintBoard(board)
	}
}