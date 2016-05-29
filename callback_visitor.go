package main

import (
	"fmt"
)

func set7_visitor(num *int) {
           fmt.Printf("I am: ")
           *num = 7
}
func print_visitor(num *int) {
        fmt.Printf("%d", *num)
}
   
func visit(visitor func(*), obj *) {
        visitor(obj)
}

func main() {
	var a,b,c = 1,2,3
	macro_visit(set7_visitor, &b)
	macro_visit(print_visitor, &a)
	macro_visit(print_visitor, &c)
	macro_visit(print_visitor, &c)
	macro_visit(print_visitor, &b)
	
}
