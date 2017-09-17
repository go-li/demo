package main

func macro_sort(  fun func(*,*), sli []) {
	var i int
	var j int
	for i = 0; i < len(sli); i++ {
	for j = 0; j < len(sli); j++ {
		fun(&sli[i], &sli[j])
	}}
}

func bubble_ints(a,b *int) {
	if *a > *b {
		*a, *b = *b, *a
	}
}

func main() {
	var stax = []int{7,5,3,47,35}
	
	macro_sort(bubble_ints, stax)
	for i := range stax {
	print(stax[i])
	print(",")
	}
	println()
}
