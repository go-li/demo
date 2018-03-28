// A doubly linked list of the same typed items
package main

const Prev = 0
const Next = 1

func insert (list **, link func(*) *[2]*, elm *) {
	if nil == *list {
		*list = elm
		(*link(elm))[Prev] = elm
		(*link(elm))[Next] = elm

	} else if (*link(elm))[Prev] == nil {
		(*link(elm))[Prev] = *list
		(*link(elm))[Next] = (*link(*list))[Next]
		(*link(*list))[Next] = elm
		(*link( (*link(elm))[Next] ))[Prev] = elm
	} else {
		panic("One link cannot be in two lists")
	}
}

// add adds element to a list another element is already member of
func add(already *, link func(*) *[2]*, elm *) {
	if (*link(already))[Prev] == nil || (*link(already))[Next] == nil {
		panic("Already is not already in the list")
	}
	(*link(elm))[Prev] = already
	(*link(elm))[Next] = (*link(already))[Next]
	(*link(already))[Next] = elm
	(*link( (*link(elm))[Next] ))[Prev] = elm
}

func remove(list **, link func(*) *[2]*, elm *) {
	if *list == elm {
		if (*link(elm))[Prev] == elm {
			*list = nil
			goto finally
		} else {
			*list = (*link(elm))[Next]
		}
	}
	(*link( (*link(elm))[Prev] ))[Next] = (*link(elm))[Next]
	(*link( (*link(elm))[Next] ))[Prev] = (*link(elm))[Prev]
finally:
	(*link(elm))[Prev] = nil
	(*link(elm))[Next] = nil
}

func empty(list **, link func(*) *[2]*) bool {
	return nil == *list
}

// do count of the list items
func length(list **, link func(*) *[2]*) (count int) {

	if nil == *list {
		return 0
	}

	var e *;
	e = (*link(*list))[Next]

	for e != *list {
		e = (*link(e))[Next]
		count++;
	}
	count++
	return count
}

// apply function to all link elements
func foreach(direction byte, list **, link func(*) *[2]*, f func(*)) {

	if nil == *list {
		return
	}

	var end *;
	end = *list

	var e *;
	e = (*link(*list))[direction]
	f(end)

	for (e != end) && ((*link(e))[direction] != nil) {
		var newe = (*link(e))[direction]
		f(e)
		e = newe
	}
	return
}


// USER-LOOSER WRITTEN CODE

type item struct {
	link [2]*item
}


func item_link(x *item) *[2]*item {
	return &(x.link)
}

func main() {

	var a *item
	var b item
	var c item
	var d item

	print("len=");println(length(&a, item_link))

	insert(&a, item_link, &b)

	print("len=");println(length(&a, item_link))
	print("a=");println(a)

	insert(&a, item_link, &c)

	print("len=");println(length(&a, item_link))
	print("a=");println(a)

	add(&c, item_link, &d)

	print("len=");println(length(&a, item_link))
	print("a=");println(a)

	remove(&a, item_link, &d)

	print("len=");println(length(&a, item_link))
	print("a=");println(a)

	remove(&a, item_link, &b)

	print("len=");println(length(&a, item_link))
	print("a=");println(a)

	remove(&a, item_link, &c)

	print("len=");println(length(&a, item_link))
	print("a=");println(a)

	insert(&a, item_link, &b)

	print("len=");println(length(&a, item_link))
	print("a=");println(a)

	insert(&a, item_link, &c)

	print("len=");println(length(&a, item_link))
	print("a=");println(a)

	add(&c, item_link, &d)

	print("len=");println(length(&a, item_link))
	print("a=");println(a)

	foreach(Next, &a, item_link, func(i *item) {
		print("i=");println(i)
	})

	// Evacuate the list
	foreach(Prev, &a, item_link, func(i *item) {
		remove(&a, item_link, i)
	})

	print("len=");println(length(&a, item_link))
	print("a=");println(a)
}
