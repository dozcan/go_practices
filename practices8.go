fonksiyonlar içerisinde dizi değişikliği yapmak isteniyorsa pointer a gerek yoktur
cunku diziler referans tiptir

package main

import (
	"fmt"
)

func swap(a []int){
    a[0],a[1]=a[1],a[0]
    fmt.Println(a)

}

func main() {
        a:=[]int{10,100}    
        swap(a) 
	fmt.Println(a)
}
*****************************************ama değer tipleri value gecirilirler onun için pointer gereklidir******************

package main

import (
	"fmt"
)

func swap(a *int,b *int){
    *a,*b=*b,*a

}

func main() {
        a,b:=10,100   
        swap(&a,&b) 
	fmt.Println(a,b)
}


**************************************interface***********************************************
package main

import (
	"fmt"
	"strings"
)

type ILowerCaser interface {
     toLowerCase()
}

type IUpperCaser interface {
     toUpperCase()
}

type IAllCaser interface {
     ILowerCaser 
     IUpperCaser 
}

type ICaser interface {
     IAllCaser 
}

type StringPair struct{ first, second string }

func(s *StringPair) toLowerCase(){
     s.first = strings.ToUpper(s.first)
     s.second = strings.ToUpper(s.second )
}

func(s *StringPair) toUpperCase(){
     s.first = strings.ToLower(s.first)
     s.second = strings.ToLower(s.second )
}

func genericICaser(i ICaser){
     i.toUpperCase()
     i.toLowerCase()

}

func main(){
        s:= StringPair{"doga","ozcan"}
        genericICaser(&s)
        fmt.Println(s)


}
