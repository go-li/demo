package main

type template struct {
	thing *
}

type foo struct {
	thing *int
}

func DoFoo(t *template, i *) {
	t.Foo(i)
}

func (hi *template) Foo(i *) {
	hi.thing = i
}

func main() {

	var x int = 99999
//	var y int = 37
	var a foo


	DoFoo(&a, &x)

	print(*(a.thing))

//	(&a).Foo(&y)

//	print(*(a.thing))
}
