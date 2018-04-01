package main

import "math"

func search(buf [], searchFromEnd bool, find func(*,int)bool, fromInclusive int, toExclusive int) int {
	if len (buf) == 0 {
		return -1
	}

	var elm *;

	elm = &buf[0]

	var ordered = find(elm, 0) != searchFromEnd
	var lo = fromInclusive;
	var hi = toExclusive;
	for (lo < hi) {
		var mid = (lo & hi) + ((lo ^ hi) >> 1);

		elm = &buf[mid]

		if (find(elm, mid)) != ordered {
			lo = mid + 1;
		} else {
			hi = mid;
		}
	}
	return hi;
}

func find(buf [], goal *, compare func(*,*)int, fromInclusive int, toExclusive int) int {
	if len (buf) == 0 {
		return -1
	}

	var elm *;
	var nod *;

	elm = &buf[0]
	nod = &buf[len(buf)-1]

	var ordered = compare(elm, nod) > 0
	var add = 0
	if ordered {
		add = -1
	}
	var lo = fromInclusive;
	var hi = toExclusive;
	for (lo < hi) {
		var mid = (lo & hi) + ((lo ^ hi) >> 1);

		elm = &buf[mid]

		if (compare(goal, elm) > add) != ordered {
			lo = mid + 1;
		} else {
			hi = mid;
		}
	}
	return hi;
}

// float compares two numbers.
func float(a, b *float32) int {
	r := *a - *b
	rr := int32(math.Float32bits(r))
	return int(rr)
}

func main() {
	var goal int = 13

	println(search([]int{5,6,7,13,15,19,25}, true, func(n *int, mid int)bool {
		println(*n)
		return *n >= 13
	}, 0, 7))

	println()

	println(find([]int{5,6,7,13,15,19,25}, &goal, func(a *int, b *int)int {
		println(*b)
		return *a - *b
	}, 0, 7))

	println()

	println(search([]int{25,19,15,13,7,6,5}, true, func(n *int, mid int)bool {
		println(*n)
		return *n <= 13
	}, 0, 7))

	println()

	println(find([]int{25,19,15,13,7,6,5}, &goal, func(a *int, b *int)int {
		println(*b)
		return *a - *b
	}, 0, 7))

	println()

	var buffer = []float32{7.3,9.5,11.2,16.8,24.4,39.7,70.2}

	var target float32 = 16.9

	var off = find(buffer, &target, float, 0, len(buffer))
	_ = off
	println(off)
	println(buffer[off])
}
