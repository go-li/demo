package main

func pad(v []byte) []byte {
	var o [31]byte
	for i := 0; i < 31; i++ {
		o[i] = ' '
	}
	for i := 0; (i < 31) && (i < len(v)); i++ {
		o[i] = v[i]
	}
	return o[:]
}

func bytecompare(a, b []byte) (o int) {
	o = 0
	for i := 0; i < 31; i++ {
		o = int(a[i]) - int(b[i])
		if o != 0 {
			return o
		}
	}
	return 0
}

func keyer(links *[3]*, key []byte) []byte {
	return key
}

func linker(links *[3]*, key []byte) *[3]* {
	return links
}

func apend(node *, tree func(node *) (links *[3]*, key []byte), key []byte, result **) {
	var y * ;   /* Top node to update balance factor, and parent. */
	var p * ;
	var q * ; /* Iterator, and parent. */
	var n * ;    /* Newly inserted node. */
	var w * ;    /* New root of rebalanced subtree. */

	var dir int /* Direction to descend. */

	y = (*linker(tree((*)(nil))))[0]

	q = nil
	for p = (*linker(tree((*)(nil))))[0]; p != nil; {
		var cmp int = bytecompare(key[:31], keyer(tree(p))[:31])
		if cmp == 0 {
			if result != nil {
				*result = p
			}
		}
		if cmp > 0 {
			dir = 1
		} else {
			dir = 0
		}

		if keyer(tree(p))[31] != 0 {
			y = p
		}

		q = p
		p = linker(tree(p))[dir]
	}

	n = node
	//	if n == nil {
	//		return nil
	//	}

	for i := 0; i < 31; i++ {
		keyer(tree(n))[i] = key[i]
	}

	linker(tree(n))[0] = nil
	linker(tree(n))[1] = nil
	linker(tree(n))[2] = q

	if q != nil {
		linker(tree(q))[dir] = n
	} else {
		(*linker(tree((*)(nil))))[0] = n
	}
	keyer(tree(n))[31] = 0
	if (*linker(tree((*)(nil))))[0] == n {
		if result != nil {
			*result = n
		}
		return
	}

	for p = n; p != y; p = q {
		q = linker(tree(p))[2]
		if linker(tree(q))[0] != p {
			dir = 1
		} else {
			dir = 0
		}
		if dir == 0 {
			keyer(tree(q))[31]--
		} else {
			keyer(tree(q))[31]++
		}
	}

	if keyer(tree(y))[31] == 254 {
		var x *;
		x = linker(tree(y))[0]
		if keyer(tree(x))[31] == 255 {
			w = x
			linker(tree(y))[0] = linker(tree(x))[1]
			linker(tree(x))[1] = y
			keyer(tree(x))[31] = 0
			keyer(tree(y))[31] = 0
			linker(tree(x))[2] = linker(tree(y))[2]
			linker(tree(y))[2] = x
			if linker(tree(y))[0] != nil {
				linker(tree(linker(tree(y))[0]))[2] = y
			}
		} else {
			if !(keyer(tree(x))[31] == +1) {
				panic("assert")
			}
			w = linker(tree(x))[1]
			linker(tree(x))[1] = linker(tree(w))[0]
			linker(tree(w))[0] = x
			linker(tree(y))[0] = linker(tree(w))[1]
			linker(tree(w))[1] = y
			if keyer(tree(w))[31] == 255 {
				keyer(tree(x))[31] = 0
				keyer(tree(y))[31] = +1
			} else if keyer(tree(w))[31] == 0 {
				keyer(tree(x))[31] = 0
				keyer(tree(y))[31] = 0
			} else { /* |keyer(tree(w))[31] == +1| */
				keyer(tree(x))[31] = 255
				keyer(tree(y))[31] = 0
			}
			keyer(tree(w))[31] = 0
			linker(tree(w))[2] = linker(tree(y))[2]
			linker(tree(x))[2] = w
			linker(tree(y))[2] = w
			if linker(tree(x))[1] != nil {
				linker(tree(linker(tree(x))[1]))[2] = x
			}
			if linker(tree(y))[0] != nil {
				linker(tree(linker(tree(y))[0]))[2] = y
			}
		}
	} else if keyer(tree(y))[31] == +2 {
		var x *;
		x = linker(tree(y))[1]
		if keyer(tree(x))[31] == +1 {
			w = x
			linker(tree(y))[1] = linker(tree(x))[0]
			linker(tree(x))[0] = y
			keyer(tree(x))[31] = 0
			keyer(tree(y))[31] = 0
			linker(tree(x))[2] = linker(tree(y))[2]
			linker(tree(y))[2] = x
			if linker(tree(y))[1] != nil {
				linker(tree(linker(tree(y))[1]))[2] = y
			}
		} else {
			if !(keyer(tree(x))[31] == 255) {
				panic("assert")
			}
			w = linker(tree(x))[0]
			linker(tree(x))[0] = linker(tree(w))[1]
			linker(tree(w))[1] = x
			linker(tree(y))[1] = linker(tree(w))[0]
			linker(tree(w))[0] = y
			if keyer(tree(w))[31] == +1 {
				keyer(tree(x))[31] = 0
				keyer(tree(y))[31] = 255
			} else if keyer(tree(w))[31] == 0 {
				keyer(tree(x))[31] = 0
				keyer(tree(y))[31] = 0
			} else { /* |keyer(tree(w))[31] == 255| */
				keyer(tree(x))[31] = +1
				keyer(tree(y))[31] = 0
			}
			keyer(tree(w))[31] = 0
			linker(tree(w))[2] = linker(tree(y))[2]
			linker(tree(x))[2] = w
			linker(tree(y))[2] = w
			if linker(tree(x))[0] != nil {
				linker(tree(linker(tree(x))[0]))[2] = x
			}
			if linker(tree(y))[1] != nil {
				linker(tree(linker(tree(y))[1]))[2] = y
			}
		}
	} else {
		if result != nil {
			*result = n
		}
		return
	}
	if linker(tree(w))[2] != nil {
		var oo = 0
		if y != linker(tree(linker(tree(w))[2]))[0] {
			oo = 1
		}
		linker(tree(linker(tree(w))[2]))[oo] = w
	} else {
		(*linker(tree((*)(nil))))[0] = w
	}

	if result != nil {
		*result = n
	}
}


