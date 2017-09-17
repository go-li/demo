package main

func b(arg string) []byte {
	return []byte(arg)
}

//

type RootBloc struct {
	fil []*FileBloc
}


// file block aka single sourcecode file
type FileBloc struct {
	pkg *PkgSment // there is one package statement
	imp []*ImpSment //there are several imports
	fun []*FunSment //there are several functions
	tds []*TypStrct //there are several typedef structs
	order []byte // the order of elements
}

// toplevel package statement
type PkgSment struct {
	pkg []byte // string
}



// toplevel import statement
type ImpSment struct {
	path []byte // string
}



// toplevel function
type FunSment struct {
	arg []*TyIDitem

	recv byte
	args byte
	rets byte

	name []byte // string
}





// typed identifier item
type TyIDitem struct {

	dadkind byte

	name []byte // string
	typ []byte // string
}

// typedef struct statement
type TypStrct struct {
	row []*TyIDitem

	rows byte

	typename []byte // string
}


//

const KindRootBloc = 0
const KindFileBloc = 1
const KindPkgSment = 2
const KindImpSment = 3
const KindFunSment = 4
const KindTyIDitem = 5
const KindTypStrct = 6

const KindShutdown = 255

type Node struct {
	kind byte
	fb *FileBloc
	rb *RootBloc
	ps *PkgSment
	is *ImpSment
	fs *FunSment
	ti *TyIDitem
	ts *TypStrct
}

///

func Step(n *Node, it *[]int) {

	var l = len(*it)

	switch (n.kind) {
	case KindRootBloc:
		var i = (*it)[l-1]

		if i >= len(n.rb.fil) {
			n.kind = KindShutdown
			return
		}

		(*it) = append((*it), 0,0,0,0)

		var v = n.rb.fil[i]
		n.kind = KindFileBloc
		n.fb = v

		(*it)[l-1]++


	case KindFileBloc:
		var i = (*it)[l-1] + (*it)[l-2] + (*it)[l-3] + (*it)[l-4]

		if i >= len(n.fb.order) {
			(*it) = (*it)[:l-4]
			n.kind = KindRootBloc
			return
		}

		var v = n.fb.order[i]
		switch (v) {
		case 0: n.kind = KindPkgSment; n.ps = n.fb.pkg
		case 1: n.kind = KindImpSment; n.is = n.fb.imp[(*it)[l-4+1]]
		case 2: n.kind = KindFunSment; n.fs = n.fb.fun[(*it)[l-4+2]] ; (*it) = append((*it), 0)
		case 3: n.kind = KindTypStrct; n.ts = n.fb.tds[(*it)[l-4+3]] ; (*it) = append((*it), 0)
		}

		(*it)[l-4+int(v)]++ 
	case KindPkgSment:

		n.kind = KindFileBloc
	case KindImpSment:

		n.kind = KindFileBloc
	case KindFunSment:
		var i = (*it)[l-1]

		if i >= len(n.fs.arg) {
			(*it) = (*it)[:l-1]
			n.kind = KindFileBloc
			return
		}
		var v = n.fs.arg[i]
		n.kind = KindTyIDitem
		n.ti = v

		(*it)[l-1]++


	case KindTypStrct:
		var i = (*it)[l-1]

		if i >= len(n.ts.row) {
			(*it) = (*it)[:l-1]
			n.kind = KindFileBloc
			return
		}
		var v = n.ts.row[i]
		n.kind = KindTyIDitem
		n.ti = v

		(*it)[l-1]++

	case KindTyIDitem:
		n.kind = n.ti.dadkind
	}
}

func main() {

	var root RootBloc

	var file FileBloc

	root.fil = append(root.fil, &file)

	file.pkg = &PkgSment{pkg:b("main")}
	file.imp = append(file.imp, &ImpSment{path:b(`"fmt"`)})
	file.fun = append(file.fun, &FunSment{name:b("Whatever1")})
	file.fun = append(file.fun, &FunSment{name:b("Something2")})
	file.fun = append(file.fun, &FunSment{name:b("Anything3")})
	file.fun = append(file.fun, &FunSment{name:b("Funfun"),recv:1,args:2,rets:2,
		arg:[]*TyIDitem{&TyIDitem{name:b("foo"),typ:b("bar"),dadkind:KindFunSment},
			&TyIDitem{name:b("boo"),typ:b("baz"),dadkind:KindFunSment},
			&TyIDitem{name:b("coo"),typ:b("caz"),dadkind:KindFunSment},
			&TyIDitem{name:b("doo"),typ:b("daz"),dadkind:KindFunSment},
			&TyIDitem{name:b("eoo"),typ:b("eaz"),dadkind:KindFunSment}}})
	file.tds = append(file.tds, &TypStrct{typename:b("Wow"),rows:5,
		row:[]*TyIDitem{&TyIDitem{name:b("foo"),typ:b("bar"),dadkind:KindTypStrct},
			&TyIDitem{name:b("boo"),typ:b("baz"),dadkind:KindTypStrct},
			&TyIDitem{name:b("coo"),typ:b("caz"),dadkind:KindTypStrct},
			&TyIDitem{name:b("doo"),typ:b("daz"),dadkind:KindTypStrct},
			&TyIDitem{name:b("eoo"),typ:b("eaz"),dadkind:KindTypStrct}}})
	file.order = []byte{0,1,3,2,2,2,2}

	// Now we print the AST

	var iter = []int{0}

	var n Node

	n.kind = KindRootBloc
	n.rb = &root

	for n.kind != KindShutdown {

		Step(&n, &iter)

//		println(n.kind)

		switch (n.kind) {
		case KindPkgSment: print("package "); println(string(n.ps.pkg))
		case KindImpSment: print("import "); println(string(n.is.path))
		case KindFunSment:
			var i = iter[len(iter)-1]

			if i == 0 {
				print("func ");
				if n.fs.recv > 0 {
					print("(");
				}
			}
			if i == int(n.fs.recv) {
				if n.fs.recv > 0 {
					print(") ");
				}
				 print(string(n.fs.name));print("(");
			}
			if (i > int(n.fs.recv)) && (i < int(n.fs.recv)+int(n.fs.args)) {
				print(", ");
			}
			if i == int(n.fs.recv)+int(n.fs.args) {
				print(") ");
				if n.fs.rets > 0 {
					print("(");
				}
			}
			if (i > int(n.fs.recv)+int(n.fs.args)) && (i < int(n.fs.recv)+int(n.fs.args)+int(n.fs.rets)) {
				print(", ");
			}
			if i == int(n.fs.recv)+int(n.fs.args)+int(n.fs.rets) {
				if n.fs.rets > 0 {
					print(") ");
				}
				println("{");println("}")
			}
		case KindTypStrct:
			var i = iter[len(iter)-1]
			if i == 0 {
				print("type ");print(string(n.ts.typename));println(" struct {");print("\t")
			} else if i == len(n.ts.row) {
				println("");println("}")
			} else {
				println("");print("\t")
			}
		case KindTyIDitem: print(string(n.ti.name)); print(" "); print(string(n.ti.typ));
		}


	}
}
