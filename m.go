package main

import(
	s "./soup"
	//"fmt"
)

func main() {
	
	//fmt.Println(s.Contains([]int{1,2,3,4,5},[]int{9,6,7,8}))
	//"CHILE","USA","CATAR","JAPON",
	m := s.MakeDic([]string{"PERU","INDIA","CHINA","UK","RUSIA","IRAN","IRAK","USA","CUBA","CHILE"})
	/*for k,v := range m {
		fmt.Println(k, v)
	}*/
	s.PrintBoard(s.FillBoard(m))
	//fmt.Println(len())
	//fmt.Println(s.Match([]int{1,2,3,4},[]int{4,5,33},"hola","abc"))
}