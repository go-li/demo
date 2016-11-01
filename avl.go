package main

type Node struct {
	links [3]*Node
	keybal [32]byte
	value *
};

func Pad(v []byte) (o [31]byte) {
	for i:= 0; i < 31; i++ {
		o[i] = ' '
	}
	for i:= 0; (i < 31) && (i < len(v)); i++ {
//		o[31-i-1] = v[i]
		o[i] = v[i]
	}
	return o
}

func bytecompare(a, b []byte) (o int) {
	o = 0
	for i := 0; i < 31; i++ {
		o = int(a[i]) - int(b[i])
		if (o != 0) {
			return o
		}
	}
	return 0
}


func Probe(value *, tree **Node, key [31]byte, result **) {


	var y *Node    /* Top node to update balance factor, and parent. */
	var p, q *Node /* Iterator, and parent. */
	var n *Node    /* Newly inserted node. */
	var w *Node    /* New root of rebalanced subtree. */

	var dir int /* Direction to descend. */

	//  assert (tree != nil && item != nil);

	y = (*tree)
	q = nil
	for p = (*tree); p != nil; {
		var cmp int = bytecompare(p.keybal[:], key[:])
		if cmp == 0 {
			if (result != nil) {
				*result = p.value
			}
			return
		}
		if cmp > 0 {
			dir = 1
		} else {
			dir = 0
		}

		if p.keybal[31] != 0 {
			y = p
		}

		q = p
		p = p.links[dir]
	}

	n = &Node{}
	//	if n == nil {
	//		return nil
	//	}

	//n = key

//	tree.pavl_count++
	n.links[0] = nil
	n.links[1] = nil
	n.links[2] = q

//	valueswap(n,key, true)

	n.value = value
	for i:= 0; i < 31; i++ {
		n.keybal[i] = key[i]
	}

	//  n.pavl_data = item;
	if q != nil {
		q.links[dir] = n
	} else {
		(*tree) = n
	}
	n.keybal[31] = 0
	if (*tree) == n {
		if (result != nil) {
			*result = n.value
		}
		return
	}

	for p = n; p != y; p = q {
		q = p.links[2]
		if q.links[0] != p {
			dir = 1
		} else {
			dir = 0
		}
		if dir == 0 {
			q.keybal[31]--
		} else {
			q.keybal[31]++
		}
	}

	if y.keybal[31] == 254 {
		var x *Node = y.links[0]
		if x.keybal[31] == 255 {
			w = x
			y.links[0] = x.links[1]
			x.links[1] = y
			x.keybal[31] = 0
			y.keybal[31] = 0
			x.links[2] = y.links[2]
			y.links[2] = x
			if y.links[0] != nil {
				y.links[0].links[2] = y
			}
		} else {
			if !(x.keybal[31] == +1) {
				panic("assert")
			}
			w = x.links[1]
			x.links[1] = w.links[0]
			w.links[0] = x
			y.links[0] = w.links[1]
			w.links[1] = y
			if w.keybal[31] == 255 {
				x.keybal[31] = 0
				y.keybal[31] = +1
			} else if w.keybal[31] == 0 {
				x.keybal[31] = 0
				y.keybal[31] = 0
			} else { /* |w.keybal[31] == +1| */
				x.keybal[31] = 255
				y.keybal[31] = 0
			}
			w.keybal[31] = 0
			w.links[2] = y.links[2]
			x.links[2] = w
			y.links[2] = w
			if x.links[1] != nil {
				x.links[1].links[2] = x
			}
			if y.links[0] != nil {
				y.links[0].links[2] = y
			}
		}
	} else if y.keybal[31] == +2 {
		var x *Node = y.links[1]
		if x.keybal[31] == +1 {
			w = x
			y.links[1] = x.links[0]
			x.links[0] = y
			x.keybal[31] = 0
			y.keybal[31] = 0
			x.links[2] = y.links[2]
			y.links[2] = x
			if y.links[1] != nil {
				y.links[1].links[2] = y
			}
		} else {
			if !(x.keybal[31] == 255) {
				panic("assert")
			}
			w = x.links[0]
			x.links[0] = w.links[1]
			w.links[1] = x
			y.links[1] = w.links[0]
			w.links[0] = y
			if w.keybal[31] == +1 {
				x.keybal[31] = 0
				y.keybal[31] = 255
			} else if w.keybal[31] == 0 {
				x.keybal[31] = 0
				y.keybal[31] = 0
			} else { /* |w.keybal[31] == -1| */
				x.keybal[31] = +1
				y.keybal[31] = 0
			}
			w.keybal[31] = 0
			w.links[2] = y.links[2]
			x.links[2] = w
			y.links[2] = w
			if x.links[0] != nil {
				x.links[0].links[2] = x
			}
			if y.links[1] != nil {
				y.links[1].links[2] = y
			}
		}
	} else {
		if (result != nil) {
			*result = n.value
		}
		return
	}
	if w.links[2] != nil {
		var oo = 0
		if y != w.links[2].links[0] {
			oo = 1
		}
		w.links[2].links[oo] = w
	} else {
		*tree = w
	}

	if (result != nil) {
		*result = n.value
	}
	return
}

