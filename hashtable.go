package main

type inserter func(int,int)
type selector int


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




func slect(hash func(*, int) uint64, key *, table []) selector {

	// next we look if the slot is occupied

	var slot = hash(key, len(table))

	var location *;
	location = &table[slot]

	var hloc = hash(location, -1)

	// if it's occupied by a key

	var hkey = hash(key, -1)

	if hloc == hkey {
		return selector(int(slot))
	}

	return -1
}

// Fast Search when you know the key exists
// If unsure, use Select
func fetch(hash func(*, int) uint64, key *, table []) selector {
	return selector(int(hash(key, len(table))))
}


func (f selector) from(values []) * {
	if f == -1 {
		return nil
	}

	return &values[f]
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
	return o % uint64(size)

}

func (f inserter) into(hash func(*, int) uint64, key *, table *[]) {

	const grow_by = 15


	if len(*table) == 0 {
		(*table) = make([], 1)
		var loc *;
		loc = &(*table)[0]
		*loc = *key

		f(1,-1)
		f(0,-2)
		f(0,-4)

		return
	}

	// next we look if the slot is occupied

	var slot = hash(key, len(*table))

	var location *;
	location = &(*table)[slot]

	var hloc = hash(location, -1)

	if hash(key, -1) == hloc {

		f(int(slot),-4)

		return
	}

	// we check if the slot is empty

	var empty [];
	empty = make([], 1)

	var void *;
	void = &empty[0]

	var hvoid = hash(void, -1)

	empty = nil

	if hvoid == hloc {

		*location = *key
		f(int(slot),-4)

		return
	}



	var oldsize = len(*table)
	var newsize = len(*table)+grow_by

	var keys [];
	keys = make([], newsize)

	f(newsize, -1)




	for i := 0; i < oldsize; i++ {

		var each *;
		each = &(*table)[i]

		var heach = hash(each, -1)

		if hvoid == heach {
			continue
		}

		heach = hash(each, newsize)

		var newkey *;
		newkey = &keys[heach]

		var hnewkey = hash(newkey, -1)

		if hvoid != hnewkey {
			newsize += grow_by
			keys = make([], newsize)
			f(newsize, -1)
			i = -1
			continue
		}

		*newkey = *each
		f(int(heach), i)
	}


	*table = keys
	keys = nil
	f(0,-2)


	f.into(hash, key, table)
}

func insert(value *, values *[]) inserter {

	var oldvalues []
	_ = oldvalues

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
		case emptyElement:
			var void [];
			void = make([], 1)
			_ = void
			srcp = &void[0]
			dstp = &(*values)[dst]
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

	var a = slect(hashfun, &[8]byte{1,3,8,5,4,9,5,20}, keys).from(values)

	var b *[2]uintptr = a

	print((*b)[0])
	println((*b)[1])

	_ = a

	if nil == slect(hashfun, &[8]byte{0,0,0,0,1,0,0,1}, keys).from(values) {
		println("not found")
	}

}
