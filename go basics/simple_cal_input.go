package main
import(
  "fmt"
)
func main(){
  var a int
  var b int
  fmt.Println("enter two numbers for calculation")
  fmt.Scan(&a)
  fmt.Scan(&b)
  fmt.Println("")
  fmt.Println(a+b)
  fmt.Println(a-b)
  fmt.Println(a*b)
  fmt.Println(a/b)
  fmt.Println(a%b)
}