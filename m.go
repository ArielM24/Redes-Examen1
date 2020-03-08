package main

import(
	s "./soup"
	//"fmt"
)

func main() {
	
	//fmt.Println(s.Contains([]int{1,2,3,4,5},[]int{9,6,7,8}))
	m := s.MakeDic([]string{"MEXICO","USA","CANADA","PORTUGAL","BRAZIL","INDIA","CHINA","JAPON","UK","PERU","CHILE","ARGENTINA","ECUADOR","COLOMBIA","RUSIA"})
	/*for k,v := range m {
		fmt.Println(k, v)
	}*/
	s.PrintBoard(s.FillBoard(m))
}