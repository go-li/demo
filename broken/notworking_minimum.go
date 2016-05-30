package main

func min(a, b *) * {
	if a < b {
		return a
	}
	return b
}

func main() {
	a1, b1 := 1, 2
	a2, b2 := 'a', 'b'
	a3, b3 := "a", "b"
	min(a1, b1)
	min(a2, b2)
	min(a3, b3)
}
