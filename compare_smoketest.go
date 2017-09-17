package main

func Int(a, b *int) int {
	return *a - *b
}

func String(a, b *string) int {
	sa := *a
	sb := *b
	for i := 0; (i < len(sa)) && (i < len(sb)); i++ {
		if sa[i] != sb[i] {
			return int(sa[i]) - int(sb[i])
		}
	}
	return 0
}

func Compare2(cmp func(*,*)(int), l *,r *) {
	println(What(cmp(l,r)))
}

func What(n int) string {
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
	
	Compare2(String, &s[0], &s[1])
	Compare2(String, &s[2], &s[3])
	Compare2(String, &s[4], &s[5])
	Compare2(String, &s[6], &s[7])
	Compare2(Int, &one, &two)
}
