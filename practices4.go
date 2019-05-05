*************************create new slice************************************
package main

import (
	"fmt"
)

func filter(x int) bool {
   if x%2==0{
     return true
   }
   return false
}

type intArr []int

func(list intArr) filterIt(f func(int) bool) intArr {
    var x intArr 
    for _,element:= range list {
         if f(element) {
            x = append(x,element)
         }
     }
     return x
}

func main() {
   x:= intArr {1,2,3,4,5,6}
   a:=x.filterIt(filter)   
   fmt.Println(a)
}



********************************update same slice*******************************
