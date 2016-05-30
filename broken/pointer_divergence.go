package main

func macro_ref_divergence(a, b *) {
	print("hello")
	macro_ref_divergence(&a, &b)
}

func main() {
	var ua, ub int
	ua = ub

	macro_ref_divergence(&ua, &ub)
}
