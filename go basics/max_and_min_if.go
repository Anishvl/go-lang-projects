package main
import(
  "fmt"
)
func main(){
  var a,b int
  fmt.Println("enter two values for max and min values")
  fmt.Scanln(&a)
  fmt.Scanln(&b)
  if a<b{
    fmt.Println(b,"is greater than",a)
  }else{
    fmt.Println(a,"is greater than",b)
  }
}