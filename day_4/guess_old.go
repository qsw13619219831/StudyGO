package main

import "fmt"
func main(){
 fmt.Println("guess")
 for {
  age := 0
  fmt.Scanf("%d\n", &age)
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