/* Deletes from |tree| and returns an item matching |item|.
   Returns a null pointer if no matching item found. */
func remove(tree func(node *) (links *[3]*, key []byte), key []byte, result **) {

	var p *; /* Traverses tree to find node to delete. */
	var q *; /* Parent of |p|. */
	var dir int /* Side of |q| on which |p| is linked. */

	if !(tree != nil) {
		panic("assert")
	}

	if (*linker(tree((*)(nil))))[0] == nil {
		if result != nil {
			*result = nil
		}
		return
	}

	p = (*linker(tree((*)(nil))))[0]
	for {
		var cmp int = bytecompare(key[:31], keyer(tree(p))[:31])
		if cmp == 0 {
			break
		}

		if cmp > 0 {
			dir = 1
		} else {
			dir = 0
		}

		p = linker(tree(p))[dir]
		if p == nil {
			if result != nil {
				*result = nil
			}
			return
		}
	}
	if result != nil {
		*result = p
	}

	q = linker(tree(p))[2]
	if q == nil {
		q = (*linker(tree((*)(nil))))[0]
		dir = 0
	}

	if linker(tree(p))[1] == nil {
		linker(tree(q))[dir] = linker(tree(p))[0]
		if linker(tree(q))[dir] != nil {
			linker(tree(linker(tree(q))[dir]))[2] = linker(tree(p))[2]
		}
	} else {
		var r *;
		r = linker(tree(p))[1]
		if linker(tree(r))[0] == nil {
			linker(tree(r))[0] = linker(tree(p))[0]
			linker(tree(q))[dir] = r
			linker(tree(r))[2] = linker(tree(p))[2]
			if linker(tree(r))[0] != nil {
				linker(tree(linker(tree(r))[0]))[2] = r
			}
			keyer(tree(r))[31] = keyer(tree(p))[31]
			q = r
			dir = 1
		} else {
			var s *;
			s = linker(tree(r))[0]
			for linker(tree(s))[0] != nil {
				s = linker(tree(s))[0]
			}
			r = linker(tree(s))[2]
			linker(tree(r))[0] = linker(tree(s))[1]
			linker(tree(s))[0] = linker(tree(p))[0]
			linker(tree(s))[1] = linker(tree(p))[1]
			linker(tree(q))[dir] = s
			if linker(tree(s))[0] != nil {
				linker(tree(linker(tree(s))[0]))[2] = s
			}
			linker(tree(linker(tree(s))[1]))[2] = s
			linker(tree(s))[2] = linker(tree(p))[2]
			if linker(tree(r))[0] != nil {
				linker(tree(linker(tree(r))[0]))[2] = r
			}
			keyer(tree(s))[31] = keyer(tree(p))[31]
			q = r
			dir = 0
		}
	}

	linker(tree(p))[0] = nil
	linker(tree(p))[1] = nil
	linker(tree(p))[2] = nil
	keyer(tree(p))[31] = 0

	//  tree.pavl_alloc.libavl_free (tree.pavl_alloc, p);

	for q != (*linker(tree((*)(nil))))[0] {
		var y *;
		y = q

		if linker(tree(y))[2] != nil {
			q = linker(tree(y))[2]
		} else {
			q = (*linker(tree((*)(nil))))[0]
		}

		if dir == 0 {
			if linker(tree(q))[0] != y {
				dir = 1
			} else {
				dir = 0
			}
			keyer(tree(y))[31]++
			if keyer(tree(y))[31] == +1 {
				break
			} else if keyer(tree(y))[31] == +2 {
				var x *;
				x = linker(tree(y))[1]
				if keyer(tree(x))[31] == 255 {
					var w *;

					if !(keyer(tree(x))[31] == 255) {
						panic("assert")
					}
					w = linker(tree(x))[0]
					linker(tree(x))[0] = linker(tree(w))[1]
					linker(tree(w))[1] = x
					linker(tree(y))[1] = linker(tree(w))[0]
					linker(tree(w))[0] = y
					if keyer(tree(w))[31] == +1 {
						keyer(tree(x))[31] = 0
						keyer(tree(y))[31] = 255
					} else if keyer(tree(w))[31] == 0 {
						keyer(tree(x))[31] = 0
						keyer(tree(y))[31] = 0
					} else { /* |keyer(tree(w))[31] == 255| */
						keyer(tree(x))[31] = +1
						keyer(tree(y))[31] = 0
					}
					keyer(tree(w))[31] = 0
					linker(tree(w))[2] = linker(tree(y))[2]
					linker(tree(x))[2] = w
					linker(tree(y))[2] = w
					if linker(tree(x))[0] != nil {
						linker(tree(linker(tree(x))[0]))[2] = x
					}
					if linker(tree(y))[1] != nil {
						linker(tree(linker(tree(y))[1]))[2] = y
					}
					linker(tree(q))[dir] = w
				} else {
					linker(tree(y))[1] = linker(tree(x))[0]
					linker(tree(x))[0] = y
					linker(tree(x))[2] = linker(tree(y))[2]
					linker(tree(y))[2] = x
					if linker(tree(y))[1] != nil {
						linker(tree(linker(tree(y))[1]))[2] = y
					}
					linker(tree(q))[dir] = x
					if keyer(tree(x))[31] == 0 {
						keyer(tree(x))[31] = 255
						keyer(tree(y))[31] = +1
						break
					} else {
						keyer(tree(x))[31] = 0
						keyer(tree(y))[31] = 0
						y = x
					}
				}
			}
		} else {
			if linker(tree(q))[0] != y {
				dir = 1
			} else {
				dir = 0
			}
			keyer(tree(y))[31]--
			if keyer(tree(y))[31] == 255 {
				break
			} else if keyer(tree(y))[31] == 254 {
				var x *;
				x = linker(tree(y))[0]
				if keyer(tree(x))[31] == +1 {
					var w *;
					if !(keyer(tree(x))[31] == +1) {
						panic("assert")
					}
					w = linker(tree(x))[1]
					linker(tree(x))[1] = linker(tree(w))[0]
					linker(tree(w))[0] = x
					linker(tree(y))[0] = linker(tree(w))[1]
					linker(tree(w))[1] = y
					if keyer(tree(w))[31] == 255 {
						keyer(tree(x))[31] = 0
						keyer(tree(y))[31] = +1
					} else if keyer(tree(w))[31] == 0 {
						keyer(tree(x))[31] = 0
						keyer(tree(y))[31] = 0
					} else { /* |keyer(tree(w))[31] == +1| */
						keyer(tree(x))[31] = 255
						keyer(tree(y))[31] = 0
					}
					keyer(tree(w))[31] = 0
					linker(tree(w))[2] = linker(tree(y))[2]
					linker(tree(x))[2] = w
					linker(tree(y))[2] = w
					if linker(tree(x))[1] != nil {
						linker(tree(linker(tree(x))[1]))[2] = x
					}
					if linker(tree(y))[0] != nil {
						linker(tree(linker(tree(y))[0]))[2] = y
					}
					linker(tree(q))[dir] = w
				} else {
					linker(tree(y))[0] = linker(tree(x))[1]
					linker(tree(x))[1] = y
					linker(tree(x))[2] = linker(tree(y))[2]
					linker(tree(y))[2] = x
					if linker(tree(y))[0] != nil {
						linker(tree(linker(tree(y))[0]))[2] = y
					}
					linker(tree(q))[dir] = x
					if keyer(tree(x))[31] == 0 {
						keyer(tree(x))[31] = +1
						keyer(tree(y))[31] = 255
						break
					} else {
						keyer(tree(x))[31] = 0
						keyer(tree(y))[31] = 0
						y = x
					}
				}
			}
		}
	}

	return

}

