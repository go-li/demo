package main

import "strconv"

type converter struct {
	f func()string
	l int
}

func slice_via_str_from(slice [], call func(*)string) (s converter) {

	var n = -1
	return converter { f: func()string {
			if n >= len(slice) {
				return ""
			} else {
				n++
				var element *;

				element = &slice[n]

				return call(element)
			}
		}, l : len(slice),
	}
}

func (f converter) to(FIXME *, call func(string, *)) (slice []) {

	slice = make([], f.l)

	for i := range slice {
		call(f.f(), &slice[i])
	}
	return slice
}

////////  TYPE-SPECIFIC BOILERPLATE //////////

func my_float64_to_string(f *float64)string {
	return strconv.FormatInt(int64(*f), 10)
}
func my_string_to_int(s string, out *int) {

	i, err := strconv.Atoi(s)
	if (err != nil) {
		*out = nil
	} else {
		*out = i
	}
}

///////////////////////////

func main() {
	var as = []float64{32.,640.,9600.}
	
	// arbitrary slice conversion from original type via string to new type
	var bs []int = slice_via_str_from(as, my_float64_to_string).to(nil, my_string_to_int)

	
	// print the result
	for _, v := range bs {
		print(v);
		print("\n")
	}
}
