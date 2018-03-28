package main

import "runtime" 

type pool struct {
    pool   chan *
}


// newPool creates a new pool with max
func newPool(_ *pool, max int) (p pool) {
	p = pool{
		pool: make(chan *, max),
	}
	return p
}    

// borrow a Item from the pool.
func borrow(p *pool) (c *) {
    select {
    case c = <-p.pool:
    default:
        c = nil
    }
    return c
}

// reclaim returns a Item to the pool.
func reclaim(p *pool, c *) {
    select {
    case p.pool <- c:
    default:
        // let it go, let it go...
	// leaky pool
    }
}

// Cl13Nt C0de //////////////////////////////////////////////////


type foo struct {
	 n  int
}

type foopool struct {
	pool   chan *foo
}

func main() {
	var l = newPool((*foopool)(nil),1)
	var r = newPool((*foopool)(nil),1)

	var left foopool = l
	var right foopool = r

	reclaim(&left, &foo{1})

	go func(){for {
		print("Take From Right\n");
		var x = borrow(&right)
		if (x != nil) {
			print("Put to left\n");
			reclaim(&left, x)
		} else {
			runtime.Gosched()
		}

	}}()
	for i := 0; i < 10; i++ {
		print("Take From Left\n");
		var x = borrow(&left)
		if (x != nil) {
			print("Put to right\n");
			reclaim(&right, x)
		} else {
			runtime.Gosched()
		}
	}
}
