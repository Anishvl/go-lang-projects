package main
import(
  "fmt"
)
func main(){
  var a,b float64
  fmt.Println("Enter two numbers here to calculate")
  fmt.Scanln(&a)
  fmt.Scanln(&b)
  fmt.Println("Enter the operation here")
  var opr string
  fmt.Scanln(&opr)
  if opr=="+"{
    fmt.Println(a,"+",b,"is",a+b)
  }else if(opr=="-"){
    fmt.Println(a,"-",b,"is",a-b)
  }else if(opr=="*"){
    fmt.Println(a,"*",b,"is",a*b)
  }else if(opr=="/"){
    fmt.Println(a,"/",b,"is",a/b)
  }else{
    fmt.Println("The Entered operations is invalid")
  }
}