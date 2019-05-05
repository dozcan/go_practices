type embedded

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
    it Item 
    catalogId int 
}

func (item *Item) Cost() float64 {
    return item.price * float64(item.quantity)
}


func main() {
     s := SpecialItem {Item{"1",2,2},3}
     fmt.Println(s.it.price)   
}

************************************************************************
bu fonksiyonalite embedded type olsa idi

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


func main() {
     s := SpecialItem {Item{"1",2,2},3}
     fmt.Println(s.price) 
     fmt.Println(s.Cost())   
}
