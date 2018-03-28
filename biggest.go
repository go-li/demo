package main

func compareInt(a, b *int) int {
	return *a - *b
}

func max(compar func(*, *) int, items [] /*<-- delete int*/) int {
	j := 0
	for i := range items {
		var l *;
		var r *;

		l = &items[i]
		r = &items[j]

		if compar(l, r) > 0 {
			j = i
		}
	}
	return j
}

func whatIs(j int) {
	print("biggest:#")
	print(j+1)
	println(".")
}

func main() {

	whatIs(max(compareInt, []int{1, 0}));
	whatIs(max(compareInt, []int{1, 2}));
	whatIs(max(compareInt, []int{1, 1, 1}));

}
