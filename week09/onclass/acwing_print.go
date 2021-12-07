package week09onclass

// 打印
type ePrint struct {
	dep int
	pre string
}

var pDep map[string]ePrint

func eightPrint(start string) []string {
	target := "12345678x"
	h = heap{
		heapList: make([]eval, 1, 10),
	}
	h.Push(eval{e: eEval(start), s: start})
	pDep = make(map[string]ePrint, 1)
	pDep[start] = ePrint{
		dep: 0,
		pre: "",
	}
	for h.Size() > 0 {
		s := h.Pop()
		pos := findX(s.s)
		if pos >= 3 {
			pExpand(s.s, pos, pos-3)
		}
		if pos < 6 {
			pExpand(s.s, pos, pos+3)
		}
		if pos%3 != 0 {
			pExpand(s.s, pos, pos-1)
		}
		if pos%3 != 2 {
			pExpand(s.s, pos, pos+1)
		}
		if _, ok := pDep[target]; ok {
			return eightPr(target)
		}
	}
	return nil
}

func eightPr(t string) []string {
	var ans = []string{t}
	for {
		if _, ok := pDep[t]; ok {
			if pDep[t].pre == "" {
				break
			}
			ans = append(ans, pDep[t].pre)
			t = pDep[t].pre
		} else {
			break
		}
	}
	return ans
}

func pExpand(s string, p1, p2 int) {
	var b = []byte(s)
	tmp := b[p1]
	b[p1] = b[p2]
	b[p2] = tmp
	if _, ok := pDep[string(b)]; ok {
		return
	}
	pDep[string(b)] = ePrint{
		dep: pDep[s].dep + 1,
		pre: s,
	}
	h.Push(eval{
		e: eEval(string(b)),
		s: string(b),
	})
}
