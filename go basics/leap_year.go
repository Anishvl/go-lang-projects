package main
import(
  "fmt"
)
func main(){
  fmt.Println("Enter the year to say its leap year or not")
  var a int
  fmt.Scanln(&a)
  if a%400 == 0 {
      fmt.Println(a, "is a leap year")
  } else if a%100 == 0 {
      fmt.Println(a, "is a leap year")
  } else if a%4 == 0 {
      fmt.Println(a, "is a leap year")
  } else {
      fmt.Println(a, "is not a leap year")
  }

}