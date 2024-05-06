package main
import(
  "fmt"
)
func main(){
  var c int
  //a:=3
  //b:=3
  fmt.Println("enter the number of rows and columns")
  var arr [3][3] int
  //var arr1 [b] int
  for i:=0;i<3;i++{
    for j:=0;j<3;j++{
      fmt.Scanln(&c)
      arr[i][j]=c
    }
  }
  for i:=0;i<3;i++{
    for j:=0;j<3;j++{
      fmt.Printf("%d",arr[i][j])
    }
    fmt.Printf("\n")
  }

}