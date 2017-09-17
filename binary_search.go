package main

import "math"

func Search(buf [], goal *, compare func(*,*)int, fromInclusive int, toExclusive int) int {
	if len (buf) == 0 {
		return -1
	}

	var elm *;
	var nod *;

	elm = &buf[0]
	nod = &buf[len(buf)-1]

	var ordered = compare(elm, nod) > 0
	var lo = fromInclusive;
	var hi = toExclusive;
	for (lo < hi) {
		var mid = (lo & hi) + ((lo ^ hi) >> 1);

		elm = &buf[mid]

		if (compare(goal, elm) > 0) != ordered {
			lo = mid + 1;
		} else {
			hi = mid;
		}
	}
	return hi;
}

// Float32 compares two numbers.
func Float32(a, b *float32) int {
	r := *a - *b
	rr := int32(math.Float32bits(r))
	return int(rr)
}

func main() {
	var buffer = []float32{7.3,9.5,11.2,16.8,24.4,39.7,70.2}

	var target float32 = 16.9

	var off = Search(buffer, &target, Float32, 0, len(buffer))
	_ = off
	println(off)
	println(buffer[off])
}
