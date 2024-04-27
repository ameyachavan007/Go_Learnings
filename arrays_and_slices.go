package hello 

import (
	"fmt"
)  

var w = [...]int{1,2,3} // array of len 3
var x = []int{0,0,0} // slice of len (3)

func do (a [3]int, b []int) []int {
  a = b // syntax error: cannot assign slice to array
  a[0] = 4 // w is unchanged
  b[0] = 3 // x is changed
  c := make([]int, 5) // []iint{0,0,0,0,0}
  c[4] = 42
  copy(c,b) // copies only 3 elements
  return c 
}

var y = do(w,x)  
// fmt.Println(w, x, y)

// [1,2,3] [3, 0, 0] [3 0 0 0 42]  
