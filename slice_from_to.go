package main

import "strconv"

type converter func()string

func slice_via_str_from(slice []*, call func(*)string) (s converter) {
	var n = -1
	return func()string {
		if n >= len(slice) {
			return ""
		} else {
			n++
			return call(slice[n])
		}
	}
}

func (f converter) to(slice []*, call func(string, **)) {
	for i := range slice {
		call(f(), &slice[i])
	}
}

////////  TYPE-SPECIFIC BOILERPLATE //////////

func my_float64_to_string(f *float64)string {
	return strconv.FormatInt(int64(*f), 10)
}
func my_string_to_int(s string, out **int) {
	i, err := strconv.Atoi(s)
	if (err != nil) {
		*out = nil
	} else {
		*out = &i
	}
}

///////////////////////////

func main() {
	n1 := 32. ;	 n2 := 640.;	 n3 := 9600.
	var as = []*float64{&n1, &n2, &n3}
	
	var bs = make ([]*int, 3)
	_ = bs
	
	// arbitrary slice conversion from original type via string to new type
	slice_via_str_from(as, my_float64_to_string).to(bs, my_string_to_int)
	
	
	// print the result
	for _, v := range bs {
		print(*v);
		print("\n")
	}
}
