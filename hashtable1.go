package main

const hops = 255

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

func hashfun(key *[8]byte, size int) (o uint64) {
	o |= uint64(key[0])
	o |= uint64(key[1]) << 8 
	o |= uint64(key[2]) << 16 
	o |= uint64(key[3]) << 24
	o |= uint64(key[4]) << 32
	o |= uint64(key[5]) << 40
	o |= uint64(key[6]) << 48
	o |= uint64(key[7]) << 56

	if size == -1 {
		return o
	}

	return o % uint64(size)

}



func hashfn(key *int, size int) (o uint64) {
	switch (size) {
	case -1: return uint64(*key)
	default: return (uint64(*key) * 0x61C8864680B583EB) % uint64(size)
	}
}


type selector int

func (f selector) Remove(table []) selector {

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

		f(int(slot),-4)

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

type keyval struct {
	key int
	val int
}

func kvhfunc(key *keyval, modulo int) (o uint64) {

	return hashfn(&(key.key), modulo)
/*
	var x = uint64((*key).key)

	if modulo == -1 {
		return x
	}

	o = (
		((((x >> +0) & 1) - 1) & 0x34b1abcc664e7d08) +
		((((x >> +1) & 1) - 1) & 0x8984b493db9046ce) +
		((((x >> +2) & 1) - 1) & 0x9257b296b7128aa1) +
		((((x >> +3) & 1) - 1) & 0x2541b069313e8711) +
		((((x >> +4) & 1) - 1) & 0x7e124a220e93b2a9) +
		((((x >> +5) & 1) - 1) & 0xb3e2573e99a87d28) +
		((((x >> +6) & 1) - 1) & 0x52d5d851cc301368) +
		((((x >> +7) & 1) - 1) & 0x4fc24b628be15ef1) +
		((((x >> +8) & 1) - 1) & 0x98462271312ae033) +
		((((x >> +9) & 1) - 1) & 0x47d7121596e4a8a7) +
		((((x >> 10) & 1) - 1) & 0xf96889941a9427ef) +
		((((x >> 11) & 1) - 1) & 0xd1ae04ebc278f81e) +
		((((x >> 12) & 1) - 1) & 0x5f6392a366627e89) +
		((((x >> 13) & 1) - 1) & 0x3e8aae727dea34f5) +
		((((x >> 14) & 1) - 1) & 0xc77d4169f0959d11) +
		((((x >> 15) & 1) - 1) & 0x9972d93c1ac5ee97) +
		((((x >> 16) & 1) - 1) & 0x502267ac8dcec52d) +
		((((x >> 17) & 1) - 1) & 0xb8216acbbba21c84) +
		((((x >> 18) & 1) - 1) & 0x4560ec8a83409f86) +
		((((x >> 19) & 1) - 1) & 0x3783a6fac074e237) +
		((((x >> 20) & 1) - 1) & 0xf23928a936e9a255) +
		((((x >> 21) & 1) - 1) & 0x6971e1827421f5b8) +
		((((x >> 22) & 1) - 1) & 0x378cf244e4525bbd) +
		((((x >> 23) & 1) - 1) & 0x270334c7b5246eaf) +
		((((x >> 24) & 1) - 1) & 0xdb00f461834296f9) +
		((((x >> 25) & 1) - 1) & 0x865b916f3af6785c) +
		((((x >> 26) & 1) - 1) & 0x0ef31c7299afae18) +
		((((x >> 27) & 1) - 1) & 0x275cccc6bcfbe655) +
		((((x >> 28) & 1) - 1) & 0xb679ed17b169952a) +
		((((x >> 29) & 1) - 1) & 0x1d9735290203b7bd) +
		((((x >> 30) & 1) - 1) & 0xa28fa59ab03336a0) +
		((((x >> 31) & 1) - 1) & 0x223f894ab56a927d) +
		((((x >> 32) & 1) - 1) & 0x7525cef820c1d003) +
		((((x >> 33) & 1) - 1) & 0x7a8a4b20547a16e1) +
		((((x >> 34) & 1) - 1) & 0x60062881f00946f2) +
		((((x >> 35) & 1) - 1) & 0x2900e59c29960321) +
		((((x >> 36) & 1) - 1) & 0xaab883262559d99b) +
		((((x >> 37) & 1) - 1) & 0x143842c43dd6d612) +
		((((x >> 38) & 1) - 1) & 0x3bfc780b1279edb7) +
		((((x >> 39) & 1) - 1) & 0x35894032901336ca) +
		((((x >> 40) & 1) - 1) & 0xda5c1457870482fe) +
		((((x >> 41) & 1) - 1) & 0x3eb30aa435cd58b8) +
		((((x >> 42) & 1) - 1) & 0x2483a66644cd97b6) +
		((((x >> 43) & 1) - 1) & 0x67f721ef981c768f) +
		((((x >> 44) & 1) - 1) & 0x574e495f1519c828) +
		((((x >> 45) & 1) - 1) & 0x88effc0102de79c1) +
		((((x >> 46) & 1) - 1) & 0xf348c37ad4130787) +
		((((x >> 47) & 1) - 1) & 0x6ff5a240eb06d7ce) +
		((((x >> 48) & 1) - 1) & 0xf4ea2abcf4a01533) +
		((((x >> 49) & 1) - 1) & 0x91ca718ac8d4dbfe) +
		((((x >> 50) & 1) - 1) & 0xdeb03e2f8404e0c7) +
		((((x >> 51) & 1) - 1) & 0x733a8f26baec9dcb) +
		((((x >> 52) & 1) - 1) & 0x58c69a5d9c02795a) +
		((((x >> 53) & 1) - 1) & 0xfb69b32a68334a13) +
		((((x >> 54) & 1) - 1) & 0x9e9137026c76a891) +
		((((x >> 55) & 1) - 1) & 0x3c990e2b9c74bc78) +
		((((x >> 56) & 1) - 1) & 0x59939482968310b0) +
		((((x >> 57) & 1) - 1) & 0xb8172843814db2fb) +
		((((x >> 58) & 1) - 1) & 0x0d403ab99a734832) +
		((((x >> 59) & 1) - 1) & 0x8cf7b79d0d9d45c8) +
		((((x >> 60) & 1) - 1) & 0x87ef5765487e2285) +
		((((x >> 61) & 1) - 1) & 0xd220b2fa846a281d) +
		((((x >> 62) & 1) - 1) & 0x0450a5a5d76cb644) +
		((((x >> 63) & 1) - 1) & 0x536dd65bd43624d8) +
	0)



	return o % uint64(modulo)
*/
}

//// VACUUM HASHTABLE //////////////////////////////////////////////////////////

// vacuum a hashtable
// this procedure is not finished, it fails on a hashtable with a deleted items



type vacuumer func(int,int)

func (f vacuumer) Table(hash func(*, int) uint64, table [], accuracy byte) (displaced int) {

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

		f(-1, nexslot)
		*void = *posptr

		var src *;
		var mynexslot = nexslot
	again2:
			displaced++
		posptr = &table[mynexslot]
		src = &table[buf[0][0]]
		f(mynexslot, buf[0][0])
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


		f(buf[0][0],-1)
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
			f(-1, nexslot)
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

			f(target, nextarget)
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

			f(target, -1)
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

		f(-1, nexslot)
		*void = *posptr

		var src *;
		var mynexslot = nexslot
	again3:
			displaced++
		posptr = &table[mynexslot]
		src = &table[buf[i][0]]
		f(mynexslot, buf[i][0])
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


		f(buf[i][0],-1)
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
			f(-1, nexslot)
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

			f(target, nextarget)
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

			f(target, -1)
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
		return vacuumer( func(dst int, src int) {

		})
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
		return func(int,int){}
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


	for i := slot + 1 ; (int(i) < len(table)) && (i < slot + hops); i++ {

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

		d(int(slot), int(i))

		*location = *item
		*item = *void


		slot = i
		location = &table[slot]
		continue

		}
	}
}

