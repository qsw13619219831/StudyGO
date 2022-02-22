package main

import "fmt"
func main(){
 fmt.Println("guess")

 for {
    var age int
  fmt.Scanf("%d", &age)
  if age > 24 {
   fmt.Println("too old", age)
  } else if age < 24 {
   fmt.Println("too young", age)
  } else {
   fmt.Println("good bye", age)
   break
  }
 }
}