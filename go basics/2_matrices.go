package main
import(
  "fmt"
)
func main(){
  var c int
  //a:=3
  //b:=3
  fmt.Println("Enter the values for the 3x3 matrix matrix 1")
  var arr [3][3] int

  //var arr1 [b] int
  for i:=0;i<3;i++{
    for j:=0;j<3;j++{
      fmt.Scanln(&c)
      arr[i][j]=c
    }
  }
  fmt.Println("Enter the values for the 3x3 matrix matrix 1")
  var arr1 [3][3] int
  //var arr1 [b] int
  for i:=0;i<3;i++{
    for j:=0;j<3;j++{
      fmt.Scanln(&c)
      arr1[i][j]=c
    }
  }
  fmt.Println("First matrix")
  for i:=0;i<3;i++{
    for j:=0;j<3;j++{
      fmt.Printf("%d",arr[i][j])
    }
    fmt.Printf("\n")
  }

  fmt.Println("Second matrix")
  for i:=0;i<3;i++{
    for j:=0;j<3;j++{
      fmt.Printf("%d",arr1[i][j])
    }
    fmt.Printf("\n")
  }
}