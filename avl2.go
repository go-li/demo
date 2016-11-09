package main

type Node struct {
	pavl_link [3]*Node /* Subtrees. */
	keybal    [32]byte
	pavl_data * /* value itself. */
}

func Pad(v []byte) []byte {
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

func Probe(value *, tree **Node, key []byte, result **) {
	var y *Node    /* Top node to update balance factor, and parent. */
	var p, q *Node /* Iterator, and parent. */
	var n *Node    /* Newly inserted node. */
	var w *Node    /* New root of rebalanced subtree. */

	var dir int /* Direction to descend. */

	//  assert (tree != nil && item != nil);

	y = *tree
	q = nil
	for p = *tree; p != nil; {
		var cmp int = bytecompare(key[:31], p.keybal[:31])
		if cmp == 0 {
			if result != nil {
				*result = p.pavl_data
			}
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
		p = p.pavl_link[dir]
	}

	n = &Node{}
	//	if n == nil {
	//		return nil
	//	}

	for i := 0; i < 31; i++ {
		n.keybal[i] = key[i]
	}

	n.pavl_link[0] = nil
	n.pavl_link[1] = nil
	n.pavl_link[2] = q
	n.pavl_data = value

	if q != nil {
		q.pavl_link[dir] = n
	} else {
		*tree = n
	}
	n.keybal[31] = 0
	if *tree == n {
		if result != nil {
			*result = n.pavl_data
		}
		return
	}

	for p = n; p != y; p = q {
		q = p.pavl_link[2]
		if q.pavl_link[0] != p {
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
		var x *Node = y.pavl_link[0]
		if x.keybal[31] == 255 {
			w = x
			y.pavl_link[0] = x.pavl_link[1]
			x.pavl_link[1] = y
			x.keybal[31] = 0
			y.keybal[31] = 0
			x.pavl_link[2] = y.pavl_link[2]
			y.pavl_link[2] = x
			if y.pavl_link[0] != nil {
				y.pavl_link[0].pavl_link[2] = y
			}
		} else {
			if !(x.keybal[31] == +1) {
				panic("assert")
			}
			w = x.pavl_link[1]
			x.pavl_link[1] = w.pavl_link[0]
			w.pavl_link[0] = x
			y.pavl_link[0] = w.pavl_link[1]
			w.pavl_link[1] = y
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
			w.pavl_link[2] = y.pavl_link[2]
			x.pavl_link[2] = w
			y.pavl_link[2] = w
			if x.pavl_link[1] != nil {
				x.pavl_link[1].pavl_link[2] = x
			}
			if y.pavl_link[0] != nil {
				y.pavl_link[0].pavl_link[2] = y
			}
		}
	} else if y.keybal[31] == +2 {
		var x *Node = y.pavl_link[1]
		if x.keybal[31] == +1 {
			w = x
			y.pavl_link[1] = x.pavl_link[0]
			x.pavl_link[0] = y
			x.keybal[31] = 0
			y.keybal[31] = 0
			x.pavl_link[2] = y.pavl_link[2]
			y.pavl_link[2] = x
			if y.pavl_link[1] != nil {
				y.pavl_link[1].pavl_link[2] = y
			}
		} else {
			if !(x.keybal[31] == 255) {
				panic("assert")
			}
			w = x.pavl_link[0]
			x.pavl_link[0] = w.pavl_link[1]
			w.pavl_link[1] = x
			y.pavl_link[1] = w.pavl_link[0]
			w.pavl_link[0] = y
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
			w.pavl_link[2] = y.pavl_link[2]
			x.pavl_link[2] = w
			y.pavl_link[2] = w
			if x.pavl_link[0] != nil {
				x.pavl_link[0].pavl_link[2] = x
			}
			if y.pavl_link[1] != nil {
				y.pavl_link[1].pavl_link[2] = y
			}
		}
	} else {
		if result != nil {
			*result = n.pavl_data
		}
		return
	}
	if w.pavl_link[2] != nil {
		var oo = 0
		if y != w.pavl_link[2].pavl_link[0] {
			oo = 1
		}
		w.pavl_link[2].pavl_link[oo] = w
	} else {
		*tree = w
	}

	if result != nil {
		*result = n.pavl_data
	}
}

/* Deletes from |tree| and returns an item matching |item|.
   Returns a null pointer if no matching item found. */
func Drop(tree **Node, key []byte, result **) {

	var p *Node /* Traverses tree to find node to delete. */
	var q *Node /* Parent of |p|. */
	var dir int /* Side of |q| on which |p| is linked. */

	if !(tree != nil) {
		panic("assert")
	}

	if *tree == nil {
		if result != nil {
			*result = nil
		}
		return
	}

	p = *tree
	for {
		var cmp int = bytecompare(key[:31], p.keybal[:31])
		if cmp == 0 {
			break
		}

		if cmp > 0 {
			dir = 1
		} else {
			dir = 0
		}

		p = p.pavl_link[dir]
		if p == nil {
			if result != nil {
				*result = nil
			}
			return
		}
	}
	if result != nil {
		*result = p.pavl_data
	}

	q = p.pavl_link[2]
	if q == nil {
		q = *tree
		dir = 0
	}

	if p.pavl_link[1] == nil {
		q.pavl_link[dir] = p.pavl_link[0]
		if q.pavl_link[dir] != nil {
			q.pavl_link[dir].pavl_link[2] = p.pavl_link[2]
		}
	} else {
		var r *Node = p.pavl_link[1]
		if r.pavl_link[0] == nil {
			r.pavl_link[0] = p.pavl_link[0]
			q.pavl_link[dir] = r
			r.pavl_link[2] = p.pavl_link[2]
			if r.pavl_link[0] != nil {
				r.pavl_link[0].pavl_link[2] = r
			}
			r.keybal[31] = p.keybal[31]
			q = r
			dir = 1
		} else {
			var s *Node = r.pavl_link[0]
			for s.pavl_link[0] != nil {
				s = s.pavl_link[0]
			}
			r = s.pavl_link[2]
			r.pavl_link[0] = s.pavl_link[1]
			s.pavl_link[0] = p.pavl_link[0]
			s.pavl_link[1] = p.pavl_link[1]
			q.pavl_link[dir] = s
			if s.pavl_link[0] != nil {
				s.pavl_link[0].pavl_link[2] = s
			}
			s.pavl_link[1].pavl_link[2] = s
			s.pavl_link[2] = p.pavl_link[2]
			if r.pavl_link[0] != nil {
				r.pavl_link[0].pavl_link[2] = r
			}
			s.keybal[31] = p.keybal[31]
			q = r
			dir = 0
		}
	}

	p.pavl_link[0] = nil
	p.pavl_link[1] = nil
	p.pavl_link[2] = nil
	p.keybal[31] = 0

	//  tree.pavl_alloc.libavl_free (tree.pavl_alloc, p);

	for q != *tree {
		var y *Node = q

		if y.pavl_link[2] != nil {
			q = y.pavl_link[2]
		} else {
			q = *tree
		}

		if dir == 0 {
			if q.pavl_link[0] != y {
				dir = 1
			} else {
				dir = 0
			}
			y.keybal[31]++
			if y.keybal[31] == +1 {
				break
			} else if y.keybal[31] == +2 {
				var x *Node = y.pavl_link[1]
				if x.keybal[31] == 255 {
					var w *Node

					if !(x.keybal[31] == 255) {
						panic("assert")
					}
					w = x.pavl_link[0]
					x.pavl_link[0] = w.pavl_link[1]
					w.pavl_link[1] = x
					y.pavl_link[1] = w.pavl_link[0]
					w.pavl_link[0] = y
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
					w.pavl_link[2] = y.pavl_link[2]
					x.pavl_link[2] = w
					y.pavl_link[2] = w
					if x.pavl_link[0] != nil {
						x.pavl_link[0].pavl_link[2] = x
					}
					if y.pavl_link[1] != nil {
						y.pavl_link[1].pavl_link[2] = y
					}
					q.pavl_link[dir] = w
				} else {
					y.pavl_link[1] = x.pavl_link[0]
					x.pavl_link[0] = y
					x.pavl_link[2] = y.pavl_link[2]
					y.pavl_link[2] = x
					if y.pavl_link[1] != nil {
						y.pavl_link[1].pavl_link[2] = y
					}
					q.pavl_link[dir] = x
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
			if q.pavl_link[0] != y {
				dir = 1
			} else {
				dir = 0
			}
			y.keybal[31]--
			if y.keybal[31] == 255 {
				break
			} else if y.keybal[31] == 254 {
				var x *Node = y.pavl_link[0]
				if x.keybal[31] == +1 {
					var w *Node
					if !(x.keybal[31] == +1) {
						panic("assert")
					}
					w = x.pavl_link[1]
					x.pavl_link[1] = w.pavl_link[0]
					w.pavl_link[0] = x
					y.pavl_link[0] = w.pavl_link[1]
					w.pavl_link[1] = y
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
					w.pavl_link[2] = y.pavl_link[2]
					x.pavl_link[2] = w
					y.pavl_link[2] = w
					if x.pavl_link[1] != nil {
						x.pavl_link[1].pavl_link[2] = x
					}
					if y.pavl_link[0] != nil {
						y.pavl_link[0].pavl_link[2] = y
					}
					q.pavl_link[dir] = w
				} else {
					y.pavl_link[0] = x.pavl_link[1]
					x.pavl_link[1] = y
					x.pavl_link[2] = y.pavl_link[2]
					y.pavl_link[2] = x
					if y.pavl_link[0] != nil {
						y.pavl_link[0].pavl_link[2] = y
					}
					q.pavl_link[dir] = x
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

	return

}

// Visits tree values in sequence
func Preorder(unused *, tree **Node, visit func(*)) {

	var p *Node /* Iterator. */

	p = *tree
	if p == nil {
		return
	}

	if p.pavl_link[0] != nil {

		p = p.pavl_link[0]
		Preorder(unused, &p, visit)
		p = p.pavl_link[2]
	}

//	var x = string(p.keybal[:31])
//	print(x)
	visit(p.pavl_data)

	if p.pavl_link[1] != nil {

		p = p.pavl_link[1]
		Preorder(unused, &p, visit)
		p = p.pavl_link[2]
	}
}

//-ClienT C0DE--------------------

type MyValue struct {
	str string
}

type StringNode struct {
	pavl_link [3]*StringNode /* Subtrees. */
	keybal    [32]byte
	pavl_data *MyValue /* value itself. */
}

func main() {

	var root *StringNode

	Probe(&MyValue{"Paul Sartorius"}, &root, Pad([]byte("composer")), nil)
	Probe(&MyValue{"Elkanah Settle"}, &root, Pad([]byte("writer")), nil)
	Probe(&MyValue{"Edie Martin"}, &root, Pad([]byte("actress")), nil)
	Probe(&MyValue{"Walter de Stapledon"}, &root, Pad([]byte("bishop")), nil)
	Probe(&MyValue{"Blake Ross"}, &root, Pad([]byte("developer")), nil)
	Probe(&MyValue{"Cicely Saunders"}, &root, Pad([]byte("nurse")), nil)
	Probe(&MyValue{"Bob Sweikert"}, &root, Pad([]byte("driver")), nil)
	Probe(&MyValue{"Peter Godfrey"}, &root, Pad([]byte("accountant")), nil)
	Probe(&MyValue{"Pam Beesley"}, &root, Pad([]byte("receptionist")), nil)
	Probe(&MyValue{"Valdemar Poulsen"}, &root, Pad([]byte("engineer")), nil)
	Probe(&MyValue{"Bucky Harris"}, &root, Pad([]byte("manager")), nil)
	Probe(&MyValue{"Carl W Scheele"}, &root, Pad([]byte("pharmacist")), nil)
	Probe(&MyValue{"Tim Cook"}, &root, Pad([]byte("cook")), nil)
	Probe(&MyValue{"David Griswold"}, &root, Pad([]byte("cashier")), nil)

	Preorder(nil, &root, func(value *MyValue) {
		print(value.str)
		print("\n")
	})
	print("\n")

	Drop(&root, Pad([]byte("accountant")), nil)
	Drop(&root, Pad([]byte("developer")), nil)

	Preorder(nil, &root, func(value *MyValue) {
		print(value.str)
		print("\n")
	})
}
