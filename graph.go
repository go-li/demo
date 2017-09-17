package main

import "fmt"
import "math"

// int8 runlength encoding

func Last(boolset []int8) byte {
	if boolset[len(boolset)-1] < 0 {
		return 0xff
	} else {
		return 0
	}
}

func Concatenate(boolset []int8, another []int8 ) []int8 {
	if len(boolset) == 0 {
		return another
	}
	if len(another) == 0 {
		return boolset
	}
	if (another[0] >= 0) != (boolset[len(boolset)-1] >= 0) {
		return append(boolset, another...)
	}
	var bigboolset = Unpack(boolset)
	var a,end = UnpackInt(another)
	if a >= 0 {
		a++
	}
	bigboolset[len(bigboolset)-1] += a
	boolset = Packed(bigboolset)

	return append(boolset, end...)
}

func RunLen(setofbools []int8) (l int) {
	var set = Unpack(setofbools)
	for i := range set {
		if set[i] >= 0 {
			l+=int(set[i])
		} else {
			l+=-int(set[i])
		}
	}
	return l
}

func Unpack(runlen []int8) (o []int64) {
	for len(runlen) > 0 {
		var i int64
		i, runlen = UnpackInt(runlen)
		o = append(o, i)
	}
	return o
}



func UnpackInt(packed []int8) (o int64, end []int8) {


	var l int

	for i := 0; i < 9; i++ {

		o <<= 7
		o |= int64(packed[i]&127)

		if i == len(packed)-1 {
//			l++
			break
		}

		if (packed[i] >= 0) == (packed[i+1] >= 0) {
			l++
		} else {
			break
		}
	}
	for i := 0; i <= l; i++ {
		o += 1 << uint(7*i)
	}


	if packed[0] < 0 {
		o = -o
	}
//	println(";;")
//	println(packed[0])
//	println(l)
//	println(o)
//

	return o, packed[l+1:]

}

func DeCode(boolset []int64) (out []byte) {
	if len(boolset) == 0 {
		return out
	}
//	out = append(out, boolset[0] < 0)
	for len(boolset) > 0 {
		if boolset[0] > 0 {
			out = append(out, 0x00)
			boolset[0]--
		} else if boolset[0] < 0 {
			out = append(out, 0xff)
			boolset[0]++
		} else {
			boolset = boolset[1:]
		}
	}
	return out
}

func RunCode(bools []byte) (out []int64) {
	var counter int
	for  {
		if (counter < len(bools)) && (bools[0] == bools[counter]) {
			counter++
			continue
		}

		if bools[0] != 0 {
			out = append(out, -int64(counter))
		} else {
			out = append(out, int64(counter-1))
		}
		if counter == len(bools) {
			return out
		}

		bools = bools[counter:]
		counter = 0

	}
	return out
}

func Abs(n int64) uint64 {
	if n < 0 {
		return uint64(n) ^ 0xffffffffffffffff
	}
	return uint64(n)
}

func PackInt(run int64) (o [9]int8, l int) {
	var s = Abs(run)
	for i := uint64(9295997013522923649); i > 0; i >>= 7  {
		l++
		if (s >= i-1) {
			s -= i
			s++
			break
		}
	}
//	println("::")
//	println(run)
//	println(s)
//	println(l)

	o[8] = int8(s)
	o[7] = int8(s >> 7)
	o[6] = int8(s >> 14)
	o[5] = int8(s >> 21)
	o[4] = int8(s >> 28)
	o[3] = int8(s >> 35)
	o[2] = int8(s >> 42)
	o[1] = int8(s >> 49)
	o[0] = int8(s >> 56)

	for i := 0; i < 9; i++ {
		if run < 0 {
			o[i] |= -128
		} else {
			o[i] &= 127
		}
	}

	return o, l-2
}

func Packed(runlen []int64) (out []int8) {
	for i := range runlen {
		var o, l = PackInt(runlen[i])
		out = append(out, o[l:9]...)
	}
	return out
}

func Gap(n uint) []int8 {
	if n == 0 {
		return nil
	}
	var o, l = PackInt(int64(n)-1)
	return o[l:9]
}

func Burst(n uint) []int8 {
	if n == 0 {
		return nil
	}
	var o, l = PackInt(-int64(n))
	return o[l:9]
}

func And64(a []int64, b []int64) []int64 {
	if (len(a) == 0) || (len(b) == 0) {
		return []int64{}
	}
	var a_dec = DeCode(a)
	var b_dec = DeCode(b)
	for i := range a_dec {
		a_dec[i] = a_dec[i] & b_dec[i]
	}
	return RunCode(a_dec)
}

