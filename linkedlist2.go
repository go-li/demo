package main

// creates a new list header
func Mklist(link func(*, byte)(**), list *) {
	*link(list, 0) = list
	*link(list, 1) = list
}

// inserts item to list
func Insert(link func(*, byte)(**), list *, elm *) {
	*link(elm, 0) = list
	*link(elm, 1) = *link(list, 1)
	*link(list, 1) = elm
	*link(*link(elm, 1), 0) = elm
}

// remove item from list.
func Remove(link func(*, byte)(**), elm *) {
	*link(*link(elm, 0), 1) = *link(elm, 1)
	*link(*link(elm, 1), 0) = *link(elm, 0)
}

// is list empty?
func Empty(link func(*, byte)(**), list *) bool {
	return *link(list, 1) == list;
}

// do count of the list items
func Len(link func(*, byte)(**), list *) (count int) {
	e := *link(list, 1)
	for e != list {
		e = *link(e, 1)
		count++;
	}
	return count
}

////////////////////
type baz struct {
	primesprevnext [2]*baz
	allprevnext [2]*baz
	somevalue int
}

func primelink(w *baz, b byte) (lr **baz) {
	return &(w.primesprevnext[b])
}

func alllink(w *baz, b byte) (lr **baz) {
	return &(w.allprevnext[b])
}
///////////////////////////////
func main() {
	var u, v, w, x, y, z, header baz
	u.somevalue = 13
	v.somevalue = 11
	w.somevalue = 15
	x.somevalue = 14
	y.somevalue = 19
	z.somevalue = 12

	// head object serves as a beginning, entry point of the 2 lists
	Mklist(primelink, &header)

//add objects to prime list
	Insert(primelink, &header, &z)
	Insert(primelink, &header, &u)
	Insert(primelink, &header, &v)
	Insert(primelink, &header, &y)
	Remove(primelink, &z)
	var primes = Len(primelink, &u)

// initialize and add objects to another list
	Mklist(alllink, &header)
	Insert(alllink, &header, &u)
	Insert(alllink, &header, &v)
	Insert(alllink, &header, &w)
	Insert(alllink, &header, &x)
	Insert(alllink, &header, &y)
	Insert(alllink, &header, &z)

// count objects in lists
	var all = Len(alllink, &header)
	print(primes)
	print(" primes\n all:")
	print(all)
}
