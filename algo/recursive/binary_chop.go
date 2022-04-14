package recursive

//BinaryChop by recursive
func BinaryChop(val []int, id, s, e int) (int, bool) {
	if s > e || val[s] > id || val[e] < id {
		return -1, false
	}
	if len(val[s:e]) == 1 {
		if val[s] == id {
			return s, true
		}
		return s, false
	}
	k := len(val[s:e]) / 2
	if id > val[k+s] {
		return BinaryChop(val, id, k+s, e)
	} else if id == val[k+s] {
		return k + s, true
	}
	return BinaryChop(val, id, s, s+k)
}

//BinaryChop2 no recursive
func BinaryChop2(val []int, id int) (int, bool) {
	if len(val) == 0 {
		return -1, false
	}
	s, e := 0, len(val)-1
	for s <= e {
		k := (s + e) / 2
		if id == val[k] {
			return s, true
		}
		if id > val[k] {
			s = k + 1
		} else {
			e = k - 1
		}
	}
	return -1, false
}
