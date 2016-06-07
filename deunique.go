package main



func Deuniqe(compare func(*,*)int, arry []*) (l int) {
 l = len(arry)
 for i:=0;i<l;i++{
  for j:=i+1;j<l;j++{
   if(compare(arry[j], arry[i])==0){
        arry[j] = arry[l-1]
        l--
        continue;
     }
   if (j==l-1) {
     arry[i] = arry[j]
     l--
     i--
     break
   }
  }
 }
  return l
}

type man struct {
name string
id byte
}

func compare_man_by_id(l *man, r *man) int {
      return int(l.id) - int(r.id)
}


func main() {
	var y = []*man{&man{"Bob",0},&man{"Pat",1},&man{"Bob",0},&man{"Rob",2},&man{"Rob",2},&man{"Bob",0}}

	y = y[:Deuniqe(compare_man_by_id, y)]

	for n := range y {
		var x man = *(y[n])
		print(x.name)
		print(" ")
		print(x.id)
		print("\n")
	}


}
