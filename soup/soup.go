package soup

import(
	"fmt"
	"math/rand"
    "time"
)


//0 -> normal
//1 -> inv
//2 -> abj
//3 -> arr
//4 -> diag
//5 -> diag inv
//6 -> diag rev
//7 -> diag inv rev
const size = 15

func MakeDic(words []string) map[string][]int{
	Rd := rand.New(rand.NewSource(time.Now().UnixNano()))
	m := make(map[string][]int)
	force := 0
	for _,w := range words {
		start := Rd.Intn(size*size)
		dir := Rd.Intn(8)
		end := 0
		//fmt.Println(w,dir)
		pos := getPos(dir,start,end,w)
		for !verifyPosition(getNumbers(pos,w),m) {
			//fmt.Println(w,"pos1:",pos,getNumbers(pos,w))
			start = Rd.Intn(size*size)
			dir = Rd.Intn(8)
			end = 0
			pos = getPos(dir,start,end,w)
			//fmt.Println(w,"pos2:",pos,getNumbers(pos,w))
			//fmt.Scanf("%d",&dir)
			force++
		}
		m[w] = pos
		//fmt.Println(w,m[w],getNumbers(m[w],w))
		force++
		if force > 75 {
			fmt.Println("rec")
			return MakeDic(words)
		}
	}
	fmt.Println("force",force)
	return m
}
func getPos(dir, start,end int, w string) []int {
	//fmt.Println(dir)
	var pos []int
	switch dir {
		case 0:
			end = start + len(w) - 1
			pos = normal(start, end)
		break
		case 1:
			end = start - len(w) + 1
			pos = inv(start, end)
		break
		case 2:
			end = start + (len(w))*size
			pos = abj(start, end)
		break
		case 3:
			end = start - (len(w))*size
			pos = arr(start, end)
		break
		case 4:
			end = start - (len(w)-1)*(size - 1)
			pos = diag(start, end, len(w))
		break
		case 5:
			end = start + (len(w)-1)*(size - 1)
			pos = diagI(start, end, len(w))
		break
		case 6:
			end = start - (len(w)-1)*(size + 1)
			pos = diagR(start, end, len(w))
		break
		case 7:
			end = start + (len(w)-1)*(size + 1)
			pos = diagRI(start, end,len(w))
		break
	}
	return pos
}

func verifyPosition(data []int, words map[string][]int) bool{
	for w,pos := range words {
		v := Contains(data,getNumbers(pos,w))
		//fmt.Println("w:",w,"v:",v)
		if v {
			return false
		}
	}
	return true
}

func Contains(a,b []int) bool {
	//fmt.Println("a:",a,"b:",b)
	for _,v1 := range a {
		for _,v2 := range b {
			if v1 == v2 {
				return true
			}
		}
	}
	return false
}

func Match(a,b []int, w1,w2 string) bool {
	match := false
	for _,v1 := range a {
		for _,v2 := range b {
			if v1 == v2 {
				if !match {
					if w1[v1] == w2[v2] {
						match = true
					}else{
						return false
					}
				}else{ 
					return false 
				}
			}
		}
	}
	return false
}

func getNumbers(pos []int, w string) []int {
	sz := len(w)
	numbers := make([]int,0)

	switch pos[2] {
		case 0:
			for i := pos[0]; i <= pos[1]; i++{
				numbers = append(numbers,i)
			}
		break
		case 1:
			for i := pos[1]; i <= pos[0]; i++{
				numbers = append(numbers,i)
			}
		break
		case 2:
			j := pos[0]
			for i := 0; i < sz; i++{
				numbers = append(numbers,j)
				j += size
			}
		break
		case 3:
			j := pos[0]
			for i := 0; i < sz; i++{
				numbers = append(numbers,j)
				j -= size
			}
		break
		case 4:
			j := pos[0]
			for i := 0; i < sz; i++{
				numbers = append(numbers,j)
				j -= (size - 1)
			}
		break
		case 5:
			j := pos[0]
			for i := 0; i < sz; i++{
				numbers = append(numbers,j)
				j += (size - 1)
			}
		break
		case 6:
			j := pos[0]
			for i := 0; i < sz; i++{
				numbers = append(numbers,j)
				j -= (size + 1)
			}
		break
		case 7:
			j := pos[0]
			for i := 0; i < sz; i++{
				numbers = append(numbers,j)
				j += (size + 1)
			}
		break
	}
	return numbers
}

