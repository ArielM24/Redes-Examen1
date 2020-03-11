package main

import(
	s "./soup"
	"fmt"
)

func main() {
	m := s.MakeDic([]string{"PERU","INDIA","CHINA","UK","RUSIA","IRAN",
		"IRAK","USA","CUBA","CHILE","CATAR","JAPON","CHAD","CONGO","FIYI"})
	m1 := s.MakeDic([]string{"c++","c#","python","assembler","java","ruby","go",
		"javascript","basic","fortran","rust","pascal","perl","swift","dart"})
	m2 := s.MakeDic([]string{"sabritas","doritos","chips","cheetos","totis","taquis","oyuki",
		"churrumais","rancheritos","crujitos","fritos","ruffles","pringles","paketaxo","runners"})
	s.PrintBoard(s.FillBoard(m))

	var x1,x2 string
	/*for {
		fmt.Printf("x1:\t")
		fmt.Scanf("%s",&x1)
		fmt.Printf("x2:\t")
		fmt.Scanf("%s",&x2)
		fmt.Println(s.IsThere(x1,x2,m))
	}*/
	s.PrintBoard(s.FillBoard(m1))
	s.PrintBoard(s.FillBoard(m2))
	for {
		fmt.Printf("x1:\t")
		fmt.Scanf("%s",&x1)
		fmt.Printf("x2:\t")
		fmt.Scanf("%s",&x2)
		fmt.Println(s.IsThere(x1,x2,m2))
	}
}