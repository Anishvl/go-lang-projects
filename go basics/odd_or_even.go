package main
import(
  "fmt"
)
func main(){
  var a int
  fmt.Println("Enter a number for odd or even")
  fmt.Scanln(&a)
  if a%2==0{
    fmt.Println(a,"is a even number")
  }else{
    fmt.Println(a,"is a odd number")

  }
}