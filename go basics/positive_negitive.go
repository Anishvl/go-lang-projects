package main
import(
  "fmt"
)
func main(){
  var a int
  fmt.Println("enter a number here to say +ve or -ve")
  fmt.Scanln(&a)
  if a<0{
    fmt.Println("Its a negative number")
  }else{
    fmt.Println("Its a postive number")
  }
}