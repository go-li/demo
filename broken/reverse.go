package main

func reverse(slice [] ) {
	var l = len(slice)
	for i := 0 ; i < l/2; i++  {
		var j = l-i-1
		slice[i], slice[j] = slice[j], slice[i]
	}
}

func main() {
	var things = []int{5,6,7,8,9}
	var bytes = []byte{4,3,2,1}
	_  = things
	_  = bytes

	reverse(things)
	reverse(bytes)

	for i := 0; i < 5; i++ {
	print(things[i])
	}
	print("\n")
	for i := 0; i < 4; i++ {
	print(bytes[i])
	}
	print("\n")

}
