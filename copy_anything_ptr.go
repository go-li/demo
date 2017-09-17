
package main

import "fmt"

func custom_copy_anything(to *, from *) {
	*to = *from
}

func main() {

	var a uint16 = 64548
	var b uint16 = 36543
	var c uint16 = 24264

	var x = [11]byte{'t','h','i','s',' ','i','s',' ','f','u','n'}
	var y = [11]byte{'g','o','l','a','n','g',' ','r','u','l','e'}

	fmt.Printf("a=%d  b=%d  c=%d \n", a,b,c)

	custom_copy_anything(&b, &c)

	fmt.Printf("a=%d  b=%d  c=%d \n", a,b,c)

	custom_copy_anything(&c, &a)

	fmt.Printf("a=%d  b=%d  c=%d \n", a,b,c)


	custom_copy_anything(&x, &y)

	fmt.Printf("x='%s'  y='%s' \n", string(x[:]), string(y[:]))

}
