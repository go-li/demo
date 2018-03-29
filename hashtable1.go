package main

import "time"

// BEGIN PRNG

var s [2]uint64

func rotl( x uint64,  k uint) uint64 {
	return (x << k) | (x >> (64 - k));
}

func next() uint64 {
	var s0 = s[0];
	var s1 = s[1];
	var result = s0 + s1;

	s1 ^= s0;
	s[0] = rotl(s0, 55) ^ s1 ^ (s1 << 14); // a, b
	s[1] = rotl(s1, 36); // c

	return result;
}

// END PRNG

// BEGIN HTABLE


const hops = 96

func forEachKey(hash func(*,int) uint64, keys [], do func(*, int)) {

	var void uint64

	if len(keys)>0 {
		var voidt [];
		voidt = make([],1)

		void = hash(&voidt[0], -1)

		voidt = nil
	}


	for i := range keys {

		var val *;
		val = &keys[i]

		if hash(val,-1) == void {
			continue
		}
		do (val, i)
	}

}


type selector int

func (f selector) remove(table []) selector {

	if f == -1 {
		return -1
	}

	var empty [];
	empty = make([], 1)

	var void *;
	void = &empty[0]

	var where *;
	where = &table[f]

	*where = *void

	return f
}

func slect(hash func(*, int) uint64, key *, table []) selector {

	if key == nil {
		return -1;
	}

	// next we look if the slot is occupied

	var slot = hash(key, len(table)-hops)

	var location *;
	location = &table[slot]

	var hloc = hash(location, -1)

	// if it's occupied by a key

	var hkey = hash(key, -1)

	if hloc == hkey {
		return selector(int(slot))
	}

	for j := 0; j < hops; j++ {

	// otherwise look at location +1

	slot++

	location = &table[slot]

	hloc = hash(location, -1)

	// if it's occupied by a key

	if hloc == hkey {
		return selector(int(slot))
	}

	}


	return -1
}

func (f selector) from(values []) * {
	if f == -1 {
		return nil
	}

	return &values[f]
}

// fetch is like slect(..)from(..) but the key table pointer is returned
// fetch cannot fetch values from the values slice
func fetch(hash func(*, int) uint64, key *, table []) * {

//	if key == nil {
///		return nil
//	}

	// next we look if the slot is occupied

	var slot = hash(key, len(table)-hops)

	var location *;
	location = &table[slot]

	var hloc = hash(location, -1)

	// if it's occupied by a key

	var hkey = hash(key, -1)

	if hloc == hkey {
		return &table[slot]
	}

	for j := 0; j < hops; j++ {

	// otherwise look at location +1

	slot++

	location = &table[slot]

	hloc = hash(location, -1)

	// if it's occupied by a key

	if hloc == hkey {
		return &table[slot]
	}

	}


	return nil
}


type inserter func(int,int)

// n is a "prime number" from the sequence 251,491,971...
func grow(n int) int {
	n = ((n + 1) * 2) - 13
	var m = byte(((n & 0x7fff) ^ 9744) % 23)
	if (2400896 >> m) & 1 == 0 {
		return n
	}
	m = (((m+1)^4)%7)&7
	return int([8]uint32{30881,0,2426871851,15451,37867961,1189361,251,0}[m])
}


