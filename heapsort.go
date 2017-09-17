package main

func CompareInt(a *int, b *int) int {
	return (*a) - (*b)
}

func Sort(base [], cmp_func func (*, *) int) {

	var tmp []
	var pt *;
	tmp = make([],1)
	pt = &tmp[0]

	var c, r = 0, 0
	var pc *;
	var pr *;

	/* heapify */
	for i := (len(base)/2 - 1) ; i >= 0; i -= 1 {
		for r = i; r * 2 + 1 < len(base); r = c {
			c = r * 2 + 1;

			pc = &base[c]
			pr = &base[c+1]

			if ((c < len(base) - 1) && (cmp_func(pc, pr) < 0)){
				c += 1;}

			pc = &base[c]
			pr = &base[r]

			if (cmp_func(pr, pc) >= 0) {
				break;}

			*pt = *pr
			*pr = *pc
			*pc = *pt
		}
	}

	/* sort */
	for i := len(base) - 1; i > 0; i -= 1 {

		pc = &base[0]
		pr = &base[i]

		*pt = *pr
		*pr = *pc
		*pc = *pt

		for r = 0; r * 2 + 1 < i; r = c {
			c = r * 2 + 1;

			pc = &base[c]
			pr = &base[c+1]

			if ((c < i - 1) && (cmp_func(pc, pr) < 0)){
				c += 1;}

			pc = &base[c]
			pr = &base[r]

			if (cmp_func(pr, pc) >= 0){
				break;}

			*pt = *pr
			*pr = *pc
			*pc = *pt
		}
	}
}

func main() {
	var buf = []int {984,135,651,897,648,412,741,151,944,841,254,825,985,981,785}


	Sort(buf, CompareInt)

	for i := range buf {
		println(buf[i])
	}
}
