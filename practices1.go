package main

import (
	"fmt"
	"errors"
)


func ModAndDivideAndMultiple(a int, b int) (func(int,int) int,error) {
   if b == 0 {
     return nil, errors.New("0 ile bölünme")
   }
   b=b%5
   return func(bolunen int,bolen int) int {
      return bolunen/bolen * a * b
  },nil
}

func main() {
      c,err := ModAndDivideAndMultiple(6,2);
      if err!= nil {
        fmt.Println("hata")
        return
      }
      fmt.Println(c(6,2))
}

********************************************************
package main

import (
	
"fmt"
)

var kisi = "benim adım:%s'dır yasım %v'dir"

func main(){
   kisi= fmt.Sprintf(kisi,"doga",21)
   fmt.Println(kisi)
}


************************************************
package main

import (
	
"fmt"
"strings"
)

func toUp(s rune) rune{
    if s == 'a' {
       return 'b'
    }
    if(s == ' '){
       return '*'
    }
    return 'c'
}

func main(){
   i := "doga ozcan"
   fmt.Println(strings.Map(toUp,i))
}

