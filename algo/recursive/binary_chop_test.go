package recursive

import "testing"

func TestName(t *testing.T) {
	s := []int{1, 2, 3, 4, 5, 6, 7, 9, 10}
	sub := 9

	id, ok := BinaryChop(s, sub, 0, len(s)-1)
	if ok {
		println("----", s[id], sub)
	}

	id, ok = BinaryChop2(s, sub)
	if ok {
		println("----", s[id], sub)
	}

	//for _, i := range s {
	//	id, ok := BinaryChop3(s, i)
	//	if ok {
	//		println(s[id], i)
	//	}
	//}

}

func TestName2(t *testing.T) {
	var s uint64 = 0xFFFFFFFFFFFFFFFF
	println(s)
}
