overriding struct

package main

import (
	"fmt"
)

type Item struct {
    id string // Named field (aggregation)
    price float64 // Named field (aggregation)
    quantity int // Named field (aggregation)
}

type SpecialItem struct {
    Item 
    catalogId int 
}

func (item *Item) Cost() float64 {
    return item.price * float64(item.quantity)
}

func (item *SpecialItem ) Cost() float64 {
    return item.Item.Cost()
}


func main() {
     s := SpecialItem {Item{"1",2,2},3}
     fmt.Println(s.price) 
     fmt.Println(s.Cost())   
}


*******************************************ınterface**************************
package main

import (
	"fmt"
)

type Iexchanger interface {
   Exchange()
}

type StringPair struct{
   first,second string
}

type Point [2]int

func(s *StringPair) Exchange() {
   s.first,s.second = s.second,s.first
}

func(p *Point) Exchange(){
   p[0],p[1]=p[1],p[0]
}

func ExchangeInterface(arg ...Iexchanger){
   for _,element:= range arg {
       element.Exchange()
  }
}


func main() {
     s:= StringPair{"first","second"}
     p:= Point{2,3}
     ExchangeInterface(&p,&s)
     fmt.Println(s,p)
}
