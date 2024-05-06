package main
import(
  "fmt"
)
func main(){
  var a,b int
  fmt.Println("enter the number of rows and columns")
  fmt.Scanln(&a)
  fmt.Scanln(&b)
  arr:=make([][]int,a)
  for i := range arr {
    arr[i] = make([]int,b)
  }
  for i:=0;i<a;i++{
    for j:=0;j<b;j++{
      var c int
      fmt.Println("enter the value of",i,j)
      fmt.Scanln(&c)
      arr[i][j]=c
    }
  }
  fmt.Println("The entered values are:")
  for i:=0;i<a;i++{
    for j:=0;j<b;j++{
      fmt.Printf("%d",arr[i][j])
    }
    fmt.Printf("\n")
  }
}