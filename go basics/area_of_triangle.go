package main
import(
  "fmt"
)
func main(){
  var b,h float64
  fmt.Println("Enter the base and height")
  fmt.Scanln(&b,&h)
  fmt.Println("The area of triangle is",0.5*b*h)
}