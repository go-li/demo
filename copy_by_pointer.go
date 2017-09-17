// the generic copy demo
package main

func nocopy0(a *, b *) {
///////////// do something anyway
	for i := 0; i < 10; i++ {
		print(".")
	}
	print("\n")
}


func docopy1( a *, b *) {
	*a = *b
}

func docopy2( a **, b **) {
	**a = **b
}

func docopy3( a ***, b ***) {
	***a = ***b
}

func main() {
	var a int = 0x1E1412A1
	var b int = 0x3723B044
	var c int = 0x22658468

	var pa = &a
	var pb = &b
	var pc = &c

	var ppa = &pa
	var ppb = &pb
	var ppc = &pc

	_,_,_ = ppa,ppb,ppc


	nocopy0(&a, &b)

	print("  a=")
	print(a)
	print("  b=")
	print(b)
	print("  c=")
	print(c)
	print("\n")

	docopy1( &b, &c)

	print("  a=")
	print(a)
	print("  b=")
	print(b)
	print("  c=")
	print(c)
	print("\n")

	docopy2( &pc, &pa)

	print("  a=")
	print(a)
	print("  b=")
	print(b)
	print("  c=")
	print(c)
	print("\n")

	docopy3( &ppa, &ppb)

	print("  a=")
	print(a)
	print("  b=")
	print(b)
	print("  c=")
	print(c)
	print("\n")
}
