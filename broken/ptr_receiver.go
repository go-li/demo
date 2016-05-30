package main

func (hi *) Foo(i int) {
	print("FOO")
}

func main() {
	(&i).Foo(3)
}