/* Deletes from |tree| and returns an item matching |item|.
   Returns a null pointer if no matching item found. */
func Drop(tree **Node, key [31]byte, result **) {


	var p *Node /* Traverses tree to find node to delete. */
	var q *Node /* Parent of |p|. */
	var dir int      /* Side of |q| on which |p| is linked. */

	if !(tree != nil) {
		panic("assert")
	}

	if tree == nil {
		*result = nil
		return
	}

	p = (*tree)
	for {
		var cmp int = bytecompare(p.keybal[:31], key[:])
		if cmp == 0 {
			break
		}

		if cmp > 0 {
			dir = 1
		} else {
			dir = 0
		}

		p = p.links[dir]
		if p == nil {
			*result = nil
			return
		}
	}
	if (result != nil) {
		*result = p.value
	}

	q = p.links[2]
	if q == nil {
		q = *tree
		dir = 0
	}

	if p.links[1] == nil {
		q.links[dir] = p.links[0]
		if q.links[dir] != nil {
			q.links[dir].links[2] = p.links[2]
		}
	} else {
		var r *Node = p.links[1]
		if r.links[0] == nil {
			r.links[0] = p.links[0]
			q.links[dir] = r
			r.links[2] = p.links[2]
			if r.links[0] != nil {
				r.links[0].links[2] = r
			}
			r.keybal[31] = p.keybal[31]
			q = r
			dir = 1
		} else {
			var s *Node = r.links[0]
			for s.links[0] != nil {
				s = s.links[0]
			}
			r = s.links[2]
			r.links[0] = s.links[1]
			s.links[0] = p.links[0]
			s.links[1] = p.links[1]
			q.links[dir] = s
			if s.links[0] != nil {
				s.links[0].links[2] = s
			}
			s.links[1].links[2] = s
			s.links[2] = p.links[2]
			if r.links[0] != nil {
				r.links[0].links[2] = r
			}
			s.keybal[31] = p.keybal[31]
			q = r
			dir = 0
		}
	}

	p.links[0] = nil
	p.links[1] = nil
	p.links[2] = nil
	p.keybal[31] = 0
	if (result != nil ){
		*result = p.value
	}
	p.value = nil

	//  tree.pavl_alloc.libavl_free (tree.pavl_alloc, p);

	for q != *tree {
		var y *Node = q

		if y.links[2] != nil {
			q = y.links[2]
		} else {
			q = *tree
		}

		if dir == 0 {
			if q.links[0] != y {
				dir = 1
			} else {
				dir = 0
			}
			y.keybal[31]++
			if y.keybal[31] == +1 {
				break
			} else if y.keybal[31] == +2 {
				var x *Node = y.links[1]
				if x.keybal[31] == 255 {
					var w *Node

					if !(x.keybal[31] == 255) {
						panic("assert")
					}
					w = x.links[0]
					x.links[0] = w.links[1]
					w.links[1] = x
					y.links[1] = w.links[0]
					w.links[0] = y
					if w.keybal[31] == +1 {
						x.keybal[31] = 0
						y.keybal[31] = 255
					} else if w.keybal[31] == 0 {
						x.keybal[31] = 0
						y.keybal[31] = 0
					} else { /* |w.keybal[31] == 255| */
						x.keybal[31] = +1
						y.keybal[31] = 0
					}
					w.keybal[31] = 0
					w.links[2] = y.links[2]
					x.links[2] = w
					y.links[2] = w
					if x.links[0] != nil {
						x.links[0].links[2] = x
					}
					if y.links[1] != nil {
						y.links[1].links[2] = y
					}
					q.links[dir] = w
				} else {
					y.links[1] = x.links[0]
					x.links[0] = y
					x.links[2] = y.links[2]
					y.links[2] = x
					if y.links[1] != nil {
						y.links[1].links[2] = y
					}
					q.links[dir] = x
					if x.keybal[31] == 0 {
						x.keybal[31] = 255
						y.keybal[31] = +1
						break
					} else {
						x.keybal[31] = 0
						y.keybal[31] = 0
						y = x
					}
				}
			}
		} else {
			if q.links[0] != y {
				dir = 1
			} else {
				dir = 0
			}
			y.keybal[31]--
			if y.keybal[31] == 255 {
				break
			} else if y.keybal[31] == 254 {
				var x *Node = y.links[0]
				if x.keybal[31] == +1 {
					var w *Node
					if !(x.keybal[31] == +1) {
						panic("assert")
					}
					w = x.links[1]
					x.links[1] = w.links[0]
					w.links[0] = x
					y.links[0] = w.links[1]
					w.links[1] = y
					if w.keybal[31] == 255 {
						x.keybal[31] = 0
						y.keybal[31] = +1
					} else if w.keybal[31] == 0 {
						x.keybal[31] = 0
						y.keybal[31] = 0
					} else { /* |w.keybal[31] == +1| */
						x.keybal[31] = 255
						y.keybal[31] = 0
					}
					w.keybal[31] = 0
					w.links[2] = y.links[2]
					x.links[2] = w
					y.links[2] = w
					if x.links[1] != nil {
						x.links[1].links[2] = x
					}
					if y.links[0] != nil {
						y.links[0].links[2] = y
					}
					q.links[dir] = w
				} else {
					y.links[0] = x.links[1]
					x.links[1] = y
					x.links[2] = y.links[2]
					y.links[2] = x
					if y.links[0] != nil {
						y.links[0].links[2] = y
					}
					q.links[dir] = x
					if x.keybal[31] == 0 {
						x.keybal[31] = +1
						y.keybal[31] = 255
						break
					} else {
						x.keybal[31] = 0
						y.keybal[31] = 0
						y = x
					}
				}
			}
		}
	}

//	tree.pavl_count--
	return

}

