package main

func ret() (*int, byte, string) {
	var i = 7
	return &i, 'a', "hello"
}

func acc(a *, b byte, c string) {
	println(b)
	println(c)

}

func main() {
	acc(ret())
}
