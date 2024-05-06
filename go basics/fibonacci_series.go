package main
import(
  "fmt"
)
func main(){
  var c,a int
  b:=0
  d:=1
  fmt.Println("enter the number of iterations")
  fmt.Scanln(&c)
  fmt.Println(b)
  fmt.Println(d)
  for a=2;a<c;a++{
    a=b+d
    fmt.Println(a)
    b=d
    d=a
  }
}