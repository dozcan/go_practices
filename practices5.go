oop

type ColoredPoint struct {
  color.Color // Anonymous field (embedding)
  x, y int // Named fields (aggregation)
}

t := ColoredPoint{}

t.x
t.y
t.Color => farklı package içinden embedded field'lara erişirken package adı ile çağrım yaparız yani
t.color.Color şeklinde çağrımda bulunmayız

***********************  pointer aritmetiği 1. tür (new ile)*************************************
package main

import (
	"fmt"
)

type DogaInt int

func(c *DogaInt) set() {  => bir referans adresi geçileceği belirtiliyor
    *c++                  => referasın içindeki deger bir artırılıyor
}

func main() {
        c:=new(DogaInt )  => bir pointer yaratılıyor(adres)
        c.set()           => set metodunu bir adres çağırıyor 
	      fmt.Println(*c)   => değeri artırıldı
}

***************************pointer aritmetiği direk atama ile ****************************
package main

import (
	"fmt"
)

type DogaInt int

func(c *DogaInt) set() {
    *c++
}

func main() {
        var c DogaInt   => value değer yaratıldı
        c.set()         => value değerin referansı metoda geçirildi
	fmt.Println(c)        => değer arttırıldı
}

***********************************************************************************************
package main

import (
	"fmt"
)

type Part struct {
    Id int // Named field (aggregation)
    Name string // Named field (aggregation)
}

func(c *Part) names(){
  c.Name = "murat"
}

func main() {
     c:= new(Part)  
     c.Name = "xcxc"
     c.names()
     fmt.Println(*c)   
}

yukarıdaki metodu 
     c:= Part{Id:1,Name:"doga"}
     c.names()
     fmt.Println(c)  
şeklinde de yazabilirdik     
