package main
import(
  "fmt"
)
func main(){
  var a, b int
  fmt.Println("Enter two numbers to swap:")
  fmt.Scanln(&a)
  fmt.Scanln(&b)
  fmt.Println("The values before swapping are:", a, b)
  a=a+b
  b=a-b
  a=a-b
  fmt.Println("The values after swapping are:", a, b)
}