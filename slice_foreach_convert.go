package main

func foreach_slice(slice [], call func(*)) {
	for i := range slice {
		call(&slice[i])
	}
}

func main() {

	var as = []string{"a", "aa", "aaa"}
	
	var bs []int
	
	// slice conversion one-liner
	foreach_slice(as, func(arg *string){bs=append(bs,len(*arg))})


	// print the result
	foreach_slice(bs, func(arg *int){
		println(*arg);
	})
	
}
