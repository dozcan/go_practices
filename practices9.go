package main

import (
	"fmt"
)


type Person struct {
     Title string 
     Forenames []string 
     Surname string 
}

type Author1 struct {
     Names Person 
     Title []string 
     YearBorn int 
}

type Author2 struct {
     Person 
     Title []string 
     YearBorn int
   

}

func main(){
        authDoga := Author1{Person{"doga",[]string{"forenames1","forenames2"},"ozcan"},[]string{"title1","title2"},1993}
        authDoga2 := Author2{Person{"doga",[]string{"forenames1","forenames2"},"ozcan"},[]string{"title1","title2"},1993}
        fmt.Println(authDoga.Names.Title)
        fmt.Println(authDoga2.Surname )
}

**************************embedding*********************

type Tasks struct {
slice []string // Named field (aggregation)
Count // Anonymous field (embedding)
}
func (tasks *Tasks) Add(task string) {
tasks.slice = append(tasks.slice, task)
tasks.Increment() // As if we had written: tasks.Count.Increment()
}
