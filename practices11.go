linked list
  package main

import (
	"fmt"
)

type NodeList []node
var list NodeList 

type data struct {
    name string
    surname string
}

type node struct {
    data
    left  *node
    right *node
}

func addToList(_node ...node) {
    if len(_node) < 2{
       list = append(list,_node[0])
    }else{
       list = append(list[0:len(list)-1],_node[0],_node[1])
    }
}

func add(_name string,_surname string) (err error){
    if len(list ) < 1 { 
        newNode := new(node)
        newNode.data = data{_name,_surname}
        newNode.right = nil
        newNode.left = nil
        addToList(*newNode)
        return nil
    }else{
        lastNode := list [len(list ) -1]
        newNode := new(node)
        newNode.data = data{_name,_surname}
        newNode.right = nil
        lastNode.right = newNode
        newNode.left  = &lastNode
        addToList(lastNode ,*newNode)
        return nil
    }
}

func showNodes(){
    for _,element := range list {
     fmt.Println(element)
    }
}

func controlAdd(_name string,_surname string) (err error){
   err = add(_name,_surname)
   if err != nil{
     panic(err)
   }
   return nil
}

func main() {
        defer func(){
         if e:=recover();e!=nil {
            fmt.Println("hatamız var",e)
         }
        }()
	controlAdd("doga","ozcan")  
        controlAdd("murat","ozcan")
        controlAdd("selda","ozcan")
        showNodes()     
}
