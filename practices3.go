hata yakalama

package main

import (
	"fmt"
)

func main() {
  i,_:= resultOfMod(7)
  fmt.Println("sayımız",i)	
}

func resultOfMod(a int) (result int,err error){
    defer func(){
       e:=recover()
       fmt.Println("errr",e)
       if  e!=nil{
         err= fmt.Errorf("hatamız %v",err)
       }
    }()     

    i:=get(a);
    fmt.Println("sonuc",i,err)
    return i,err

}

func get(a int) (int) {
    if a%3 == 0 {
       return a/3
    } 
    panic(fmt.Sprintf("3 ile bölünemeyen sayımız %v'dır",a))
}

***********************variadic functions:****************************
package main

import (
	"fmt"
)

func main() {
        a:= []int{1,2,3,4}
   	resultOfMod(4,a...)
}

func resultOfMod(a int,b ...int){
   fmt.Println(a)
   for _,element := range b {
    fmt.Println(element)
   }

}

*******************generic minumum fonksiyonu******************************
i := Minimum(4, 3, 8, 2, 9).(int)
fmt.Printf("%T %v\n", i, i)
f := Minimum(9.4, -5.4, 3.8, 17.0, -3.1, 0.0).(float64)
fmt.Printf("%T %v\n", f, f)
s := Minimum("K", "X", "B", "C", "CC", "CA", "D", "M").(string)
fmt.Printf("%T %q\n", s, s)



