////////////////////////////////////////////////////////////////////////////////

func main() {

	var keys [][8]byte
	var values [][2]uintptr
	_ = values


	var test = [2]uintptr{1337,7331}

	insert(&test, &values).into(hashfun, &[8]byte{2,6,7,8,3,2,1,0}, &keys)
	insert(&test, &values).into(hashfun, &[8]byte{2,6,7,8,3,2,1,0}, &keys)
	insert(&test, &values).into(hashfun, &[8]byte{2,6,7,8,3,2,1,1}, &keys)
	insert(&test, &values).into(hashfun, &[8]byte{2,6,7,8,3,2,1,2}, &keys)
	insert(&test, &values).into(hashfun, &[8]byte{2,6,7,8,3,2,1,3}, &keys)
	insert(&test, &values).into(hashfun, &[8]byte{2,6,7,8,3,2,1,4}, &keys)
	insert(&test, &values).into(hashfun, &[8]byte{1,3,8,5,4,9,5,3}, &keys)

	for i := 0; i < 250; i++ {

		var oldsize = len(keys)

		insert(&test, &values).into(hashfun, &[8]byte{1,3,8,5,4,9,5,byte(i)}, &keys)

		if oldsize != len(keys) {

			print("Reallocated to size ");
			println(len(keys))
		}
	}



	println("Here.")

	forEachKey(hashfun, keys, func(key *[8]byte, i int) {

		print(i)

		print(" ")
		print(values[i][0])
		print("  ")
	})
	println("")
	println("Searching.")

	var a *[2]uintptr = slect(hashfun, &[8]byte{1,3,8,5,4,9,5,20}, keys).from(values)

	print((*a)[0])
	println((*a)[1])

	if nil == slect(hashfun, &[8]byte{0,0,0,0,1,0,0,1}, keys).from(values) {
		println("not found")
	}

	///////////

	println("///////")

	var keyz []int
	var vals []int

	for i := 5; i < 1000; i++ {

	var lenkeys = len(keyz)
	_ = lenkeys

	var j int = i

	insert(&j, &vals).into(hashfn, &j, &keyz)

	if lenkeys != len(keyz) {


	print(i-4)
	print("|")
	print(len(keyz))
	print("|")



//	for j :=range keyz {
//		print(keyz[j])
//		print(" ")
//	}

		println("--")
	}
	}

	println(vacuum(vals).Table(hashfn, keyz, 4))

	delet(vals).from(hashfn, &[]int{10}[0], keyz)

	for i := 5; i < 15; i++ {
		var v *int = slect(hashfn, &i, keyz).from(vals)
		print(i)
		print(" ")
		if v == nil {println("nil");} else {
		println(*v)
		}
	}


	///////////
/*
	println("///////")

	var gomap = make( map[int]int)

	for i := 2; i < 1000000; i++ {

		gomap[i] = i
	}

	for j := 500000; j < 1000000; j++ {
		delete(gomap, j)
	}

	for j := 500000; j < 1439470; j++ {

		gomap[j] = j

	}

	var total = 0

	for q := 0; q < 10; q++ {
	for i := 2; i < 1439470; i++ {


		var a, ok = gomap[i]

		if !ok {
			println(i)
		}

		total += a

	}
	}

	println(total)

	return
*/
	///////////

	println("///////")

	var kval []keyval

	for i := 2; i < 1000000; i++ {

	var lenkeys = len(kval)
	_ = lenkeys

	insert(nil, &kval).into(kvhfunc, &keyval{i,i}, &kval)

	if lenkeys != len(kval) {


	print(i-1)
	print("|")
	print(len(kval))
	print("|")
	println("--")
	}

	}
/*
	for i := range kval {
		print("[")
		print(i)
		print("][")
		print(kvhfunc(&kval[i], len(kval)-hops))
		print("]")
		println(kval[i].key)
	}
*/
/*
	slect(kvhfunc, &keyval{key:11}, kval).Remove(kval)

	for i := 2; i < 300; i++ {


		var a *keyval = slect(kvhfunc, &keyval{key:i}, kval).from(kval)

		if a == nil {
		println("- -")
		continue
		}

		print(a.key)
		print(" ")
		println(a.val)
	}
*/
	// VACUUM


//	println(vacuum([]struct{}{}).Table(kvhfunc, kval))
//	println(vacuum([]struct{}{}).Table(kvhfunc, kval))

/*
	for i := range kval {
		print("[")
		print(i)
		print("][")
		print(kvhfunc(&kval[i], len(kval)-hops))
		print("]")
		println(kval[i].key)
	}
*/


	for j := 500000; j < 1000000; j++ {

		delet([]struct{}{}).from(kvhfunc, &keyval{key:j}, kval)
	}

	println(vacuum([]struct{}{}).Table(kvhfunc, kval, 1))
/*	{
		var tmp []keyval

		forEachKey(kvhfunc, kval, func(p *keyval, o int) {
			insert(nil, &tmp).into(kvhfunc, p, &tmp)
		})
		kval = tmp
		tmp = nil
	}
*/
	for j := 500000; j < 1439470; j++ {

	insert(nil, &kval).into(kvhfunc, &keyval{j,j}, &kval)

	}

	println(vacuum([]struct{}{}).Table(kvhfunc, kval, 4))
/*
	{
		var tmp []keyval

		forEachKey(kvhfunc, kval, func(p *keyval, o int) {
			insert(nil, &tmp).into(kvhfunc, p, &tmp)
		})
		kval = tmp
		tmp = nil
	}
*/
	var key keyval

	for q := 0; q < 10; q++ {
	for i := 2; i < 1439470; i++ {

		key.key = i

		var a *keyval = fetch(kvhfunc, &key, kval)

		if a == nil {
			println(i)
		}

	}
	}
/*
	println("///")

	////
	for i := 1; i > 0; i = grow(i) {
		print(i)
		print(" ")
		println(byte((((((((i-1)<<1)+19)&0x1fff)^104)%29))))
	}
*/
}
