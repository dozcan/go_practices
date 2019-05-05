generix interface den fixed type a dönüşüm

package main

import (
	"fmt"
)
func main() {
        var t interface{} = "4"
        i := t.(string)
      	fmt.Println(t,i)
	      var s interface{} = []string{"left", "right"}
	      k := s.([]string)
	      fmt.Println(s,k)
}

**************************************************************
package main

import (
	"fmt"
)
func main() {
     a := "4"
     b := "doga"
     c :=[]string {"sd","qeww"}      
     get(a,b,c)	 
}

func get(arg ...interface{}){

	for _,element := range arg {
	     switch element.(type){
	        case bool :
	            fmt.Println("bool")
	        case string:
	            fmt.Println("string")
	        case []string:
	            fmt.Println("[]string")
	        default:
	            fmt.Println("default")
	     }	
	}
}