func preorder(node *, tree func(node *) (links *[3]*, key []byte), callback func(*)) {
	if node != nil {
		preorder(linker(tree(node))[0], tree, callback)
		callback(node)
		preorder(linker(tree(node))[1], tree, callback)
	}
}

func previsit(unused *, tree func(node *) (links *[3]*, key []byte), callback func(*)) {
	preorder((*linker(tree((*)(nil))))[0], tree, callback)
}

//-GluE CodE------------------------

type node struct {
	pavl_link [3]*node /* Subtrees. */
	keybal    [32]byte
	pavl_data * /* value itself. */
}

func probe(value *, tree **node, key []byte, result **node) {
	var pseudoRoot node
	pseudoRoot.pavl_link[0] = *tree

	var newlyCreated *node = &node{pavl_data: value}

	var RootLinkKeyer = func(node *node) (*[3]*node, []byte) {
		if (node == nil) {
			return &(pseudoRoot.pavl_link), nil
		}
		return &node.pavl_link, node.keybal[:]
	}

	apend(newlyCreated, RootLinkKeyer, key, result)

	*tree = pseudoRoot.pavl_link[0]

	_ = RootLinkKeyer
}

func drop(tree **node, key []byte, result **node) {

	if (nil == *tree) {
		if (result != nil) {
			*result = nil
		}
		return
	}

	var pseudoRoot node
	pseudoRoot.pavl_link[0] = *tree

	var RootLinkKeyer = func(node *node) (links *[3]*node, key []byte) {
		if (node == nil) {
			return &pseudoRoot.pavl_link, nil
		}
		return &node.pavl_link, node.keybal[:]
	}

	remove(RootLinkKeyer, key, result)

	*tree = pseudoRoot.pavl_link[0]

	_ = RootLinkKeyer

}