// Visits tree values in sequence
func Preorder(unused *, tree **Node, visit func(*)) {

	var p *Node /* Iterator. */

	p = *tree
	if (p == nil) {
		return
	}

	if (p.links[0] != nil) {

		var x = string(p.links[0].keybal[:31])
		print(x)
		visit(p.links[0].value)

		p = p.links[0]
		Preorder(unused, &p, visit)
	}

		var x = string(p.keybal[:31])
		print(x)
		visit(p.value)


	if (p.links[1] != nil) {

		var x = string(p.links[1].keybal[:31])
		print(x)
		visit(p.links[1].value)

		p = p.links[1]
		Preorder(unused, &p, visit)
	}
}


//---------------------

type MyValue struct {
	str	string
}

type StringNode struct {
	links [3]*Node
	keybal [32]byte
	value *MyValue
};



func main() {

	var root *StringNode



	Probe(&MyValue{"Paul Sartorius"}, &root, Pad([]byte("composer")), nil)



	Preorder(&MyValue{}, &root, func(value *MyValue){
		print(value.str)
		print("\n");
	})
		print("\n");

	Probe(&MyValue{"Elkanah Settle"}, &root, Pad([]byte("writer")), nil)


	Preorder(&MyValue{}, &root, func(value *MyValue){
		print(value.str)
		print("\n");
	})
		print("\n");


	Probe(&MyValue{"Edie Martin"}, &root, Pad([]byte("actress")), nil)




	Preorder(&MyValue{}, &root, func(value *MyValue){
		print(value.str)
		print("\n");
	})
		print("\n");



	Probe(&MyValue{"Walter de Stapledon"}, &root, Pad([]byte("bishop")), nil)
	Probe(&MyValue{"Blake Ross"}, &root, Pad([]byte("developer")), nil)
	Probe(&MyValue{"Cicely Saunders"}, &root, Pad([]byte("nurse")), nil)
	Probe(&MyValue{"Bob Sweikert"}, &root, Pad([]byte("driver")), nil)
	Probe(&MyValue{"Peter Godfrey"}, &root, Pad([]byte("accountant")), nil)
	Probe(&MyValue{"Pam Beesley"}, &root, Pad([]byte("receptionist")), nil)

	Preorder(&MyValue{}, &root, func(value *MyValue){
		print(value.str)
		print("\n");
	})


/*
	Probe(&root, Pad([]byte("engineer")), &MyValue{"Valdemar Poulsen"}, nil)
	Probe(&root, Pad([]byte("manager")), &MyValue{"Bucky Harris"}, nil)
	Probe(&root, Pad([]byte("pharmacist")), &MyValue{"Carl W Scheele"}, nil)
	Probe(&root, Pad([]byte("cook")), &MyValue{"Tim Cook"}, nil)
	Probe(&root, Pad([]byte("cashier")), &MyValue{"David Griswold"}, nil)
*/

	Drop(&root, Pad([]byte("accountant")), nil)
	Drop(&root, Pad([]byte("receptionist")), nil)

	Preorder(&MyValue{}, &root, func(value *MyValue){
		print(value.str)
		print("\n");
	})
}
