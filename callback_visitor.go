package main

func set7_visitor(num *int) {
           print("I am: ")
           *num = 7
}
func print_visitor(num *int) {
        print(*num)
}

func newline_visitor(num *int) {
	println("")
}
   
func visit(visitor func(*), obj *) {
        visitor(obj)
}

func main() {
	var a,b,c = 1,2,3
	visit(set7_visitor, &b)
	visit(print_visitor, &a)
	visit(print_visitor, &c)
	visit(print_visitor, &c)
	visit(print_visitor, &b)
	visit(newline_visitor, &b)
}