func Or64(a []int64, b []int64) []int64 {
	if (len(a) == 0) || (len(b) == 0) {
		return []int64{}
	}
	var a_dec = DeCode(a)
	var b_dec = DeCode(b)
	for i := range a_dec {
		a_dec[i] = a_dec[i] | b_dec[i]
	}
	return RunCode(a_dec)
}

func Pow2( exp int, length int) (o []int8) {
	o = Gap(uint(exp))
	o = append(o, -128)
	o = append(o, Gap(uint(length-exp-1))...)
	return o
}

// graph code

type Graph struct {
	Nodes []*struct{}
	Edges interface{} 
	NodeId func(*struct{}) *int
}

func (g Graph) NodeID(s *struct{}) int {
	if g.NodeId == nil {
		for i := range g.Nodes {
			if g.Nodes[i] == s {
				return i
			}
		}
		return -1
	}
	return *g.NodeId(s)
	fmt.Printf("")
	return 0
}

const Edge float64 = 1
const Loop = Edge
const Void = float64(uint64(9223372036854775808));

type Mat struct {
	upper []byte
	diagonal []byte	// equal to nil if whole diagonal is empty
	lower []byte	// equal to upper in undirected graphs

}

func Undirectify1(m Mat) Mat {
	m.lower = m.upper
	return m
}

func Undirectify2(m Mat) Mat {
	m.upper = m.lower
	return m
}

func Directify1(m Mat) Mat {
	copy(m.lower, m.upper)
	return m
}

func Directify2(m Mat) Mat {
	copy(m.upper, m.lower)
	return m
}

func Unloopify(m Mat) Mat {
	m.diagonal = nil
	return m
}

func Directed(m *Mat) bool {
	if m == nil {
		return true
	}
	if m.upper == nil {
		return true
	}
	if m.lower == nil {
		return true
	}
	if len(m.upper) == 0 {
		return true
	}
	if len(m.lower) == 0 {
		return true
	}
	return &m.upper[0] != &m.lower[0]
}

func Undirected(m *Mat) bool {
	if m == nil {
		return true
	}
	if m.upper == nil {
		return true
	}
	if m.lower == nil {
		return true
	}
	if len(m.upper) == 0 {
		return true
	}
	if len(m.lower) == 0 {
		return true
	}
	return &m.upper[0] == &m.lower[0]
}


func Len(m Mat) int {
	if m.diagonal != nil {
		return len(m.diagonal)+0
	}
	if m.upper == nil {
		return 0
	}
	return int(0.5*math.Sqrt(1+8*float64(len(m.upper)+0)))+1
}


func Addnode(m *Mat, rowcolumn []int8) {
	AddnodeBidi(m, rowcolumn, rowcolumn)
}

func AddnodeBidi(m *Mat, row []int8, column []int8) {
	var length = RunLen(row)
	var length2 = RunLen(column)
	var l = Len(*m)
	var lastrow = Last(row)

	_ = l

	if length != length2 {
		panic("Row must have as many elements as column.")
	}

	if length != l+1 {
		println(length)
		println(l+1)
		panic("Must specify all edges to/from all previous nodes and to/from self.")
	}

	if (len(m.diagonal) != 0) || ((lastrow!=0) && (m.upper == nil)) {
		m.diagonal = append(m.diagonal, lastrow)
	} else if (lastrow!=0) && (m.diagonal == nil) {
		m.diagonal = make([]byte, l+1,l+1)
		m.diagonal[len(m.diagonal)-1] = 0xff
	}

//	print("adding ")
//	println(length)
//	print("last ")
//	println(lastrow)
//	print("dir ")
//	println(Directed(m))
//	print("spoj ")
//	println(Undirected(m))

	if Undirected(m) && (&row[0] == &column[0]) {

//	fmt.Println(":%v", column)
//	fmt.Println(":%v", (Unpack(column)))
//	fmt.Println(":%v", DeCode(Unpack(column)))



		m.upper = append(m.upper, DeCode(Unpack(column))...)
		if len(m.upper) > 0 {
			m.upper = m.upper[:len(m.upper)-1]
		}
		m.lower = m.upper

	} else {
		m.upper = append(m.upper, DeCode(Unpack(column))...)
		if len(m.upper) > 0 {
			m.upper = m.upper[:len(m.upper)-1]
		}
		m.lower = append(m.lower, DeCode(Unpack(row))...)
		if len(m.lower) > 0 {
			m.lower = m.lower[:len(m.lower)-1]
		}
	}

}
func ForEachOutEdge(m *Mat, fromnode uint, tonode uint, yes func(uint,uint,float64), no func(uint,uint)) {
	ForEachEdge(m, m.lower, m.upper, fromnode, tonode, yes, no)
}

