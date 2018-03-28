package main

func integer(a, b *int) int {
	return *a - *b
}

func str(a, b *string) int {
	sa := *a
	sb := *b
	for i := 0; (i < len(sa)) && (i < len(sb)); i++ {
		if sa[i] != sb[i] {
			return int(sa[i]) - int(sb[i])
		}
	}
	return 0
}

func compare2(cmp func(*,*)(int), l *,r *) {
	println(what(cmp(l,r)))
}

func what(n int) string {
	if n == 0 {
		return "same"
	} else if n > 0 {
		return "l > r"
	}
	return "l < r"
}

func main() {
	s := []string{"foo","bar","gg","zz","aa","bb","cc","cc"}

	var one = 1
	var two = 2
	
	compare2(str, &s[0], &s[1])
	compare2(str, &s[2], &s[3])
	compare2(str, &s[4], &s[5])
	compare2(str, &s[6], &s[7])
	compare2(integer, &one, &two)
}
