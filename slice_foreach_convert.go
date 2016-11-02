package main

func foreach_slice(slice []*, call func(*)) {
	for i := range slice {
		call(slice[i])
	}
}

func main() {
	 a := "a" ;	 aa := "aa";	 aaa := "aaa"
	var as = []*string{&a, &aa, &aaa}
	
	var bs []*int
	
	// slice conversion one-liner
	foreach_slice(as, func(arg *string){var i int; i = len(*arg); bs = append(bs,&i)})


	// print the result
	foreach_slice(bs, func(arg *int){
		print(*arg);
		print("\n")
	})
	
}
