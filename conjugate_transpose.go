package main

import (
"fmt"
)

func ConjCmplx(n *complex128) {
	*n = complex(real(*n), -imag(*n))
}

func ConjugateTranspose2x2(conj func(*), mat *[2][2]complex128) {
	(*mat)[0][1], (*mat)[1][0] = (*mat)[1][0], (*mat)[0][1]
	
	if conj!=nil {
		conj(&(*mat)[0][0])
		conj(&(*mat)[0][1])
		conj(&(*mat)[1][0])
		conj(&(*mat)[1][1])
	}
}

func main() {
	//example from mathworks com
	var matrix = [2][2]complex128{ 
	{ 0.0 - 1.0i ,  2.0 + 1.0i},
        { 4.0 + 2.0i ,  0.0 - 2.0i}}
	_ = matrix

	ConjugateTranspose2x2(ConjCmplx, &matrix)

	fmt.Printf("%v %v\n %v %v", matrix[0][0], matrix[0][1],matrix[1][0],matrix[1][1])
}
