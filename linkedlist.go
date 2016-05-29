package main

//////////////////////////////////////////////////////////////
const Next = byte(0)
const Prev = byte(1)

/*
//CALLBACK SIGNATURES:
set func(*, byte, *)
get func(*, byte) (*)
*/

// creates a new list
func Mklist(set func(*, byte, *), list *) {
	set(list, Prev, list)
	set(list, Next, list)
}

// inserts item to list
func Insert(set func(*, byte, *), get func(*, byte)(*), list *, elm *) {
	set(elm, Prev, list)
	set(elm, Next, get(list, Next))
	set(list, Next, elm)
	set(get(elm, Next), Prev, elm)
}

// remove item from list.
func Remove(set func(*, byte, *), get func(*, byte)(*), elm *) {
	set(get(elm, Prev), Next, get(elm, Next))
	set(get(elm, Next), Prev, get(elm, Prev))
}

// is list empty?
func Empty(get func(*, byte)(*), list *) bool {
	return get(list, Next) == list;
}

// do count of the list items
func Len(get func(*, byte)(*), list *) (count int) {
	e := get(list, Next)
	for e != list {
		e = get(e, Next)
		count++;
	}
	return count
}

//////////////////////////////////////////////////////////////

type foo struct {
	prev *foo
	next *foo
	foo int
}

//////////////////////////////////////////////////////////////
// begin of type specific linklist code

func set(l *foo, i byte, to *foo) {
	if i == Next {
		l.next = to
		return
	}
	l.prev = to
}

func get(l *foo, i byte) (link *foo) {
	if i == Next {
		return l.next
	}
	return l.prev
}
// end of type specific linklist code
//////////////////////////////////////////////////////////////

func main() {
	var x, y, z, w foo
	x = y
	z = w

	Mklist(set, &x)
	
	print(Empty(get, &x))
	Insert(set, get, &x, &y)

	print(Empty(get, &x))
	Insert(set, get, &x, &z)
	

	print(Empty(get, &x))
	Insert(set, get, &x, &w)


	print(Empty(get, &x))
	Remove(set, get, &y)
	
	print("hello")
}