func normal(start, end int) []int {
	lim := int(start/size)*size + (size - 1)
	for end > lim {
		end--
		start--
	}
	return []int{start,end,0}
}

func inv(start, end int) []int {
	lim := int(start/size)*size
	for end < lim {
		end++
		start++
	}
	return []int{start,end,1}
}

func abj(start, end int) []int {
	for end > (size*size - 1) {
		end -= size
		start -= size
	}
	return []int{start,end,2}
}

func arr(start, end int) []int {
	for end < 0 {
		end += size
		start += size
	}
	return []int{start,end,3}
}

func diag(start, end, sz int) []int {
	aux := start - sz*size
	for aux < 0 {
		aux += size
		end += size
		start += size
	}
	aux = start + sz - 1
	lim := int(start/size)*size + (size - 1)

	for aux > lim {
		aux--
		end--
		start--
	}
	return []int{start,end,4}
}

func diagI(start, end, sz int) []int {
	aux := start + sz*size
	for aux > (size*size - 1) {
		aux -= size
		end -= size
		start -= size
	}
	aux = start - sz + 1
	lim := int(start/size)*size
	for aux < lim {
		aux++
		end++
		start++
	}
	return []int{start,end,5}

}


func diagR(start, end, sz int) []int {
	aux := start - sz*size
	for aux < 0 {
		aux += 15
		end += size
		start += size
	}
	aux = start - sz + 1
	lim := int(start/size)*size
	for aux < lim {
		aux++
		end++
		start++
	}
	return []int{start,end,6}
}

func diagRI(start, end,sz int) []int {
	aux := start + sz*size
	for aux > (size*size - 1) {
		aux -= size
		end -= size
		start -= size
	}
	aux = start + sz - 1
	lim := int((start)/size)*size + (size - 1)
	for aux > lim {
		aux--
		end--
		start--
	}
	return []int{start,end,7}
}

func FillBoard(words map[string][]int) []int {
	board := make([]int,size*size)
	for i := 0; i < size; i++ {
		for j := 0; j < size; j++ {
			board[i*size + j] = '0'
		}
	}
	for k,v := range words {
		fmt.Println(k,v)
		board = fillWord(board,k,v)
	}
	return board
}

func PrintBoard(board []int) {
	fmt.Printf("    ")
	for i := 0; i < size; i++{
		fmt.Printf("%c  ",97 + i)
	}
	fmt.Println()
	for i := 0; i < size; i++ {
		fmt.Printf("%2d  ",i)
		for j := 0; j < size; j++ {
			fmt.Printf("%c  ",board[i*size + j])
		}
		fmt.Println()
	}
}

func fillWord(board []int, word string, lim []int) []int{
	switch lim[2] {
	case 0:
		j := lim[0]
		for i:= 0; i < len(word); i++{
			board[j] = int(word[i])
			j++
		}
	break
	case 1:
		j := lim[0]
		for i:= 0; i < len(word); i++{
			board[j] = int(word[i])
			j--
		}
	break
	case 2:
		j := lim[0]
		for i:= 0; i < len(word); i ++{
			board[j] = int(word[i])
			j += size
		}
	break
	case 3:
		j := lim[0]
		for i:= 0; i < len(word); i ++{
			board[j] = int(word[i])
			j -= size
		}
	break
	case 4:
		j := lim[0]
		for i:= 0; i < len(word); i ++{
			board[j] = int(word[i])
			j -= (size - 1)
		}
	break
	case 5:
		j := lim[0]
		for i:= 0; i < len(word); i ++{
			board[j] = int(word[i])
			j += (size - 1)
		}
	break
	case 6:
		j := lim[0]
		for i:= 0; i < len(word); i ++{
			board[j] = int(word[i])
			j -= (size + 1)
		}
	break
	case 7:
		j := lim[0]
		for i:= 0; i < len(word); i ++{
			board[j] = int(word[i])
			j += (size + 1)
		}
	break
	}
	return board
}