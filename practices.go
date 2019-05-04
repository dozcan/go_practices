package main

import (
	"fmt"
	"errors"
)

type stack []interface{}

func(s *stack) push(y interface{}) {
     *s = append(*s,y)
}

func(s *stack) pop() (result interface {},err error) {
   defer func() {
      if err == nil {
         fmt.Println("hata yok")
         return
     }
     fmt.Println("hata var")
   }()
			
   temp := *s
   if len(temp) == 0 {
     result = nil
     err = errors.New("hata")
     return
   }
     result = temp[0]
     err = nil
     temp = temp[1:]
     *s = temp
     return
}

func main() {
        s := stack{}
        s.push(4)
        s.push("kisi") 
        s.pop()
	fmt.Println(s)
}
