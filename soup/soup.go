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
func MakeDic(words []string) map[string][]int{
	Rd := rand.New(rand.NewSource(time.Now().UnixNano()))
	m := make(map[string][]int)
	for _,w := range words {
		start := Rd.Intn(225)
		dir := Rd.Intn(8)
		end := 0
		fmt.Println(w,dir)
		pos := getPos(dir,start,end,w)
		/*for !verifyPosition(pos,m) {
			start := Rd.Intn(225)
			dir := Rd.Intn(8)
			end := 0
			pos = getPos(dir,start,end,w)
		}*/
		m[w] = pos
	}

	return m
}
func getPos(dir, start,end int, w string) []int {
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
			end = start + (len(w))*15
			pos = abj(start, end)
		break
		case 3:
			end = start - (len(w))*15
			pos = arr(start, end)
		break
		case 4:
			end = start - (len(w)-1)*14
			pos = diag(start, end, len(w))
		break
		case 5:
			end = start + (len(w)-1)*14
			pos = diagI(start, end, len(w))
		break
		case 6:
			end = start - (len(w)-1)*16
			pos = diagR(start, end, len(w))
		break
		case 7:
			end = start + (len(w)-1)*16
			pos = diagRI(start, end,len(w))
		break
	}
	return pos
}

func verifyPosition(data []int, words map[string][]int) bool{
	for _,pos := range words {
		if Contains(data,pos) {
			return false
		}
	}
	return true
}

func Contains(a,b []int) bool {
	for _,v1 := range a {
		for _,v2 := range b {
			if v1 == v2 {
				return true
			}
		}
	}
	return false
}

func getNumbers(pos []int, w string) []int {
	sz := len(w)
	numbers := make([]int, sz)

	switch pos[3] {
		case 0:
			for i := pos[0]; i < pos[1]; i++{
				numbers = append(numbers,i)
			}
		break
		case 1:
			for i := pos[1]; i < pos[0]; i++{
				numbers = append(numbers,i)
			}
		break
		case 2:
			j := pos[0]
			for i := 0; i < sz; i++{
				numbers = append(numbers,j)
				j += 15
			}
		break
		case 3:
			j := pos[0]
			for i := 0; i < sz; i++{
				numbers = append(numbers,j)
				j -= 15
			}
		break
		case 4:
			j := pos[0]
			for i := 0; i < sz; i++{
				numbers = append(numbers,j)
				j -= 14
			}
		break
		case 5:
			j := pos[0]
			for i := 0; i < sz; i++{
				numbers = append(numbers,j)
				j += 15
			}
		break
		case 6:
			j := pos[0]
			for i := 0; i < sz; i++{
				numbers = append(numbers,j)
				j -= 16
			}
		break
		case 7:
			j := pos[0]
			for i := 0; i < sz; i++{
				numbers = append(numbers,j)
				j += 16
			}
		break
	}
	return numbers
}

func normal(start, end int) []int {
	lim := int(start/15)*15 + 14
	for end > lim {
		end--
		start--
	}
	return []int{start,end,0}
}

func inv(start, end int) []int {
	lim := int(start/15)*15
	for end < lim {
		end++
		start++
	}
	return []int{start,end,1}
}

func abj(start, end int) []int {
	for end > 224 {
		end -= 15
		start -= 15
	}
	return []int{start,end,2}
}

func arr(start, end int) []int {
	for end < 0 {
		end += 15
		start += 15
	}
	return []int{start,end,3}
}

func diag(start, end, sz int) []int {
	for end < 0 {
		end += 15
		start += 15
	}
	lim := int((start + (sz-1))/15)*15 + 14
	for end > lim {
		end--
		start--
	}
	return []int{start,end,4}
}

func diagI(start, end, sz int) []int {
	for end > 224 {
		end -= 15
		start -= 15
	}
	lim := int(start/15)*15
	for end < lim {
		end++
		start++
	}
	return []int{start,end,5}

}


func diagR(start, end, sz int) []int {
	fmt.Println("6",start, end, sz)
	for end < 0 {
		end += 15
		start += 15
	}
	lim := int(start/15)*15
	for end < lim {
		end++
		start++
	}
	fmt.Println("66",start, end, sz)
	return []int{start,end,6}
}

func diagRI(start, end,sz int) []int {
	fmt.Println("7",start, end, sz)
	for end > 224 {
		end -= 15
		start -= 15
	}
	lim := int((start + (sz-1))/15)*15 + 14
	for end > lim {
		end--
		start--
	}
	fmt.Println("77",start, end, sz)
	return []int{start,end,7}
}

func FillBoard(words map[string][]int) []int {
	board := make([]int,225)
	for i := 0; i < 15; i++ {
		for j := 0; j < 15; j++ {
			board[i*15 + j] = '0'
		}
	}
	for k,v := range words {
		fmt.Println(k,v)
		board = fillWord(board,k,v)
	}
	return board
}

func PrintBoard(board []int) {
	for i := 0; i < 15; i++ {
		for j := 0; j < 15; j++ {
			fmt.Printf("%c  ",board[i*15 + j])
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
			j += 15
		}
	break
	case 3:
		j := lim[0]
		for i:= 0; i < len(word); i ++{
			board[j] = int(word[i])
			j -= 15
		}
	break
	case 4:
		j := lim[0]
		for i:= 0; i < len(word); i ++{
			board[j] = int(word[i])
			j -= 14
		}
	break
	case 5:
		j := lim[0]
		for i:= 0; i < len(word); i ++{
			board[j] = int(word[i])
			j += 14
		}
	break
	case 6:
		j := lim[0]
		for i:= 0; i < len(word); i ++{
			board[j] = int(word[i])
			j -= 16
		}
	break
	case 7:
		j := lim[0]
		for i:= 0; i < len(word); i ++{
			board[j] = int(word[i])
			j += 16
		}
	break
	}
	return board
}