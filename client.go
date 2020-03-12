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
		readWords(conn,diff,typ)
		readBoard(conn)
	}
}

func readWords(conn net.Conn, diff,typ byte) []string{
	fmt.Println("Words:")
	words := make([]string,15)
	for i := 0; i < 15; i++ {
		aux := make([]byte,15)
		conn.Read(aux)
		w := strings.Replace(string(aux)," ","",-1)
		if diff == 0 {
			if typ == 0 {
				fmt.Println(i,":",w)
			}else if typ == 1{
				fmt.Println(i,":",soup.Anagram(w))
			}
		}
		if diff == 2 {
			fmt.Println(i,":",len(w))
		}
		words[i] = w
	}
	return words
}

func readBoard(conn net.Conn) {

}