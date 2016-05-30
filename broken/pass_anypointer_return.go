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
	
	foo = a
	bar = b
	baz = c
	
}