// Visits tree values in sequence
func inorder(unused *, tree **node, visit func(*)) {
	var pseudoRoot node
	pseudoRoot.pavl_link[0] = *tree

	var RootLinkKeyer = func(node *node) (links *[3]*node, key []byte) {
		if (node == nil) {
			return &pseudoRoot.pavl_link, nil
		}
		return &node.pavl_link, node.keybal[:]
	}


	var Visitor = func(node *node) {
		visit(node.pavl_data)
	}

	previsit(&pseudoRoot, RootLinkKeyer, Visitor)

}


//-ClienT C0DE--------------------

type myValue struct {
	str string
}

type stringnode struct {
	pavl_link [3]*stringnode /* Subtrees. */
	keybal    [32]byte
	pavl_data *myValue /* value itself. */
}



func main() {

	var root *stringnode

	probe(&myValue{"Paul Sartorius"}, &root, pad([]byte("composer")), nil)
	probe(&myValue{"Elkanah Settle"}, &root, pad([]byte("writer")), nil)
	probe(&myValue{"Edie Martin"}, &root, pad([]byte("actress")), nil)
	probe(&myValue{"Walter de Stapledon"}, &root, pad([]byte("bishop")), nil)
	probe(&myValue{"Blake Ross"}, &root, pad([]byte("developer")), nil)
	probe(&myValue{"Cicely Saunders"}, &root, pad([]byte("nurse")), nil)
	probe(&myValue{"Bob Sweikert"}, &root, pad([]byte("driver")), nil)
	probe(&myValue{"Peter Godfrey"}, &root, pad([]byte("accountant")), nil)
	probe(&myValue{"Pam Beesley"}, &root, pad([]byte("receptionist")), nil)
	probe(&myValue{"Valdemar Poulsen"}, &root, pad([]byte("engineer")), nil)
	probe(&myValue{"Bucky Harris"}, &root, pad([]byte("manager")), nil)
	probe(&myValue{"Carl W Scheele"}, &root, pad([]byte("pharmacist")), nil)
	probe(&myValue{"Tim Cook"}, &root, pad([]byte("cook")), nil)
	probe(&myValue{"David Griswold"}, &root, pad([]byte("cashier")), nil)

	inorder(&myValue{}, &root, func(value *myValue) {
		print(value.str)
		print("\n")
	})
	print("\n")

	drop(&root, pad([]byte("accountant")), nil)
	drop(&root, pad([]byte("developer")), nil)

	inorder(nil, &root, func(value *myValue) {
		print(value.str)
		print("\n")
	})

}
