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
