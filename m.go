package main

import(
	s "./soup"
	"fmt"
)

func main() {
	m := s.MakeDic([]string{"PERU","INDIA","CHINA","UK","RUSIA","IRAN",
		"IRAK","USA","CUBA","CHILE","CATAR","JAPON","CHAD","CONGO","FIYI"})
	s.PrintBoard(s.FillBoard(m))
	var x1,x2 string
	for {
		fmt.Printf("x1:\t")
		fmt.Scanf("%s",&x1)
		fmt.Printf("x2:\t")
		fmt.Scanf("%s",&x2)
		fmt.Println(s.IsThere(x1,x2,m))
	}
}