func (f inserter) into(hash func(*, int) uint64, key *, table *[]) {


	if len(*table) == 0 {
		(*table) = make([], 1+hops)
		var loc *;
		loc = &(*table)[0]
		*loc = *key

		if f != nil {
		f(1+hops,-1)
		f(0,-2)
		f(0,-4)
		}

		return
	}

	var empty [];
	empty = make([], 1)

	var void *;
	void = &empty[0]

	var hvoid = hash(void, -1)

	void = nil
	empty = nil

	again:

	// next we look if the slot is occupied

	var slot = hash(key, len(*table)-hops)

	var location *;
	location = &(*table)[slot]

	var hloc = hash(location, -1)

	var hkey = hash(key, -1)

	if hkey == hvoid {
		print("[")
		print(hvoid)
		print(" ")
		print(hkey)
		println("] Cannot insert null keyed object. Buggy hashfunction?")
		return
	}

	if hkey == hloc {

		if f != nil {
		f(int(slot),-4)
		}

		return
	}

	// we check if the slot is empty





	if hvoid == hloc {

		*location = *key
		if f != nil {
		f(int(slot),-4)
		}
		return
	}

	for j := 0; j < hops; j++ {

	// we look if the slot +1 is occupied

	slot ++
	location = &(*table)[slot]
	hloc = hash(location, -1)

	if hkey == hloc {

		if f != nil {
		f(int(slot),-4)
		}

		return
	}

	// we check if the slot+1 is empty

	if hvoid == hloc {

		*location = *key
		if f != nil {
		f(int(slot),-4)
		}

		return
	}

	}


	// not found so we grow the table

	var oldsize = len(*table)
	var newsize = grow(len(*table)-hops)+hops

//	println("")
//	println(newsize)

	var keys [];
	keys = make([], newsize)

	if f != nil {
	f(newsize, -1)
	}


	for i := 0; i < oldsize; i++ {

		var each *;
		each = &(*table)[i]

		var heach = hash(each, -1)

		if hvoid == heach {
			continue
		}

		heach = hash(each, newsize-hops)

		var newkey *;
		newkey = &keys[heach]

		var hnewkey = hash(newkey, -1)


		if hvoid == hnewkey {

			*newkey = *each
			if f != nil {
			f(int(heach), i)
			}
			continue

		}

		var j int
		for j = 0; j < hops; j++ {

		heach++
		newkey = &keys[heach]

		hnewkey = hash(newkey, -1)

		if hvoid == hnewkey {


			*newkey = *each
			if f != nil {
			f(int(heach), i)
			}
			break

		}

		}

		if j != hops {
			continue
		}


		newsize = grow(newsize-hops)+hops

		keys = make([], newsize)
		if f != nil {
		f(newsize, -1)
		}
		i = -1
	}



	*table = keys
	keys = nil
	if f != nil {
	f(0,-2)
	}

	goto again
}


func insert(value *, values *[]) inserter {

	if value == nil {
		return nil
	}

	var oldvalues []


	const resizeValues = -1	// save the old and make a new values array
	const deleteValues = -2	// clean the old (saved) values array
	const emptyElement = -3 // puts an empty element to slot in values array
	const addedElement = -4 // puts the inserted value to a slot in values


	return inserter( func(dst int, src int) {



		var dstp *;
		var srcp *; 

		switch (src) {
		case resizeValues:
			if oldvalues == nil {
				oldvalues = *values
			}
			*values = make([], dst)
			return
		case deleteValues:
			oldvalues = nil
			return
//		case emptyElement:
//			var void [];
//			void = make([], 1)
//			srcp = &void[0]
//			dstp = &(*values)[dst]
		case addedElement:
			srcp = value
			dstp = &(*values)[dst]
		default:
			srcp = &oldvalues[src]
			dstp = &(*values)[dst]
		}





		*dstp = *srcp
	})


}


//// VACUUM HASHTABLE //////////////////////////////////////////////////////////

// vacuum a hashtable



type vacuumer func(int,int)