func ForEachInEdge(m *Mat, fromnode uint, tonode uint, yes func(uint,uint,float64), no func(uint,uint)) {
	ForEachEdge(m, m.upper, m.lower, fromnode, tonode, yes, no)
}

func ForEachEdge(m *Mat, lower []byte, upper []byte, fromnode uint, tonode uint, yes func(uint,uint,float64), no func(uint,uint)) {
	if tonode == 0 {
		tonode = uint(Len(*m))
	}


	for y := int(fromnode); y < int(tonode) ;y++ {


	var tmp = y-1
	var n = (tmp*tmp + tmp )/2
	for j := n ; j < n+y; j ++ {
		if lower[j] != 0 {
		if yes != nil {
			yes(uint(j-n),uint(y),Edge)
		}
		} else {
		if no != nil {
			no(uint(j-n),uint(y))
		}
		}
	}
	if (m.diagonal == nil) {
		if no != nil {
			no(uint(y),uint(y))
		}
	} else {
		if m.diagonal[y] != 0 {
		if yes != nil {
			yes(uint(y),uint(y),Edge)
		}
		} else {
		if no != nil {
			no(uint(y),uint(y))
		}
		}
	}

	x := y
	for j := (y*y+y)/2 ; j+y < len(upper); j += x {
//		print("[")
//		print(j+y)
//		print("]")

		if upper[j+y] != 0 {
		if yes != nil {
			yes(uint(x+1),uint(y),Edge)
		}
		} else {
		if no != nil {
			no(uint(x+1),uint(y))
		}
		}

		x++
	}

	}
}

func setLine(m *Mat, node uint, line []int64) (change byte) {
	var decoded = DeCode(line)
	var y = int(node)


	var tmp = y-1
	var n = (tmp*tmp + tmp )/2
	for j := n ; j < n+y; j ++ {
		change |= m.lower[j] ^ decoded[j-n]
		m.lower[j] = decoded[j-n]

	}
	if (m.diagonal == nil) {
		panic("FIXME:no diagonal")
	} else {
		change |= m.diagonal[y] ^ decoded[y]
		m.diagonal[y] = decoded[y]
	}

	x := y
	for j := (y*y+y)/2 ; j+y < len(m.upper); j += x {
		change |= m.upper[j+y] ^ decoded[x+1]
		m.upper[j+y] = decoded[x+1]

		x++
	}
	return change
}

func getLine(m *Mat, node uint) (line []int64) {
	if Len(*m) == 0 {
		return line
	}
	line = append(line, 0)

	ForEachOutEdge(m, node, node+1, func(x uint, y uint, d float64){
		if line[len(line)-1] <= 0 {
			line[len(line)-1]--
		} else {
			line = append(line, -1)
		}
	}, func(x uint, y uint){
		if line[len(line)-1] >= 0 {
			line[len(line)-1]++
		} else {
			line = append(line, 1)
		}
	})
	return line
}

func dump(m *Mat) {
	for y := 0; y < Len(*m); y++ {
		ForEachOutEdge(m, uint(y), uint(y+1), func(x uint, y uint, d float64){
			print("1")
		}, func(x uint, y uint){
			print("0")
		})
		println("")
	}
}

