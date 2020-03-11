package main

import(
	s "./soup"
	//"fmt"
)

func main() {
	
	//fmt.Println(s.Contains([]int{72,57,42,27},[]int{56,42,28,14,0}))
	//"CHILE","USA","CATAR","JAPON",
	m := s.MakeDic([]string{"PERU","INDIA","CHINA","UK","RUSIA","IRAN","IRAK","USA","CUBA","CHILE","CATAR","JAPON","CHAD","CONGO","FIYI"})
	/*for k,v := range m {
		fmt.Println(k, v)
	}*/
	s.PrintBoard(s.FillBoard(m))
	//fmt.Println(len())
	//fmt.Println(s.Match([]int{1,2,3,4},[]int{4,5,33},"hola","abc"))
}