func (f vacuumer) table(hash func(*, int) uint64, table [], accuracy byte) (displaced int) {

	if (accuracy < 1) || (accuracy > 8) {
		panic("vacuum accuracy should be small: 1,2,3,4 etc")
	}
	accuracy = (1 << accuracy) - 1

	var empty [];
	empty = make([], 1)

	var void *;
	void = &empty[0]

	var hvoid = hash(void, -1)
//	var hvacant = hash(void, len(table)-hops)


	var buf [][2]int
//	var mapa = make( map[int]int)
	const mapbsize = 1023
	var mapb [mapbsize+1]int
	const mapblank = -1
	for i := 0; i <= mapbsize; i++ {
		mapb[i] = mapblank
	}

	var nexslot = 0

	for i := 0; i < len(table); i++ {
		var posptr *;
		posptr = &table[i]



		var vacancy = hash(posptr, -1)

		if (hvoid == vacancy) { continue;}

		var pozition = hash(posptr, len(table)-hops)

//		println("here")

		if len(buf) < int(accuracy) {

		buf = append(buf, [2]int{i, int(pozition)})

		if len(buf) == int(accuracy) {
		sort(buf)
//		check(buf)
		}

		} else {
///// begin algorithm
		if nexslot < buf[0][1] {
			nexslot = buf[0][1]
		}

		if buf[0][0] != nexslot {

		posptr = &table[nexslot]
		vacancy = hash(posptr, -1)
/*
		print("moving ")
		print(buf[0][0])
		print(" to ")
		println(nexslot)
*/
		if vacancy == hvoid {

		if f != nil {
			f(-1, nexslot)
		}
		*void = *posptr

		var src *;
		var mynexslot = nexslot
	again2:
			displaced++
		posptr = &table[mynexslot]
		src = &table[buf[0][0]]
		if f != nil {
			f(mynexslot, buf[0][0])
		}
		*posptr = *src
/*
		print(mynexslot);print(" <- ");println(buf[0][0])
*/
		if buf[0][0] < mynexslot - 1000 {
			panic("not a burst")
		}

		if mapb[buf[0][0]&mapbsize] != mapblank {
			var tmp = mapb[buf[0][0]&mapbsize]
			mapb[buf[0][0]&mapbsize] = mapblank
			mynexslot = buf[0][0]
			buf[0][0] = tmp


			goto again2
		} else {
/*
			println("done")
*/
		}

		if f != nil {
			f(buf[0][0],-1)
		}
		src = &table[buf[0][0]]
		*src = *void


		mapb[buf[0][0]&mapbsize] = mapblank


		} else {


		var target = buf[0][0]
//		mapa[nexslot] = target
		mapb[nexslot&mapbsize] = target


		var ok bool

		for {
//			target, ok = mapa[target]
			target = mapb[target&mapbsize]; ok = target != mapblank
			if (!ok) || (target == buf[0][0]) {break;}
		}


		if ok {
			target = nexslot

			var ptr *;
			var dst *;
/*
			print("put away ");println(nexslot)
*/
			if f != nil {
				f(-1, nexslot)
			}
			ptr = &table[nexslot]
			*void = *ptr


//			var nextarget = mapa[target]
			var nextarget = mapb[target&mapbsize]
			for nextarget != nexslot {


/*
			print(target);print(" <- ");println(nextarget)
*/
			if nextarget < target - 1000 {
				panic("not a loop")
			}

			if f != nil {
				f(target, nextarget)
			}
			dst = &table[target]
			ptr = &table[nextarget]
			*dst = *ptr

			displaced++
/*
			print(target)
			println("deleted")
*/
//			delete(mapa,target)
			mapb[target&mapbsize] = mapblank
			target = nextarget
//			nextarget = mapa[nextarget]
			nextarget = mapb[nextarget&mapbsize]
			}
/*
			print(target);println(" from away ");

			print(target)
			println("deleted")
*/
//			delete(mapa,target)
			mapb[target&mapbsize] = mapblank

			if f != nil {
				f(target, -1)
			}
			dst = &table[target]
			*dst = *void

		}

		}

		}
// end algorithm
		buf[0][0] = i
		buf[0][1] = int(pozition)
		bubble(buf)
//		check(buf)
/*		for i := range buf {
		print("moving ")
		print(buf[i][0])
		print(" to ")
		println(buf[i][1])



		}
*/
		nexslot++
//		if nexslot > 50 {return;}

		}

		
	}


//	println("///////////////")

	for i := range buf {
		var posptr *;
		var vacancy uint64

		if nexslot < buf[i][1] {
			nexslot = buf[i][1]
		}

		if buf[i][0] != nexslot {

		posptr = &table[nexslot]
		vacancy = hash(posptr, -1)
/*
		print("moving ")
		print(buf[i][0])
		print(" to ")
		println(nexslot)
*/
		if vacancy == hvoid {

		if f != nil {
			f(-1, nexslot)
		}
		*void = *posptr

		var src *;
		var mynexslot = nexslot
	again3:
			displaced++
		posptr = &table[mynexslot]
		src = &table[buf[i][0]]
		if f != nil {
			f(mynexslot, buf[i][0])
		}
		*posptr = *src
/*
		print(mynexslot);print(" <- ");println(buf[i][0])
*/
		if buf[i][0] < mynexslot - 1000 {
			panic("not a burst")
		}

		if mapb[buf[i][0]&mapbsize] != mapblank {
			var tmp = mapb[buf[i][0]&mapbsize]
			mapb[buf[i][0]&mapbsize] = mapblank
			mynexslot = buf[i][0]
			buf[i][0] = tmp


			goto again3
		} else {
/*
			println("done")
*/
		}

		if f != nil {
			f(buf[i][0],-1)
		}
		src = &table[buf[i][0]]
		*src = *void

		} else {


		var target = buf[i][0]
//		mapa[nexslot] = target
		mapb[nexslot&mapbsize] = target


		var ok bool

		for {
//			target, ok = mapa[target]
			target = mapb[target&mapbsize]; ok = target != mapblank
			if (!ok) || (target == buf[i][0]) {break;}
		}


		if ok {
			target = nexslot

			var ptr *;
			var dst *;
/*
			print("put away ");println(nexslot)
*/
			if f != nil {
				f(-1, nexslot)
			}
			ptr = &table[nexslot]
			*void = *ptr


//			var nextarget = mapa[target]
			var nextarget = mapb[target&mapbsize]
			for nextarget != nexslot {


/*
			print(target);print(" <- ");println(nextarget)
*/
			if nextarget < target - 1000 {
				panic("not a loop")
			}

			if f != nil {
				f(target, nextarget)
			}
			dst = &table[target]
			ptr = &table[nextarget]
			*dst = *ptr

			displaced++
/*
			print(target)
			println("deleted")
*/
//			delete(mapa,target)
			mapb[target&mapbsize] = mapblank
			target = nextarget
//			nextarget = mapa[nextarget]
			nextarget = mapb[nextarget&mapbsize]
			}
/*
			print(target);println(" from away ");

			print(target)
			println("deleted")
*/
//			delete(mapa,target)
			mapb[target&mapbsize] = mapblank

			if f != nil {
				f(target, -1)
			}
			dst = &table[target]
			*dst = *void

		}

		}

		}

		nexslot++
	}


	return displaced
}