func main() {

/*
	fmt.Println("%v", RunLen(Packed(RunCode([]byte{0}))))
	fmt.Println("%v", DeCode(Unpack(Packed(RunCode([]byte{1,1,1,0,0,1})))))
*/

/*
	var g Mat

	Addnode(&g, Packed(RunCode([]byte{0})))



	Addnode(&g, Packed(RunCode([]byte{1,0})))



	Addnode(&g, Packed(RunCode([]byte{1,1,0})))



	Addnode(&g, Packed(RunCode([]byte{0,1,1,1})))



	Addnode(&g, Packed(RunCode([]byte{1,0,1,1,0})))



	Addnode(&g, Packed(RunCode([]byte{1,0,1,1,0,0})))


	dump(&g)

	AddnodeBidi(&g, Packed(RunCode([]byte{0,1,1,0,1,0,1})), Packed(RunCode([]byte{1,0,1,1,0,0,1})))

	dump(&g)

	g = Undirectify1(g)

	println("")

	dump(&g)
*/

	var tarjan Mat

	AddnodeBidi(&tarjan, Packed(RunCode([]byte{0})), Packed(RunCode([]byte{0})))
	AddnodeBidi(&tarjan, Packed(RunCode([]byte{0,0})), Packed(RunCode([]byte{1,0})))
	AddnodeBidi(&tarjan, Packed(RunCode([]byte{0,1,0})), Packed(RunCode([]byte{1,0,0})))
	AddnodeBidi(&tarjan, Packed(RunCode([]byte{0,0,0,0})), Packed(RunCode([]byte{1,0,0,0})))
	AddnodeBidi(&tarjan, Packed(RunCode([]byte{0,0,0,0,0})), Packed(RunCode([]byte{0,1,1,0,0})))
	AddnodeBidi(&tarjan, Packed(RunCode([]byte{0,0,0,0,0,0})), Packed(RunCode([]byte{0,0,1,0,0,0})))
	AddnodeBidi(&tarjan, Packed(RunCode([]byte{0,0,0,0,0,0,0})), Packed(RunCode([]byte{0,0,0,1,0,0,0})))
	AddnodeBidi(&tarjan, Packed(RunCode([]byte{0,0,0,0,0,0,0,0})), Packed(RunCode([]byte{0,0,0,1,0,0,0,0})))
	AddnodeBidi(&tarjan, Packed(RunCode([]byte{0,0,0,0,0,1,0,0,0})), Packed(RunCode([]byte{0,0,0,0,0,1,0,0,0})))
	AddnodeBidi(&tarjan, Packed(RunCode([]byte{0,0,0,0,0,0,0,0,0,0})), Packed(RunCode([]byte{0,0,0,0,0,0,1,1,0,0})))
	AddnodeBidi(&tarjan, Packed(RunCode([]byte{0,0,0,0,0,0,0,0,0,1,0})), Packed(RunCode([]byte{0,0,0,0,0,0,0,1,0,0,0})))
	AddnodeBidi(&tarjan, Packed(RunCode([]byte{1,0,0,0,0,0,0,0,0,1,0,0})), Packed(RunCode([]byte{0,0,0,0,0,0,0,0,1,1,0,0})))
	AddnodeBidi(&tarjan, Packed(RunCode([]byte{0,0,0,0,0,0,0,0,1,0,0,0,0})), Packed(RunCode([]byte{0,0,0,0,1,0,0,0,0,0,0,0,0})))
	println("")

	dump(&tarjan)

	ForEachOutEdge(&tarjan, 0, 0, func(a uint, b uint, c float64) {
		print(b)
		print(" to ")
		print(a)
		println("")

	}, nil)

//	fmt.Println("%v", DeCode(getLine(&tarjan, 0)))

	// begin dominators algorithm

	var dommap Mat

	const entry = 0

	AddnodeBidi(&dommap, Burst(1), Burst(1))


	for i := 1; i < 13; i++ {
		AddnodeBidi(&dommap, Burst(uint(i)+1), Concatenate(Gap(1),Burst(uint(i))))

	}

//	println("")
//	dump(&dommap)

	// while change
	var change = true
	for change {
		change = false
		for vertex := 1; vertex < Len(tarjan); vertex ++ {

			var now = Burst(uint(Len(tarjan)))
			_ = now

			ForEachInEdge(&tarjan, uint(vertex), uint(vertex+1), func(parent uint, b uint, c float64) {

				var dommap_of_parent = getLine(&dommap, parent)

//				print(parent)
//				print(" to ")
//				print(b)
//				println("")

				now = Packed(And64(Unpack(now), dommap_of_parent))

			}, nil)


			now = Packed(Or64(Unpack(now), Unpack(Pow2(vertex,Len(tarjan)))))

//			fmt.Println("::%d %v", vertex, DeCode(Unpack(now)))

			if (0 != setLine(&dommap, uint(vertex), Unpack(now))) {

				change = true
			}

//			println("")
//			dump(&dommap)
		}
	}

	println("")
	dump(&dommap)

	return
/*
	g = Undirectify1(g)

//	Append(&g, []float64{Void, Void, Void}...)

	fmt.Println("%v", Packed([]int64{126,-127,127,-128,128,-129,129,-130}))
//	fmt.Println("%v", Packed([]int64{9223372036854775807,-9223372036854775808,9223372036854775807,-9223372036854775808,}))

	for i := int64(-1); i < 0; i-- {

	var pp = Packed([]int64{i})

//	fmt.Println("%v", pp)

	var qq,_ = UnpackInt(pp)

	if qq != i {
		println(i)
		println(qq)

		panic("!=")
	}

	}
*/
}
