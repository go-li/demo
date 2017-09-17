package main

func unbox(input []*, zero *) (output []) {

	output = make([], len(input))

	var in *;
	var out *;

	for i := range input {

		in = input[i]
		out = &output[i]

		if in != nil {
			*out = *in
		} else if zero != nil {
			*out = *zero
		}
	}
	return output
}

func box(input [], empty func(*) bool) (output []*) {

	output = make([]*, len(input))

	for i := range input {

		if (empty == nil) || !empty(&input[i]) {

			output[i] = &input[i]

		}
	}
	return output
}

func main() {

	println("experiment 1 we box a bunch of integers")
	println("box means create a slice of pointers to elements")

	var foo []*int = box([]int{0,1,0,2,3}, nil)

	for _,v := range foo {
		if v == nil {
			print("nil ")
		} else {
			print(*v)
			print(" ")
		}
		println(v)
	}

	println("experiment 2 we box a bunch of bytes")
	println("function represents some bytes, e.g. zeroes with NULL")

	var bar []*uint8 = box([]uint8{0,1,0,2,3}, func(a *uint8)bool {
		return 0 == *a
	})

	for _,v := range bar {
		if v == nil {
			print("NULL! ")
		} else {
			print("VALUE:")
			print(*v)
			print(" ")
		}
		println(v)
	}

	println("experiment 3 we unbox integers foo")

	var unboxed_foo []int = unbox(foo, nil)

	for _,v := range unboxed_foo {
		println(v)
	}

	println("experiment 4 we unbox integers bar")
	println("for fun we replace nils with 255")

	var nilvalue uint8 = 255

	var unboxed_bar []uint8 = unbox(bar, &nilvalue)

	for _,v := range unboxed_bar {
		println(v)
	}
}
