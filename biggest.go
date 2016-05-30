package main

func CompareInt(a, b *int) int {
	return *a - *b
}

func Max(compar func(*int, *int) int, items ...*int /*<-- delete int*/) int {
	j := 0
	for i := range items {
		if compar(items[i], items[j]) > 0 {
			j = i
		}
	}
	return j
}

func whatIs(j int) {
	print("biggest:#")
	print(j+1)
	print(". ")
}

func main() {
	var w, x, y, z int
	w = 14
	x = 65
	y = 74
	z = 96
	_,_,_,_ = w,x,y,z


	whatIs(Max(CompareInt, &w, &x));
	whatIs(Max(CompareInt, &w, &x, &y));
	whatIs(Max(CompareInt, &w, &x, &y, &z));
	
}
