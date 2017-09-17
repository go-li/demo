package main

import "runtime" 

type Pool struct {
    pool   chan *
}


// NewPool creates a new pool from one optional item.
func NewPool(initial *, extra int) (p Pool) {
	if (initial == nil ) {
		return Pool{
			pool: make(chan *, extra),
		}
	}
	p = Pool{
		pool: make(chan *, 1+extra),
	}
	Return(&p, initial)
	return p
}    

// Borrow a Item from the pool.
func Borrow(p *Pool) (c *) {
    select {
    case c = <-p.pool:
    default:
        c = nil
    }
    return c
}

// Return returns a Item to the pool.
func Return(p *Pool, c *) {
    select {
    case p.pool <- c:
    default:
        // let it go, let it go...
	// leaky pool
    }
}

// Cl13Nt C0de //////////////////////////////////////////////////


type Foo struct {
	 n  int
}

type FooPool struct {
	pool   chan *Foo
}

func main() {
	var left = NewPool(&Foo{1},0)
	var right = NewPool(nil,1)

	go func(){for {
		print("Take From Right\n");
		var x = Borrow(&right)
		if (x != nil) {
			print("Put to left\n");
			Return(&left, x)
		} else {
			runtime.Gosched()
		}

	}}()
	for i := 0; i < 10; i++ {
		print("Take From Left\n");
		var x = Borrow(&left)
		if (x != nil) {
			print("Put to right\n");
			Return(&right, x)
		} else {
			runtime.Gosched()
		}
	}
}