func vacuum(values []) vacuumer {

	if len(values) == 0 {
		return nil
	}

	var putaway []
	putaway = make([], 1)

	return vacuumer( func(dst int, src int) {
		var dstp *;
		var srcp *;

		if src == -1 {
			srcp = &putaway[0]
			dstp = &values[dst]
		} else if dst == -1 {
			srcp = &values[src]
			dstp = &putaway[0]
		} else {
			srcp = &values[src]
			dstp = &values[dst]
		}

		*dstp = *srcp
	})
}

func sort(array [][2]int) {
	for i := 0; i < len(array); i++ {
	for j := 0; j < i; j++ {
		if array[i][1] < array[j][1] {
			array[i], array[j] = array[j], array[i]
		}
	}}
}
/*
func check(array [][2]int) {
	for i := 1; i < len(array); i++ {
		if array[i][1] < array[i-1][1] {
			panic("not sorted\n")
		}
	}
}
*/
func bubble(array [][2]int) {
	for i := 1; i < len(array); i++ {
		if array[i][1] < array[i-1][1] {
			array[i], array[i-1] = array[i-1], array[i]
		} else {
			return
		}
	}
}

type deletor func(int,int)

func delet(values []) deletor {

	if len(values) == 0 {
		return nil
	}

	return func(i int, j int) {

		var vi *;
		var vj *;

		vi = &values[i]
		vj = &values[j]

		*vi = *vj

//		values[i] = values[j]
	}
}

