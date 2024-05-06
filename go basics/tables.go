package main
import(
  "fmt"
)
func main(){
  var a int
  fmt.Println("enter the number for the table")
  fmt.Scanln(&a)
  fmt.Println("")
  for i:=0;i<=10;i++{
    fmt.Println(a,"*",i,"=",a*i)
  }
}