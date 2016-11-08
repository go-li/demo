package main

func macro_pass(ptr *) (ret *) {
	return ptr
}

func main() {

	var foo int
	var bar string
	var baz byte

	var a *int = macro_pass(&foo)
	var b *string = macro_pass(&bar)
	var c *byte = macro_pass(&baz)

	foo = *a
	bar = *b
	baz = *c

	var x = macro_pass(&foo)
	var y = macro_pass(&bar)
	var z = macro_pass(&baz)

	foo = *x
	bar = *y
	baz = *z

	var u = *macro_pass(&foo)
	var v = *macro_pass(&bar)
	var w = *macro_pass(&baz)

	_ = x
	_ = y
	_ = z
	_ = u
	_ = v
	_ = w
}