func (d deletor) from(hash func(*, int) uint64, key *, table []) {

	var empty [];
	empty = make([], 1)

	var void *;
	void = &empty[0]

	var hvoid = hash(void, -1)
	var hvacant = hash(void, len(table)-hops)


	// next we look if the slot is occupied

	var slot = hash(key, len(table)-hops)

	var location *;
	location = &table[slot]

	var hloc = hash(location, -1)

	// if it's occupied by a key

	var hkey = hash(key, -1)

	if hloc == hkey {
		*location = *void

	} else {

	for j := 0; j < hops; j++ {

	// otherwise look at location +1

	slot++

	location = &table[slot]

	hloc = hash(location, -1)

	// if it's occupied by a key

	if hloc == hkey {
		*location = *void
		break
	}

	if hloc == hvoid {
		return
	}

	}
	}


	for i := slot + 1 ; (int(i) < len(table)) && (i < slot + uint64(hops)); i++ {

		var item *;

		item = &table[i]

		var itemloc = hash(item, len(table)-hops)

		if itemloc == hvacant {
			var itemhash = hash(item, -1)
			if itemhash == hvoid {
				return
			}
		}

		if itemloc <= slot {


//		print("slot=")
//		print(slot)
//		print(" i=")
//		print(i)
//		print(" itemloc=")
//		print(itemloc)
//		println(" FIXING BURST")

		if d != nil {
			d(int(slot), int(i))
		}

		*location = *item
		*item = *void


		slot = i
		location = &table[slot]
		continue

		}
	}
}

// --- 128 slice based hashtables

func (f inserter) into128(hash func(*, int) uint64, key *, table *[128][]) {
	var k = hash(key, -1)

	f.into(hash, key, &(*table)[k&127])
}

func (f vacuumer) table128(hash func(*, int) uint64, table *[128][], accuracy byte) (displaced int) {
	for i := 0; i < 128; i++ {
		displaced += f.table(hash, (*table)[i], accuracy)
	}
	return displaced
}

func (d deletor) from128(hash func(*, int) uint64, key *, table *[128][]) {
	var k = hash(key, -1)

	d.from(hash, key, (*table)[k&127])
}


//---

func hash32(ptr *uint32, mod int) uint64 {
	if mod == -1 {
		return uint64(*ptr)
	}
	var p = uint64(*ptr)
	var m = uint64(mod)
	return (uint64(uint32(p * m)) * m) >> 32
}

func hash16(ptr *uint16, mod int) uint64 {
	var p = uint64(*ptr)
	var m = uint64(mod)
	return (uint64(uint16(p * m)) * m) >> 16
}


func hash16mod(ptr *uint16, mod int) uint64 {
	return uint64(uint64(*ptr) % uint64(mod))
}

func hash32mod(ptr *uint32, mod int) uint64 {
	return uint64(uint64(*ptr) % uint64(mod))
}

//

func main() {


	s = [2]uint64{0xbeac0467eba5facd, 0xd86b048b86aa9922}


	var hashmap [128][]uint32

	var begin = time.Now().UnixNano()





	for i := 0; i < 1000000; i++ {
		var k = uint32(next())
		inserter(nil).into128(hash32, &k, &hashmap)
	}



//	vacuumer(nil).table128(hash32, &hashmap, 6)


	s = [2]uint64{0xbeac0467eba5facd, 0xd86b048b86aa9922}

	for i := 0; i < 1000000; i++ {
		var k = uint32(next())
		deletor(nil).from128(hash32, &k, &hashmap)
	}


	var end = time.Now().UnixNano()
	println(end - begin)

	for i := 0; i < len(hashmap); i++ {
	forEachKey(hash32, hashmap[i], func(*uint32, int) {
		panic("Something is in emptied table");
	})
	}



}
